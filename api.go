package cpanel

import (
	"errors"
	"fmt"

	"github.com/lucidcube/go-cpanel/uapi"
	"github.com/lucidcube/go-cpanel/whm"
)

// Connection is a single cpanel connection
type Connection struct {
	WHM  whm.Connection
	UAPI uapi.Connection
}

// New createa a new cpanel connection instance
func New(token, user, host string) (conn Connection, err error) {
	// Home »Development »Manage API Tokens
	if token == "" || user == "" || host == "" {
		err = errors.New("invalid connection params")
		return
	}

	conn = Connection{}
	conn.WHM, err = whm.New(token, user, host)
	if err != nil {
		return Connection{}, fmt.Errorf("Failed to init WHM connection %s", err.Error())
	}
	conn.UAPI, err = uapi.New(token, user, host, conn.WHM)
	if err != nil {
		return Connection{}, fmt.Errorf("Failed to init UAPI connection %s", err.Error())
	}
	return
}
