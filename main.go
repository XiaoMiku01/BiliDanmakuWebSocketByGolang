package main

import (
	"flag"
	"log"
	"sync"

	"github.com/XiaoMiku01/BiliDanmakuWebSocket/danmaku"
)

func main() {
	var roomId = flag.String("r", "", "直播间号")
	flag.Parse()
	if *roomId == "" {
		log.Println("请输入直播间号")
		return
	}
	dm := danmaku.NewBiliRoom(*roomId)
	dm.Start()
	var w1 sync.WaitGroup
	w1.Add(1)
	w1.Wait()
}
