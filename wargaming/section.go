package wargaming

import "fmt"

type section interface {
	ApiUrl(realm Realm) string
}

type sectionS struct {
	apiUrlFormat string
}

func (s sectionS) ApiUrl(realm Realm) string {
	return fmt.Sprintf(s.apiUrlFormat, realm.TLD())
}
