package conn

import (
	"fmt"
	"time"

	"github.com/yydaily/games-engine/common"
)

func HeartBeat() {
	go func() {
		i := 1
		for {
			select {
			case <-time.After(time.Second * 5):
				common.NewMsg(common.HeartBeat, fmt.Sprintf("%d", i), "")
			}
		}
	}()
}
