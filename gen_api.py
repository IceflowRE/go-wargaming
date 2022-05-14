#!/usr/bin/env python3

import re
from dataclasses import dataclass, field
from pathlib import Path
from typing import Optional

import requests


def get(path):
    return requests.get(f'https://developers.wargaming.net/api/{path}/?realm=all',
                        headers={'X-Requested-With': 'XMLHttpRequest'}).json()


realmToGo = {
    "asia": "RealmAsia",
    "eu": "RealmEu",
    "na": "RealmNa",
    "ru": "RealmRu",
    "wgcb": "RealmWgcb",
}

typeToGo = {
    "associative array": "map[string]string",
    "block_header": "struct",
    "boolean": "bool",
    "float": "float32",
    "list of integers": "[]int",
    "numeric": "int",
    "string": "string",
    "timestamp": "UnixTime",
    "list of strings": "[]string",
    "list of timestamps": "[]UnixTime",
    "list of dicts": "map[string]string",
    "object": "struct",
    "timestamp/date": "UnixTime",

    # used in parameters
    "numeric, list": "[]int",
    "string, list": "[]string",
}

rxHtml = re.compile("<.*?>")


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


@dataclass
class Field:
    description: str = ""
    fields: dict[str, 'Field'] = field(default_factory=dict)
    typ: str = ""

    def __setitem__(self, key, value):
        self.fields[key] = value

    def __getitem__(self, item):
        return self.fields[item]

    @classmethod
    def from_list(cls, fields: list) -> 'Field':
        go_fields = cls()
        for cur_field in fields:
            typ = cur_field['type']
            if typ == 'empty_line':
                continue
            go_field: Field = Field(
                description=rxHtml.sub("", cur_field['description']).rstrip("\n").replace('\n', "\n// "),
                fields={},
                typ=typeToGo[typ]
            )
            cur_go_field: Field = go_fields
            for idx in range(len(cur_field['name'])):
                name: str = cur_field['name'][idx]
                if idx == len(cur_field['name']) - 1:
                    cur_go_field.fields[name] = go_field
                    continue
                if name not in cur_go_field.fields:
                    cur_go_field.fields[name] = Field()
                cur_go_field = cur_go_field.fields[name]
        return go_fields

    @staticmethod
    def to_str(base: 'Field', lines: Optional[list[str]] = None, depth: int = 1) -> str:
        if lines is None:
            lines = []
        tabs: str = '\t' * depth
        for name, cur_field in base.fields.items():
            doc = cur_field.description.replace("\n", f"\n{tabs}")
            lines.append(f"{tabs}// {doc}")
            if "struct" in cur_field.typ:
                lines.append(f"{tabs}{name_to_camel(name)} {cur_field.typ} {'{'}")
                Field.to_str(cur_field, lines, depth + 1)
                lines.append(f"{tabs}{'}'} `json:\"{name},omitempty\"`")
            else:
                lines.append(f"{tabs}{name_to_camel(name)} {cur_field.typ} `json:\"{name},omitempty\"`")
        if depth != 1:
            return ""
        return '\n'.join(lines)

    def fix(self, method: str) -> 'Field':
        match method:
            case "wot_account_info":
                self.fields['statistics']['frags'].typ = "map[string]int"
                self.fields['private']['boosters'].typ = "map[string]struct"
            case "wot_encyclopedia_vehicles":
                self.fields['default_profile']['ammo']['stun']['duration'].typ = "[]int"
        return self


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
        params = Parameters._from_list_dict(parameters)

        self.doc: str = Parameters._to_doc(params)
        self.imports: str = Parameters._to_imports(params)
        self.method_params: str = Parameters._to_method_params(params)
        self.req_map: str = Parameters._to_req_map(params)

    @staticmethod
    def _from_list_dict(parameters: list[dict]) -> list['Parameter']:
        params = []
        for cur_param in parameters:
            if cur_param['name'] == "application_id":
                continue
            params.append(Parameter(
                description=rxHtml.sub("", cur_param['description']).rstrip("\n").replace('\n', "\n//     "),
                name=cur_param['name'],
                required=cur_param.get('required', False),
                typ=typeToGo[cur_param['type']]
            ))
        params.sort()
        return params

    @staticmethod
    def _to_method_params(params: list['Parameter']) -> str:
        parts = []
        for param in params:
            parts.append(f"{name_to_camel_lower(param.name)} {param.typ}")
        if parts:
            return ', ' + ', '.join(parts)
        return ""

    @staticmethod
    def _to_imports(params: list['Parameter']) -> str:
        imports = set()
        for param in params:
            match param.typ:
                case ("int" | "UnixTime"):
                    imports.add('"strconv"')
                case "[]string":
                    imports.add('"strings"')
                case "[]int":
                    imports.add('"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"')
        if imports:
            return '\t' + '\n\t'.join(imports)
        return ""

    @staticmethod
    def _to_req_map(params: list['Parameter']) -> str:
        lines = []
        for param in params:
            match param.typ:
                case "int":
                    lines.append(f'\t\t"{param.name}": strconv.Itoa({name_to_camel_lower(param.name)}),')
                case "[]string":
                    lines.append(f'\t\t"{param.name}": strings.Join({name_to_camel_lower(param.name)}, ","),')
                case "[]int":
                    lines.append(f'\t\t"{param.name}": utils.SliceIntToString({name_to_camel_lower(param.name)}, ","),')
                case "UnixTime":
                    lines.append(f'\t\t"{param.name}": strconv.FormatInt({name_to_camel_lower(param.name)}.Unix(), 10),')
                case _:
                    lines.append(f'\t\t"{param.name}": {name_to_camel_lower(param.name)},')
        return '\n'.join(lines)

    @staticmethod
    def _to_doc(params: list['Parameter']) -> str:
        lines = []
        for param in params:
            lines.append(f'// {param.name}:')
            lines.append(f"//     {param.description}")
        return '\n'.join(lines)


@dataclass
class GoFile:
    name: str
    realms: str
    documentation: str
    fields: str
    params: Parameters
    http_methods: list
    return_type: str


def function_to_go(key: str, base_return_type: str, data: dict) -> GoFile:
    go_file = GoFile(
        name=key,
        realms=', '.join([realmToGo[realm] for realm in data['available_display_indices']]),
        documentation=rxHtml.sub("", data['description']).rstrip("\n").replace('\n', "\n// "),
        fields=Field.to_str(Field.from_list(data.get('fields', [])).fix(key)),
        params=Parameters(data['parameters']),
        http_methods=data['allowed_http_methods'],
        return_type="",
    )
    if data.get('deprecated', False):
        go_file.documentation = "Deprecated: Attention! The method is deprecated.\n// " + go_file.documentation
    match base_return_type:
        case 'list':
            go_file.return_type = f"[]*{name_to_camel(go_file.name)}"
        case _:
            go_file.return_type = f"*{name_to_camel(go_file.name)}"
    return go_file


def gen_api(output: Path):
    games = get('methods')['data']
    for game_data in games:
        for section in game_data['sections']:
            for method in section['methods']:
                # method = {'key': 'wot_ratings_accounts', 'slug': 'list'}
                print(f"Generating {method['key']}")

                data = get(f"methods/{method['key']}")['data']
                go: GoFile = function_to_go(method['key'], method['slug'], data)

                meth_return_type = f"({go.return_type}, error)" if go.fields else "error"
                return_struct = ""
                if go.fields:
                    return_struct = f"""type {name_to_camel(go.name)} struct {'{'}
{go.fields}
{'}'}
"""
                imports = ""
                if go.params.imports:
                    imports = f"""
import (
{go.params.imports}
)
"""
                if go.fields:
                    req_method = "doGetDataRequest" if "GET" in go.http_methods else "doPostDataRequest"
                    req_part = f"""
\tvar result {go.return_type}
\terr := client.{req_method}(realm, "/{go.name.replace('_', '/')}/", reqParam, &result)
\treturn result, err"""
                else:
                    req_method = "doGetRequest" if "GET" in go.http_methods else "doPostRequest"
                    req_part = f"""
\terr := client.{req_method}(realm, "/{go.name.replace('_', '/')}/", reqParam)
\treturn err"""

                content = f"""package wargaming
{imports}
{return_struct}
// {name_to_camel(go.name)} {go.documentation}
//
// https://developers.wargaming.net/reference/all/{go.name.replace('_', '/')}
//
{go.params.doc}
func (client *Client) {name_to_camel(go.name)}(realm Realm{go.params.method_params}) {meth_return_type} {'{'}
\tif err := ValidateRealm(realm, []Realm{'{'}{go.realms}{'}'}); err != nil {'{'}
\t\t{'return nil, err' if go.fields else 'return nil'}
\t{'}'}

\treqParam := map[string]string{'{'}
{go.params.req_map}
\t{'}'}
{req_part}
{'}'}
"""
                output.joinpath(f"{go.name}.go").write_text(content, encoding='utf-8')


if __name__ == '__main__':
    output = Path("./pkg/wargaming/")
    # remove files
    for path in output.glob('*.go'):
        if any([path.name.startswith(prefix) for prefix in ['wgn', 'wot', 'wotb', 'wotx', 'wowp', 'wows']]):
            path.unlink()
    # gen api files
    gen_api(output)
