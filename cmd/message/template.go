package message

import "onebot12_client/cmd/util"

type Message struct {
	Type string `json:"type"`
	Data struct {
		Text      string  `json:"text,omitempty"`
		FileId    string  `json:"file_id,omitempty"`
		UserId    string  `json:"user_id,omitempty"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Title     string  `json:"title"`
		Content   string  `json:"content"`
		MessageId string  `json:"message_id"`
	} `json:"data"`
}

type RevMessage struct {
	Id         string    `json:"id"`
	Time       float64   `json:"time"`
	Type       string    `json:"type"`
	DetailType string    `json:"detail_type"`
	SubType    string    `json:"sub_type"`
	MessageId  string    `json:"message_id"`
	Self       util.Self `json:"self"`
	Message    []Message `json:"message"`
	AltMessage string    `json:"alt_message"`
	UserId     string    `json:"user_id"`
	GroupId    string    `json:"group_id"`
}
