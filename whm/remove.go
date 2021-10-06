package whm

import (
	"encoding/json"
	"github.com/lucidcube/go-cpanel/response"
	"net/url"
)

// RemoveUserResponse response from removing a user account
type RemoveUserResponse struct {
	response.ExtendedBaseWhmApiResponse
}

// RemoveAccount removes a users account
func (c *Connection) RemoveAccount(username string) (*RemoveUserResponse, error) {
	params := url.Values{}
	params.Add("user", username)
	body, err := c.WHMCall("removeacct", params)
	if err != nil {
		return nil, err
	}
	resp := &RemoveUserResponse{}
	err = json.Unmarshal(body, resp)
	return resp, err
}
