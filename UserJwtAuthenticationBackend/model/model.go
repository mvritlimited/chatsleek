package model

// User ..
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

// ResponseResult ..
type ResponseResult struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}
