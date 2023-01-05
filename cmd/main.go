package cmd

import (
	"bufio"
	"fmt"
	"log"
	"onebot12_client/cmd/util"
	"os"
	"strings"
)

func GetCmd(self util.Self, msgChan chan []byte) {
	CmdHandleFunc := map[string]func([]string, util.Self) []byte{
		"private_message":        SendPrivateMsg,
		"pm":                     SendPrivateMsg,
		"group_message":          SendGroupMsg,
		"gm":                     SendGroupMsg,
		"friend_increase":        NoticeFriendIncrease,
		"fi":                     NoticeFriendIncrease,
		"friend_decrease":        NoticeFriendDecrease,
		"fd":                     NoticeFriendDecrease,
		"group_increase":         NoticeGroupIncrease,
		"gi":                     NoticeGroupIncrease,
		"group_decrease":         NoticeGroupDecrease,
		"gd":                     NoticeGroupDecrease,
		"private_message_delete": PrivateMsgDelete,
		"pmd":                    PrivateMsgDelete,
		"group_message_delete":   GroupMsgDelete,
		"gmd":                    GroupMsgDelete,
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.ReplaceAll(text, "\r", "")
		text = strings.ReplaceAll(text, "\n", "")
		args := strings.Split(text, " ")
		if len(args) == 0 {
			continue
		}
		if _, ok := CmdHandleFunc[args[0]]; ok {
			c := CmdHandleFunc[args[0]](args, self)
			if len(c) != 0 {
				msgChan <- c
			} else {
				log.Printf("Get empty message")
			}
		}
	}
}
