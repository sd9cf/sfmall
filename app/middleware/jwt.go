package middleware

import (
	jwt "github.com/gogf/gf-jwt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"net/http"
	"sfmall/app/model"
	"sfmall/app/service"
	"time"
)

type JwtConfig struct {
	ExpireTime    int    // 过期时间
	RefreshTime   int    // 刷新时间
	SignKey       string // 签名
	Realm         string // 域
	IdentityKey   string // 鉴权中用户唯一标识
	TokenLookup   string // token在请求中位置
	TokenHeadName string // token前缀
}

var (
	// The underlying JWT middleware.
	Auth *jwt.GfJWTMiddleware
	JwtCfg   *JwtConfig
)

// Initialization function,
// rewrite this function to customized your own JWT settings.
func init() {
	// 初始化配置
	JwtCfg = new(JwtConfig)
	if err := g.Cfg().GetStruct("jwt", JwtCfg); err != nil {
		g.Log().Errorf("JWT加载配置错误: %v", err)
	}
	authMiddleware, err := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           JwtCfg.Realm,
		Key:             []byte(JwtCfg.SignKey),
		Timeout:         time.Minute * time.Duration(JwtCfg.ExpireTime),
		MaxRefresh:      time.Minute * time.Duration(JwtCfg.RefreshTime),
		IdentityKey:     JwtCfg.IdentityKey,
		TokenLookup:     JwtCfg.TokenLookup,
		TokenHeadName:   JwtCfg.TokenHeadName,
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,
		LoginResponse:   LoginResponse,
		RefreshResponse: RefreshResponse,
		LogoutResponse:  LogoutResponse,
		Unauthorized:    Unauthorized,
		PayloadFunc:     PayloadFunc,
		IdentityHandler: IdentityHandler,
	})
	if err != nil {
		glog.Fatal("JWT Error:" + err.Error())
	}
	Auth = authMiddleware
}

// PayloadFunc is a callback function that will be called during login.
// Using this function it is possible to add additional payload data to the webtoken.
// The data is then made available during requests via c.Get("JWT_PAYLOAD").
// Note that the payload is not encrypted.
// The attributes mentioned on jwt.io can't be used as keys for the map.
// Optional, by default no additional data will be set.
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// IdentityHandler get the identity from JWT and set the identity for every request
// Using this function, by r.GetParam("id") get identity
func IdentityHandler(r *ghttp.Request) interface{} {
	claims := jwt.ExtractClaims(r)
	return claims[Auth.IdentityKey]
}

// Unauthorized is used to define customized Unauthorized callback function.
func Unauthorized(r *ghttp.Request, code int, message string) {
	r.Response.WriteJson(g.Map{
		"code": code,
		"msg":  message,
	})
	r.ExitAll()
}

// LoginResponse is used to define customized login-successful callback function.
func LoginResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	r.Response.WriteJson(g.Map{
		"code":   http.StatusOK,
		"token":  token,
		"expire": expire.Format(time.RFC3339),
	})
	r.ExitAll()
}

// RefreshResponse is used to get a new token no matter current token is expired or not.
func RefreshResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	r.Response.WriteJson(g.Map{
		"code":   http.StatusOK,
		"token":  token,
		"expire": expire.Format(time.RFC3339),
	})
	r.ExitAll()
}

// LogoutResponse is used to set token blacklist.
func LogoutResponse(r *ghttp.Request, code int) {
	r.Response.WriteJson(g.Map{
		"code":    code,
		"message": "success",
	})
	r.ExitAll()
}

// Authenticator is used to validate login parameters.
// It must return user data as user identifier, it will be stored in Claim Array.
// if your identityKey is 'id', your user data must have 'id'
// Check error (e) to determine the appropriate error message.
func Authenticator(r *ghttp.Request) (interface{}, error) {
	var (
		apiReq     *model.AuthApiLoginReq
		serviceReq *model.AuthServiceLoginReq
		err error
		user g.Map
	)
	if err = r.Parse(&apiReq); err != nil {
		return "", err
	}
	if err = gconv.Struct(apiReq, &serviceReq); err != nil {
		return "", err
	}
	if user, err = service.User.GetUser(serviceReq); err != nil {
		return nil, err
	}

	return user, nil
}

func MiddlewareAuth(r *ghttp.Request) {
	Auth.MiddlewareFunc()(r)
	r.Middleware.Next()
}