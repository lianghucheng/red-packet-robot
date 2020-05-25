package robot

import (
	"redpacket-sweep-robot/conf"
	"redpacket-sweep-robot/msg"
)


func (a *Agent) wechatLogin() {
	mu.Lock()
	defer mu.Unlock()
	//a.playerData.Unionid = unionids[count]
	//a.playerData.Nickname = nicknames[count]
	//a.playerData.Nickname = "我是测试人"

	a.writeMsg(&msg.C2S_TokenAuthorize{
		Token:conf.GetCfgGameInfo().Token[count],
	})

	count++
}

func (a *Agent) sendHeartbeat() {
	a.writeMsg(&msg.C2S_Heartbeat{})
}

func (a *Agent) StartMatching(itemType int){
	a.writeMsg(&msg.C2SL_StartMatch{
		RoomType:1,
		ItemType:itemType,
	})
}