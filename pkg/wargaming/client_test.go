package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
	"net/http"
	"os"
	"testing"
	"time"
)

var apiId = os.Getenv("WARGAMING_API_ID")
var client = NewClient(apiId).
	WithHttpClient(
		&http.Client{
			Timeout: 2 * time.Second,
		},
	)

const (
	reasonTooComplex          = "Unable to test method, too complex."
	reasonRequiresAccessToken = "Unable to test method, requires access token."
)

func skipTest(name string, reason string, test *testing.T) {
	test.Run(name, func(test *testing.T) {
		test.Skip(reason)
	})
}

func checkErr(test *testing.T, err error) {
	if err != nil {
		test.Fatal(err.Error())
	}
}

func TestApi(test *testing.T) {
	if apiId == "" {
		test.Fatal("Api ID is empty.")
	}
	test.Run("WotAccountList", func(test *testing.T) {
		_, err := client.WotAccountList(RealmEu, "Licht", []string{}, "", 3, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotAccountInfo", func(test *testing.T) {
		_, err := client.WotAccountInfo(RealmEu, []int{534339383}, "", []string{}, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotAccountTanks", func(test *testing.T) {
		_, err := client.WotAccountTanks(RealmEu, []int{534339383}, "", []string{}, "", []int{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotAccountAchievements", func(test *testing.T) {
		_, err := client.WotAccountAchievements(RealmEu, []int{534339383}, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	skipTest("WotAuthLogin", reasonTooComplex, test)
	skipTest("WotAuthProlongate", reasonRequiresAccessToken, test)
	skipTest("WotAuthLogout", reasonRequiresAccessToken, test)
	test.Run("WotStrongholdClaninfo", func(test *testing.T) {
		_, err := client.WotStrongholdClaninfo(RealmEu, []int{500150958}, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	skipTest("WotStrongholdClanreserves", reasonRequiresAccessToken, test)
	skipTest("WotStrongholdActivateclanreserve", reasonTooComplex, test)
	test.Run("WotGlobalmapFronts", func(test *testing.T) {
		_, err := client.WotGlobalmapFronts(RealmEu, []string{}, []string{}, "", 3, 0)
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotGlobalmapProvinces", func(test *testing.T) {
		_, err := client.WotGlobalmapProvinces(RealmEu, "", 0, 0, "season_18_bg", []string{}, "", "", 3, "", 0, 0, []string{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotGlobalmapClaninfo", func(test *testing.T) {
		_, err := client.WotGlobalmapClaninfo(RealmEu, []int{500150958}, "", []string{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotGlobalmapClanprovinces", func(test *testing.T) {
		_, err := client.WotGlobalmapClanprovinces(RealmEu, []int{500150958}, "", []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotGlobalmapClanbattles", func(test *testing.T) {
		_, err := client.WotGlobalmapClanbattles(RealmEu, 500150958, []string{}, "", 3, 0)
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotGlobalmapSeasons", func(test *testing.T) {
		_, err := client.WotGlobalmapSeasons(RealmEu, []string{}, "", 3, 0, "", "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotGlobalmapSeasonclaninfo", func(test *testing.T) {
		_, err := client.WotGlobalmapSeasonclaninfo(RealmEu, 500150958, "season_17", []string{"8"}, []string{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotGlobalmapSeasonaccountinfo", func(test *testing.T) {
		_, err := client.WotGlobalmapSeasonaccountinfo(RealmEu, 534339383, "season_17", []string{"8"}, []string{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotGlobalmapSeasonrating", func(test *testing.T) {
		_, err := client.WotGlobalmapSeasonrating(RealmEu, "season_17", "8", []string{}, 3, 0)
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotGlobalmapSeasonratingneighbors", func(test *testing.T) {
		_, err := client.WotGlobalmapSeasonratingneighbors(RealmEu, 500150958, 3, "season_17", "8", []string{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotGlobalmapEvents", func(test *testing.T) {
		_, err := client.WotGlobalmapEvents(RealmEu, "", []string{}, "", 3, 0, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotGlobalmapEventclaninfo", func(test *testing.T) {
		_, err := client.WotGlobalmapEventclaninfo(RealmEu, 500150958, "season_18", []string{"season_18_bg"}, []string{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotGlobalmapEventaccountinfo", func(test *testing.T) {
		_, err := client.WotGlobalmapEventaccountinfo(RealmEu, 534339383, 0, "season_18", []string{"season_18_bg"}, []string{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotGlobalmapEventaccountratings", func(test *testing.T) {
		_, err := client.WotGlobalmapEventaccountratings(RealmEu, "season_18", "season_18_bg", []string{}, 0, 3, 0)
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotGlobalmapEventaccountratingneighbors", func(test *testing.T) {
		_, err := client.WotGlobalmapEventaccountratingneighbors(RealmEu, 534339383, "season_18", "season_18_bg", []string{}, 3, 0, 0)
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotGlobalmapEventrating", func(test *testing.T) {
		_, err := client.WotGlobalmapEventrating(RealmEu, "season_18", "season_18_bg", []string{}, 3, 0)
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotGlobalmapEventratingneighbors", func(test *testing.T) {
		_, err := client.WotGlobalmapEventratingneighbors(RealmEu, 500150958, "season_18", "season_18_bg", []string{}, 3)
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotGlobalmapInfo", func(test *testing.T) {
		_, err := client.WotGlobalmapInfo(RealmEu, []string{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotEncyclopediaVehicles", func(test *testing.T) {
		_, err := client.WotEncyclopediaVehicles(RealmEu, []string{}, "", 3, []string{}, 0, []int{}, []int{}, []string{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotEncyclopediaVehicleprofile", func(test *testing.T) {
		_, err := client.WotEncyclopediaVehicleprofile(RealmEu, 0, 49, []string{}, 0, "", "", 0, 0, 0)
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotEncyclopediaAchievements", func(test *testing.T) {
		_, err := client.WotEncyclopediaAchievements(RealmEu, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotEncyclopediaInfo", func(test *testing.T) {
		_, err := client.WotEncyclopediaInfo(RealmEu, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotEncyclopediaArenas", func(test *testing.T) {
		_, err := client.WotEncyclopediaArenas(RealmEu, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotEncyclopediaProvisions", func(test *testing.T) {
		_, err := client.WotEncyclopediaProvisions(RealmEu, []string{}, "", 3, 0, []int{}, []string{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotEncyclopediaPersonalmissions", func(test *testing.T) {
		_, err := client.WotEncyclopediaPersonalmissions(RealmEu, []int{}, []string{}, "", []int{}, []int{}, []string{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotEncyclopediaBoosters", func(test *testing.T) {
		_, err := client.WotEncyclopediaBoosters(RealmEu, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotEncyclopediaVehicleprofiles", func(test *testing.T) {
		_, err := client.WotEncyclopediaVehicleprofiles(RealmEu, 49, []string{}, "", "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotEncyclopediaModules", func(test *testing.T) {
		_, err := client.WotEncyclopediaModules(RealmEu, []string{}, []string{}, "", 3, []int{}, []string{}, 0, []string{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotEncyclopediaBadges", func(test *testing.T) {
		_, err := client.WotEncyclopediaBadges(RealmEu, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotEncyclopediaCrewroles", func(test *testing.T) {
		_, err := client.WotEncyclopediaCrewroles(RealmEu, []string{}, "", []string{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotEncyclopediaCrewskills", func(test *testing.T) {
		_, err := client.WotEncyclopediaCrewskills(RealmEu, []string{}, "", "", []string{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotClanratingsTypes", func(test *testing.T) {
		_, err := client.WotClanratingsTypes(RealmEu)
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotClanratingsDates", func(test *testing.T) {
		_, err := client.WotClanratingsDates(RealmEu, 3)
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotClanratingsClans", func(test *testing.T) {
		_, err := client.WotClanratingsClans(RealmEu, []int{500150958}, wgnTime.UnixTime{}, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotClanratingsNeighbors", func(test *testing.T) {
		_, err := client.WotClanratingsNeighbors(RealmEu, 500150958, "global_rating_avg", wgnTime.UnixTime{}, []string{}, "", 3)
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotClanratingsTop", func(test *testing.T) {
		_, err := client.WotClanratingsTop(RealmEu, 3, 0, "global_rating_avg", wgnTime.UnixTime{}, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotTanksStats", func(test *testing.T) {
		_, err := client.WotTanksStats(RealmEu, 534339383, "", []string{}, []string{}, "", "", []int{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotTanksAchievements", func(test *testing.T) {
		_, err := client.WotTanksAchievements(RealmEu, 534339383, "", []string{}, "", "", []int{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotTanksMastery", func(test *testing.T) {
		_, err := client.WotTanksMastery(RealmEu, "xp", []int{3}, []string{}, "", []int{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotClansList", func(test *testing.T) {
		_, err := client.WotClansList(RealmEu, []string{}, "", 3, 0, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotClansInfo", func(test *testing.T) {
		_, err := client.WotClansInfo(RealmEu, []int{500150958}, "", []string{}, []string{}, "", "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotClansAccountinfo", func(test *testing.T) {
		_, err := client.WotClansAccountinfo(RealmEu, []int{534339383}, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotClansGlossary", func(test *testing.T) {
		_, err := client.WotClansGlossary(RealmEu, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	skipTest("WotClansMessageboard", reasonRequiresAccessToken, test)
	test.Run("WotClansMemberhistory", func(test *testing.T) {
		_, err := client.WotClansMemberhistory(RealmEu, 534339383, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotbAccountList", func(test *testing.T) {
		_, err := client.WotbAccountList(RealmEu, "Licht", []string{}, "", 3, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotbAccountInfo", func(test *testing.T) {
		_, err := client.WotbAccountInfo(RealmEu, []int{529712743}, "", []string{}, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotbAccountAchievements", func(test *testing.T) {
		_, err := client.WotbAccountAchievements(RealmEu, []int{529712743}, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotbAccountTankstats", func(test *testing.T) {
		_, err := client.WotbAccountTankstats(RealmEu, []int{529712743}, 49, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotbEncyclopediaVehicles", func(test *testing.T) {
		_, err := client.WotbEncyclopediaVehicles(RealmEu, []string{}, "", []string{}, []int{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotbEncyclopediaVehicleprofile", func(test *testing.T) {
		_, err := client.WotbEncyclopediaVehicleprofile(RealmEu, 0, 49, []string{}, 0, "", "", 0, 0)
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotbEncyclopediaModules", func(test *testing.T) {
		_, err := client.WotbEncyclopediaModules(RealmEu, []string{}, "", []int{}, "", "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotbEncyclopediaProvisions", func(test *testing.T) {
		_, err := client.WotbEncyclopediaProvisions(RealmEu, []string{}, "", []int{}, []int{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotbEncyclopediaInfo", func(test *testing.T) {
		_, err := client.WotbEncyclopediaInfo(RealmEu, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotbEncyclopediaAchievements", func(test *testing.T) {
		_, err := client.WotbEncyclopediaAchievements(RealmEu, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotbEncyclopediaCrewskills", func(test *testing.T) {
		_, err := client.WotbEncyclopediaCrewskills(RealmEu, []string{}, "", []string{}, []string{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotbEncyclopediaVehicleprofiles", func(test *testing.T) {
		_, err := client.WotbEncyclopediaVehicleprofiles(RealmEu, []int{49}, []string{}, "", "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotbClansList", func(test *testing.T) {
		_, err := client.WotbClansList(RealmEu, []string{}, "", 3, 0, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotbClansInfo", func(test *testing.T) {
		_, err := client.WotbClansInfo(RealmEu, []int{31045}, []string{}, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotbClansAccountinfo", func(test *testing.T) {
		_, err := client.WotbClansAccountinfo(RealmEu, []int{529712743}, []string{}, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotbClansGlossary", func(test *testing.T) {
		_, err := client.WotbClansGlossary(RealmEu, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotbTanksStats", func(test *testing.T) {
		_, err := client.WotbTanksStats(RealmEu, 529712743, "", []string{}, "", "", []int{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotbTanksAchievements", func(test *testing.T) {
		_, err := client.WotbTanksAchievements(RealmEu, 529712743, "", []string{}, "", "", []int{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	skipTest("WotbClanmessagesMessages", reasonRequiresAccessToken, test)
	skipTest("WotbClanmessagesCreate", reasonTooComplex, test)
	skipTest("WotbClanmessagesDelete", reasonTooComplex, test)
	skipTest("WotbClanmessagesLike", reasonTooComplex, test)
	skipTest("WotbClanmessagesLikes", reasonTooComplex, test)
	skipTest("WotbClanmessagesUpdate", reasonTooComplex, test)
	test.Run("WotbTournamentsList", func(test *testing.T) {
		_, err := client.WotbTournamentsList(RealmEu, []string{}, "", 3, 0, "", []string{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	skipTest("WotbTournamentsInfo", reasonTooComplex, test)
	skipTest("WotbTournamentsTeams", reasonTooComplex, test)
	skipTest("WotbTournamentsStages", reasonTooComplex, test)
	skipTest("WotbTournamentsMatches", reasonTooComplex, test)
	skipTest("WotbTournamentsStandings", reasonTooComplex, test)
	skipTest("WotbTournamentsTables", reasonTooComplex, test)
	test.Run("WotxAccountList", func(test *testing.T) {
		_, err := client.WotxAccountList(RealmWgcb, "Licht", []string{}, "", 3, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotxAccountInfo", func(test *testing.T) {
		_, err := client.WotxAccountInfo(RealmWgcb, []int{2738768}, "", []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotxAccountAchievements", func(test *testing.T) {
		_, err := client.WotxAccountAchievements(RealmWgcb, []int{2738768}, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	skipTest("WotxAccountXuidinfo", reasonTooComplex, test)
	skipTest("WotxAccountPsninfo", reasonTooComplex, test)
	skipTest("WotxAuthLogin", reasonTooComplex, test)
	skipTest("WotxAuthProlongate", reasonRequiresAccessToken, test)
	skipTest("WotxAuthLogout", reasonRequiresAccessToken, test)
	test.Run("WotxClansList", func(test *testing.T) {
		_, err := client.WotxClansList(RealmWgcb, []string{}, "", 3, 0, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotxClansInfo", func(test *testing.T) {
		_, err := client.WotxClansInfo(RealmWgcb, []int{5743}, []string{}, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotxClansAccountinfo", func(test *testing.T) {
		_, err := client.WotxClansAccountinfo(RealmWgcb, []int{2738768}, []string{}, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotxClansGlossary", func(test *testing.T) {
		_, err := client.WotxClansGlossary(RealmWgcb, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotxEncyclopediaCrewroles", func(test *testing.T) {
		_, err := client.WotxEncyclopediaCrewroles(RealmWgcb, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotxEncyclopediaVehicles", func(test *testing.T) {
		_, err := client.WotxEncyclopediaVehicles(RealmWgcb, []string{}, "", []string{}, []int{}, []int{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotxEncyclopediaVehicleupgrades", func(test *testing.T) {
		_, err := client.WotxEncyclopediaVehicleupgrades(RealmWgcb, []int{49}, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotxEncyclopediaAchievements", func(test *testing.T) {
		_, err := client.WotxEncyclopediaAchievements(RealmWgcb, []string{}, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotxEncyclopediaInfo", func(test *testing.T) {
		_, err := client.WotxEncyclopediaInfo(RealmWgcb, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotxEncyclopediaModules", func(test *testing.T) {
		_, err := client.WotxEncyclopediaModules(RealmWgcb, []string{}, []string{}, "", 3, []int{}, []string{}, 0, []string{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotxEncyclopediaArenas", func(test *testing.T) {
		_, err := client.WotxEncyclopediaArenas(RealmWgcb, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotxEncyclopediaVehicleprofile", func(test *testing.T) {
		_, err := client.WotxEncyclopediaVehicleprofile(RealmWgcb, 0, 49, []string{}, 0, "", "", 0, 0, 0)
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotxTanksStats", func(test *testing.T) {
		_, err := client.WotxTanksStats(RealmWgcb, 2738768, "", []string{}, "", "", []int{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WotxTanksAchievements", func(test *testing.T) {
		_, err := client.WotxTanksAchievements(RealmWgcb, 2738768, "", []string{}, "", "", []int{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowsAccountList", func(test *testing.T) {
		_, err := client.WowsAccountList(RealmEu, "Licht", []string{}, "", 3, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowsAccountInfo", func(test *testing.T) {
		_, err := client.WowsAccountInfo(RealmEu, []int{534339383}, "", []string{}, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowsAccountAchievements", func(test *testing.T) {
		_, err := client.WowsAccountAchievements(RealmEu, []int{534339383}, "", []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowsAccountStatsbydate", func(test *testing.T) {
		_, err := client.WowsAccountStatsbydate(RealmEu, 534339383, "", []string{}, []string{}, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowsEncyclopediaInfo", func(test *testing.T) {
		_, err := client.WowsEncyclopediaInfo(RealmEu, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowsEncyclopediaShips", func(test *testing.T) {
		_, err := client.WowsEncyclopediaShips(RealmEu, []string{}, "", 3, []string{}, 0, []int{}, []string{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowsEncyclopediaAchievements", func(test *testing.T) {
		_, err := client.WowsEncyclopediaAchievements(RealmEu, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowsEncyclopediaShipprofile", func(test *testing.T) {
		_, err := client.WowsEncyclopediaShipprofile(RealmEu, 0, 0, 0, 0, 3751786480, []string{}, 0, 0, 0, "", 0, 0)
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowsEncyclopediaModules", func(test *testing.T) {
		_, err := client.WowsEncyclopediaModules(RealmEu, []string{}, "", 3, []int{}, 0, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowsEncyclopediaAccountlevels", func(test *testing.T) {
		_, err := client.WowsEncyclopediaAccountlevels(RealmEu, []string{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowsEncyclopediaCrews", func(test *testing.T) {
		_, err := client.WowsEncyclopediaCrews(RealmEu, []int{}, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowsEncyclopediaCrewskills", func(test *testing.T) {
		_, err := client.WowsEncyclopediaCrewskills(RealmEu, []string{}, "", []int{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowsEncyclopediaCrewranks", func(test *testing.T) {
		_, err := client.WowsEncyclopediaCrewranks(RealmEu, []string{}, "", "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowsEncyclopediaBattletypes", func(test *testing.T) {
		_, err := client.WowsEncyclopediaBattletypes(RealmEu, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowsEncyclopediaConsumables", func(test *testing.T) {
		_, err := client.WowsEncyclopediaConsumables(RealmEu, []int{}, []string{}, "", 3, 0, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowsEncyclopediaCollections", func(test *testing.T) {
		_, err := client.WowsEncyclopediaCollections(RealmEu, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowsEncyclopediaCollectioncards", func(test *testing.T) {
		_, err := client.WowsEncyclopediaCollectioncards(RealmEu, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowsEncyclopediaBattlearenas", func(test *testing.T) {
		_, err := client.WowsEncyclopediaBattlearenas(RealmEu, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowsShipsStats", func(test *testing.T) {
		_, err := client.WowsShipsStats(RealmEu, 534339383, "", []string{}, []string{}, "", "", []int{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowsSeasonsInfo", func(test *testing.T) {
		_, err := client.WowsSeasonsInfo(RealmEu, []string{}, "", []int{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowsSeasonsShipstats", func(test *testing.T) {
		_, err := client.WowsSeasonsShipstats(RealmEu, 534339383, "", []string{}, "", []int{}, []int{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowsSeasonsAccountinfo", func(test *testing.T) {
		_, err := client.WowsSeasonsAccountinfo(RealmEu, []int{534339383}, "", []string{}, "", []int{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowsClansList", func(test *testing.T) {
		_, err := client.WowsClansList(RealmEu, []string{}, "", 3, 0, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowsClansInfo", func(test *testing.T) {
		_, err := client.WowsClansInfo(RealmEu, []int{500195107}, []string{}, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowsClansAccountinfo", func(test *testing.T) {
		_, err := client.WowsClansAccountinfo(RealmEu, []int{534339383}, []string{}, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowsClansGlossary", func(test *testing.T) {
		_, err := client.WowsClansGlossary(RealmEu, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowsClansSeason", func(test *testing.T) {
		_, err := client.WowsClansSeason(RealmEu, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowpAccountList", func(test *testing.T) {
		_, err := client.WowpAccountList(RealmEu, "Licht", []string{}, "", 3, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowpAccountInfo2", func(test *testing.T) {
		_, err := client.WowpAccountInfo2(RealmEu, []int{529033710}, "", []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowpAccountAchievements", func(test *testing.T) {
		_, err := client.WowpAccountAchievements(RealmEu, []int{529033710}, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowpEncyclopediaPlanes", func(test *testing.T) {
		_, err := client.WowpEncyclopediaPlanes(RealmEu, []string{}, "", []string{}, []string{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowpEncyclopediaPlaneinfo", func(test *testing.T) {
		_, err := client.WowpEncyclopediaPlaneinfo(RealmEu, []int{1203}, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowpEncyclopediaPlanemodules", func(test *testing.T) {
		_, err := client.WowpEncyclopediaPlanemodules(RealmEu, []int{1203}, []string{}, "", []string{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowpEncyclopediaPlaneupgrades", func(test *testing.T) {
		_, err := client.WowpEncyclopediaPlaneupgrades(RealmEu, []int{1203}, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowpEncyclopediaPlanespecification", func(test *testing.T) {
		_, err := client.WowpEncyclopediaPlanespecification(RealmEu, []int{}, []int{}, 1203, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowpEncyclopediaAchievements", func(test *testing.T) {
		_, err := client.WowpEncyclopediaAchievements(RealmEu, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowpEncyclopediaInfo", func(test *testing.T) {
		_, err := client.WowpEncyclopediaInfo(RealmEu)
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowpRatingsTypes", func(test *testing.T) {
		_, err := client.WowpRatingsTypes(RealmEu, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowpRatingsAccounts", func(test *testing.T) {
		_, err := client.WowpRatingsAccounts(RealmEu, []int{529033710}, wgnTime.UnixTime{}, "all", []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowpRatingsNeighbors", func(test *testing.T) {
		_, err := client.WowpRatingsNeighbors(RealmEu, 529033710, wgnTime.UnixTime{}, 3, "wins_ratio", "all", []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowpRatingsTop", func(test *testing.T) {
		_, err := client.WowpRatingsTop(RealmEu, wgnTime.UnixTime{}, 3, 0, "wins_ratio", "all", []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowpRatingsDates", func(test *testing.T) {
		_, err := client.WowpRatingsDates(RealmEu, []int{}, "all", []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowpPlanesStats", func(test *testing.T) {
		_, err := client.WowpPlanesStats(RealmEu, 529033710, "", []string{}, "", "", []int{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowpPlanesAchievements", func(test *testing.T) {
		_, err := client.WowpPlanesAchievements(RealmEu, 529033710, []string{}, "", []int{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowpClansList", func(test *testing.T) {
		_, err := client.WowpClansList(RealmEu, []string{}, "", 3, 0, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowpClansInfo", func(test *testing.T) {
		_, err := client.WowpClansInfo(RealmEu, []int{500094562}, []string{}, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowpClansGlossary", func(test *testing.T) {
		_, err := client.WowpClansGlossary(RealmEu, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WowpClansAccountinfo", func(test *testing.T) {
		_, err := client.WowpClansAccountinfo(RealmEu, []int{529033710}, []string{}, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WgnAccountList", func(test *testing.T) {
		_, err := client.WgnAccountList(RealmEu, "Licht", []string{}, []string{}, "", 3, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WgnAccountInfo", func(test *testing.T) {
		_, err := client.WgnAccountInfo(RealmEu, []int{534339383}, "", []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WgnWargagContent", func(test *testing.T) {
		_, err := client.WgnWargagContent(RealmRu, "", 0, 0, 0, wgnTime.UnixTime{}, []string{}, "", 0, 0, 0, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WgnWargagSearch", func(test *testing.T) {
		_, err := client.WgnWargagSearch(RealmRu, "wot", "", 0, []string{}, 0, 0, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WgnWargagComments", func(test *testing.T) {
		_, err := client.WgnWargagComments(RealmRu, 581891, []string{}, 0)
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WgnWargagCategories", func(test *testing.T) {
		_, err := client.WgnWargagCategories(RealmRu, 0, "picture", []string{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WgnWargagTags", func(test *testing.T) {
		_, err := client.WgnWargagTags(RealmRu, []string{}, 0)
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	skipTest("WgnWargagRate", reasonTooComplex, test)
	skipTest("WgnWargagNewcomment", reasonTooComplex, test)
	skipTest("WgnWargagDeletecomment", reasonTooComplex, test)
	test.Run("WgnWgtvTags", func(test *testing.T) {
		_, err := client.WgnWgtvTags(RealmEu, []string{}, "")
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WgnWgtvVideos", func(test *testing.T) {
		_, err := client.WgnWgtvVideos(RealmEu, []int{}, wgnTime.UnixTime{}, []string{}, 0, "", 3, 0, []int{}, []int{}, "", []int{}, []string{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WgnWgtvVehicles", func(test *testing.T) {
		_, err := client.WgnWgtvVehicles(RealmEu, []int{}, []int{}, []int{})
		checkErr(test, err)
	})
	time.Sleep(100 * time.Millisecond)
	test.Run("WgnServersInfo", func(test *testing.T) {
		_, err := client.WgnServersInfo(RealmEu, []string{}, []string{}, "")
		checkErr(test, err)
	})
}
