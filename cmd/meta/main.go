package meta

import (
	"encoding/json"
	"onebot12_client/cmd/util"
	"time"
)

func GetConnectEvent(Impl string, Version string) []byte {
	c := Connect{
		ID:         util.GetUUID(),
		Time:       float64(time.Now().Unix()),
		Type:       "meta",
		DetailType: "connect",
	}
	c.Version.Impl = Impl
	c.Version.Version = Version
	c.Version.OnebotVersion = "12"
	j, err := json.Marshal(c)
	if err != nil {
		return []byte{}
	}
	return j
}

func GetHeartBeatEvent(Interval int) []byte {
	c := HeartBeat{
		ID:         util.GetUUID(),
		Time:       float64(time.Now().Unix()),
		Type:       "meta",
		DetailType: "heartbeat",
		Interval:   Interval,
	}
	j, err := json.Marshal(c)
	if err != nil {
		return []byte{}
	}
	return j
}
