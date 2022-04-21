package serializer

// 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// 返回带token的Response
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}
