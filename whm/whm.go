package whm

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Connection is a single user connection to WHM API
type Connection struct {
	client *http.Client
	token  string
	user   string
	host   string
}

// WHMCall makes the call to the Web Host Manager
func (c *Connection) WHMCall(call string, params url.Values) ([]byte, error) {
	uri := "https://" + c.host + ":2087/json-api/" + call + "?api.version=1&" + params.Encode()
	req, err := http.NewRequest(http.MethodGet, uri, strings.NewReader(""))
	if err != nil {
		return []byte(""), err
	}

	req.Header.Add("Authorization", "WHM "+c.user+":"+c.token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return []byte(""), err
	}
	return ioutil.ReadAll(resp.Body)
}

// New returns a new WHM api instance
func New(token, user, host string) (conn Connection, err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	conn = Connection{client: &http.Client{Timeout: time.Second * 10, Transport: tr},
		token: token,
		user:  user,
		host:  host,
	}
	return
}
