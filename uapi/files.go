package uapi

import (
	"encoding/json"
	"net/url"

	"github.com/lucidcube/go-cpanel/response"
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
	File      string         `json:"file"`
	Exists    response.Cbool `json:"exists"`
	Ctime     int64          `json:"ctime"`
	Type      string         `json:"type"`
	Path      string         `json:"path"`
	HumanSize string         `json:"humansize"`
	GID       int64          `json:"gid"`
	Mode      int64          `json:"mode"`
	AbsDir    string         `json:"absdir"`
	FullPath  string         `json:"fullpath"`
	Mtime     int64          `json:"mtime"`
	NiceMode  string         `json:"nicemode"`
	UID       int64          `json:"uid"`
	Size      string         `json:"size"`
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
