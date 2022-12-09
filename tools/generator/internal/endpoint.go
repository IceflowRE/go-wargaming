package internal

import (
	"sort"
	"strings"
)

type endpoint struct {
	Id   string
	Game string
	Url  string

	Name        string
	Parameters  []*goType
	OptionsType *goType
	DataType    *goType
	// has to be added via patches, not retrievable from the api documentation
	MetaType      *goType
	AllowedRealms []string
	Documentation string
	Deprecated    bool
	HttpMethods   []string
	// additional imports required by the endpoint method
	OtherImports map[string]struct{}
	// post-processing code has to include return
	PostProcessing string
}

func newEndpointFromDoc(doc *wgMethodDoc, game string, id string) (*endpoint, error) {
	ep := endpoint{
		Id:   id,
		Game: snakeToCamelLower(game),
		Url:  doc.Url,

		Name:          snakeToCamel(id),
		MetaType:      nil,
		Documentation: strings.TrimPrefix(cleanDocumentation(doc.Description), "Method "),
		Deprecated:    doc.Deprecated,
		HttpMethods:   doc.AllowedHttpMethods,
		OtherImports:  map[string]struct{}{"context": {}},
		OptionsType:   nil,
	}
	if !strings.HasPrefix(ep.Url, "/") {
		ep.Url = "/" + ep.Url
	}

	if strings.HasPrefix(id, game+"_") {
		ep.Name = snakeToCamel(id[len(game)+1:])
	} else {
		ep.Name = snakeToCamel(id)
	}
	ep.AllowedRealms = doc.AvailableDisplayIndices
	sort.Strings(ep.AllowedRealms)
	ep.DataType = wgDocToGoDataType(strings.TrimPrefix(ep.Id, ep.Game+"_"), doc.Fields)
	// created required Parameters
	// create optional struct from not required once
	optionsStruct := &goType{
		Name:    ep.Name + "Options",
		TypeStr: "struct",
		Imports: make(map[string]struct{}),
	}
	for _, wgParam := range doc.Parameters {
		if wgParam.Name == "application_id" {
			continue
		}
		if wgParam.Required {
			param := wgToGoType(wgParam.Type)
			param.Name = snakeToCamelLower(wgParam.Name)
			param.Description = cleanDocumentation(wgParam.Description)
			param.JsonTag = wgParam.Name
			ep.Parameters = append(ep.Parameters, param)
		} else {
			field := wgToGoType(wgParam.Type)
			field.TypeStr = toPointerType(field.TypeStr)
			field.Name = snakeToCamel(wgParam.Name)
			field.Description = cleanDocumentation(wgParam.Description)
			field.JsonTag = wgParam.Name
			optionsStruct.Fields = append(optionsStruct.Fields, field)
		}
	}
	sortParameters(ep.Parameters)
	if len(optionsStruct.Fields) > 0 {
		sortFields(optionsStruct)
		ep.OptionsType = optionsStruct
	}

	return &ep, nil
}

func (ep *endpoint) MethodImports() []string {
	tmp := map[string]struct{}{}
	for _, param := range ep.Parameters {
		for imp := range param.Imports {
			tmp[imp] = struct{}{}
		}
		imps := valueConvImports(param.TypeStr)
		for _, imp := range imps {
			tmp[imp] = struct{}{}
		}
	}
	if ep.OptionsType != nil {
		for _, field := range ep.OptionsType.Fields {
			imps := valueConvImports(field.TypeStr)
			for _, imp := range imps {
				tmp[imp] = struct{}{}
			}
		}
	}
	for imp := range ep.OtherImports {
		tmp[imp] = struct{}{}
	}

	imports := make([]string, 0, len(tmp)+1)
	for imp := range tmp {
		imports = append(imports, imp)
	}
	if ep.OptionsType != nil || ep.DataType != nil && ep.DataType.IsStruct() {
		imports = append(imports, wgModuleName+"/wargaming/"+ep.Game)
	}
	sort.Strings(imports)
	return imports
}

func wgToGoType(name string) *goType {
	return map[string]*goType{
		"associative array": {TypeStr: "map[string]string", Imports: make(map[string]struct{})},
		"block_header":      {TypeStr: "struct", Imports: make(map[string]struct{})},
		"boolean":           {TypeStr: "bool", Imports: make(map[string]struct{})},
		"float":             {TypeStr: "float32", Imports: make(map[string]struct{})},
		"list of integers":  {TypeStr: "[]int", Imports: make(map[string]struct{})},
		"numeric":           {TypeStr: "int", Imports: make(map[string]struct{})},
		"string":            {TypeStr: "string", Imports: make(map[string]struct{})},
		"timestamp":         {TypeStr: "wgnTime.UnixTime", Imports: map[string]struct{}{wgModuleName + "/wargaming/wgnTime": {}}},

		"list of strings":    {TypeStr: "[]string", Imports: make(map[string]struct{})},
		"list of timestamps": {TypeStr: "[]wgnTime.UnixTime", Imports: map[string]struct{}{wgModuleName + "/wargaming/wgnTime": {}}},
		"list of dicts":      {TypeStr: "map[string]string", Imports: make(map[string]struct{})},
		"object":             {TypeStr: "struct", Imports: make(map[string]struct{})},
		"timestamp/date":     {TypeStr: "wgnTime.UnixTime", Imports: map[string]struct{}{wgModuleName + "/wargaming/wgnTime": {}}},

		// used in Parameters
		"numeric, list": {TypeStr: "[]int", Imports: make(map[string]struct{})},
		"string, list":  {TypeStr: "[]string", Imports: make(map[string]struct{})},
	}[name]
}

func wgDocToGoDataType(idWoPkg string, docFields []*wgFieldDoc) *goType {
	if len(docFields) == 0 {
		return nil
	}
	root := &goType{
		Name:    snakeToCamel(idWoPkg),
		TypeStr: "*struct",
		JsonTag: idWoPkg,
		Imports: map[string]struct{}{},
	}
	for _, dField := range docFields {
		typ := dField.Type
		if typ == "empty_line" {
			continue
		}
		// get field on root struct, create missing fields on the way down
		nestedField := root
		for _, curName := range dField.Name {
			curName = snakeToCamel(curName)
			curField := nestedField.F(curName)
			if curField == nil {
				curField = &goType{Name: curName}
				nestedField.Fields = append(nestedField.Fields, curField)
			}
			nestedField = curField
		}
		// add field values
		tmpType := wgToGoType(dField.Type)
		nestedField.TypeStr = toPointerType(tmpType.TypeStr)
		nestedField.Name = snakeToCamel(dField.Name[len(dField.Name)-1])
		nestedField.Description = cleanDocumentation(dField.Description)
		nestedField.JsonTag = dField.Name[len(dField.Name)-1]
		nestedField.Imports = tmpType.Imports
	}
	sortFields(root)
	return root
}
