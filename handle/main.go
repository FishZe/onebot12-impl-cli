package handle

import (
	"log"
	"onebot12_client/cmd/action"
	"time"
)

func getAction(msg []byte) string {
	var layer = 0
	for i, v := range msg {
		if v == '{' || v == '[' {
			layer++
		} else if v == '}' || v == ']' {
			layer--
		} else if layer == 1 && v == '"' {
			if i+10 < len(msg) && string(msg[i+1:i+7]) == "action" && msg[i+7] == '"' {
				var from = i + 10
				var to int
				for to = from + 1; to < len(msg); to++ {
					if msg[to] == '"' {
						break
					}
				}
				return string(msg[from:to])
			}
		}
	}
	return ""
}

func Handler(MsgChan chan []byte) {
	actionHandleFunc := map[string]func([]byte) string{
		"send_message":   action.GetSendMessage,
		"delete_message": action.GetDeleteMessage,
	}
	for {
		select {
		case msg := <-MsgChan:
			action := getAction(msg)
			if _, ok := actionHandleFunc[action]; ok {
				log.Println(actionHandleFunc[action](msg))
			} else {
				log.Println("UnSupported Action: ", action)
			}
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}
