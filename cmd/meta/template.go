package meta

type Connect struct {
	ID         string  `json:"id"`
	Time       float64 `json:"time"`
	Type       string  `json:"type"`
	DetailType string  `json:"detail_type"`
	SubType    string  `json:"sub_type"`
	Version    struct {
		Impl          string `json:"impl"`
		Version       string `json:"version"`
		OnebotVersion string `json:"onebot_version"`
	} `json:"version"`
}

type HeartBeat struct {
	ID         string  `json:"id"`
	Time       float64 `json:"time"`
	Type       string  `json:"type"`
	DetailType string  `json:"detail_type"`
	SubType    string  `json:"sub_type"`
	Interval   int     `json:"interval"`
}
