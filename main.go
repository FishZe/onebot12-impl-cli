package main

import (
	"log"
	"onebot12_client/client"
	"onebot12_client/cmd"
	"onebot12_client/cmd/util"
	"onebot12_client/handle"
)

func main() {
	f, err := readConfig()
	if err != nil {
		makeConfig()
		log.Printf("Config file not found, please edit the config file and restart the program.")
		return
	}
	msgChan := make(chan []byte, 1)
	ActionChan := make(chan []byte, 1)
	self := util.Self{UserId: f.Self.UserId, Platform: f.Self.Platform}
	c := client.Client{
		Url:               f.Onebot.Host,
		Port:              f.Onebot.Port,
		UserAgent:         f.Onebot.UserAgent,
		Impl:              f.Onebot.Impl,
		Version:           f.Onebot.Version,
		ReconnectInterval: f.Onebot.ReconnectInterval,
		Self:              self,
		MsgChan:           msgChan,
		ActionChan:        ActionChan}
	go c.Connect()
	go handle.Handler(ActionChan)
	cmd.GetCmd(self, msgChan)
}
