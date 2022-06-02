package wargaming

import "fmt"

type section interface {
	ApiUrl(realm Realm) string
}

var (
	// sectionWot World of Tanks.
	sectionWot sectionWotS
	// sectionWotb World of Tanks Blitz.
	sectionWotb sectionWotbS
	// sectionWotx World of Tanks Console.
	sectionWotx sectionWotxS
	// sectionWowp World of Warplanes.
	sectionWowp sectionWowpS
	// sectionWows World of Warships.
	sectionWows sectionWowsS
	// sectionWgn Wargaming.net.
	sectionWgn sectionWgnS
)

type sectionWotS struct{}

func (section sectionWotS) ApiUrl(realm Realm) string {
	return fmt.Sprintf("https://api.worldoftanks.%s/wot", realm.TLD())
}

type sectionWotbS struct{}

func (section sectionWotbS) ApiUrl(realm Realm) string {
	return fmt.Sprintf("https://api.wotblitz.%s/wotb", realm.TLD())
}

type sectionWotxS struct{}

func (section sectionWotxS) ApiUrl(realm Realm) string {
	return fmt.Sprintf("https://api-console.worldoftanks.%s/wotx", realm.TLD())
}

type sectionWowpS struct{}

func (section sectionWowpS) ApiUrl(realm Realm) string {
	return fmt.Sprintf("https://api.worldofwarplanes.%s/wowp", realm.TLD())
}

type sectionWowsS struct{}

func (section sectionWowsS) ApiUrl(realm Realm) string {
	return fmt.Sprintf("https://api.worldofwarships.%s/wows", realm.TLD())
}

type sectionWgnS struct{}

func (section sectionWgnS) ApiUrl(realm Realm) string {
	return fmt.Sprintf("https://api.worldoftanks.%s/wgn", realm.TLD())
}
