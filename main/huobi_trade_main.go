package main

import "exchange_websocket/huobi_websocket"

func main() {
	hb := huobi_websocket.HbWebsocketInit()
	hb.HbTradeWebsocket()
	for {
		hb.Subscribe()
		hb.ReadMessage()
	}
}
