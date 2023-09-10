package internal

import (
	"cmp"
	"slices"
	"strings"
)

type goType struct {
	// Name may be empty in case of root structs
	Name        string
	TypeStr     string
	Description string
	// only filled if type contains struct like []struct or map[string]struct.
	Fields []*goType
	// will be ignored if root struct
	JsonTag string
	Imports map[string]struct{}
}

func (t *goType) AllImports() map[string]struct{} {
	imports := map[string]struct{}{}
	for imp := range t.Imports {
		imports[imp] = struct{}{}
	}
	for _, field := range t.Fields {
		for imp := range field.AllImports() {
			imports[imp] = struct{}{}
		}
	}
	return imports
}

func (t *goType) IsStruct() bool {
	return strings.Contains(t.TypeStr, "struct")
}

// ToPkgAccessName returns type string with "struct" replaced by pkgName.Name
func (t *goType) ToPkgAccessName(pkgName string) string {
	var name string
	if pkgName == "" {
		name = t.Name
	} else {
		name = pkgName + "." + t.Name
	}
	return strings.ReplaceAll(t.TypeStr, "struct", name)
}

// returns type as pointer as string
func toPointerType(tStr string) string {
	if strings.HasPrefix(tStr, "map[") || strings.HasPrefix(tStr, "[]") || strings.HasPrefix(tStr, "*") {
		return tStr
	}
	return "*" + tStr
}

func (t *goType) AddImport(newImport string) {
	if _, ok := t.Imports[newImport]; !ok {
		t.Imports[newImport] = struct{}{}
	}
}

func (t *goType) F(name string) *goType {
	for _, field := range t.Fields {
		if field.Name == name {
			return field
		}
	}
	return nil
}

func sortFields(typ *goType) {
	for _, field := range typ.Fields {
		sortFields(field)
	}
	slices.SortFunc(typ.Fields, func(lhs, rhs *goType) int {
		return cmp.Compare(lhs.Name, rhs.Name)
	})
}

func sortParameters(params []*goType) {
	slices.SortFunc(params, func(lhs *goType, rhs *goType) int {
		return cmp.Compare(lhs.Name, rhs.Name)
	})
}
