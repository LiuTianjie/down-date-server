/*
 * @Descripttion: Danmu api
 * @version: 1.0
 * @Author: Nickname4th
 * @Date: 2021-05-10 14:34:32
 * @LastEditors: Nickname4th
 * @LastEditTime: 2021-05-12 09:55:05
 */
package api

import (
	"down-date-server/src/danmu"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// initialize danmu servers

func DanmuHandler(c *gin.Context) {
	var (
		wsConn *websocket.Conn
		err    error
		dan    *danmu.Danmu
		data   []byte
	)
	if wsConn, err = upGrader.Upgrade(c.Writer, c.Request, nil); err != nil {
		return
	}
	if dan, err = danmu.InitDanmu(wsConn); err != nil {
		goto ERR
	}
	// go func() {
	// 	var (
	// 		err error
	// 	)
	// 	for {
	// 		if err = dan.WriteMessage([]byte("heart beat")); err != nil {
	// 			return
	// 		}
	// 		log.Println("send heart beat")
	// 		time.Sleep(2 * time.Second)
	// 	}
	// }()
	for {
		if data, err = dan.ReadMessage(); err != nil {
			goto ERR
		}
		if err = dan.WriteMessage(data); err != nil {
			goto ERR
		}

	}
ERR:
	dan.Close()
}
