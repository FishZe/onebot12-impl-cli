package notice

import "onebot12_client/cmd/util"

type FriendChange struct {
	Id         string    `json:"id"`
	Self       util.Self `json:"self"`
	Time       float64   `json:"time"`
	Type       string    `json:"type"`
	DetailType string    `json:"detail_type"`
	SubType    string    `json:"sub_type"`
	UserId     string    `json:"user_id"`
}

type MsgDelete struct {
	Id         string    `json:"id"`
	Self       util.Self `json:"self"`
	Time       float64   `json:"time"`
	Type       string    `json:"type"`
	DetailType string    `json:"detail_type"`
	SubType    string    `json:"sub_type"`
	MessageId  string    `json:"message_id"`
	UserId     string    `json:"user_id"`
}

type GroupMemberChange struct {
	Id         string    `json:"id"`
	Self       util.Self `json:"self"`
	Time       float64   `json:"time"`
	Type       string    `json:"type"`
	DetailType string    `json:"detail_type"`
	SubType    string    `json:"sub_type"`
	UserId     string    `json:"user_id"`
	GroupId    string    `json:"group_id"`
	OperatorId string    `json:"operator_id"`
}

type GroupMsgDelete struct {
	Id         string    `json:"id"`
	Self       util.Self `json:"self"`
	Time       float64   `json:"time"`
	Type       string    `json:"type"`
	DetailType string    `json:"detail_type"`
	SubType    string    `json:"sub_type"`
	GroupId    string    `json:"group_id"`
	MessageId  string    `json:"message_id"`
	UserId     string    `json:"user_id"`
	OperatorId string    `json:"operator_id"`
}
