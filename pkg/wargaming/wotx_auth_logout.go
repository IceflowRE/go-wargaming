package wargaming


// WotxAuthLogout Method deletes user's access_token.
// After this method is called, access_token becomes invalid.
//
// https://developers.wargaming.net/reference/all/wotx/auth/logout
//
// access_token:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
func (client *Client) WotxAuthLogout(realm Realm, accessToken string) error {
	if err := ValidateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil
	}

	reqParam := map[string]string{
		"access_token": accessToken,
	}

	err := client.doPostRequest(realm, "/wotx/auth/logout/", reqParam)
	return err
}
