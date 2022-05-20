package bilistruct

type MessageInfo struct {
	Cmd  string      `json:"cmd"`
	Info interface{} `json:"info"`
}

type SuperChatInfo struct {
	Cmd  string `json:"cmd"`
	Data struct {
		BackgroundBottomColor string  `json:"background_bottom_color"`
		BackgroundColor       string  `json:"background_color"`
		BackgroundColorEnd    string  `json:"background_color_end"`
		BackgroundColorStart  string  `json:"background_color_start"`
		BackgroundIcon        string  `json:"background_icon"`
		BackgroundImage       string  `json:"background_image"`
		BackgroundPriceColor  string  `json:"background_price_color"`
		ColorPoint            float64 `json:"color_point"`
		Dmscore               int     `json:"dmscore"`
		EndTime               int     `json:"end_time"`
		Gift                  struct {
			GiftID   int    `json:"gift_id"`
			GiftName string `json:"gift_name"`
			Num      int    `json:"num"`
		} `json:"gift"`
		ID          int `json:"id"`
		IsRanked    int `json:"is_ranked"`
		IsSendAudit int `json:"is_send_audit"`
		MedalInfo   struct {
			AnchorRoomid     int    `json:"anchor_roomid"`
			AnchorUname      string `json:"anchor_uname"`
			GuardLevel       int    `json:"guard_level"`
			IconID           int    `json:"icon_id"`
			IsLighted        int    `json:"is_lighted"`
			MedalColor       string `json:"medal_color"`
			MedalColorBorder int    `json:"medal_color_border"`
			MedalColorEnd    int    `json:"medal_color_end"`
			MedalColorStart  int    `json:"medal_color_start"`
			MedalLevel       int    `json:"medal_level"`
			MedalName        string `json:"medal_name"`
			Special          string `json:"special"`
			TargetID         int    `json:"target_id"`
		} `json:"medal_info"`
		Message          string `json:"message"`
		MessageFontColor string `json:"message_font_color"`
		MessageTrans     string `json:"message_trans"`
		Price            int    `json:"price"`
		Rate             int    `json:"rate"`
		StartTime        int    `json:"start_time"`
		Time             int    `json:"time"`
		Token            string `json:"token"`
		TransMark        int    `json:"trans_mark"`
		Ts               int    `json:"ts"`
		UID              int    `json:"uid"`
		UserInfo         struct {
			Face       string `json:"face"`
			FaceFrame  string `json:"face_frame"`
			GuardLevel int    `json:"guard_level"`
			IsMainVip  int    `json:"is_main_vip"`
			IsSvip     int    `json:"is_svip"`
			IsVip      int    `json:"is_vip"`
			LevelColor string `json:"level_color"`
			Manager    int    `json:"manager"`
			NameColor  string `json:"name_color"`
			Title      string `json:"title"`
			Uname      string `json:"uname"`
			UserLevel  int    `json:"user_level"`
		} `json:"user_info"`
	} `json:"data"`
	Roomid int `json:"roomid"`
}

type CrewInfo struct {
	Cmd  string `json:"cmd"`
	Data struct {
		AnchorShow       bool   `json:"anchor_show"`
		Color            string `json:"color"`
		Dmscore          int    `json:"dmscore"`
		EffectID         int    `json:"effect_id"`
		EndTime          int    `json:"end_time"`
		GuardLevel       int    `json:"guard_level"`
		IsShow           int    `json:"is_show"`
		Num              int    `json:"num"`
		OpType           int    `json:"op_type"`
		PayflowID        string `json:"payflow_id"`
		Price            int    `json:"price"`
		RoleName         string `json:"role_name"`
		StartTime        int    `json:"start_time"`
		SvgaBlock        int    `json:"svga_block"`
		TargetGuardCount int    `json:"target_guard_count"`
		ToastMsg         string `json:"toast_msg"`
		UID              int    `json:"uid"`
		Unit             string `json:"unit"`
		UserShow         bool   `json:"user_show"`
		Username         string `json:"username"`
	} `json:"data"`
}

type BlockInfo struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Dmscore  int    `json:"dmscore"`
		Operator int    `json:"operator"`
		UID      int    `json:"uid"`
		Uname    string `json:"uname"`
	} `json:"data"`
	UID   string `json:"uid"`
	Uname string `json:"uname"`
}

type PopInfo struct {
	Cmd   string `json:"cmd"`
	Count int    `json:"count"`
}

type GiftInfo struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Action            string      `json:"action"`
		BatchComboID      string      `json:"batch_combo_id"`
		BatchComboSend    interface{} `json:"batch_combo_send"`
		BeatID            string      `json:"beatId"`
		BizSource         string      `json:"biz_source"`
		BlindGift         interface{} `json:"blind_gift"`
		BroadcastID       int         `json:"broadcast_id"`
		CoinType          string      `json:"coin_type"`
		ComboResourcesID  int         `json:"combo_resources_id"`
		ComboSend         interface{} `json:"combo_send"`
		ComboStayTime     int         `json:"combo_stay_time"`
		ComboTotalCoin    int         `json:"combo_total_coin"`
		CritProb          int         `json:"crit_prob"`
		Demarcation       int         `json:"demarcation"`
		DiscountPrice     int         `json:"discount_price"`
		Dmscore           int         `json:"dmscore"`
		Draw              int         `json:"draw"`
		Effect            int         `json:"effect"`
		EffectBlock       int         `json:"effect_block"`
		Face              string      `json:"face"`
		FloatScResourceID int         `json:"float_sc_resource_id"`
		GiftID            int         `json:"giftId"`
		GiftName          string      `json:"giftName"`
		GiftType          int         `json:"giftType"`
		Gold              int         `json:"gold"`
		GuardLevel        int         `json:"guard_level"`
		IsFirst           bool        `json:"is_first"`
		IsSpecialBatch    int         `json:"is_special_batch"`
		Magnification     int         `json:"magnification"`
		MedalInfo         struct {
			AnchorRoomid     int    `json:"anchor_roomid"`
			AnchorUname      string `json:"anchor_uname"`
			GuardLevel       int    `json:"guard_level"`
			IconID           int    `json:"icon_id"`
			IsLighted        int    `json:"is_lighted"`
			MedalColor       int    `json:"medal_color"`
			MedalColorBorder int    `json:"medal_color_border"`
			MedalColorEnd    int    `json:"medal_color_end"`
			MedalColorStart  int    `json:"medal_color_start"`
			MedalLevel       int    `json:"medal_level"`
			MedalName        string `json:"medal_name"`
			Special          string `json:"special"`
			TargetID         int    `json:"target_id"`
		} `json:"medal_info"`
		NameColor         string      `json:"name_color"`
		Num               int         `json:"num"`
		OriginalGiftName  string      `json:"original_gift_name"`
		Price             int         `json:"price"`
		Rcost             int         `json:"rcost"`
		Remain            int         `json:"remain"`
		Rnd               string      `json:"rnd"`
		SendMaster        interface{} `json:"send_master"`
		Silver            int         `json:"silver"`
		Super             int         `json:"super"`
		SuperBatchGiftNum int         `json:"super_batch_gift_num"`
		SuperGiftNum      int         `json:"super_gift_num"`
		SvgaBlock         int         `json:"svga_block"`
		TagImage          string      `json:"tag_image"`
		Tid               string      `json:"tid"`
		Timestamp         int         `json:"timestamp"`
		TopList           interface{} `json:"top_list"`
		TotalCoin         int         `json:"total_coin"`
		UID               int         `json:"uid"`
		Uname             string      `json:"uname"`
	} `json:"data"`
}
