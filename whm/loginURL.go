package whm

import (
	"encoding/json"
	"net/url"

	"github.com/LucidCube/go-cpanel/response"
)

// CreateUserSessionResponse response from creation of a user session through WHMAPI
type CreateUserSessionResponse struct {
	response.BaseWhmAPIResponse
	Data struct {
		Session       string `json:"session"`
		Service       string `json:"service"`
		URL           string `json:"url"`
		SecurityToken string `json:"cp_security_token"`
		Expires       int64  `json:"expires"`
	} `json:"data"`
}

// GetLoginURL retrieves c-panel 'magic link' to log user directly into the control panel
func (c *Connection) GetLoginURL() (string, error) {
	params := url.Values{}
	params.Add("user", c.user)
	params.Add("service", "cpaneld")
	body, err := c.WHMCall("create_user_session", params)
	if err != nil {
		return "", err
	}
	response := &CreateUserSessionResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return "", err
	}
	return response.Data.URL, nil
}
