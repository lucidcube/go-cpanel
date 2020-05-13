package whm

import (
	"encoding/json"
	"github.com/LucidCube/go-cpanel/response"
	"net/url"
)

// SuspendUserResponse response from suspending a user account
type SuspendUserResponse struct {
	response.ExtendedBaseWhmApiResponse
}

// Suspend suspends a users account
func (c *Connection) Suspend(username, reason string) (*SuspendUserResponse, error) {
	params := url.Values{}
	params.Add("user", username)
	params.Add("reason", reason)
	body, err := c.WHMCall("suspendacct", params)
	if err != nil {
		return nil, err
	}
	resp := &SuspendUserResponse{}
	err = json.Unmarshal(body, resp)
	return resp, err
}

// UnSuspendUserResponse response from un-suspending a users account
type UnSuspendUserResponse struct {
	response.ExtendedBaseWhmApiResponse
}

// UnSuspend un-suspends a users account
func (c *Connection) UnSuspend(username string) (*UnSuspendUserResponse, error) {
	params := url.Values{}
	params.Add("user", username)
	body, err := c.WHMCall("unsuspendacct", params)
	if err != nil {
		return nil, err
	}
	resp := &UnSuspendUserResponse{}
	err = json.Unmarshal(body, resp)
	return resp, err
}
