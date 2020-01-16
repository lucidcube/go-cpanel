package cpanel

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/LucidCube/go-cpanel/uapi"
	"github.com/LucidCube/go-cpanel/whm"
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
	r, err := conn.WHM.GetLoginURL()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("\nLogin: %s\n", r)

	// Account Stats
	r2, err := conn.UAPI.GetAccountStats(uapi.AccountStatCollection{uapi.FTPAccounts, uapi.EmailAccounts})
	if err != nil {
		t.Fatal(err)
	}

	gotFTP := false
	gotEmail := false

	for _, i := range r2 {
		switch i.StatType {
		case uapi.FTPAccounts:
			gotFTP = true
		case uapi.EmailAccounts:
			gotEmail = true
		}
	}

	if !gotFTP || !gotEmail {
		t.Fatal("Did not result in expected stats")
	}

	fmt.Printf("\nAccount Stats: %v\n", r2)

	// Email

	r3, err := conn.WHM.GetEmailAccountList()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("\nEmail accounts: %v\n", r3)

	// Files
	r4, err := conn.UAPI.GetDirectoryFileListing("")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Files: %d\n", len(r4))

	// Disk usage
	r5, err := conn.WHM.GetDiskUsage()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Disk Usage contents: %d\n", len(r5.CPanelResult.Data))

	// Site stats
	r6, err := conn.WHM.GetStatSites()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("AWSites: %d\n", len(r6))

	// Create email account
	r7, err := conn.UAPI.CreateEmailAccount("qwe@provtest.com", "WQE1242!#!@")
	if err != nil {
		if err.Error() != "The account qwe@provtest.com already exists!" {
			t.Fatal(err)
		}
	}

	fmt.Printf("Created email address: %v\n", r7)

	// Domain listing
	r8, err := conn.UAPI.GetDomainListing()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Domains: %d\n", len(r8))
}

func TestCreateAccountResponse(t *testing.T) {
	rawResponse := `{"metadata":{"command":"createacct","output":{"raw":"RAWOUTPUT"},"version":1,"result":1,"reason":"Account Creation Ok"},"data":{"nameserver2":"ns2.lucidcube.com","nameserver4":"",
	"nameservera3":null,"nameservera2":null,"package":"default","nameserverentry":null,"ip":"123.123.123.123","nameserverentry3":null,"nameserver3":"","nameservera4":null,"nameservera":null,
	"nameserverentry4":null,"nameserverentry2":null,"nameserver":"ns1.lucidcube.com"}}`

	resp := &whm.CreateAccountResponse{}
	decodeErr := json.Unmarshal([]byte(rawResponse), resp)
	if decodeErr != nil {
		t.Fail()
	}

	if resp.MetaData.Command != "createacct" {
		t.Log("Incorrect command")
		t.Fail()
	}

	if resp.MetaData.Version != 1 {
		t.Log("Incorrect version")
		t.Fail()
	}

	if resp.MetaData.Output.Raw != "RAWOUTPUT" {
		t.Log("Incorrect raw output")
		t.Fail()
	}

	if resp.MetaData.Result != 1 {
		t.Log("Incorrect result")
		t.Fail()
	}

	if resp.MetaData.Reason != "Account Creation Ok" {
		t.Log("Incorrect reason")
		t.Fail()
	}

	if resp.Data.Nameserver != "ns1.lucidcube.com" {
		t.Log("Incorrect nameserver")
		t.Fail()
	}

	if resp.Data.Nameserver2 != "ns2.lucidcube.com" {
		t.Log("Incorrect nameserver")
		t.Fail()
	}

	if resp.Data.Package != "default" {
		t.Log("Incorrect package")
		t.Fail()
	}

	if resp.Data.IP != "123.123.123.123" {
		t.Log("Incorrect IP")
		t.Fail()
	}

}
