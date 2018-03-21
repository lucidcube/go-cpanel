package cpanel

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

// StatsResponse is response from UAPI when querying statistics
type StatsResponse struct {
	Stats []StatResponse `json:"data"`
}

// StatResponse is a single UAPI statistic
type StatResponse struct {
	ZeroIsUnlimited bool   `json:"zeroisunlimited"`
	Percent20       int    `json:"percent20"`
	Percent10       int    `json:"percent10"`
	Percent5        int    `json:"percent5"`
	Percent         int    `json:"percent"`
	Item            string `json:"item"`
	Max             string `json:"max"`
	LangKey         string `json:"langkey"`
	ID              string `json:"id"`
	Module          string `json:"module"`
	Count           string `json:"count"`
	Name            string `json:"name"`
	Normalized      int    `json:"normalized"`
	Units           string `json:"units"`
	NearLimitPhrase string `json:"near_limit_phrase"`
	MaxedPhrase     string `json:"maxed_phrase"`
}
