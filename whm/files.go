package whm

import (
	"encoding/json"
	"net/url"

	"github.com/LucidCube/go-cpanel/response"
)

// DiskUsageResponse response from disk usage query
type DiskUsageResponse struct {
	response.BaseWhmAPIResponse
	CPanelResult struct {
		Data []struct {
			MailArchives  int64  `json:"mailarchives"`
			SkipMailMan   string `json:"skipMailman"`
			MailAccounts  int64  `json:"mailaccounts"`
			MailMan       int64  `json:"mailman"`
			QuotaUsed     int64  `json:"quotaused"`
			HomeDirectory struct {
				Usage              int64          `json:"usage"`
				ContainedUsage     int64          `json:"contained_usage"`
				Owner              string         `json:"owner"`
				UserContainedUsage int64          `json:"user_contained_usage"`
				Name               string         `json:"name"`
				Traversible        response.Cbool `json:"traversible"`
				Type               string         `json:"dir"`
				Contents           []struct {
					Usage              int64          `json:"usage"`
					ContainedUsage     int64          `json:"contained_usage"`
					Owner              string         `json:"owner"`
					UserContainedUsage int64          `json:"user_contained_usage"`
					Contents           int64          `json:"contents"`
					Traversible        response.Cbool `json:"traversible"`
					Name               string         `json:"name"`
					Type               string         `json:"type"`
				}
			} `json:"homedir"`
		} `json:"data"`
	} `json:"cpanelresult"`
}

// GetDiskUsage retrieves the account's disk space usage data.
// The results include content outside of the home directory,
// such as databases, mailing lists, and mail archives.
func (c *Connection) GetDiskUsage() (DiskUsageResponse, error) {
	params := url.Values{}
	params.Add("cpanel_jsonapi_user", c.user)
	params.Add("cpanel_jsonapi_version", "2")
	params.Add("cpanel_jsonapi_module", "DiskUsage")
	params.Add("cpanel_jsonapi_func", "fetchdiskusagewithextras")
	body, err := c.WHMCall("cpanel", params)
	if err != nil {
		return DiskUsageResponse{}, err
	}

	var resp DiskUsageResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return DiskUsageResponse{}, err
	}

	return resp, nil
}
