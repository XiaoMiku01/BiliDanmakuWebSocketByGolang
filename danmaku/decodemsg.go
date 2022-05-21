package danmaku

import (
	"bytes"
	"compress/zlib"
	// "encoding/json"
	"fmt"
	"io"
	// "log"
	"strconv"
	// "strings"

	"github.com/andybalholm/brotli"
)

var user_danmaku = map[float64]string{}

func decodeMessage(message []byte, outMsg chan []byte) {
	packet_len, _ := strconv.ParseUint(fmt.Sprintf("%x", message[:4]), 16, 64)
	ver, _ := strconv.ParseInt(fmt.Sprintf("%x", message[6:8]), 16, 64)
	op, _ := strconv.ParseInt(fmt.Sprintf("%x", message[8:12]), 16, 64)
	for int64(len(message)) > int64(packet_len) {
		decodeMessage(message[packet_len:], outMsg)
		message = message[:packet_len]
	}
	switch ver {
	case 0:
		outMsg <- message[16:]
		// log.Printf(string(message[16:]))

	case 1:
		if op == 3 {
			popular, _ := strconv.ParseInt(fmt.Sprintf("%x", message[16:]), 16, 64)
			// log.Printf("[人气]  %d", popular)
			popmsg := fmt.Sprintf(`{"cmd":"POP","count": %d}`, popular)
			// fmt.Println(popmsg)
			outMsg <- []byte(popmsg)
		}
	case 2:
		b := bytes.NewReader(message[16:])
		var out bytes.Buffer
		r, _ := zlib.NewReader(b)
		io.Copy(&out, r)
		decodeMessage(out.Bytes(), outMsg)
	case 3:
		b := bytes.NewReader(message[16:])
		r := brotli.NewReader(b)
		bytess, _ := io.ReadAll(r)
		decodeMessage(bytess, outMsg)
	default:
		// log.Printf("未知数据包")
		// log.Printf(string(message))
		// log.Printf("packet_len: %d, ver: %d, op: %d", packet_len, ver, op)
	}
}
