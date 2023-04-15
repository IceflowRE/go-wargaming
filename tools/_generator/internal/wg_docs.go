package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type wgMethodsDoc struct {
	Name     string
	Slug     string
	Sections []struct {
		Slug    string
		Methods []struct {
			Url        string
			Name       string
			Key        string
			Slug       string
			Deprecated bool
		}
	}
}

type wgParameter struct {
	Name        string
	Type        string
	Description string
	Required    bool
}

type wgFieldDoc struct {
	Name        []string
	Description string
	Type        string
}

type wgMethodDoc struct {
	AvailableDisplayIndices []string `json:"available_display_indices"`
	AllowedProtocols        []string `json:"allowed_protocols"`
	AllowedHttpMethods      []string `json:"allowed_http_methods"`
	Name                    string
	Description             string
	Deprecated              bool
	Realm                   string
	DisplayIndex            string `json:"display_index"`
	ApiUrl                  string `json:"api_url"`
	Url                     string
	Section                 string
	Errors                  []struct {
		Message     string
		Code        int
		Description string
	}
	Parameters []*wgParameter
	Fields     []*wgFieldDoc
}

func getWgDoc(path string, realm string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://developers.wargaming.net/api/%s/?realm=%s", path, realm), nil)
	if err != nil {
		return nil, nil
	}
	req.Header.Set("User-Agent", "go-wargaming-api-generator")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status was not OK")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//os.WriteFile("method.json", body, 0644)
	return body, nil
}

type realmDoc struct {
	Index string
	Name  string
}

func getWgMethodsDoc() ([]*wgMethodsDoc, []*realmDoc, error) {
	body, err := getWgDoc("methods", "all")
	if err != nil {
		return nil, nil, err
	}
	tmp := struct {
		Data   []*wgMethodsDoc
		Realms []*realmDoc
	}{}
	err = json.Unmarshal(body, &tmp)
	if err != nil {
		return nil, nil, err
	}
	return tmp.Data, tmp.Realms, nil
}

func getWgMethodDoc(id string, realm string) (*wgMethodDoc, error) {
	body, err := getWgDoc("methods/"+id, realm)
	if err != nil {
		return nil, err
	}
	tmp := struct {
		Data *wgMethodDoc
	}{}
	err = json.Unmarshal(body, &tmp)
	if err != nil {
		return nil, err
	}
	return tmp.Data, nil
}

var wgModuleName, _ = getModuleName()
