package wargaming

import (
	"fmt"
)

type Realm interface {
	Name() string
	Index() string
	TLD() string
}

type realmS struct {
	name  string
	index string
	tld   string
}

func (realm realmS) Name() string {
	return realm.name
}
func (realm realmS) Index() string {
	return realm.index
}
func (realm realmS) TLD() string {
	return realm.tld
}

type InvalidRealm struct {
	Realm Realm
}

func (err InvalidRealm) Error() string {
	return fmt.Sprintf("invalid realm '%s' (%s)", err.Realm.Index(), err.Realm.Name())
}

func validateRealm(realm Realm, allowedRealms []Realm) error {
	for _, elem := range allowedRealms {
		if realm.Index() == elem.Index() {
			return nil
		}
	}
	return InvalidRealm{realm}
}
