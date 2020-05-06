package response

// BaseWhmAPIResponse base response when calling WHMAPI
type BaseWhmAPIResponse struct {
	Metadata struct {
		Version int    `json:"version"`
		Command string `json:"command"`
		Reason  string `json:"reason"`
	} `json:"metadata"`
}

type ExtendedBaseWhmApiResponse struct {
	MetaData struct {
		Command string `json:"command"`
		Output  struct {
			Raw string `json:"raw"`
		} `json:"output"`
		Version int    `json:"version"`
		Result  int    `json:"result"`
		Reason  string `json:"reason"`
	} `json:"metadata"`
}
