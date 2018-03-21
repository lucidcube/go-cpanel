package cpanel

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Connection is a single cpanel connection
type Connection struct {
	client *http.Client
	token  string
	user   string
	host   string
}

// New createa a new cpanel connection instance
func New(token, user, host string) (conn Connection, err error) {
	//Home »Development »Manage API Tokens
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	if token == "" || user == "" || host == "" {
		err = errors.New("invalid connection params")
		return
	}

	conn = Connection{
		token:  token,
		user:   user,
		host:   host,
		client: &http.Client{Timeout: time.Second * 10, Transport: tr},
	}
	return
}

// WHMCall makes the call to the Web Host Manager
func (c *Connection) WHMCall(call string, params url.Values) ([]byte, error) {
	uri := "https://" + c.host + ":2087/json-api/" + call + "?api.version=1&" + params.Encode()
	req, err := http.NewRequest(http.MethodGet, uri, strings.NewReader(""))
	if err != nil {
		return []byte(""), err
	}
	req.Header.Add("Authorization", "whm root:"+c.token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return []byte(""), err
	}
	return ioutil.ReadAll(resp.Body)
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
		log.Print(err)
		return "", err
	}
	return response.Data.URL, nil
}

func (c *Connection) GetStats() ([]StatResponse, error) {
	params := url.Values{}
	params.Add("user", c.user)
	params.Add("service", "cpaneld")

	params.Add("display", "bandwidthusage")
	body, err := c.MakeUAPICall("StatsBar", "get_stats", params)
	if err != nil {
		return []StatResponse{}, err
	}
	response := &StatsResponse{}

	fmt.Printf("BODY:%s\n", body)
	err = json.Unmarshal(body, response)
	if err != nil {
		log.Print(err)
		return []StatResponse{}, err
	}
	return response.Stats, nil
}

// MakeUAPICall creates calls to UAPI
func (c *Connection) MakeUAPICall(module, function string, args url.Values) ([]byte, error) {
	args.Add("cpanel_jsonapi_user", c.user)
	args.Add("cpanel_jsonapi_module", module)
	args.Add("cpanel_jsonapi_func", function)
	args.Add("cpanel_jsonapi_apiversion", "3")
	return c.WHMCall("cpanel", args)
}
