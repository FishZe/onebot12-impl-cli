package cmd

import (
	"log"
	"onebot12_client/cmd/message"
	"onebot12_client/cmd/notice"
	"onebot12_client/cmd/util"
)

func SendPrivateMsg(args []string, self util.Self) []byte {
	// private_message [user_id] [message]
	if len(args) < 3 {
		log.Printf("Lack of arguments")
		return []byte{}
	}
	if len(args) > 3 {
		for i := 3; i < len(args); i++ {
			args[2] += " " + args[i]
		}
	}
	msg := message.DecodeAltMessage(args[2])
	if len(msg) == 0 {
		log.Printf("Decode message error")
		return []byte{}
	}
	c := message.GetPrivateMsg(self, args[1], msg)
	return c
}

func SendGroupMsg(args []string, self util.Self) []byte {
	// group_message [group_id] [user_id] [message]
	if len(args) < 4 {
		log.Printf("Lack of arguments")
		return []byte{}
	}
	if len(args) > 4 {
		for i := 4; i < len(args); i++ {
			args[2] += " " + args[i]
		}
	}
	msg := message.DecodeAltMessage(args[3])
	if len(msg) == 0 {
		log.Printf("Decode message error")
		return []byte{}
	}
	c := message.GetGroupMsg(self, args[1], args[2], msg)
	return c
}

func NoticeFriendIncrease(args []string, self util.Self) []byte {
	// friend_increase [user_id]
	if len(args) < 2 {
		log.Printf("Lack of arguments")
		return []byte{}
	}
	c := notice.GetFriendChangeNotice(self, args[1], "friend_increase")
	return c
}

func NoticeFriendDecrease(args []string, self util.Self) []byte {
	// friend_decrease [user_id]
	if len(args) < 2 {
		log.Printf("Lack of arguments")
		return []byte{}
	}
	c := notice.GetFriendChangeNotice(self, args[1], "friend_decrease")
	return c
}

func NoticeGroupIncrease(args []string, self util.Self) []byte {
	// group_increase [group_id] [user_id] [operator_id]
	if len(args) < 4 {
		log.Printf("Lack of arguments")
		return []byte{}
	}
	c := notice.GetGroupChangeNotice(self, "group_increase", args[1], args[2], args[3])
	return c
}

func NoticeGroupDecrease(args []string, self util.Self) []byte {
	// group_decrease [group_id] [user_id] [operator_id]
	if len(args) < 4 {
		log.Printf("Lack of arguments")
		return []byte{}
	}
	c := notice.GetGroupChangeNotice(self, "group_decrease", args[1], args[2], args[3])
	return c
}

func PrivateMsgDelete(args []string, self util.Self) []byte {
	// private_message_delete [message_id] [user_id]
	if len(args) < 3 {
		log.Printf("Lack of arguments")
		return []byte{}
	}
	c := notice.GetMsgDelete(self, args[1], args[2])
	return c
}

func GroupMsgDelete(args []string, self util.Self) []byte {
	// group_message_delete [message_id] [group_id] [user_id] [operator_id]
	if len(args) < 5 {
		log.Printf("Lack of arguments")
		return []byte{}
	}
	c := notice.GetGroupMsgDelete(self, args[1], args[3], args[2], args[4])
	return c
}
