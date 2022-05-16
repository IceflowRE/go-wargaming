#!/usr/bin/env python3

import re
import shutil
from dataclasses import dataclass, field
from pathlib import Path

import requests


def get(path):
    return requests.get(f'https://developers.wargaming.net/api/{path}/?realm=all',
                        headers={'X-Requested-With': 'XMLHttpRequest'}).json()


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
    return rxHtml.sub("", text).strip("\n")


def comment_documentation(text: str) -> str:
    return "// " + text.replace('\n', "\n// ")


def get_base_return_type(method: dict) -> str:
    if method['key'].endswith("_account_info"):
        return "map"
    if method['name'].startswith("List of") or method['slug'] == "list":
        return "list"
    return "struct"


@dataclass
class Struct:
    name: str = ""
    description: str = ""
    fields: list = field(default_factory=list)
    typ: str = ""
    json_tag: str = ""

    def __lt__(self, other):
        return self.name < other.name

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

    # fields may contain nested structure in an unsorted order!
    # You may get the deepest item before the items above.
    @classmethod
    def from_doc(cls, name: str, fields: list) -> 'Struct':
        base = cls(name=name, typ="struct", json_tag=camel_to_snake(name))
        if base.name == "type":
            base.name = "typ"
        for cur_field in fields:
            typ = cur_field['type']
            if typ == 'empty_line':
                continue
            go_field: Struct = Struct(
                name=name_to_camel(cur_field['name'][-1]),
                json_tag=cur_field['name'][-1],
                description=clean_documentation(cur_field['description']),
                typ=type_to_go(typ),
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
        return base

    def to_code(self, is_root: bool = False) -> tuple[str, set[str]]:
        imports = set()
        if "wgnTime.UnixTime" in self.typ:
            imports.add('"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"')
        lines = []
        if self.description:
            documentation = self.description.replace('\n', "\n// ")
            lines.append(f"// {documentation}")
        json_tag = f'`json:"{self.json_tag},omitempty"`'
        if is_root:
            json_tag = ""
        if "struct" in self.typ:
            if is_root:
                lines.append(f"type {self.name} {self.typ} {'{'}")
            else:
                lines.append(f"{self.name} {self.typ} {'{'}")
            fields = self.fields
            fields.sort(reverse=True)
            for field in fields:
                field_code, field_imports = field.to_code()
                field_str = "\t" + field_code.replace("\n", "\n\t")
                lines.append(field_str)
                imports.update(field_imports)
            lines.append(f'{"}"} {json_tag}')
        else:
            lines.append(f'{self.name} {self.typ} {json_tag}')
        if is_root and imports:
            imports_l = list(imports)
            imports_l.sort()
            lines.insert(0, "import (\n\t" + "\n\t".join(imports) + "\n)\n")
        return '\n'.join(lines), imports

    # Fix inaccurate wargaming documentation.
    def fix(self, method: str) -> 'Struct':
        match method:
            case "wot_account_info":
                self['Statistics']['Frags'].typ = "map[string]int"
                self['Private']['Boosters'].typ = "map[string]struct"
            case "wot_encyclopedia_vehicles":
                self['DefaultProfile']['Ammo']['Stun']['Duration'].typ = "[]int"
        return self


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
        self._params = []
        for cur_param in parameters:
            if cur_param['name'] == "application_id":
                continue
            self._params.append(Parameter(
                description=rxHtml.sub("", cur_param['description']).strip("\n").replace('\n', "\n    "),
                name=cur_param['name'],
                required=cur_param.get('required', False),
                typ=type_to_go(cur_param['type']),
            ))
        self._params.sort()

    def method_params(self) -> str:
        parts = []
        for param in self._params:
            parts.append(f"{name_to_camel_lower(param.name)} {param.typ}")
        if parts:
            return ', ' + ', '.join(parts)
        return ""

    def imports(self) -> list[str]:
        imports = set()
        for param in self._params:
            match param.typ:
                case ("int"):
                    imports.add('"strconv"')
                case "[]int":
                    imports.add('"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"')
                case "[]string":
                    imports.add('"strings"')
                case "wgnTime.UnixTime":
                    imports.add('"strconv"')
                    imports.add('"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"')

        imports = list(imports)
        imports.sort()
        return imports

    def req_map(self) -> str:
        lines = []
        for param in self._params:
            match param.typ:
                case "int":
                    lines.append(f'\t\t"{param.name}": strconv.Itoa({name_to_camel_lower(param.name)}),')
                case "[]int":
                    lines.append(f'\t\t"{param.name}": utils.SliceIntToString({name_to_camel_lower(param.name)}, ","),')
                case "[]string":
                    lines.append(f'\t\t"{param.name}": strings.Join({name_to_camel_lower(param.name)}, ","),')
                case "wgnTime.UnixTime":
                    lines.append(
                        f'\t\t"{param.name}": strconv.FormatInt({name_to_camel_lower(param.name)}.Unix(), 10),')
                case _:
                    lines.append(f'\t\t"{param.name}": {name_to_camel_lower(param.name)},')
        return '\n'.join(lines)

    def doc(self) -> str:
        lines = []
        for param in self._params:
            lines.append(f'{param.name}:')
            lines.append(f"    {param.description}")
        return '\n'.join(lines)


class GoApi:
    def __init__(self, game: str, method: str, base_return_type: str, data: dict):
        params = Parameters(data['parameters'])

        self.file_name: str = method
        self.allowed_realms: str = ', '.join([realm_to_go(realm) for realm in data['available_display_indices']])
        self.documentation: str = clean_documentation(data['description'])
        self.param_documentation = params.doc()
        self.game: str = game
        self.struct: Struct = Struct.from_doc(name_to_camel(method[len(f"{game}_"):]), data.get('fields', [])).fix(
            method)
        self.method_params: str = params.method_params()
        self.http_req_map: str = params.req_map()
        self.http_methods: list = data['allowed_http_methods']
        self.return_type: str = ""
        self.imports: str = ""
        self.deprecated: bool = data.get('deprecated', False)

        imports = params.imports()
        # has fields (returns something)
        if self.struct.fields:
            imports.append(f'"github.com/IceflowRE/go-wargaming/pkg/wargaming/{self.game}"')
        imports.sort()
        if imports:
            self.imports = "import (\n\t" + "\n\t".join(imports) + "\n)"
        match base_return_type:
            case 'list':
                self.return_type = f"[]*{self.game}.{self.struct.name}"
            case 'map':
                self.return_type = f"map[string]*{self.game}.{self.struct.name}"
            case _:
                self.return_type = f"*{self.game}.{self.struct.name}"

    def method_file(self) -> str:
        meth_return_type = f"({self.return_type}, error)" if self.struct.fields else "error"
        if self.struct.fields:
            req_method = "doGetDataRequest" if "GET" in self.http_methods else "doPostDataRequest"
            req_part = f"""
\tvar result {self.return_type}
\terr := client.{req_method}(realm, "/{self.file_name.replace('_', '/')}/", reqParam, &result)
\treturn result, err"""
        else:
            req_method = "doGetRequest" if "GET" in self.http_methods else "doPostRequest"
            req_part = f"""
\terr := client.{req_method}(realm, "/{self.file_name.replace('_', '/')}/", reqParam)
\treturn err"""
        documentation = f"""\n// {name_to_camel(self.file_name)} {comment_documentation(self.documentation)[3:]}
//
// https://developers.wargaming.net/reference/all/{self.file_name.replace('_', '/')}"""
        if self.deprecated:
            documentation += "\n//\n// Deprecated: Attention! The method is deprecated."
        if self.param_documentation:
            documentation += "\n//\n" + comment_documentation(self.param_documentation)
        return f"""package wargaming

{self.imports}
{documentation}
func (client *Client) {name_to_camel(self.file_name)}(realm Realm{self.method_params}) {meth_return_type} {'{'}
\tif err := ValidateRealm(realm, []Realm{'{'}{self.allowed_realms}{'}'}); err != nil {'{'}
\t\t{'return nil, err' if self.struct.fields else 'return nil'}
\t{'}'}

\treqParam := map[string]string{'{'}
{self.http_req_map}
\t{'}'}
{req_part}
{'}'}
"""

    def struct_file(self) -> str:
        code, _ = self.struct.to_code(True)
        return f"package " + str(self.game) + "\n\n" + code + "\n"


def gen_go_api(game_name: str, method: str, base_return_type: str) -> GoApi:
    data = get(f"methods/{method}")['data']
    return GoApi(game_name, method, base_return_type, data)


def gen_api(output: Path, games: dict):
    counter = 0
    max_methods = 0
    for game in games:
        for section in game['sections']:
            max_methods += len(section['methods'])

    for game in games:
        game_name = game['slug']
        if game['sections']:
            output.joinpath(game_name).mkdir(parents=True, exist_ok=True)
        for section in game['sections']:
            for method in section['methods']:
                counter += 1
                # method = {'key': 'wot_ratings_accounts', 'slug': 'list'}
                print(f"Generating ({str(counter).zfill(len(str(max_methods)))}/{max_methods}) {method['key']}")

                go: GoApi = gen_go_api(game_name, method['key'], get_base_return_type(method))
                output.joinpath(f"{go.file_name}.go").write_text(go.method_file(), encoding='utf-8')
                if go.struct.fields:
                    output.joinpath(game_name).joinpath(f"{go.file_name[len(f'{go.game}_'):]}.go").write_text(
                        go.struct_file(), encoding='utf-8')


if __name__ == '__main__':
    output = Path("./pkg/wargaming/")

    games = get('methods')['data']
    prefixes = [game['slug'] for game in games]
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
