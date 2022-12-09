package internal

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/template"
)

func formatWriteFile(buf []byte, outputPath string) error {
	content, err := format.Source(buf)
	if err != nil {
		fmt.Println(string(buf))
		return err
	}
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()
	_, err = outFile.Write(content)
	if err != nil {
		return err
	}
	return nil
}

//go:embed templates/method.tmpl
var methodTempl string

func genMethod(ep *endpoint, outputPath string) error {
	templ, err := template.New("method").Funcs(template.FuncMap{
		"camelLowerToCamel": camelLowerToCamel,
		"requestMethod": func(availHttpMethods []string) string {
			if contains(availHttpMethods, "GET") {
				return "getRequest"
			}
			return "postRequest"
		},
		"metaReturnType": func(ep *endpoint) string {
			return toPointerType(ep.MetaType.ToPkgAccessName(ep.Game))
		},
		"dataReturnType": func(ep *endpoint) string {
			return toPointerType(ep.DataType.ToPkgAccessName(ep.Game))
		},
		"realmList": func(realmIdcs []string) string {
			realms := make([]string, 0, len(realmIdcs))
			for _, realmIdx := range realmIdcs {
				realms = append(realms, "Realm"+strings.Title(realmIdx))
			}
			return strings.Join(realms, ", ")
		},
		"replaceAll": strings.ReplaceAll,
		"summaryDoc": func(doc string) string {
			return commentDocumentation(doc)[3:]
		},
		"paramDoc": func(param *goType) string {
			res := param.Name + ":\n"
			res += "    " + strings.ReplaceAll(param.Description, "\n", "\n    ")
			res = strings.TrimSuffix(res, "    ")
			return commentDocumentation(res)
		},
		"optionsType": func(ep *endpoint) string {
			return toPointerType(ep.OptionsType.ToPkgAccessName(ep.Game))
		},
		"valueToStringConv": valueToStringConv,
		"valueToStringConvOptions": func(typ string, value string) string {
			return valueToStringConv(typ, "options."+value)
		},
	}).Parse(methodTempl)
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	err = templ.Execute(buf, ep)

	return formatWriteFile(buf.Bytes(), outputPath+ep.Id+".go")
}

//go:embed templates/structBase.tmpl
var structBaseTempl string

//go:embed templates/struct.tmpl
var structTempl string

func genStructs(ep *endpoint, outputPath string) error {
	templ, err := template.New("struct").Funcs(template.FuncMap{
		"imports": func(ep *endpoint) []string {
			tmp := map[string]struct{}{}
			if ep.OptionsType != nil {
				for imp := range ep.OptionsType.AllImports() {
					tmp[imp] = struct{}{}
				}
			}
			if ep.MetaType != nil {
				for imp := range ep.MetaType.AllImports() {
					tmp[imp] = struct{}{}
				}
			}
			if ep.DataType != nil {
				for imp := range ep.DataType.AllImports() {
					tmp[imp] = struct{}{}
				}
			}
			imports := make([]string, 0, len(tmp))
			for imp := range tmp {
				imports = append(imports, imp)
			}
			sort.Strings(imports)
			return imports
		},
		"commentIndentDoc": func(doc string, indent int) string {
			return strings.ReplaceAll(commentDocumentation(doc), "//", "\t//")
		},
		"indentation": func(depth int) string {
			return strings.Repeat("\t", depth)
		},
		"summaryDoc": func(doc string) string {
			return commentDocumentation(doc)[3:]
		},
		"inc": func(val int) int {
			return val + 1
		},
		// used to track the nested indentation
		"combine": func(t *goType, depth int) struct {
			T     *goType
			Depth int
		} {
			return struct {
				T     *goType
				Depth int
			}{t, depth}
		},
	}).Parse(structTempl)
	if err != nil {
		return err
	}
	templ, err = templ.New("structBase").Parse(structBaseTempl)
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	err = templ.ExecuteTemplate(buf, "structBase", ep)

	filename := strings.TrimPrefix(ep.Id, ep.Game+"_")
	return formatWriteFile(buf.Bytes(), outputPath+ep.Game+"/"+filename+".go")
}

//go:embed templates/services.tmpl
var servicesTempl string

type serviceData struct {
	Name     string
	LongName string
}

func genServices(services []serviceData, outputPath string) error {
	templ, err := template.New("services").Funcs(template.FuncMap{
		"snakeToCamel": snakeToCamel,
	}).Parse(servicesTempl)
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	err = templ.Execute(buf, services)

	return formatWriteFile(buf.Bytes(), outputPath+"services.go")
}

//go:embed templates/sections.tmpl
var sectionsTempl string

type sectionData struct {
	Name         string
	LongName     string
	ApiUrlFormat string
}

var rxApiUrl = regexp.MustCompile(`(.*\.)([a-z]+)(.*)`)

// sectionData key is name
func genSections(sections []*sectionData, outputPath string) error {
	templ, err := template.New("sections").Funcs(template.FuncMap{
		"formatApiUrl": func(url string) string {
			return rxApiUrl.ReplaceAllString(url, "$1%s$3")
		},
	}).Parse(sectionsTempl)
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	err = templ.Execute(buf, sections)

	return formatWriteFile(buf.Bytes(), outputPath+"sections.go")
}

//go:embed templates/realms.tmpl
var realmsTempl string

type realmData struct {
	LongName string
	Tld      string
	Index    string
}

// sectionData key is name
func genRealms(realms []*realmData, outputPath string) error {
	templ, err := template.New("realms").Funcs(template.FuncMap{
		"title": strings.Title,
	}).Parse(realmsTempl)
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	err = templ.Execute(buf, realms)

	return formatWriteFile(buf.Bytes(), outputPath+"realms.go")
}
