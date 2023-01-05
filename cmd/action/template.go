package action

import "onebot12_client/cmd/message"

type SendMessage struct {
	Action string `json:"action"`
	Params struct {
		DetailType string            `json:"detail_type"`
		GroupId    string            `json:"group_id"`
		UserId     string            `json:"user_id"`
		Message    []message.Message `json:"message"`
	} `json:"params"`
}

type DeleteMessage struct {
	Action string `json:"action"`
	Params struct {
		MessageId string `json:"message_id"`
	} `json:"params"`
}
