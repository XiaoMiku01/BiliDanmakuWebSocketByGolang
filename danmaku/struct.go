package danmaku

type RoomInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Group            string  `json:"group"`
		BusinessID       int     `json:"business_id"`
		RefreshRowFactor float64 `json:"refresh_row_factor"`
		RefreshRate      int     `json:"refresh_rate"`
		MaxDelay         int     `json:"max_delay"`
		Token            string  `json:"token"`
		HostList         []struct {
			Host    string `json:"host"`
			Port    int    `json:"port"`
			WssPort int    `json:"wss_port"`
			WsPort  int    `json:"ws_port"`
		} `json:"host_list"`
	} `json:"data"`
}

type RealIdInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		RoomID          int           `json:"room_id"`
		ShortID         int           `json:"short_id"`
		UID             int           `json:"uid"`
		NeedP2P         int           `json:"need_p2p"`
		IsHidden        bool          `json:"is_hidden"`
		IsLocked        bool          `json:"is_locked"`
		IsPortrait      bool          `json:"is_portrait"`
		LiveStatus      int           `json:"live_status"`
		HiddenTill      int           `json:"hidden_till"`
		LockTill        int           `json:"lock_till"`
		Encrypted       bool          `json:"encrypted"`
		PwdVerified     bool          `json:"pwd_verified"`
		LiveTime        int           `json:"live_time"`
		RoomShield      int           `json:"room_shield"`
		IsSp            int           `json:"is_sp"`
		SpecialType     int           `json:"special_type"`
		PlayURL         interface{}   `json:"play_url"`
		AllSpecialTypes []interface{} `json:"all_special_types"`
	} `json:"data"`
}
