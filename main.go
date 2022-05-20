package main

import (
	"encoding/json"
	"flag"
	"log"
	// "sync"

	. "github.com/XiaoMiku01/BiliDanmakuWebSocket/bilistruct"
	"github.com/XiaoMiku01/BiliDanmakuWebSocket/danmaku"
)

func main() {
	var roomId = flag.String("r", "", "直播间号")
	var user_danmaku = map[int]string{}
	flag.Parse()
	if *roomId == "" {
		log.Println("请输入直播间号")
		return
	}
	dm := danmaku.NewBiliRoom(*roomId)
	go dm.Start()
	for message := range dm.OutMsg {
		mi := new(MessageInfo)
		json.Unmarshal(message[16:], &mi)
		switch mi.Cmd {
		case "POP":
			// 人气
			pi := new(PopInfo)
			json.Unmarshal(message[16:], &pi)
			log.Printf("[人气] %d", pi.Count)

		case "DANMU_MSG":
			// 弹幕
			log.Printf("[弹幕] %s: %s", mi.Info.([]interface{})[2].([]interface{})[1], mi.Info.([]interface{})[1])
			user_danmaku[int(mi.Info.([]interface{})[2].([]interface{})[0].(interface{}).(float64))] = mi.Info.([]interface{})[1].(string)
		case "SUPER_CHAT_MESSAGE":
			// SC
			sci := new(SuperChatInfo)
			json.Unmarshal(message[16:], &sci)
			log.Printf("[SC] %d元 %s: %s %d", sci.Data.Price, sci.Data.UserInfo.Uname, sci.Data.Message, sci.Data.ID)
		case "SUPER_CHAT_MESSAGE_JPN":

		case "SUPER_CHAT_MESSAGE_DELETE":
			// SC被删除
			log.Printf("%s", string(message[16:]))

		case "SEND_GIFT":
			// 礼物
			// log.Printf("%s", string(message[16:]))
			gf := new(GiftInfo)
			json.Unmarshal(message[16:], &gf)
			log.Printf("[礼物] %s 赠送 %d个 %s %.1f元", gf.Data.Uname, gf.Data.Num, gf.Data.GiftName, float64(gf.Data.Price)/1000*float64(gf.Data.Num))
		case "COMBO_SEND":
			// 连击礼物
		case "GUARD_BUY":
			// 大航海
		case "USER_TOAST_MSG":
			// 大航海
			ci := new(CrewInfo)
			json.Unmarshal(message[16:], &ci)
			// log.Printf("[大航海] %s 开通 %s * %d%s %d元", ci.Data.Username, ci.Data.RoleName, ci.Data.Num, ci.Data.Unit, ci.Data.Price/1000)
		case "ONLINE_RANK_V2":

		case "ONLINE_RANK_TOP3":

		case "INTERACT_WORD":

		case "ENTRY_EFFECT":

		case "ROOM_REAL_TIME_MESSAGE_UPDATE":

		case "ONLINE_RANK_COUNT":

		case "HOT_RANK_CHANGED_V2":

		case "LIVE":
			log.Printf("开播了")
			// fmt.Println(string(message[16:]))
		case "PREPARING":
			log.Printf("已下播")
			// fmt.Println(string(message[16:]))
		case "ROOM_CHANGE":

		case "WATCHED_CHANGE":

		case "STOP_LIVE_ROOM_LIST":

		case "HOT_ROOM_NOTIFY":

		case "HOT_RANK_CHANGED":

		case "HOT_RANK_SETTLEMENT":

		case "HOT_RANK_SETTLEMENT_V2":

		case "LIVE_INTERACTIVE_GAME":

		case "VOICE_JOIN_LIST":
			// 连麦请求
			// {"cmd":"VOICE_JOIN_LIST","data":{"cmd":"","room_id":23606554,"category":1,"apply_count":6,"red_point":1,"refresh":1},"room_id":23606554}
		case "VOICE_JOIN_ROOM_COUNT_INFO":
			// 连麦消息
			// {"cmd":"VOICE_JOIN_ROOM_COUNT_INFO","data":{"cmd":"","room_id":23606554,"root_status":1,"room_status":1,"apply_count":5,"notify_count":0,"red_point":1},"room_id":23606554}
		case "ROOM_BLOCK_MSG":
			// 禁言个人
			// {"cmd":"ROOM_BLOCK_MSG","data":{"dmscore":30,"operator":2,"uid":1772442517,"uname":"晓小轩iAvA"},"uid":"1772442517","uname":"晓小轩iAvA"}
			bi := new(BlockInfo)
			json.Unmarshal(message[16:], &bi)
			log.Printf("[禁言] %s 被禁言", bi.Data.Uname)
			log.Printf("上一条弹幕是: %s", user_danmaku[bi.Data.UID])
			// log.Printf("%s", string(message[16:]))
		case "ROOM_SILENT_ON":
			// 开启禁言
		case "ROOM_SILENT_OFF":
			// 关闭禁言
		case "WIDGET_BANNER":

		case "COMMON_NOTICE_DANMAKU":

		case "ANCHOR_LOT_START":
			// 天选
			// {"cmd":"ANCHOR_LOT_START","data":{"asset_icon":"https://i0.hdslb.com/bfs/live/627ee2d9e71c682810e7dc4400d5ae2713442c02.png","award_image":"","award_name":"年度大会员","award_num":40,"cur_gift_num":0,"current_time":1653042849,"danmu":"哔哩哔哩 (゜-゜)つロ 干杯~","gift_id":31039,"gift_name":"牛哇牛哇","gift_num":1,"gift_price":100,"goaway_time":180,"goods_id":15,"id":2656528,"is_broadcast":1,"join_type":1,"lot_status":0,"max_time":600,"require_text":"关注主播","require_type":1,"require_value":0,"room_id":25059330,"send_gift_ensure":0,"show_panel":1,"start_dont_popup":0,"status":1,"time":599,"url":"https://live.bilibili.com/p/html/live-lottery/anchor-join.html?is_live_half_webview=1\u0026hybrid_biz=live-lottery-anchor\u0026hybrid_half_ui=1,5,100p,100p,000000,0,30,0,0,1;2,5,100p,100p,000000,0,30,0,0,1;3,5,100p,100p,000000,0,30,0,0,1;4,5,100p,100p,000000,0,30,0,0,1;5,5,100p,100p,000000,0,30,0,0,1;6,5,100p,100p,000000,0,30,0,0,1;7,5,100p,100p,000000,0,30,0,0,1;8,5,100p,100p,000000,0,30,0,0,1","web_url":"https://live.bilibili.com/p/html/live-lottery/anchor-join.html"}}
		case "ANCHOR_LOT_AWARD":
			// 中奖名单

		case "POPULARITY_RED_POCKET_NEW":
			// 新红包

		case "POPULARITY_RED_POCKET_START":
			// 红包开始

		case "POPULARITY_RED_POCKET_WINNER_LIST":
			// 红包中奖名单

		case "NOTICE_MSG":
			// 通知消息
			// if strings.Contains(string(message[16:]), "舰长") || strings.Contains(string(message[16:]), "提督") || strings.Contains(string(message[16:]), "总督") {
			// 	return
			// }
		default:

			log.Printf("%s", string(message[16:]))
		}
	}
	// var w1 sync.WaitGroup
	// w1.Add(1)
	// w1.Wait()
}
