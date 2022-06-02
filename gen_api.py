#!/usr/bin/env python3

import re
import shutil
from dataclasses import dataclass, field
from pathlib import Path
from typing import Optional

import requests


def get(path: str):
    return requests.get(
        f'https://developers.wargaming.net/api/{path}/?realm=all', headers={'X-Requested-With': 'XMLHttpRequest'}
    ).json()


def realm_to_go(realm: str) -> str:
    return {
        "asia": "RealmAsia",
        "eu": "RealmEu",
        "na": "RealmNa",
        "ru": "RealmRu",
        "wgcb": "RealmWgcb",
    }[realm]


def type_to_go(typ: str) -> str:
    return {
        "associative array": "map[string]string",
        "block_header": "struct",
        "boolean": "bool",
        "float": "float32",
        "list of integers": "[]int",
        "numeric": "int",
        "string": "string",
        "timestamp": "wgnTime.UnixTime",
        "list of strings": "[]string",
        "list of timestamps": "[]wgnTime.UnixTime",
        "list of dicts": "map[string]string",
        "object": "struct",
        "timestamp/date": "wgnTime.UnixTime",

        # used in parameters
        "numeric, list": "[]int",
        "string, list": "[]string",
    }[typ]


def go_pointer(typ: str) -> str:
    if typ.startswith("map[") or typ.startswith("[]") or typ.startswith("*"):
        return typ
    return f"*{typ}"


def go_imports(imports: set[str]) -> str:
    if imports:
        import_list = list(imports)
        import_list.sort()
        return "import (\n\t" + "\n\t".join(import_list) + "\n)\n"
    return ""


def go_dereference(typ: str, val: str) -> str:
    if typ.startswith("*"):
        return f'*{val}'
    return val


def go_value_to_string(typ: str, val: str) -> str:
    match typ:
        case "int" | "*int":
            return f'strconv.Itoa({go_dereference(typ, val)})'
        case "[]int":
            return f'internal.SliceIntToString({go_dereference(typ, val)}, ",")'
        case "[]string":
            return f'strings.Join({go_dereference(typ, val)}, ",")'
        case "wgnTime.UnixTime" | "*wgnTime.UnixTime":
            return f'strconv.FormatInt({val}.Unix(), 10)'
    return go_dereference(typ, val)


def name_to_camel(name: str) -> str:
    new_name = name.replace("_", " ").title().replace(" ", "")
    if new_name == "Type":
        return "Type_"
    return new_name


def name_to_camel_lower(name: str) -> str:
    if len(name) == 0:
        return name
    new_name = name_to_camel(name)
    return new_name[0].lower() + new_name[1:]


def camel_to_snake(name: str) -> str:
    return ''.join(['_' + char.lower() if char.isupper() else char for char in name]).lstrip('_')


rxHtml = re.compile("<.*?>")


def clean_documentation(text: str) -> str:
    return rxHtml.sub("", text).strip("\n").replace("&mdash;", "-")


def comment_documentation(text: str) -> str:
    return "// " + text.replace('\n', "\n// ")


# Api method parameter.
@dataclass
class Parameter:
    description: str = ""
    name: str = ""
    required: bool = False
    typ: str = ""

    def __lt__(self, other):
        return self.required > other.required or self.name < other.name


class Parameters:
    def __init__(self, parameters: list[dict]):
        self.params = []
        for cur_param in parameters:
            if cur_param['name'] == "application_id":
                continue
            self.params.append(Parameter(
                description=clean_documentation(cur_param['description']),
                name=name_to_camel_lower(cur_param['name']),
                required=cur_param.get('required', False),
                typ=type_to_go(cur_param['type']),
            ))
        self.params.sort()

    def filter(self, required: bool) -> 'Parameters':
        params = Parameters([])
        for param in self.params:
            if param.required == required:
                params.params.append(param)
        return params

    def method_params(self) -> str:
        parts = []
        for param in self.params:
            parts.append(f"{param.name} {param.typ}")
        if parts:
            return ', ' + ', '.join(parts)
        return ""

    def imports(self) -> set[str]:
        imports = set()
        for param in self.params:
            match param.typ:
                case "context.Context":
                    imports.add('"context"')
                case "int":
                    imports.add('"strconv"')
                case "[]int":
                    imports.add('"github.com/IceflowRE/go-wargaming/wargaming/internal"')
                case "[]string":
                    imports.add('"strings"')
                case "wgnTime.UnixTime":
                    imports.add('"strconv"')
                    imports.add('"github.com/IceflowRE/go-wargaming/wargaming/wgnTime"')
        return imports

    def req_map(self) -> str:
        lines = []
        for param in self.params:
            lines.append(f'\t\t"{camel_to_snake(param.name).rstrip("_")}": {go_value_to_string(param.typ, param.name)},')
        return '\n'.join(lines)

    def doc(self) -> str:
        lines = []
        for param in self.params:
            lines.append(f'{name_to_camel_lower(camel_to_snake(param.name))}:')
            lines.append("    " + param.description.replace("\n", "\n    "))
        return '\n'.join(lines)


@dataclass
class Struct:
    name: str = ""
    description: str = ""
    fields: list = field(default_factory=list)
    typ: str = ""
    json_tag: Optional[str] = None

    def __lt__(self, other):
        return self.name > other.name

    def __hash__(self):
        return hash(self.name)

    def __eq__(self, other):
        if isinstance(other, self.__class__):
            return self.name == other.name and self.description == other.description and self.fields == other.fields and self.typ == other.typ
        elif isinstance(other, str):
            return self.name == other
        return False

    def __ne__(self, other):
        return not self.__eq__(other)

    def __getitem__(self, item):
        for struct in self.fields:
            if struct.name == item:
                return struct
        raise ValueError()

    def __contains__(self, item):
        for struct in self.fields:
            if struct.name == item:
                return True
        return False

    def empty(self) -> bool:
        return len(self.fields) == 0

    def get(self, item, default=None):
        for struct in self.fields:
            if struct.name == item:
                return struct
        return default

    def _get_idx(self, name):
        for idx, struct in enumerate(self.fields):
            if struct.name == name:
                return idx
        return -1

    @classmethod
    def from_parameters(cls, name: str, params: Parameters) -> 'Struct':
        base = cls(name=name, typ="struct", description=f"{name} options.")
        for param in params.params:
            field = Struct(
                name=name_to_camel(camel_to_snake(param.name)),
                description=param.description,
                typ=go_pointer(param.typ),
            )
            base.fields.append(field)
        base.fields.sort()
        return base

    # fields may contain nested structure in an unsorted order!
    # You may get the deepest item before the items above.
    @classmethod
    def from_response_doc(cls, name: str, fields: list) -> 'Struct':
        base = cls(name=name, typ="struct", json_tag=camel_to_snake(name))
        for cur_field in fields:
            typ = cur_field['type']
            if typ == 'empty_line':
                continue
            go_field: Struct = Struct(
                name=name_to_camel(cur_field['name'][-1]),
                json_tag=cur_field['name'][-1],
                description=clean_documentation(cur_field['description']),
                typ=go_pointer(type_to_go(typ)),
            )
            # get nested struct
            cur_struct: Struct = base
            max_depth: int = len(cur_field['name']) - 1
            for idx, field_name in enumerate(cur_field['name']):
                field_name = name_to_camel(field_name)
                if idx == max_depth:
                    break
                new_struct = cur_struct.get(field_name)
                if new_struct is None:
                    new_struct = Struct(name=field_name)
                    cur_struct.fields.append(new_struct)
                cur_struct = new_struct
            # if already existing update values
            new_struct: Struct = cur_struct.get(go_field.name)
            if new_struct is None:
                cur_struct.fields.append(go_field)
            else:
                new_struct.description = go_field.description
                new_struct.typ = go_field.typ
                new_struct.json_tag = go_field.json_tag
        base.fields.sort()
        return base

    def to_code(self, is_root: bool = False) -> tuple[set[str], str]:
        imports = set()
        if "wgnTime.UnixTime" in self.typ:
            imports.add('"github.com/IceflowRE/go-wargaming/wargaming/wgnTime"')
        lines = []
        if self.description:
            documentation = self.description.replace('\n', "\n// ")
            lines.append(f"// {documentation}")
        json_tag = ""
        if not is_root and self.json_tag is not None:
            json_tag = f' `json:"{self.json_tag},omitempty"`'
        if "struct" in self.typ:
            if is_root:
                lines.append(f"type {self.name} {self.typ} {'{'}")
            else:
                lines.append(f"{self.name} {self.typ} {{")
            fields = self.fields
            fields.sort(reverse=True)
            for field in fields:
                field_imports, field_code = field.to_code()
                field_str = "\t" + field_code.replace("\n", "\n\t")
                lines.append(field_str)
                imports.update(field_imports)
            lines.append(f'}}{json_tag}')
        else:
            lines.append(f'{self.name} {self.typ}{json_tag}')
        return imports, '\n'.join(lines)


@dataclass
class ServiceMethod:
    name: str
    allowed_realms: list[str]
    documentation: str
    http_methods: list[str]
    # method return type
    return_type: str
    # base imports
    imports: set[str] = field(default_factory=set)
    # type which is used to unmarshal
    # if empty return_type is used
    unmarshal_type: str = ""
    # post-processing code has to include the return part if used
    post_processing: str = ""


class GoApi:
    def __init__(self, game: str, method_id: str, data: dict):
        self.method_id: str = method_id
        self.game: str = name_to_camel_lower(game)
        self.params: Parameters = Parameters(data['parameters'])
        self.url = data['url']
        if not self.url.startswith("/"):
            self.url = "/" + self.url

        self.service_method: ServiceMethod = ServiceMethod(
            name=name_to_camel(self.method_id),
            allowed_realms=[realm_to_go(realm) for realm in data['available_display_indices']],
            documentation=clean_documentation(data['description']),
            http_methods=data['allowed_http_methods'],
            return_type="",
            imports={'"context"'}
        )
        if self.method_id.startswith(f"{game}_"):
            self.service_method.name = name_to_camel(self.method_id[len(game) + 1:])

        self.response_struct: Struct = Struct.from_response_doc(self.service_method.name, data.get('fields', []))
        self.options_struct: Struct = Struct.from_parameters(f"{self.service_method.name}Options",
                                                             self.params.filter(required=False))

        # has fields (returns something)
        if not self.response_struct.empty():
            self.service_method.return_type = f"*{self.game}.{self.response_struct.name}"
        self.deprecated: bool = data.get('deprecated', False)

    def method_imports(self) -> set[str]:
        imports = self.params.imports()
        imports.update(self.service_method.imports)
        if not self.response_struct.empty():
            imports.add(f'"github.com/IceflowRE/go-wargaming/wargaming/{self.game}"')
        success = False
        for param in self.params.params:
            if param.typ == "wgnTime.UnixTime" and param.required:
                success = True
                break
        if not success:
            imports.discard('"github.com/IceflowRE/go-wargaming/wargaming/wgnTime"')
        return imports

    def method_file(self) -> str:
        documentation = f"""// {self.service_method.name} {comment_documentation(self.service_method.documentation)[3:]}
//
// https://developers.wargaming.net/reference/all/{self.method_id.replace('_', '/')}"""
        if self.deprecated:
            documentation += "\n//\n// Deprecated: Attention! The method is deprecated."
        params_doc: str = self.params.filter(required=True).doc()
        if self.service_method.allowed_realms or params_doc:
            documentation += "\n//"
        if self.service_method.allowed_realms:
            valid_realms: str = ", ".join([realm for realm in self.service_method.allowed_realms])
            documentation += f"\n// realm:\n//     Valid realms: {valid_realms}"
        if params_doc:
            documentation += "\n" + comment_documentation(params_doc)

        meth_return_type = f"({self.service_method.return_type}, error)" if self.service_method.return_type else "error"
        allowed_realms: str = ', '.join(self.service_method.allowed_realms)
        param_options: str = "" if self.options_struct.empty() else f", options *{self.game}.{self.options_struct.name}"
        opt_req_map: str = ""
        for field in self.options_struct.fields:
            opt_req_map += f'\t\tif options.{field.name} != nil {{\n\t\t\treqParam["{camel_to_snake(field.name).rstrip("_")}"] = {go_value_to_string(field.typ, f"options.{field.name}")}\n\t\t}}\n'
        if opt_req_map:
            opt_req_map = "\tif options != nil {\n" + opt_req_map + "\t}\n"

        req_method = "getRequest" if "GET" in self.service_method.http_methods else "postRequest"
        if self.service_method.post_processing:
            post_processing = self.service_method.post_processing
        elif self.service_method.return_type:
            post_processing = "\treturn data, err"
        else:
            post_processing = "\treturn err"

        if self.service_method.return_type:
            unmarshal_type: str = self.service_method.unmarshal_type if self.service_method.unmarshal_type else self.service_method.return_type
            request_part = f"""
\tvar data {unmarshal_type}
\terr := service.client.{req_method}(ctx, section{name_to_camel(camel_to_snake(self.game))}, realm, "{self.url}/", reqParam, &data)
{post_processing}"""
        else:
            request_part = f"""
\terr := service.client.{req_method}(ctx, section{name_to_camel(camel_to_snake(self.game))}, realm, "{self.url}/", reqParam, nil)
{post_processing}"""

        return f"""package wargaming

{go_imports(self.method_imports())}
{documentation}
func (service *{self.game}Service) {self.service_method.name}(ctx context.Context, realm Realm{self.params.filter(True).method_params()}{param_options}) {meth_return_type} {{
\tif err := validateRealm(realm, []Realm{{{allowed_realms}}}); err != nil {{
\t\t{'return nil' if self.response_struct.empty() else 'return nil, err'}
\t}}

\treqParam := map[string]string{{
{self.params.filter(True).req_map()}
\t}}
{opt_req_map}{request_part}
}}
"""


def check_unit_test(unit_test: str, go: GoApi):
    # ignore deprecated methods
    test_name: str = name_to_camel(go.method_id)
    is_tested: bool = test_name in unit_test
    if go.deprecated:
        if is_tested:
            print("    > Method is deprecated, but still tested. Remove the test.")
        return
    # ignore to complex methods
    if go.method_id in [
        "wot_stronghold_activateclanreserve",
        "wot_auth_login",
        "wotb_clanmessages_create", "wotb_clanmessages_delete", "wotb_clanmessages_like", "wotb_clanmessages_likes",
        "wotb_clanmessages_update",
        "wgn_wargag_rate", "wgn_wargag_newcomment", "wgn_wargag_deletecomment",
        "wotb_tournaments_info", "wotb_tournaments_teams", "wotb_tournaments_stages", "wotb_tournaments_matches",
        "wotb_tournaments_standings", "wotb_tournaments_tables",
        "wotx_account_xuidinfo", "wotx_account_psninfo",
        "wotx_auth_login",
    ]:
        if not is_tested:
            print(
                f'    > Method is not tested and complex. Add a test case with content \'skipTest("{test_name}", reasonTooComplex, test)\'.')
            return
        if f'skipTest("{test_name}"' not in unit_test:
            print("    > Method is marked as complex and tested. Remove it from the complex list.")
        return
    if not is_tested:
        print("    > Method is not tested. Add a test case.")


# Fix inaccurate wargaming documentation.
def fix_api(go: GoApi):
    def fix_struct(struct: Struct):
        for field in struct.fields:
            if field.typ == "struct":
                fix_struct(field)
            if field.name == "ExpiresAt":
                if field.typ == "int":
                    field.typ = "wgnTime.UnixTime"
                if field.typ == "*int":
                    field.typ = "*wgnTime.UnixTime"

    fix_struct(go.response_struct)
    fix_struct(go.options_struct)
    for param in go.params.params:
        if param.name == "expiresAt":
            if param.typ == "int":
                param.typ = "wgnTime.UnixTime"
            if param.typ == "*int":
                param.typ = "*wgnTime.UnixTime"

    match go.method_id:
        case "wot_account_info":
            go.response_struct['Statistics']['Frags'].typ = "map[string]int"
            go.response_struct['Private']['Boosters'].typ = "map[string]*struct"
        case "wot_encyclopedia_vehicles":
            go.response_struct['DefaultProfile']['Ammo']['Stun']['Duration'].typ = "[]int"
        case "wot_globalmap_events":
            go.response_struct['Fronts'].typ = "[]*struct"
        case "wot_globalmap_fronts":
            go.response_struct['AvailableExtensions'].typ = "[]*struct"
        case "wot_globalmap_seasons":
            go.response_struct['Fronts'].typ = "[]*struct"
        case "wot_tanks_mastery":
            go.response_struct['Distribution'].typ = "map[string]map[string]int"
        case "wotb_encyclopedia_modules":
            go.response_struct['Engines'].typ = "[]*struct"
            go.response_struct['Guns'].typ = "[]*struct"
            go.response_struct['Guns']['Shells'].typ = "[]*struct"
            go.response_struct['Suspensions'].typ = "[]*struct"
            go.response_struct['Turrets'].typ = "[]*struct"
        case "wgn_wgtv_tags":
            go.response_struct['Categories'].typ = "[]*struct"
            go.response_struct['Projects'].typ = "[]*struct"
            go.response_struct['Programs'].typ = "[]*struct"
        case 'wot_clans_messageboard' | 'wotb_clanmessages_messages':
            go.service_method.return_type = f"map[string][]{go.service_method.return_type}"
        case 'wotb_tournaments_teams':
            go.response_struct['Players'].typ = "[]*struct"
    if go.method_id in [
        'wotb_tournaments_info',
        'wowp_encyclopedia_planes',
    ]:
        go.service_method.return_type = f"map[string]{go.service_method.return_type}"
    if go.method_id in [
        'wgn_account_list', 'wgn_wgtv_videos',
        'wot_account_list', 'wot_clans_list',
        'wot_clanratings_neighbors', 'wot_clanratings_top', 'wot_globalmap_clanbattles',
        'wot_globalmap_eventaccountratingneighbors', 'wot_globalmap_eventaccountratings', 'wot_globalmap_eventrating',
        'wot_globalmap_eventratingneighbors', 'wot_globalmap_events', 'wot_globalmap_fronts', 'wot_globalmap_provinces',
        'wot_globalmap_seasonrating', 'wot_globalmap_seasonratingneighbors', 'wot_globalmap_seasons',
        'wot_stronghold_clanreserves',
        'wotb_account_list', 'wotb_clans_list', 'wotb_clanmessages_likes', 'wotb_tournaments_list',
        'wotb_tournaments_teams', 'wotb_tournaments_stages', 'wotb_tournaments_standings', 'wotb_tournaments_tables',

        'wotx_account_list', 'wotx_clans_list', 'wotx_account_xuidinfo', 'wotx_account_psninfo',
        'wowp_account_list', 'wowp_clans_list', 'wowp_ratings_neighbors',
        'wowp_ratings_top',
        'wows_account_list', 'wows_clans_list',
    ]:
        go.service_method.return_type = f"[]{go.service_method.return_type}"


def gen_api(output: Path, games: dict):
    counter: int = 0
    max_methods: int = 0
    for game in games:
        for section in game['sections']:
            max_methods += len(section['methods'])

    api: list[GoApi] = []
    unit_test: str = output.joinpath("client_test.go").read_text(encoding='utf-8')
    # collect api methods
    for game in games:
        if not game['sections']:
            continue
        game_name = game['slug']
        output.joinpath(game_name).mkdir(parents=True, exist_ok=True)
        for section in game['sections']:
            for method in section['methods']:
                counter += 1
                print(f"Collecting ({str(counter).zfill(len(str(max_methods)))}/{max_methods}) {method['key']}")

                raw_data = get(f"methods/{method['key']}")['data']
                go = GoApi(game_name, method['key'], raw_data)
                fix_api(go)
                api.append(go)
    # write source code
    # write service methods and structs
    counter = 0
    for go in api:
        counter += 1
        print(f"Generating ({str(counter).zfill(len(str(max_methods)))}/{max_methods}) {go.method_id}")
        output.joinpath(f"{go.method_id}.go").write_text(go.method_file(), encoding='utf-8')
        if not go.response_struct.empty() or not go.options_struct.empty():
            imports, opt_code = go.options_struct.to_code(True)
            imports_tmp, resp_code = go.response_struct.to_code(True)
            imports.update(imports_tmp)
            if go.options_struct.empty():
                opt_code = ""
            if go.response_struct.empty():
                resp_code = ""

            if not go.response_struct.empty() and not go.options_struct.empty():
                opt_code += "\n\n"
            output.joinpath(go.game).joinpath(f"{go.method_id[len(f'{go.game}_'):]}.go").write_text(
                f"package {go.game}\n\n{go_imports(imports)}\n{opt_code}{resp_code}\n",
                encoding='utf-8'
            )
        check_unit_test(unit_test, go)
    # service stuff
    service_types: list[str] = []
    service_fields: list[str] = []
    service_init: list[str] = []
    for game in games:
        game_name = name_to_camel(game['slug'])
        service_name: str = f"{name_to_camel_lower(game['slug'])}Service"
        service_types.append(f"// {game['name']} service.\ntype {service_name} service\n")
        service_fields.append(f"\t{game_name} *{service_name}")
        service_init.append(f"\tclient.{game_name} = (*{service_name})(&client.common)")
    service_types.sort()
    service_fields.sort()
    service_init.sort()
    # write service types
    tmp_text: str = '\n'.join(service_types)
    output.joinpath(f"services.go").write_text(f"package wargaming\n\n{tmp_text}", encoding='utf-8')
    # write service fields and initialization
    client: str = output.joinpath(f"client.go").read_text(encoding='utf-8')
    tmp_text: str = '\n'.join(service_fields)
    delimiter = ["// AUTO GENERATION START FIELDS\n", "\n\t// AUTO GENERATION END FIELDS"]
    client = re.sub(f'{delimiter[0]}.*{delimiter[1]}', f"{delimiter[0]}{tmp_text}{delimiter[1]}", client,
                    flags=re.DOTALL)
    tmp_text: str = '\n'.join(service_init)
    delimiter = ["// AUTO GENERATION START INIT\n", "\n\t// AUTO GENERATION END INIT"]
    client = re.sub(f'{delimiter[0]}.*{delimiter[1]}', f"{delimiter[0]}{tmp_text}{delimiter[1]}", client,
                    flags=re.DOTALL)
    output.joinpath(f"client.go").write_text(client, encoding='utf-8')

    deprecated_files = []
    for go in api:
        if go.deprecated:
            deprecated_files.append(f'    - "wargaming/{go.method_id}"')
    if deprecated_files:
        ignore = '\n'.join(deprecated_files)
        Path("./codecov.yml").write_text(f"ignore:\n{ignore}\n")


if __name__ == '__main__':
    output = Path("./wargaming/")

    games = get('methods')['data']
    prefixes = [f"{game['slug']}" for game in games]
    prefixes.append("services")
    # remove files
    print("Remove existing files.")
    for path in output.iterdir():
        if path.name == "wgnTime":
            continue
        if any([path.name.startswith(prefix) for prefix in prefixes]):
            if path.is_dir():
                shutil.rmtree(path)
            else:
                path.unlink()
    # gen api files
    gen_api(output, games)
