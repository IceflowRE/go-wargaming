// This file only checks if the API endpoint responds and unmarshalling. It does not check the content.

package wargaming

import (
	"context"
	"errors"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wgn"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wot"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wotx"
	"net/http"
	"os"
	"testing"
	"time"
)

var apiId = os.Getenv("WARGAMING_API_ID")
var client = NewClient(apiId, &ClientOptions{
	HTTPClient: &http.Client{
		Timeout: 5 * time.Second,
	},
})

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
	var respErr *ResponseError
	if errors.As(err, &respErr) {
		if respErr.Code == 504 {
			test.Skip("Method seems to be deprecated.")
		}
		if respErr.Code == 404 {
			test.Skip("Method seems to be temporarily unavailable.")
		}
	}
	if err != nil {
		test.Fatal(err.Error())
	}
}

func delay() {
	time.Sleep(100 * time.Millisecond)
}

const (
	wotAccountId  = 534339383
	wotClanId     = 500150958
	wotbAccountId = 529712743
	wotbClanId    = 31045
	wotxAccountId = 2738768
	wotxClanId    = 5743
	wowpAccountId = 582581733
	wowpClanId    = 500174897
	wowsAccountId = 543737361
	wowsClanId    = 500195107
)

func TestApi(test *testing.T) {
	if apiId == "" {
		test.Fatal("Api ID is empty.")
	}
	test.Run("WotAccountList", func(test *testing.T) {
		_, err := client.Wot.AccountList(context.Background(), RealmEu, "Licht", nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotAccountInfo", func(test *testing.T) {
		_, err := client.Wot.AccountInfo(context.Background(), RealmEu, []int{wotAccountId}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotAccountTanks", func(test *testing.T) {
		_, err := client.Wot.AccountTanks(context.Background(), RealmEu, []int{wotAccountId}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotAccountAchievements", func(test *testing.T) {
		_, err := client.Wot.AccountAchievements(context.Background(), RealmEu, []int{wotAccountId}, nil)
		checkErr(test, err)
	})
	delay()
	skipTest("WotAuthLogin", reasonTooComplex, test)
	skipTest("WotAuthProlongate", reasonRequiresAccessToken, test)
	skipTest("WotAuthLogout", reasonRequiresAccessToken, test)
	test.Run("WotStrongholdClaninfo", func(test *testing.T) {
		_, err := client.Wot.StrongholdClaninfo(context.Background(), RealmEu, []int{wotClanId}, nil)
		checkErr(test, err)
	})
	delay()
	skipTest("WotStrongholdClanreserves", reasonRequiresAccessToken, test)
	skipTest("WotStrongholdActivateclanreserve", reasonTooComplex, test)
	test.Run("WotGlobalmapFronts", func(test *testing.T) {
		_, err := client.Wot.GlobalmapFronts(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotGlobalmapProvinces", func(test *testing.T) {
		_, err := client.Wot.GlobalmapProvinces(context.Background(), RealmEu, "season_19_bg", nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotGlobalmapClaninfo", func(test *testing.T) {
		_, err := client.Wot.GlobalmapClaninfo(context.Background(), RealmEu, []int{wotClanId}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotGlobalmapClanprovinces", func(test *testing.T) {
		_, err := client.Wot.GlobalmapClanprovinces(context.Background(), RealmEu, []int{wotClanId}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotGlobalmapClanbattles", func(test *testing.T) {
		_, err := client.Wot.GlobalmapClanbattles(context.Background(), RealmEu, wotClanId, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotGlobalmapSeasons", func(test *testing.T) {
		_, err := client.Wot.GlobalmapSeasons(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotGlobalmapSeasonclaninfo", func(test *testing.T) {
		_, err := client.Wot.GlobalmapSeasonclaninfo(context.Background(), RealmEu, wotClanId, "season_17", []string{"8"}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotGlobalmapSeasonaccountinfo", func(test *testing.T) {
		_, err := client.Wot.GlobalmapSeasonaccountinfo(context.Background(), RealmEu, wotAccountId, "season_17", []string{"8"}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotGlobalmapSeasonrating", func(test *testing.T) {
		_, err := client.Wot.GlobalmapSeasonrating(context.Background(), RealmEu, "season_17", "8", nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotGlobalmapSeasonratingneighbors", func(test *testing.T) {
		_, err := client.Wot.GlobalmapSeasonratingneighbors(context.Background(), RealmEu, wotClanId, "season_17", "8", nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotGlobalmapEvents", func(test *testing.T) {
		_, err := client.Wot.GlobalmapEvents(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotGlobalmapEventclaninfo", func(test *testing.T) {
		_, err := client.Wot.GlobalmapEventclaninfo(context.Background(), RealmEu, wotAccountId, "season_19", []string{"season_19_bg"}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotGlobalmapEventaccountinfo_accountId", func(test *testing.T) {
		_, err := client.Wot.GlobalmapEventaccountinfo(context.Background(), RealmEu, "season_19", []string{"season_19_bg"},
			&wot.GlobalmapEventaccountinfoOptions{
				AccountId: Int(wotAccountId),
			},
		)
		checkErr(test, err)
	})
	delay()
	test.Run("WotGlobalmapEventaccountinfo_clan", func(test *testing.T) {
		_, err := client.Wot.GlobalmapEventaccountinfo(context.Background(), RealmEu, "season_19", []string{"season_19_bg"},
			&wot.GlobalmapEventaccountinfoOptions{
				ClanId: Int(wotClanId),
			},
		)
		checkErr(test, err)
	})
	delay()
	test.Run("WotGlobalmapEventaccountratings", func(test *testing.T) {
		_, err := client.Wot.GlobalmapEventaccountratings(context.Background(), RealmEu, "season_19", "season_19_bg", nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotGlobalmapEventaccountratingneighbors", func(test *testing.T) {
		_, err := client.Wot.GlobalmapEventaccountratingneighbors(context.Background(), RealmEu, wotAccountId, "season_19", "season_19_bg", nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotGlobalmapEventrating", func(test *testing.T) {
		_, err := client.Wot.GlobalmapEventrating(context.Background(), RealmEu, "season_19", "season_19_bg", nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotGlobalmapEventratingneighbors", func(test *testing.T) {
		_, err := client.Wot.GlobalmapEventratingneighbors(context.Background(), RealmEu, wotClanId, "season_19", "season_19_bg", nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotGlobalmapInfo", func(test *testing.T) {
		_, err := client.Wot.GlobalmapInfo(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotEncyclopediaVehicles", func(test *testing.T) {
		_, err := client.Wot.EncyclopediaVehicles(context.Background(), RealmEu, &wot.EncyclopediaVehiclesOptions{
			Limit:  Int(3),
			PageNo: Int(1),
		})
		checkErr(test, err)
	})
	delay()
	test.Run("WotEncyclopediaVehicleprofile", func(test *testing.T) {
		_, err := client.Wot.EncyclopediaVehicleprofile(context.Background(), RealmEu, 49, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotEncyclopediaAchievements", func(test *testing.T) {
		_, err := client.Wot.EncyclopediaAchievements(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotEncyclopediaInfo", func(test *testing.T) {
		_, err := client.Wot.EncyclopediaInfo(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotEncyclopediaArenas", func(test *testing.T) {
		_, err := client.Wot.EncyclopediaArenas(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotEncyclopediaProvisions", func(test *testing.T) {
		_, err := client.Wot.EncyclopediaProvisions(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotEncyclopediaPersonalmissions", func(test *testing.T) {
		_, err := client.Wot.EncyclopediaPersonalmissions(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotEncyclopediaBoosters", func(test *testing.T) {
		_, err := client.Wot.EncyclopediaBoosters(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotEncyclopediaVehicleprofiles", func(test *testing.T) {
		_, err := client.Wot.EncyclopediaVehicleprofiles(context.Background(), RealmEu, 49, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotEncyclopediaModules", func(test *testing.T) {
		_, err := client.Wot.EncyclopediaModules(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotEncyclopediaBadges", func(test *testing.T) {
		_, err := client.Wot.EncyclopediaBadges(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotEncyclopediaCrewroles", func(test *testing.T) {
		_, err := client.Wot.EncyclopediaCrewroles(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotEncyclopediaCrewskills", func(test *testing.T) {
		_, err := client.Wot.EncyclopediaCrewskills(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotClanratingsTypes", func(test *testing.T) {
		_, err := client.Wot.ClanratingsTypes(context.Background(), RealmEu)
		checkErr(test, err)
	})
	delay()
	test.Run("WotClanratingsDates", func(test *testing.T) {
		_, err := client.Wot.ClanratingsDates(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotClanratingsClans", func(test *testing.T) {
		_, err := client.Wot.ClanratingsClans(context.Background(), RealmEu, []int{wotClanId}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotClanratingsNeighbors", func(test *testing.T) {
		_, err := client.Wot.ClanratingsNeighbors(context.Background(), RealmEu, wotClanId, "global_rating_avg", nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotClanratingsTop", func(test *testing.T) {
		_, err := client.Wot.ClanratingsTop(context.Background(), RealmEu, "global_rating_avg", nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotTanksStats", func(test *testing.T) {
		_, err := client.Wot.TanksStats(context.Background(), RealmEu, wotAccountId, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotTanksAchievements", func(test *testing.T) {
		_, err := client.Wot.TanksAchievements(context.Background(), RealmEu, wotAccountId, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotTanksMastery", func(test *testing.T) {
		_, err := client.Wot.TanksMastery(context.Background(), RealmEu, "xp", []int{3}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotClansList", func(test *testing.T) {
		_, err := client.Wot.ClansList(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotClansInfo", func(test *testing.T) {
		_, err := client.Wot.ClansInfo(context.Background(), RealmEu, []int{wotClanId}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotClansAccountinfo", func(test *testing.T) {
		_, err := client.Wot.ClansAccountinfo(context.Background(), RealmEu, []int{wotAccountId}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotClansGlossary", func(test *testing.T) {
		_, err := client.Wot.ClansGlossary(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	skipTest("WotClansMessageboard", reasonRequiresAccessToken, test)
	test.Run("WotClansMemberhistory", func(test *testing.T) {
		_, err := client.Wot.ClansMemberhistory(context.Background(), RealmEu, wotAccountId, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotbAccountList", func(test *testing.T) {
		_, err := client.Wotb.AccountList(context.Background(), RealmEu, "Licht", nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotbAccountInfo", func(test *testing.T) {
		_, err := client.Wotb.AccountInfo(context.Background(), RealmEu, []int{wotbAccountId}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotbAccountAchievements", func(test *testing.T) {
		_, err := client.Wotb.AccountAchievements(context.Background(), RealmEu, []int{wotbAccountId}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotbAccountTankstats", func(test *testing.T) {
		_, err := client.Wotb.AccountTankstats(context.Background(), RealmEu, []int{wotbAccountId}, 49, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotbEncyclopediaVehicles", func(test *testing.T) {
		_, err := client.Wotb.EncyclopediaVehicles(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotbEncyclopediaVehicleprofile", func(test *testing.T) {
		_, err := client.Wotb.EncyclopediaVehicleprofile(context.Background(), RealmEu, 49, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotbEncyclopediaModules", func(test *testing.T) {
		_, err := client.Wotb.EncyclopediaModules(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotbEncyclopediaProvisions", func(test *testing.T) {
		_, err := client.Wotb.EncyclopediaProvisions(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotbEncyclopediaInfo", func(test *testing.T) {
		_, err := client.Wotb.EncyclopediaInfo(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotbEncyclopediaAchievements", func(test *testing.T) {
		_, err := client.Wotb.EncyclopediaAchievements(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotbEncyclopediaCrewskills", func(test *testing.T) {
		test.Skip("Method seems to be deprecated.")
		_, err := client.Wotb.EncyclopediaCrewskills(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotbEncyclopediaVehicleprofiles", func(test *testing.T) {
		_, err := client.Wotb.EncyclopediaVehicleprofiles(context.Background(), RealmEu, []int{49}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotbClansList", func(test *testing.T) {
		_, err := client.Wotb.ClansList(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotbClansInfo", func(test *testing.T) {
		_, err := client.Wotb.ClansInfo(context.Background(), RealmEu, []int{wotbClanId}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotbClansAccountinfo", func(test *testing.T) {
		_, err := client.Wotb.ClansAccountinfo(context.Background(), RealmEu, []int{wotbAccountId}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotbClansGlossary", func(test *testing.T) {
		_, err := client.Wotb.ClansGlossary(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotbTanksStats", func(test *testing.T) {
		_, err := client.Wotb.TanksStats(context.Background(), RealmEu, wotbAccountId, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotbTanksAchievements", func(test *testing.T) {
		_, err := client.Wotb.TanksAchievements(context.Background(), RealmEu, wotbAccountId, nil)
		checkErr(test, err)
	})
	delay()
	skipTest("WotbClanmessagesMessages", reasonRequiresAccessToken, test)
	skipTest("WotbClanmessagesCreate", reasonTooComplex, test)
	skipTest("WotbClanmessagesDelete", reasonTooComplex, test)
	skipTest("WotbClanmessagesLike", reasonTooComplex, test)
	skipTest("WotbClanmessagesLikes", reasonTooComplex, test)
	skipTest("WotbClanmessagesUpdate", reasonTooComplex, test)
	var tournamentId []int
	test.Run("WotbTournamentsList", func(test *testing.T) {
		res, err := client.Wotb.TournamentsList(context.Background(), RealmEu, nil)
		checkErr(test, err)
		if err == nil && len(res) > 0 {
			tournamentId = append(tournamentId, *res[0].TournamentId)
		}
	})
	if len(tournamentId) > 0 {
		delay()
		test.Run("WotbTournamentsInfo", func(test *testing.T) {
			_, err := client.Wotb.TournamentsInfo(context.Background(), RealmEu, tournamentId, nil)
			checkErr(test, err)
		})
		delay()
		test.Run("WotbTournamentsTeams", func(test *testing.T) {
			_, err := client.Wotb.TournamentsTeams(context.Background(), RealmEu, tournamentId[0], nil)
			checkErr(test, err)
		})
		delay()
		test.Run("WotbTournamentsStages", func(test *testing.T) {
			_, err := client.Wotb.TournamentsStages(context.Background(), RealmEu, tournamentId[0], nil)
			checkErr(test, err)
		})
	} else {
		skipTest("WotbTournamentsInfo", reasonTooComplex, test)
		skipTest("WotbTournamentsTeams", reasonTooComplex, test)
		skipTest("WotbTournamentsStages", reasonTooComplex, test)
	}
	skipTest("WotbTournamentsMatches", reasonTooComplex, test)
	skipTest("WotbTournamentsStandings", reasonTooComplex, test)
	skipTest("WotbTournamentsTables", reasonTooComplex, test)
	delay()
	test.Run("WotxAccountList", func(test *testing.T) {
		_, err := client.Wotx.AccountList(context.Background(), RealmWgcb, "Licht", nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotxAccountInfo", func(test *testing.T) {
		_, err := client.Wotx.AccountInfo(context.Background(), RealmWgcb, []int{wotxAccountId}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotxAccountAchievements", func(test *testing.T) {
		_, err := client.Wotx.AccountAchievements(context.Background(), RealmWgcb, []int{wotxAccountId}, nil)
		checkErr(test, err)
	})
	delay()
	skipTest("WotxAccountXuidinfo", reasonTooComplex, test)
	skipTest("WotxAccountPsninfo", reasonTooComplex, test)
	skipTest("WotxAuthLogin", reasonTooComplex, test)
	skipTest("WotxAuthProlongate", reasonRequiresAccessToken, test)
	skipTest("WotxAuthLogout", reasonRequiresAccessToken, test)
	test.Run("WotxClansList", func(test *testing.T) {
		_, err := client.Wotx.ClansList(context.Background(), RealmWgcb, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotxClansInfo", func(test *testing.T) {
		_, err := client.Wotx.ClansInfo(context.Background(), RealmWgcb, []int{wotxClanId}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotxClansAccountinfo", func(test *testing.T) {
		_, err := client.Wotx.ClansAccountinfo(context.Background(), RealmWgcb, []int{wotxAccountId}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotxClansGlossary", func(test *testing.T) {
		_, err := client.Wotx.ClansGlossary(context.Background(), RealmWgcb, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotxEncyclopediaCrewroles", func(test *testing.T) {
		_, err := client.Wotx.EncyclopediaCrewroles(context.Background(), RealmWgcb, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotxEncyclopediaVehicles", func(test *testing.T) {
		_, err := client.Wotx.EncyclopediaVehicles(context.Background(), RealmWgcb, &wotx.EncyclopediaVehiclesOptions{
			TankId: []int{49},
		})
		checkErr(test, err)
	})
	delay()
	test.Run("WotxEncyclopediaVehicleupgrades", func(test *testing.T) {
		_, err := client.Wotx.EncyclopediaVehicleupgrades(context.Background(), RealmWgcb, []int{49}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotxEncyclopediaAchievements", func(test *testing.T) {
		_, err := client.Wotx.EncyclopediaAchievements(context.Background(), RealmWgcb, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotxEncyclopediaInfo", func(test *testing.T) {
		_, err := client.Wotx.EncyclopediaInfo(context.Background(), RealmWgcb, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotxEncyclopediaModules", func(test *testing.T) {
		_, err := client.Wotx.EncyclopediaModules(context.Background(), RealmWgcb, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotxEncyclopediaArenas", func(test *testing.T) {
		_, err := client.Wotx.EncyclopediaArenas(context.Background(), RealmWgcb, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotxEncyclopediaVehicleprofile", func(test *testing.T) {
		_, err := client.Wotx.EncyclopediaVehicleprofile(context.Background(), RealmWgcb, 49, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotxTanksStats", func(test *testing.T) {
		_, err := client.Wotx.TanksStats(context.Background(), RealmWgcb, wotxAccountId, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WotxTanksAchievements", func(test *testing.T) {
		_, err := client.Wotx.TanksAchievements(context.Background(), RealmWgcb, wotxAccountId, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowpAccountList", func(test *testing.T) {
		_, err := client.Wowp.AccountList(context.Background(), RealmEu, "Licht", nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowpAccountInfo2", func(test *testing.T) {
		_, err := client.Wowp.AccountInfo2(context.Background(), RealmEu, []int{wowpAccountId}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowpAccountAchievements", func(test *testing.T) {
		_, err := client.Wowp.AccountAchievements(context.Background(), RealmEu, []int{wowpAccountId}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowpEncyclopediaPlanes", func(test *testing.T) {
		_, err := client.Wowp.EncyclopediaPlanes(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowpEncyclopediaPlaneinfo", func(test *testing.T) {
		_, err := client.Wowp.EncyclopediaPlaneinfo(context.Background(), RealmEu, []int{wowpClanId}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowpEncyclopediaPlanemodules", func(test *testing.T) {
		_, err := client.Wowp.EncyclopediaPlanemodules(context.Background(), RealmEu, []int{1203}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowpEncyclopediaPlaneupgrades", func(test *testing.T) {
		_, err := client.Wowp.EncyclopediaPlaneupgrades(context.Background(), RealmEu, []int{1203}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowpEncyclopediaPlanespecification", func(test *testing.T) {
		_, err := client.Wowp.EncyclopediaPlanespecification(context.Background(), RealmEu, 1203, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowpEncyclopediaAchievements", func(test *testing.T) {
		_, err := client.Wowp.EncyclopediaAchievements(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowpEncyclopediaInfo", func(test *testing.T) {
		_, err := client.Wowp.EncyclopediaInfo(context.Background(), RealmEu)
		checkErr(test, err)
	})
	delay()
	test.Run("WowpRatingsTypes", func(test *testing.T) {
		_, err := client.Wowp.RatingsTypes(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowpRatingsAccounts", func(test *testing.T) {
		_, err := client.Wowp.RatingsAccounts(context.Background(), RealmEu, []int{wowpAccountId}, "all", nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowpRatingsNeighbors", func(test *testing.T) {
		_, err := client.Wowp.RatingsNeighbors(context.Background(), RealmEu, wowpAccountId, "wins_ratio", "all", nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowpRatingsTop", func(test *testing.T) {
		_, err := client.Wowp.RatingsTop(context.Background(), RealmEu, "wins_ratio", "all", nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowpRatingsDates", func(test *testing.T) {
		_, err := client.Wowp.RatingsDates(context.Background(), RealmEu, "all", nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowpPlanesStats", func(test *testing.T) {
		_, err := client.Wowp.PlanesStats(context.Background(), RealmEu, wowpAccountId, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowpPlanesAchievements", func(test *testing.T) {
		_, err := client.Wowp.PlanesAchievements(context.Background(), RealmEu, wowpAccountId, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowpClansList", func(test *testing.T) {
		_, err := client.Wowp.ClansList(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowpClansInfo", func(test *testing.T) {
		_, err := client.Wowp.ClansInfo(context.Background(), RealmEu, []int{wowpClanId}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowpClansGlossary", func(test *testing.T) {
		_, err := client.Wowp.ClansGlossary(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowpClansAccountinfo", func(test *testing.T) {
		_, err := client.Wowp.ClansAccountinfo(context.Background(), RealmEu, []int{wowpAccountId}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowsAccountList", func(test *testing.T) {
		_, err := client.Wows.AccountList(context.Background(), RealmEu, "Licht", nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowsAccountInfo", func(test *testing.T) {
		_, err := client.Wows.AccountInfo(context.Background(), RealmEu, []int{wowsAccountId}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowsAccountAchievements", func(test *testing.T) {
		_, _, err := client.Wows.AccountAchievements(context.Background(), RealmEu, []int{wowsAccountId}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowsAccountStatsbydate", func(test *testing.T) {
		_, err := client.Wows.AccountStatsbydate(context.Background(), RealmEu, wowsAccountId, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowsEncyclopediaInfo", func(test *testing.T) {
		_, err := client.Wows.EncyclopediaInfo(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowsEncyclopediaShips", func(test *testing.T) {
		_, _, err := client.Wows.EncyclopediaShips(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowsEncyclopediaAchievements", func(test *testing.T) {
		_, err := client.Wows.EncyclopediaAchievements(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowsEncyclopediaShipprofile", func(test *testing.T) {
		_, err := client.Wows.EncyclopediaShipprofile(context.Background(), RealmEu, 3751786480, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowsEncyclopediaModules", func(test *testing.T) {
		_, err := client.Wows.EncyclopediaModules(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowsEncyclopediaAccountlevels", func(test *testing.T) {
		_, err := client.Wows.EncyclopediaAccountlevels(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowsEncyclopediaCrews", func(test *testing.T) {
		_, err := client.Wows.EncyclopediaCrews(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowsEncyclopediaCrewskills", func(test *testing.T) {
		_, err := client.Wows.EncyclopediaCrewskills(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowsEncyclopediaCrewranks", func(test *testing.T) {
		_, err := client.Wows.EncyclopediaCrewranks(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowsEncyclopediaBattletypes", func(test *testing.T) {
		_, err := client.Wows.EncyclopediaBattletypes(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowsEncyclopediaConsumables", func(test *testing.T) {
		_, err := client.Wows.EncyclopediaConsumables(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowsEncyclopediaCollections", func(test *testing.T) {
		_, err := client.Wows.EncyclopediaCollections(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowsEncyclopediaCollectioncards", func(test *testing.T) {
		_, err := client.Wows.EncyclopediaCollectioncards(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowsEncyclopediaBattlearenas", func(test *testing.T) {
		_, err := client.Wows.EncyclopediaBattlearenas(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowsShipsStats", func(test *testing.T) {
		_, _, err := client.Wows.ShipsStats(context.Background(), RealmEu, wowsAccountId, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowsSeasonsInfo", func(test *testing.T) {
		_, err := client.Wows.SeasonsInfo(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowsSeasonsShipstats", func(test *testing.T) {
		_, _, err := client.Wows.SeasonsShipstats(context.Background(), RealmEu, wowsAccountId, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowsSeasonsAccountinfo", func(test *testing.T) {
		_, _, err := client.Wows.SeasonsAccountinfo(context.Background(), RealmEu, []int{wowsAccountId}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowsClansList", func(test *testing.T) {
		_, err := client.Wows.ClansList(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowsClansInfo", func(test *testing.T) {
		_, err := client.Wows.ClansInfo(context.Background(), RealmEu, []int{wowsClanId}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowsClansAccountinfo", func(test *testing.T) {
		_, err := client.Wows.ClansAccountinfo(context.Background(), RealmEu, []int{wowsAccountId}, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowsClansGlossary", func(test *testing.T) {
		_, err := client.Wows.ClansGlossary(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WowsClansSeason", func(test *testing.T) {
		_, err := client.Wows.ClansSeason(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WgnAccountList", func(test *testing.T) {
		_, err := client.Wgn.AccountList(context.Background(), RealmEu, "Licht", nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WgnAccountInfo", func(test *testing.T) {
		_, err := client.Wgn.AccountInfo(context.Background(), RealmEu, []int{wotAccountId}, nil)
		checkErr(test, err)
	})
	delay()
	skipTest("WgnWargagRate", reasonTooComplex, test)
	skipTest("WgnWargagNewcomment", reasonTooComplex, test)
	skipTest("WgnWargagDeletecomment", reasonTooComplex, test)
	test.Run("WgnWgtvTags", func(test *testing.T) {
		_, err := client.Wgn.WgtvTags(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WgnWgtvVideos", func(test *testing.T) {
		_, err := client.Wgn.WgtvVideos(context.Background(), RealmEu, &wgn.WgtvVideosOptions{
			Limit:  Int(3),
			PageNo: Int(1),
		})
		checkErr(test, err)
	})
	delay()
	test.Run("WgnWgtvVehicles", func(test *testing.T) {
		_, err := client.Wgn.WgtvVehicles(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
	delay()
	test.Run("WgnServersInfo", func(test *testing.T) {
		_, err := client.Wgn.ServersInfo(context.Background(), RealmEu, nil)
		checkErr(test, err)
	})
}
