package robot

import (
	"fmt"
	"github.com/name5566/leaf/log"
	"redpacket-sweep-robot/msg"
	"time"
)

var isSend = false
var Item *int

func (a *Agent) handleMsg(jsonMap map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Release("检测到异常")
		}
	}()

	if _, ok := jsonMap[`S2C_Heartbeat`]; ok {
		a.sendHeartbeat()
	} else if _, ok := jsonMap[`S2C_Authorize`]; ok {
		time.Sleep(1.*time.Second)

		a.StartMatching(*Item)
	} else if _, ok := jsonMap[`SL2C_EnterRoom`]; ok {
		a.writeMsg(&msg.C2SL_TakenRedPacket{})

		time.Sleep(1 * time.Second)
		if !isSend {
			isSend = true
			a.writeMsg(&msg.C2SL_SendRedPacket{
				RedPacketMetaData:msg.RedPacketMetaData{
					Quota:80,
					Num:1,
					Boom:5,
				},
			})
		}
	} else if _, ok := jsonMap[`SL2C_StartGame`]; ok {
		time.AfterFunc(5 * time.Second, func() {
			a.writeMsg(&msg.C2SL_TakenRedPacket{})
		})
	} else if res, ok := jsonMap[`SL2C_TakenRedPacket`];ok {
		fmt.Println("SL2C_TakenRedPacket",res)
	} else if _, ok := jsonMap[`SL2C_EndGame`]; ok {
		fmt.Println("【游戏结束】")
	} else if res, ok := jsonMap[`SL2C_RoundResult`]; ok {
		fmt.Println("SL2C_RoundResult",res)
	} else if res, ok := jsonMap[`SL2C_ExitRoom`]; ok {
		fmt.Println("SL2C_ExitRoom",res)
	}
}