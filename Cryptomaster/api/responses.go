package api

// Convert JSON to Go struct
// https://mholt.github.io/json-to-go/
type CEXResponse struct {
	Timestamp             string  `json:"timestamp"`
	Low                   string  `json:"low"`
	High                  string  `json:"high"`
	Last                  string  `json:"last"`
	Volume                string  `json:"volume"`
	Volume30D             string  `json:"volume30d"`
	Bid                   float64 `json:"bid"`
	Ask                   float64 `json:"ask"`
	PriceChange           string  `json:"priceChange"`
	PriceChangePercentage string  `json:"priceChangePercentage"`
	Pair                  string  `json:"pair"`
}

// `json:"timestamp"`
// This ^ is metadata that has no meaning to go itself
// BUT packages can use reflection to read this metadata and do something with it

// here this metadata is meant to be used by the json package therefore json:xxxx
// could also look like this `json:"timestamp" someOterPackage:YYY`

// The backticks are used in go to use double quotes in a string.
