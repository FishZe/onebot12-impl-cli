package client

import (
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
	"onebot12_client/cmd/meta"
	"onebot12_client/cmd/util"
	"strconv"
	"sync"
	"time"
)

type Client struct {
	Url               string
	Port              int
	UserAgent         string
	Impl              string
	Version           string
	AccessToken       string
	ReconnectInterval int
	Self              util.Self
	MsgChan           chan []byte
	ActionChan        chan []byte
	connect           *websocket.Conn
	connected         bool
	lock              sync.RWMutex
}

func (c *Client) ConnectSDK() error {
	err := error(nil)
	u := url.URL{Scheme: "ws", Host: c.Url + ":" + strconv.Itoa(c.Port), Path: "/"}
	var header = make(http.Header)
	header.Add("User-Agent", c.UserAgent)
	header.Add("Sec-WebSocket-Protocol", "12."+c.Impl)
	if c.AccessToken != "" {
		header.Add("Authorization", "Bearer "+c.AccessToken)
	}
	c.connect, _, err = websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) receiveWsMsg() error {
	for {
		_, message, err := c.connect.ReadMessage()
		if err != nil {
			c.connected = false
			return err
		}
		c.ActionChan <- message
	}
}

func (c *Client) heartBeat() {
	for {
		if c.connected {
			time.Sleep(time.Second * 5)
			c.lock.Lock()
			err := c.connect.WriteMessage(websocket.TextMessage, meta.GetHeartBeatEvent(5))
			c.lock.Unlock()
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func (c *Client) SendMessage(Msg []byte) error {
	if c.connected {
		c.lock.Lock()
		err := c.connect.WriteMessage(websocket.TextMessage, Msg)
		c.lock.Unlock()
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("onebot SDK not connected")
	}
}

func (c *Client) WaitForSend() {
	for {
		select {
		case msg, ok := <-c.MsgChan:
			if ok {
				err := c.SendMessage(msg)
				if err != nil {
					log.Printf("Send message failed: %v\n", err)
				}
			}
		default:
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func (c *Client) Connect() {
	for {
		c.connected = false
		err := c.ConnectSDK()
		go c.heartBeat()
		go c.WaitForSend()
		if err != nil {
			log.Printf("Connect to OneBot SDK failed: %v\n", err)
		} else {
			c.connected = true
			err = c.SendMessage(meta.GetConnectEvent(c.Impl, c.Version))
			if err != nil {
				log.Printf("Send connect event failed: %v\n", err)
			}
			log.Printf("Connect to OneBot SDK successfully\n")
			err = c.receiveWsMsg()
			if err != nil {
				log.Printf("Receive message failed: %v\n", err)
			}
		}
		time.Sleep(time.Duration(c.ReconnectInterval) * time.Second)
	}
}
