package shared

// 必须与业务服务一致
type WsChatSign struct {
	UserID    string `json:"userId"`
	Timestamp int64  `json:"timestamp"`
	Expire    int64  `json:"expire"`
	Sign      string `json:"sign"`
}

const SignUserIDKey = "userID"
