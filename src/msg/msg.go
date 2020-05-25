package msg

type C2S_TokenAuthorize struct {
	Token   string
}

type C2S_Heartbeat struct {
}

type C2S_SetRobotData struct {
	LoginIP string
}

type C2S_EnterRoom struct {
}

type C2S_GetAllPlayers struct {
}

type C2S_Prepare struct {
}

type C2S_FakeWXPay struct {
	TotalFee int
}

type C2SL_StartMatch struct {
	RoomType	int
	ItemType	int
}

type C2SL_SendRedPacket struct {
	RedPacketMetaData
}

type RedPacketMetaData struct {
	Quota				int
	Num 				int
	Boom				int
}

type C2SL_TakenRedPacket struct {

}

type C2SL_ExitRoom struct {

}