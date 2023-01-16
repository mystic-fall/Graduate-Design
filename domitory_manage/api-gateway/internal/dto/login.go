package dto

// 直接和前端对接的结构体 需要和前端json字段统一

type UserLoginReq struct {
	Username string `json:"account"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

type UserLoginResp struct {
	Token string `json:"token"`
}
