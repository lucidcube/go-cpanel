package response

// BaseWhmAPIResponse base response when calling WHMAPI
type BaseWhmAPIResponse struct {
	Metadata struct {
		Version int    `json:"version"`
		Command string `json:"command"`
		Reason  string `json:"reason"`
	} `json:"metadata"`
}
