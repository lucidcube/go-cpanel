package uapi

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"

	"github.com/LucidCube/go-cpanel/whm"
)

// Connection is a single user connection to UAPI
type Connection struct {
	client    *http.Client
	parentWHM whm.Connection
	token     string
	user      string
	host      string
}

// MakeUAPICall creates calls to UAPI
func (c *Connection) MakeUAPICall(module, function string, args url.Values) ([]byte, error) {
	args.Add("cpanel_jsonapi_user", c.user)
	args.Add("cpanel_jsonapi_module", module)
	args.Add("cpanel_jsonapi_func", function)
	args.Add("cpanel_jsonapi_apiversion", "3")
	return c.parentWHM.WHMCall("cpanel", args)
}

// New returns a new UAPI connection instance
func New(token, user, host string, parent whm.Connection) (conn Connection, err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	conn = Connection{client: &http.Client{Timeout: time.Second * 10, Transport: tr},
		token:     token,
		user:      user,
		host:      host,
		parentWHM: parent,
	}
	return
}
