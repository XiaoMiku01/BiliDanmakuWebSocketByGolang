package danmaku

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/andybalholm/brotli"
)

func decodeMessage(message []byte) {
	packet_len, _ := strconv.ParseUint(fmt.Sprintf("%x", message[:4]), 16, 64)
	ver, _ := strconv.ParseInt(fmt.Sprintf("%x", message[6:8]), 16, 64)
	op, _ := strconv.ParseInt(fmt.Sprintf("%x", message[8:12]), 16, 64)
	for int64(len(message)) > int64(packet_len) {
		decodeMessage(message[packet_len:])
		message = message[:packet_len]
	}
	switch ver {
	case 0:
		log.Printf(string(message[16:]))
	case 1:
		if op == 3 {
			popular, _ := strconv.ParseInt(fmt.Sprintf("%x", message[16:]), 16, 64)
			log.Printf("人气: %d", popular)
		}
	case 2:
		b := bytes.NewReader(message[16:])
		var out bytes.Buffer
		r, _ := zlib.NewReader(b)
		io.Copy(&out, r)
		decodeMessage(out.Bytes())
	case 3:
		b := bytes.NewReader(message[16:])
		r := brotli.NewReader(b)
		bytess, _ := io.ReadAll(r)
		decodeMessage(bytess)
	default:
		// log.Printf("未知数据包")
		// log.Printf(string(message))
		// log.Printf("packet_len: %d, ver: %d, op: %d", packet_len, ver, op)
	}
}
