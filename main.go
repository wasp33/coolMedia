package main

import (
	"fmt"
	"log"
//	"time"
	"github.com/gorilla/websocket"
)

func main() {
	connect, _, err := websocket.DefaultDialer.Dial("wss://stream.yshyqxx.com/stream?streams=btcusdt@aggTrade", nil)
	if nil != err {
		log.Println(err)
		return
	}
	defer connect.Close()

	//sendPong
//  	t1 := time.Now()
	//go tickWriter(connect)

	for {
		//从 websocket 
		//messageType ，websocket 
		//messageData 
		messageType, messageData, err := connect.ReadMessage()
		if nil != err {
			log.Println(err)
			break
		}
		switch messageType {
		case websocket.TextMessage: 
			fmt.Println(string(messageData))
		case websocket.BinaryMessage: 
			fmt.Println(messageData)
		case websocket.CloseMessage: 
		case websocket.PingMessage: //Ping
			err := connect.WriteMessage(websocket.TextMessage, []byte("pong"))
			if nil != err {
				log.Println(err)
				break
			}
		case websocket.PongMessage: //Pong
		default:

		}


	}
}

