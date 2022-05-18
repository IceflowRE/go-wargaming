package wargaming

import "github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"

type Realm string

const (
	RealmAsia = Realm("asia")
	RealmEu   = Realm("eu")
	RealmNa   = Realm("na")
	RealmRu   = Realm("ru")
	RealmWgcb = Realm("wgcb")
)

// ApiUrl defaults to EU
func (realm Realm) ApiUrl() string {
	switch realm {
	case RealmAsia:
		return "https://api.worldoftanks.asia"
	case RealmEu:
		return "https://api.worldoftanks.eu"
	case RealmNa:
		return "https://api.worldoftanks.com"
	case RealmRu:
		return "https://api.worldoftanks.ru"
	case RealmWgcb:
		return "https://api-console.worldoftanks.com"
	}
	return "https://api.worldoftanks.eu"
}

func (realm Realm) String() string {
	return string(realm)
}

type InvalidRealm string

func (err InvalidRealm) Error() string {
	return string(err)
}

func validateRealm(realm Realm, allowedRealms []Realm) error {
	if !utils.Contains(allowedRealms, realm) {
		return InvalidRealm(realm)
	}
	return nil
}
