package wargaming



// WotAuthLogout Method deletes user's access_token.
// After this method is called, access_token becomes invalid.
//
// https://developers.wargaming.net/reference/all/wot/auth/logout
//
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
func (client *Client) WotAuthLogout(realm Realm, accessToken string) error {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil
	}

	reqParam := map[string]string{
		"access_token": accessToken,
	}

	err := client.doPostRequest(realm, "/wot/auth/logout/", reqParam)
	return err
}
