package cpanel

import (
	"net/http"
	"net/url"
	"strings"
	"io/ioutil"
	"encoding/json"
	"log"
	"crypto/tls"
	"time"
)

var (
	client *http.Client
	token  string
	User   = ""
	Host   = ""
)

func Init() {
	//Home »Development »Manage API Tokens
	token = "PMIK472JO3JNYT6NCOA9W3V5C9UFNGBB"

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Timeout: time.Second * 10, Transport: tr}
}

func WhmCall(call string, params url.Values) ([]byte, error) {

	uri := "https://" + Host + ":2087/json-api/" + call + "?api.version=1&" + params.Encode()
	req, err := http.NewRequest(http.MethodGet, uri, strings.NewReader(""))
	if err != nil {
		return []byte(""), err
	}
	req.Header.Add("Authorization", "whm root:"+token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return []byte(""), err
	}
	return ioutil.ReadAll(resp.Body)
}

func GetLoginUrl(user string) (string, error) {
	params := url.Values{}
	params.Add("user", user)
	params.Add("service", "cpaneld")
	body, err := WhmCall("create_user_session", params)
	log.Print(err, string(body))
	if err != nil {
		log.Print(err, string(body))
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

func MakeUAPICall(user, module, function string, args *url.Values) ([]byte, error) {
	if args == nil {
		args = &url.Values{}
	}
	args.Add("cpanel_jsonapi_user", user)
	args.Add("cpanel_jsonapi_module", module)
	args.Add("cpanel_jsonapi_func", function)
	args.Add("cpanel_jsonapi_apiversion", "3")
	return WhmCall("cpanel", *args)
}

type BaseWhmApiResponse struct {
	Metadata struct {
		Version int    `json:"version"`
		Command string `json:"command"`
		Reason  string `json:"reason"`
	} `json:"metadata"`
}

type CreateUserSessionResponse struct {
	BaseWhmApiResponse
	Data struct {
		Session       string `json:"session"`
		Service       string `json:"service"`
		URL           string `json:"url"`
		SecurityToken string `json:"cp_security_token"`
		Expires       int64  `json:"expires"`
	} `json:"data"`
}
