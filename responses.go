package cpanel

import (
	"fmt"
)

// Cbool is bool type used when unmarshaling values such as '1' or "TRUE" to Go bools
type Cbool bool

// UnmarshalJSON unmarshals ConvertibleBoolean to Go bool type
func (b *Cbool) UnmarshalJSON(d []byte) error {
	s := string(d)
	if s == "1" || s == "true" {
		*b = true
	} else if s == "0" || s == "false" || s == "\"\"" {
		*b = false
	} else {
		return fmt.Errorf("Cannot unmarshal %s to Cbool", s)
	}
	return nil
}

// BaseWhmAPIResponse base response when calling WHMAPI
type BaseWhmAPIResponse struct {
	Metadata struct {
		Version int    `json:"version"`
		Command string `json:"command"`
		Reason  string `json:"reason"`
	} `json:"metadata"`
}

// CreateUserSessionResponse response from creation of a user session through WHMAPI
type CreateUserSessionResponse struct {
	BaseWhmAPIResponse
	Data struct {
		Session       string `json:"session"`
		Service       string `json:"service"`
		URL           string `json:"url"`
		SecurityToken string `json:"cp_security_token"`
		Expires       int64  `json:"expires"`
	} `json:"data"`
}
