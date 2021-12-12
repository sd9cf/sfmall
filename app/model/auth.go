package model

type AuthApiLoginReq struct {
	Phone    string `v:"required|phone#手机号不能为空|请输入正确的手机号"`
	Password string `v:"required|password:6,16#请输入密码|密码长度应在6~18之间"`
}

type AuthServiceLoginReq struct {
	Phone    string 
	Password string 
}