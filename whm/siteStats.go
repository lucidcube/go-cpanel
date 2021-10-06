package whm

import (
	"encoding/json"
	"net/url"

	"github.com/lucidcube/go-cpanel/response"
)

// AWStatDomainsResponse response from list for AWDomain listing
type AWStatDomainsResponse struct {
	response.BaseWhmAPIResponse
	CPanelResult struct {
		Data []AWDomain `json:"data"`
	} `json:"cpanelresult"`
}

// AWDomain is a single domain that has stats available
type AWDomain struct {
	SSL    response.Cbool `json:"ssl"`
	Lang   string         `json:"lang"`
	TXT    string         `json:"txt"`
	Domain string         `json:"domain"`
}

// GetStatSites retrieves listing of websites that have available stats
// site stats come from 'Awstats' module of cpanel
func (c *Connection) GetStatSites() ([]AWDomain, error) {
	params := url.Values{}
	params.Add("cpanel_jsonapi_user", c.user)
	params.Add("cpanel_jsonapi_version", "2")
	params.Add("cpanel_jsonapi_module", "Stats")
	params.Add("cpanel_jsonapi_func", "listawstats")
	body, err := c.WHMCall("cpanel", params)
	if err != nil {
		return nil, err
	}

	var resp AWStatDomainsResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return resp.CPanelResult.Data, nil
}
