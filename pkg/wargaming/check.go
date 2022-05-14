package wargaming

import (
	"fmt"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

var validGames = []string{"wotb", "wot", "wows", "wowp"}
var validLanguages = []string{"cs", "de", "en", "es", "fr", "ko", "pl", "ru", "th", "tr", "vi", "zh-cn", "zh-tw"}

func CheckLanguage(lang string) *InvalidParameter {
	if !utils.Contains(validLanguages, lang) {
		return NewInvalidParameter("language", lang, fmt.Sprintf("must be one of the following '%v'", validLanguages))
	}
	return nil
}

func CheckResponseFields(responseFields string, validResponseFields []string) *InvalidParameter {
	if len(responseFields) > 100 {
		return NewInvalidParameter("fields", responseFields, "cannot be longer than 100 chars")
	}
	for _, value := range strings.Split(responseFields, ",") {
		if len(value) == 0 {
			return NewInvalidParameter("fields", responseFields, "contains an empty field")
		}
		normValue := value
		if value[0] == '-' {
			normValue = value[1:]
		}
		if !utils.Contains(validResponseFields, normValue) {
			return NewInvalidParameter("fields", responseFields, fmt.Sprintf("contains an invalid field '%s'", value))
		}
	}
	return nil
}

func CheckLimit(limit int) *InvalidParameter {
	if limit < 0 || limit > 100 {
		return NewInvalidParameter("limit", limit, "must between 0 and 100")
	}
	return nil
}

func CheckType(typ string) *InvalidParameter {
	if !utils.Contains([]string{"startswith", "exact"}, typ) {
		return NewInvalidParameter("type", typ, "must be empty or 'startswith' or 'exact'")
	}
	return nil
}
