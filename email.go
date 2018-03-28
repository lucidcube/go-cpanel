package cpanel

import (
	"encoding/json"
	"errors"
	"net/url"
)

// AccountEmailsResponse response from list user email accounts query
type AccountEmailsResponse struct {
	BaseWhmAPIResponse
	Data struct {
		Pops []string `json:"pops"`
	} `json:"data"`
}

// EmailCreateResponse response from creating a new email address
type EmailCreateResponse struct {
	Response struct {
		Errors []string `json:"errors"`
	} `json:"result"`
}

// CreateEmailAccount creates a new email address with a fully qualified address and strong password
// using a weak password will return a cpanel error
func (c *Connection) CreateEmailAccount(address, password string) (bool, error) {
	params := url.Values{}
	params.Add("user", c.user)
	params.Add("service", "cpaneld")
	params.Add("email", address)
	params.Add("password", password)

	body, err := c.MakeUAPICall("Email", "add_pop", params)
	if err != nil {
		return false, err
	}

	var response EmailCreateResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return false, err
	}

	if len(response.Response.Errors) > 0 {
		return false, errors.New(response.Response.Errors[0])
	}

	return false, err
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
