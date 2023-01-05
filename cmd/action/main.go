package action

import (
	"encoding/json"
	"log"
	"onebot12_client/cmd/message"
)

func GetSendMessage(msg []byte) string {
	var a SendMessage
	err := json.Unmarshal(msg, &a)
	if err != nil {
		log.Println("Get Send Message Action Error: ", err)
		return ""
	}
	ret := "send_message: "
	if a.Params.DetailType == "group" {
		ret += "group[" + a.Params.GroupId + "] :"
	} else {
		ret += "private[" + a.Params.UserId + "] :"
	}
	encodeMsg := message.EncodeAltMessage(a.Params.Message)
	ret += encodeMsg
	return ret
}

func GetDeleteMessage(msg []byte) string {
	var a DeleteMessage
	err := json.Unmarshal(msg, &a)
	if err != nil {
		log.Println("Get Delete Message Action Error: ", err)
		return ""
	}
	ret := "delete_message: " + a.Params.MessageId
	return ret
}
