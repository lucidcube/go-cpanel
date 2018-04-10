package response

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
