package cpanel

import (
	"encoding/json"
	"net/url"
)

// DomainListingResponse response from listing available domains
type DomainListingResponse struct {
	Response struct {
		Data struct {
			Main   string   `json:"main_domain"`
			Addon  []string `json:"addon_domains"`
			Parked []string `json:"parked_domains"`
			Sub    []string `json:"sub_domains"`
		} `json:"data"`
	} `json:"result"`
}

// GetDomainListing retrieves cpanel instance domain listing
func (c *Connection) GetDomainListing() ([]string, error) {
	body, err := c.MakeUAPICall("DomainInfo", "list_domains", url.Values{})
	if err != nil {
		return []string{}, err
	}

	var response DomainListingResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return []string{}, err
	}

	result := []string{}
	if response.Response.Data.Main != "" {
		result = append(result, response.Response.Data.Main)
	}

	result = append(result, response.Response.Data.Addon...)
	result = append(result, response.Response.Data.Parked...)
	result = append(result, response.Response.Data.Sub...)
	return result, err
}
