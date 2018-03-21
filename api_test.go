package cpanel

import (
	"fmt"
	"testing"
)

const (
	testUser  = "provtester"
	testToken = "PMIK472JO3JNYT6NCOA9W3V5C9UFNGBB"
	testHost  = "cpanel.lucidcube.com"
	testKey   = "test"
)

// TestAPICall is test function for iterative testing
func TestAPICall(t *testing.T) {
	conn, err := New(testToken, testUser, testHost)
	if err != nil {
		t.Fatal(err)
	}
	r, err := conn.GetLoginURL()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Login: %s\n", r)

	r2, err := conn.GetStats(StatCollection{FTPAccounts, EmailAccounts})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Stats: %v\n", r2)
}
