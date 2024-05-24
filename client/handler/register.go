package handler

import (
	"encoding/json"
	"fmt"

	"github.com/yydaily/games-engine/client/conn"
)

type RegisterMsg struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func newRegisterMsg(userName, password string) []byte {
	msg := RegisterMsg{
		UserName: userName,
		Password: password,
	}
	m, e := json.Marshal(msg)
	if e != nil {
		panic(e)
	}
	return m
}

func Register(userName, password string) {
	msg := newRegisterMsg(userName, password)
	err := conn.Client.SendToServer(msg)
	if err != nil {
		fmt.Println("注册失败，请重试")
	}
}
