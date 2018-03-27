package cpanel

import (
	"encoding/json"
	"net/url"
)

// DirectoryListingResponse is response from UAPI when querying file
// listing for a specific directory
type DirectoryListingResponse struct {
	Result struct {
		Items []DirectoryItem `json:"data"`
	} `json:"result"`
}

// DirectoryItem is a single item in a directory
type DirectoryItem struct {
	File      string `json:"file"`
	Exists    Cbool  `json:"exists"`
	Ctime     int64  `json:"ctime"`
	Type      string `json:"type"`
	Path      string `json:"path"`
	HumanSize string `json:"humansize"`
	GID       int64  `json:"gid"`
	Mode      int64  `json:"mode"`
	AbsDir    string `json:"absdir"`
	FullPath  string `json:"fullpath"`
	Mtime     int64  `json:"mtime"`
	NiceMode  string `json:"nicemode"`
	UID       int64  `json:"uid"`
	Size      string `json:"size"`
}

// GetDirectoryFileListing retrieves file listing for given directory
func (c *Connection) GetDirectoryFileListing(directory string) ([]DirectoryItem, error) {
	params := url.Values{}
	params.Add("user", c.user)
	params.Add("service", "cpaneld")
	params.Add("dirs", directory)

	body, err := c.MakeUAPICall("Fileman", "list_files", params)
	if err != nil {
		return []DirectoryItem{}, err
	}

	var response DirectoryListingResponse

	err = json.Unmarshal(body, &response)
	if err != nil {
		return []DirectoryItem{}, err
	}
	return response.Result.Items, nil
}

// DiskUsageResponse response from disk usage query
type DiskUsageResponse struct {
	BaseWhmAPIResponse
	CPanelResult struct {
		Data []struct {
			MailArchives  int64  `json:"mailarchives"`
			SkipMailMan   string `json:"skipMailman"`
			MailAccounts  int64  `json:"mailaccounts"`
			MailMan       int64  `json:"mailman"`
			QuotaUsed     int64  `json:"quotaused"`
			HomeDirectory struct {
				Usage              int64  `json:"usage"`
				ContainedUsage     int64  `json:"contained_usage"`
				Owner              string `json:"owner"`
				UserContainedUsage int64  `json:"user_contained_usage"`
				Name               string `json:"name"`
				Traversible        Cbool  `json:"traversible"`
				Type               string `json:"dir"`
				Contents           []struct {
					Usage              int64  `json:"usage"`
					ContainedUsage     int64  `json:"contained_usage"`
					Owner              string `json:"owner"`
					UserContainedUsage int64  `json:"user_contained_usage"`
					Contents           int64  `json:"contents"`
					Traversible        Cbool  `json:"traversible"`
					Name               string `json:"name"`
					Type               string `json:"type"`
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

	var response DiskUsageResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return DiskUsageResponse{}, err
	}

	return response, nil
}
