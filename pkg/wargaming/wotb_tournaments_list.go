package wargaming

import (
	"strconv"
	"strings"
)

type WotbTournamentsList struct {
	// Tournament description
	Description string `json:"description,omitempty"`
	// Tournament end date and time
	EndAt UnixTime `json:"end_at,omitempty"`
	// Matches start date and time
	MatchesStartAt UnixTime `json:"matches_start_at,omitempty"`
	// Registration end date and time
	RegistrationEndAt UnixTime `json:"registration_end_at,omitempty"`
	// Registration start date and time
	RegistrationStartAt UnixTime `json:"registration_start_at,omitempty"`
	// Tournament start date and time
	StartAt UnixTime `json:"start_at,omitempty"`
	// Tournament status
	Status string `json:"status,omitempty"`
	// Tournament name
	Title string `json:"title,omitempty"`
	// Tournament id
	TournamentId int `json:"tournament_id,omitempty"`
	// Award for participating in tournament
	Award struct {
		// Award amount
		Amount int `json:"amount,omitempty"`
		// Award currency: Free XP, gold or credits
		Currency string `json:"currency,omitempty"`
	} `json:"award,omitempty"`
	// Fee for participating in tournament
	Fee struct {
		// Fee amount
		Amount int `json:"amount,omitempty"`
		// Fee currency: Free XP, gold or credits
		Currency string `json:"currency,omitempty"`
	} `json:"fee,omitempty"`
	// Tournament Logo
	Logo struct {
		// Link to logo
		Original string `json:"original,omitempty"`
		// Link to preview
		Preview string `json:"preview,omitempty"`
	} `json:"logo,omitempty"`
	// Award for winning tournament
	WinnerAward struct {
		// Winner Award amount
		Amount int `json:"amount,omitempty"`
		// Winner Award currency: Free XP, gold or credits
		Currency string `json:"currency,omitempty"`
	} `json:"winner_award,omitempty"`
}

// WotbTournamentsList Method returns list of tournaments.
//
// https://developers.wargaming.net/reference/all/wotb/tournaments/list
//
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "ru" &mdash; Русский (by default)
// limit:
//     Page limit. Number of tournaments in one page. Default is 10. Min value is 1. Maximum value: 25.
// page_no:
//     Result page number. Default is 1. Min value is 1.
// search:
//     First letters in tournament name for search. Minimum length: 2 characters. Maximum length: 50 characters.
// status:
//     Tournament status. Maximum limit: 100. Valid values:
//     
//     "upcoming" &mdash; Tournament is planned 
//     "registration_started" &mdash; Players can apply to join the tournament 
//     "registration_finished" &mdash; Registration has finished 
//     "running" &mdash; The first match has started 
//     "finished" &mdash; The last match among all stages has been played 
//     "complete" &mdash; Tournament has been completed
func (client *Client) WotbTournamentsList(realm Realm, fields []string, language string, limit int, pageNo int, search string, status []string) ([]*WotbTournamentsList, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
		"limit": strconv.Itoa(limit),
		"page_no": strconv.Itoa(pageNo),
		"search": search,
		"status": strings.Join(status, ","),
	}

	var result []*WotbTournamentsList
	err := client.doGetDataRequest(realm, "/wotb/tournaments/list/", reqParam, &result)
	return result, err
}
