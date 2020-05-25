package robot

import (
	"encoding/json"
	"math/rand"
	"redpacket-sweep-robot/common"
	"redpacket-sweep-robot/conf"
	"redpacket-sweep-robot/netC"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network"
	"github.com/name5566/leaf/timer"
)

var (
	//addr = "ws://gdmj.shenzhouxing.com:3660"
	addr        = conf.GetCfgGameInfo().HallAddress
	clients     []*netC.Client
	unionids    []string
	nicknames   []string
	headimgurls []string
	loginIPs    []string
	count       = 0
	mu          sync.Mutex
	Play        = true

	RobotNumber *int // 机器人数量

	dispatcher *timer.Dispatcher
)

func init() {
	rand.Seed(time.Now().UnixNano())
	names, ips := make([]string, 0), make([]string, 0)
	names, _ = common.ReadFile("conf/robot_nickname.txt")
	//names = common.Shuffle2(names)

	ips, _ = common.ReadFile("conf/robot_ip.txt")
	nicknames = names
	loginIPs = ips
	//ips = common.Shuffle2(ips)
	/*
		if err == nil {
			nicknames = append(nicknames, names[conf.GetCfgGameInfo().RobotIndex:conf.GetCfgGameInfo().RobotIndex+conf.GetCfgGameInfo().RobotNumber]...)
			loginIPs = append(loginIPs, ips[conf.GetCfgGameInfo().RobotIndex:conf.GetCfgGameInfo().RobotIndex+conf.GetCfgGameInfo().RobotNumber]...)
		} else {
			log.Debug("read file error: %v", err)
		}
	*/
	//temp := rand.Perm(*RobotNumber)
	for i := conf.GetCfgGameInfo().RobotIndex; i < conf.GetCfgGameInfo().RobotIndex+conf.GetCfgGameInfo().RobotNumber; i++ {
		unionids = append(unionids, strconv.Itoa(i))
		headimgurls = append(headimgurls, "http://111.231.252.178:8088/static/robotImg/"+strconv.Itoa(i)+".jpg")
	}
	dispatcher = timer.NewDispatcher(0)
}

func Init() {
	/*
		Play = flag.Bool("Play", false, "control robot enter game")
		flag.Parse()
		log.Debug("Play: %v", *Play)
	*/
	count = 0
	client := new(netC.Client)
	client.Addr = addr
	client.ConnNum = conf.GetCfgGameInfo().RobotNumber
	client.ConnectInterval = 3 * time.Second
	client.HandshakeTimeout = 10 * time.Second
	client.PendingWriteNum = 100
	client.MaxMsgLen = 4096
	client.NewAgent = newAgent

	client.Start()
	clients = append(clients, client)
}

func Destroy() {
	for _, client := range clients {
		client.Close()
	}
}

type Agent struct {
	conn       *netC.MyConn
	playerData *PlayerData
	once       *sync.Once
	timer      *time.Timer
	cronTimer  *timer.Cron
}

func newAgent(conn *netC.MyConn) network.Agent {
	a := new(Agent)
	a.conn = conn
	a.playerData = newPlayerData()
	a.once = new(sync.Once)
	return a
}

func newPlayerData() *PlayerData {
	playerData := new(PlayerData)
	playerData.Position = -1

	return playerData
}

func (a *Agent) writeMsg(msg interface{}) {
	a.conn.WriteMsg2(msg)
	return
}

func (a *Agent) readMsg() {
	for {
		msg, err := a.conn.ReadMsg()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Debug("error: %v", err)
			}
			break
		}
		log.Debug("%s", msg)
		jsonMap := map[string]interface{}{}
		err = json.Unmarshal(msg, &jsonMap)
		if err == nil {
			a.handleMsg(jsonMap)
		} else {
			log.Error("%v", err)
		}
	}
}

func (a *Agent) Run() {
	go func() {
		for {
			(<-dispatcher.ChanTimer).Cb()
		}
	}()

	go a.wechatLogin()

	a.readMsg()
}

func (a *Agent) OnClose() {

}
