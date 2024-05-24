package common

import (
	"encoding/json"
)

type MsgType int

const (
	// client -> server 上行消息
	HeartBeat MsgType = 1 // 心跳包
	Action    MsgType = 2 // client产生了某种行为，给server发送执行动作的消息，由server来判断动作的合法性，并扭转游戏局面

	// server -> client 下行消息
	BroadCast MsgType = 101 // 广播消息，一般为游戏局面发生变化时，由server下发给client
	Notify    MsgType = 102 // 通知消息，一般为游戏在中间状态的临时通知
)

type Msg struct {
	MsgType MsgType `json:"msg_type"`
	Content string  `json:"content"` // 消息体，通常为一个json / yaml，解析后转换为具体的结构
	Sender  string  `json:"sender"`  // 消息的发送方是谁
}

func NewMsg(msgType MsgType, content string, sender string) []byte {
	m := Msg{
		MsgType: msgType,
		Content: content,
		Sender:  sender,
	}

	data, _ := json.Marshal(m)
	return data
}

func DecodeMsg(msg []byte) Msg {
	var m Msg
	_ = json.Unmarshal(msg, &m)
	return m
}
