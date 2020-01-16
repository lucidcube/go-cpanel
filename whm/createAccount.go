package whm

import (
	"encoding/json"
	"net/url"
)

// CreateAccountResponse raw response from CreateAccount
type CreateAccountResponse struct {
	Raw      string
	MetaData struct {
		Command string `json:"command"`
		Output  struct {
			Raw string `json:"raw"`
		} `json:"output"`
		Version int    `json:"version"`
		Result  int    `json:"result"`
		Reason  string `json:"reason"`
	} `json:"metadata"`
	Data struct {
		IP               string `json:"ip"`
		Package          string `json:"package"`
		Nameserver       string `json:"nameserver"`
		NameserverA      string `json:"nameservera"`
		NameserverEntry  string `json:"nameserverentry"`
		Nameserver2      string `json:"nameserver2"`
		NameserverA2     string `json:"nameservera2"`
		NameserverEntry2 string `json:"nameserverentry2"`
		Nameserver3      string `json:"nameserver3"`
		NameserverA3     string `json:"nameservera3"`
		NameserverEntry3 string `json:"nameserverentry3"`
		Nameserver4      string `json:"nameserver4"`
		NameserverA4     string `json:"nameservera4"`
		NameserverEntry4 string `json:"nameserverentry4"`
	} `json:"data"`
}

// CreateAccountOptions are account creation parameters used when calling CreateAccount
type CreateAccountOptions struct {
	Plan     string
	Password string
}

// CreateAccount through WHM
func (c *Connection) CreateAccount(username, domain string, options CreateAccountOptions) (*CreateAccountResponse, error) {
	params := url.Values{}
	params.Add("username", username)
	params.Add("domain", domain)
	if options.Password != "" {
		params.Add("password", options.Password)
	}
	if options.Plan != "" {
		params.Add("plan", options.Plan)
	}

	raw, callErr := c.WHMCall("createacct", params)
	if callErr != nil {
		return nil, callErr
	}

	resp := &CreateAccountResponse{}
	resp.Raw = string(raw)
	decodeErr := json.Unmarshal(raw, resp)
	return resp, decodeErr
}
