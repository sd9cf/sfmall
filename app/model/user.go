package model

type UserApiSignUpReq struct {
	Username string `v:"required|length:6,16#用户名不能为空|账号长度应当在:min到:max之间"`
	Phone    string `v:"required|phone#手机号不能为空|请输入正确的手机号"`
	Password string `v:"required|password#请输入密码|密码长度应在6~18之间"`
}


type UserApiCheckUsernameReq struct {
	Username string `v:"required#用户名不能为空"`
}

type UserApiCheckPhoneReq struct {
	Phone string `v:"required#手机号不能为空"`
}

type UserServiceSignUpReq struct {
	Username string
	Phone    string
	Password string
}

type UserProfile struct {
	Username string
	Phone    string
	RealName string
	Balance  uint64
}