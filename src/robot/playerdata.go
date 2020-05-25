package robot

type PlayerData struct {
	Unionid       string
	Nickname      string
	AccountID     int
	RoomType      int // 房间类型 0 练习 1 房卡匹配 2 私人
	RoomCards     int
	MaxPlayers    int
	RedPacketType int
	Position      int
	Role          int
}