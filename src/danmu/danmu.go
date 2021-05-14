/*
 * @Descripttion: define danmu server
 * @version: 1.0
 * @Author: Nickname4th
 * @Date: 2021-05-10 14:38:37
 * @LastEditors: Nickname4th
 * @LastEditTime: 2021-05-12 10:04:55
 */
package danmu

import (
	"errors"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

// define danmu server.
// 1. Upgrade http protocol into websocket.
// 2. Define three chans, InChan receive message, OutChan send message
// and CloseChan send close signal to InChan and OutChan.
// 3. Define methods used to write and send.
type Danmu struct {
	WsDanmu *websocket.Conn
	InChan  chan []byte
	OutChan chan []byte
	// inform other chan to close
	CloseChan chan byte
	isClosed  bool
	mutex     sync.Mutex
}

/**
 * @name: ReadMessage
 * @msg: A thread safe message reader.
 * @param {[]byte} data
 * @param {error} err
 * @return {*}
 */
func (danmu *Danmu) ReadMessage() (data []byte, err error) {
	select {
	case data = <-danmu.InChan:
	case <-danmu.CloseChan:
		err = errors.New("danmu connection is closed!")
	}
	return
}

/**
 * @name: WriteMessage
 * @msg: A thread sage message writer.
 * @param {[]byte} data
 * @return {*}
 */
func (danmu *Danmu) WriteMessage(data []byte) (err error) {
	select {
	case danmu.OutChan <- data:
	case <-danmu.CloseChan:
		err = errors.New("danmu connection is closed!")
	}
	return
}

/**
 * @name: Close
 * @msg: A thread safe close, used to close the danmu connection once.
 * @param {*}
 * @return {*}
 */
func (danmu *Danmu) Close() {
	log.Println("close!!!")
	danmu.WsDanmu.Close()
	danmu.mutex.Lock()
	if !danmu.isClosed {
		close(danmu.CloseChan)
		danmu.isClosed = true
	}
	danmu.mutex.Unlock()
}

/**
 * @name: WriteLoop
 * @msg: wtire to the OutChan circularly.
 * @param {*}
 * @return {*}
 */
func (danmu *Danmu) WriteLoop() {
	log.Println("write loop start!")
	var (
		data []byte
		err  error
	)
	for {
		select {
		case data = <-danmu.OutChan:
		case <-danmu.CloseChan:
			goto ERR

		}
		if err = danmu.WsDanmu.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
		// danmu.OutChan <- data
	}
ERR:
	danmu.Close()
}

func (danmu *Danmu) ReadLoop() {
	log.Println("read loop start!")
	var (
		data []byte
		err  error
	)
	for {
		if _, data, err = danmu.WsDanmu.ReadMessage(); err != nil {
			goto ERR
		}
		select {
		case danmu.InChan <- data:
		case <-danmu.CloseChan:
			goto ERR
		}
	}
ERR:
	danmu.Close()
}

func InitDanmu(wsConn *websocket.Conn) (conn *Danmu, err error) {
	log.Println("iniatizaling!")
	conn = &Danmu{
		WsDanmu:   wsConn,
		InChan:    make(chan []byte, 1000),
		OutChan:   make(chan []byte, 1000),
		CloseChan: make(chan byte, 1),
	}
	go conn.ReadLoop()
	go conn.WriteLoop()
	return
}
