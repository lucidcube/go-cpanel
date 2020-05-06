package whm

import (
	"encoding/json"
	"github.com/LucidCube/go-cpanel/response"
	"net/url"
)

// SuspendUserResponse response from creation of a user session through WHMAPI
type SuspendUserResponse struct {
	response.ExtendedBaseWhmApiResponse
}

// Suspend suspends a users account
func (c *Connection) Suspend(reason string) (*SuspendUserResponse, error) {
	params := url.Values{}
	params.Add("user", c.user)
	params.Add("reason", reason)
	body, err := c.WHMCall("suspendacct", params)
	if err != nil {
		return nil, err
	}
	resp := &SuspendUserResponse{}
	err = json.Unmarshal(body, resp)
	return resp, err
}
