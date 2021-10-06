package whm

import (
	"encoding/json"
	"net/url"

	"github.com/lucidcube/go-cpanel/response"
)

// AccountEmailsResponse response from list user email accounts query
type AccountEmailsResponse struct {
	response.BaseWhmAPIResponse
	Data struct {
		Pops []string `json:"pops"`
	} `json:"data"`
}

// GetEmailAccountList retrieves cPanel account’s email accounts listing
func (c *Connection) GetEmailAccountList() ([]string, error) {
	params := url.Values{}
	params.Add("user", c.user)
	body, err := c.WHMCall("list_pops_for", params)
	if err != nil {
		return []string{}, err
	}

	resp := &AccountEmailsResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return []string{}, err
	}
	return resp.Data.Pops, nil
}
