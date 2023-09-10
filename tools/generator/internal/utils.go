package internal

import (
	"fmt"
	"html"
	"regexp"
	"runtime/debug"
	"strings"
	"unicode"
)

func getModuleName() (string, error) {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return "", fmt.Errorf("could not read build info")
	}

	return bi.Main.Path, nil
}

func snakeToCamel(name string) string {
	newName := strings.ReplaceAll(name, "_", " ")
	newName = strings.Title(newName) // lint:ignore title is enough as we do not have special cases
	return strings.ReplaceAll(newName, " ", "")
}

func snakeToCamelLower(name string) string {
	if len(name) == 0 {
		return ""
	}
	newName := snakeToCamel(name)
	return strings.ToLower(string(newName[0])) + newName[1:]
}

func camelLowerToCamel(name string) string {
	if len(name) == 0 {
		return ""
	}
	return string(unicode.ToUpper(rune(name[0]))) + name[1:]
}

var rxHtml = regexp.MustCompile(`<.*?>`)

func cleanDocumentation(text string) string {
	doc := html.UnescapeString(strings.TrimPrefix(rxHtml.ReplaceAllString(text, ""), "\n"))
	doc = strings.ReplaceAll(doc, "—", "-")
	doc = strings.ReplaceAll(doc, "“", "\"")
	doc = strings.ReplaceAll(doc, "”", "\"")
	return doc
}

func commentDocumentation(text string) string {
	doc := "// " + strings.ReplaceAll(text, "\n", "\n// ")
	return strings.TrimSuffix(doc, "\n// ")
}

// return convert code
func valueToStringConv(typ string, value string) string {
	deref := ""
	if strings.HasPrefix(typ, "*") {
		deref = "*"
	}
	switch typ {
	case "int", "*int":
		return "strconv.Itoa(" + deref + value + ")"
	case "[]int":
		return "internal.SliceIntToString(" + deref + value + ", \",\")"
	case "[]string":
		return "strings.Join(" + deref + value + ", \",\")"
	case "wgnTime.UnixTime", "*wgnTime.UnixTime":
		return "strconv.FormatInt(" + value + ".Unix(), 10)"
	}
	return deref + value
}

// returns imports required by valueToStringConv
func valueConvImports(typ string) []string {
	switch typ {
	case "int", "*int":
		return []string{"strconv"}
	case "[]int":
		return []string{wgModuleName + "/wargaming/internal"}
	case "[]string":
		return []string{"strings"}
	case "wgnTime.UnixTime", "*wgnTime.UnixTime":
		return []string{"strconv"}
	}
	return []string{}
}
