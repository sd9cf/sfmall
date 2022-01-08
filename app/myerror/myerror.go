package myerror

import (
	"github.com/gogf/gf/errors/gerror"
)

var (
	AUTHFAILERROR = gerror.New("认证失败")
	DATABASEERROR = gerror.New("数据库内部错误")
	MAPPINGERROR  = gerror.New("数据映射错误")
)