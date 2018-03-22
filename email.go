package cpanel

import (
	"encoding/json"
	"net/url"
)

// AccountEmailsResponse response from list user email accounts query
type AccountEmailsResponse struct {
	BaseWhmAPIResponse
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

	response := &AccountEmailsResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return []string{}, err
	}
	return response.Data.Pops, nil
}
