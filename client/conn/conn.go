package conn

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/yydaily/games-engine/client/config"
	"github.com/yydaily/games-engine/common"
)

type TCPClient struct {
	Conn     *net.TCPConn
	HawkAddr *net.TCPAddr
}

var Client *TCPClient

func ConnectServer() {
	hostPort := fmt.Sprintf("%s:%d", config.Conf.Server.Host, config.Conf.Server.Port)

	hawkAddr, err := net.ResolveTCPAddr("tcp", hostPort)
	if err != nil {
		log.Printf("resolve tcp addr failed, host_port=%s, err=%s", hostPort, err)
		time.Sleep(time.Second)
		ConnectServer() // try connect again
		return
	}
	Client.HawkAddr = hawkAddr

	connection, err := net.DialTCP("tcp", nil, hawkAddr)
	if err != nil {
		log.Printf("dial tcp failed, host_port=%s, err=%s", hostPort, err)
		time.Sleep(time.Second)
		ConnectServer() // try connect again
		return
	}
	Client.Conn = connection

	log.Printf("dial tcp success, host_port=%s", hostPort)
}

func (c *TCPClient) SendToServer(data []byte) error {
	_, err := c.Conn.Write(data)
	return err
}

func (c *TCPClient) ReadFromServer() chan common.Msg {
	dataChan := make(chan common.Msg, 10)
	data := make([]byte, config.Conf.Server.MaxDataBody)
	go func() {
		for {
			n, err := Client.Conn.Read(data)
			if err != nil {
				if err == io.EOF {
					// disconnect
					Client.Conn.Close()
					log.Printf("client disconnect, host_port=%s", Client.HawkAddr)
				}
			}
			if n > 0 && n < 1025 {
				dataChan <- common.DecodeMsg(data[:n])
			}
		}
	}()

	return dataChan
}
