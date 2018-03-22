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
func TestAPICalls(t *testing.T) {
	conn, err := New(testToken, testUser, testHost)
	if err != nil {
		t.Fatal(err)
	}

	// Login URL
	r, err := conn.GetLoginURL()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("\nLogin: %s\n", r)

	// Stats
	r2, err := conn.GetStats(StatCollection{FTPAccounts, EmailAccounts})
	if err != nil {
		t.Fatal(err)
	}

	gotFTP := false
	gotEmail := false

	for _, i := range r2 {
		switch i.StatType {
		case FTPAccounts:
			gotFTP = true
		case EmailAccounts:
			gotEmail = true
		}
	}

	if !gotFTP || !gotEmail {
		t.Fatal("Did not result in expected stats")
	}

	fmt.Printf("\nStats: %v\n", r2)

	// Email

	r3, err := conn.GetEmailAccountList()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("\nEmail accounts: %v\n", r3)
}
