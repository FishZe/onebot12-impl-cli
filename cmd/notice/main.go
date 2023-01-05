package notice

import (
	"encoding/json"
	"onebot12_client/cmd/util"
	"time"
)

func GetFriendChangeNotice(self util.Self, changeType string, userId string) []byte {
	c := FriendChange{
		Id:         util.GetUUID(),
		Time:       float64(time.Now().Unix()),
		Type:       "notice",
		DetailType: changeType,
		Self:       self,
		UserId:     userId,
	}
	j, err := json.Marshal(c)
	if err != nil {
		return []byte{}
	}
	return j
}

func GetGroupChangeNotice(self util.Self, changeType string, groupId string, UserId string, OperatorId string) []byte {
	c := GroupMemberChange{
		Id:         util.GetUUID(),
		Time:       float64(time.Now().Unix()),
		Type:       "notice",
		DetailType: changeType,
		Self:       self,
		UserId:     UserId,
		GroupId:    groupId,
		OperatorId: OperatorId,
	}
	j, err := json.Marshal(c)
	if err != nil {
		return []byte{}
	}
	return j
}

func GetMsgDelete(self util.Self, messageId string, UserId string) []byte {
	c := MsgDelete{
		Id:         util.GetUUID(),
		Time:       float64(time.Now().Unix()),
		Type:       "notice",
		DetailType: "private_message_delete",
		Self:       self,
		UserId:     UserId,
		MessageId:  messageId,
	}
	j, err := json.Marshal(c)
	if err != nil {
		return []byte{}
	}
	return j
}

func GetGroupMsgDelete(self util.Self, messageId string, UserId string, groupId string, OperatorId string) []byte {
	c := GroupMsgDelete{
		Id:         util.GetUUID(),
		Time:       float64(time.Now().Unix()),
		Type:       "notice",
		DetailType: "group_message_delete",
		Self:       self,
		UserId:     UserId,
		GroupId:    groupId,
		MessageId:  messageId,
		OperatorId: OperatorId,
	}
	j, err := json.Marshal(c)
	if err != nil {
		return []byte{}
	}
	return j
}
