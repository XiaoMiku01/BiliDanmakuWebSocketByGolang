package danmaku

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/XiaoMiku01/BiliDanmakuWebSocket/utils"
	"github.com/gorilla/websocket"
)

type BiliRoom struct {
	roomId  string
	address string
	token   string
	conn    *websocket.Conn
	recMsg  chan []byte
	isAlive bool
	timeout int
}

func NewBiliRoom(roomId string) *BiliRoom {
	var recMsg = make(chan []byte, 10)
	var conn *websocket.Conn
	return &BiliRoom{
		roomId:  roomId,
		conn:    conn,
		recMsg:  recMsg,
		isAlive: false,
		timeout: 3,
	}
}

func (b *BiliRoom) Start() {
	for {
		if !b.isAlive {
			if err := b.getRoomInfo(); err != nil {
				log.Println(err)
				continue
			}
			if err := b.connect(); err != nil {
				log.Println(err)
				continue
			}
			if err := b.verify(); err != nil {
				log.Println("房间连接失败:", err)
				continue
			}
			go b.readMessage()
			go b.decodeMsg()
			go b.heartBeat()
		}
		time.Sleep(time.Second * time.Duration(b.timeout))
	}
}

func (b *BiliRoom) connect() error {
	var err error
	b.conn, _, err = websocket.DefaultDialer.Dial(fmt.Sprintf("wss://%s/%s", b.address, "sub"), nil)
	if err != nil {
		return err
	}
	b.isAlive = true
	return nil
}

func (b *BiliRoom) verify() error {
	// 发送房间验证包
	roomInfo := fmt.Sprintf(`{"uid": 0, "roomid": %s, "protover": 3, "platform": "web", "type": 2, "key": "%s"}`, b.roomId, b.token)
	err := b.conn.WriteMessage(websocket.BinaryMessage, __pack(roomInfo, 1, 7))
	if err != nil {
		log.Println("write:", err)
		return err
	}
	_, _, err = b.conn.ReadMessage()
	if err != nil {
		b.isAlive = false
		return err
	}
	log.Println("房间连接成功")
	return nil
}

func (b *BiliRoom) readMessage() error {
	for {
		if !b.isAlive {
			return errors.New("房间已断开连接！")
		}
		_, message, err := b.conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			b.isAlive = false
			return err
		}
		b.recMsg <- message
	}
}

func (b *BiliRoom) decodeMsg() {
	for msg := range b.recMsg {
		if !b.isAlive {
			return
		}
		decodeMessage(msg)
	}
}

func (b *BiliRoom) heartBeat() error {
	// 心跳包
	for {
		time.Sleep(time.Second * time.Duration(30))
		err := b.conn.WriteMessage(websocket.BinaryMessage, __pack("[object Object]", 1, 2))
		if err != nil {
			log.Println(err)
			b.isAlive = false
			return err
		}
		log.Println("心跳包发送成功")

	}
}

func (b *BiliRoom) getRoomInfo() error {
	// 获取房间弹幕地址
	ri := new(RoomInfo)
	res, err := http.Get("https://api.live.bilibili.com/xlive/web-room/v1/index/getDanmuInfo?id=" + b.roomId)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	err = json.Unmarshal(data, &ri)
	if ri.Code != 0 || len(ri.Data.HostList) == 0 {
		err = errors.New("获取房间信息失败！")
	}
	if err != nil {
		return err
	}
	b.address = ri.Data.HostList[0].Host
	b.token = ri.Data.Token
	return nil
}

func __pack(s string, i int, j int) []byte {
	// 字节流打包
	format := []string{"H", "H", "I", "I"}
	values := []interface{}{16, i, j, 1}
	bp := new(utils.BinaryPack)
	data, _ := bp.Pack(format, values)
	data = append(data, []byte(s)...)
	bp2 := new(utils.BinaryPack)
	data2, _ := bp2.Pack([]string{"I"}, []interface{}{len(data) + 4})
	data = append(data2, data...)
	return data
}
