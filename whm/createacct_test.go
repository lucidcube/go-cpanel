package whm_test

import (
	"testing"
	"encoding/json"
	"github.com/lucidcube/go-cpanel/whm"
)

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
