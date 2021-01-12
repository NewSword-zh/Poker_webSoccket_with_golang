package main

import (
	"encoding/json"
	"fmt"
	"time"
)

//定时向广播站法送信息
func testPush(hub *Hub){
	ticker := time.NewTicker(time.Second*60)
	count := 0
	for{
		select{
		case <-ticker.C:
			count += 1
			message := Message{time.Now().Format("2006-01-02 15:04:05")}
			messageByte,_ := json.Marshal(message)
			hub.broadcast<- messageByte
			fmt.Println(message)

		default:

		}
	}


}
