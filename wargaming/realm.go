package wargaming

import (
	"fmt"
	"slices"
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

func containsRealm(allowedRealms []Realm, realm Realm) bool {
	return slices.ContainsFunc(allowedRealms, func(elem Realm) bool {
		return realm.Index() == elem.Index()
	})
}
