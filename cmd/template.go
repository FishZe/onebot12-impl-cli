package cmd

type Event struct {
	ID         string  `json:"id"`
	Time       float64 `json:"time"`
	Type       string  `json:"type"`
	DetailType string  `json:"detail_type"`
	SubType    string  `json:"sub_type"`
}
