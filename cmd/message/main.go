package message

import (
	"encoding/json"
	"math/rand"
	"onebot12_client/cmd/util"
	"strconv"
	"strings"
	"time"
)

func GetPrivateMsg(self util.Self, userId string, message []Message) []byte {
	c := RevMessage{
		Id:         util.GetUUID(),
		Time:       float64(time.Now().Unix()),
		Type:       "message",
		DetailType: "private",
		MessageId:  strconv.Itoa(rand.Intn(100000) + 1000),
		Self:       self,
		Message:    message,
		UserId:     userId,
	}
	for _, v := range message {
		if v.Type == "text" {
			c.AltMessage += v.Data.Text
		} else {
			c.AltMessage += "[" + v.Type + "]"
		}
	}
	j, err := json.Marshal(c)
	if err != nil {
		return []byte{}
	}
	return j
}

func GetGroupMsg(self util.Self, GroupId string, UserId string, message []Message) []byte {
	c := RevMessage{
		Id:         util.GetUUID(),
		Time:       float64(time.Now().Unix()),
		Type:       "message",
		DetailType: "group",
		MessageId:  strconv.Itoa(rand.Intn(100000) + 1000),
		Self:       self,
		Message:    message,
		UserId:     UserId,
		GroupId:    GroupId,
	}
	for _, v := range message {
		if v.Type == "text" {
			c.AltMessage += v.Data.Text
		} else {
			c.AltMessage += "[" + v.Type + "]"
		}
	}
	j, err := json.Marshal(c)
	if err != nil {
		return []byte{}
	}
	return j
}

func DecodeAltMessage(message string) []Message {
	var m []Message
	now := 0
	for strings.Index(message, "[") != -1 {
		if strings.Index(message, "[") != now {
			nowMsg := Message{Type: "text"}
			nowMsg.Data.Text = message[now:strings.Index(message, "[")]
			m = append(m, nowMsg)
			message = message[strings.Index(message, "["):]
			now = 0
		}
		nowMsg := Message{Type: message[strings.Index(message, "[")+1 : strings.Index(message, "]")]}
		if nowMsg.Type == "image" || nowMsg.Type == "voice" || nowMsg.Type == "audio" || nowMsg.Type == "video" || nowMsg.Type == "file" {
			nowMsg.Data.FileId = util.GetUUID()
		} else if nowMsg.Type != "mention_all" {
			return []Message{}
		}
		m = append(m, nowMsg)
		message = message[strings.Index(message, "]")+1:]
		now = 0
	}
	if len(message) != 0 {
		nowMsg := Message{Type: "text"}
		nowMsg.Data.Text = message
		m = append(m, nowMsg)
	}
	return m
}

func EncodeAltMessage(message []Message) string {
	var m string
	for _, v := range message {
		if v.Type == "text" {
			m += v.Data.Text
		} else {
			m += "[" + v.Type + "]"
		}
	}
	return m
}
