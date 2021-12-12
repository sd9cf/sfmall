package service

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/guuid"
	"sfmall/app/dao"
	"sfmall/app/model"
)

// 中间件管理服务
var User = userService{}

type userService struct{}

func (s *userService) GetUser(req *model.AuthServiceLoginReq) (g.Map, error) {
	var user *model.User
	err :=  dao.User.Ctx(nil).Where("phone=? and password=?", req.Phone, req.Password).Scan(user)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("手机号或密码错误")
	}
	return g.Map{
		"id": user.Id,
		"username": user.Username,
		"roleID": user.RoleId,
	}, nil
}

func (s *userService) SignUp(req *model.UserServiceSignUpReq) error {
	if !s.checkUsername(req.Username) {
		return errors.New(fmt.Sprintf("账号 %s 已经存在", req.Username))
	}
	if !s.checkPhone(req.Phone) {
		return errors.New(fmt.Sprintf("手机号 %s 已经注册", req.Phone))
	}
	var user *model.User
	user.Username = req.Username
	user.Phone = req.Phone
	user.Password = req.Password
	user.Id = guuid.New().String()
	user.Status = 1
	user.RoleId = 1
	user.Balance = 0
	user.RealName = ""
	if _, err := dao.User.Ctx(nil).Save(user); err != nil {
		return err
	}
	return nil
}

func(s *userService) checkUsername(username string) bool {
	if i, err := dao.User.Ctx(nil).FindCount("username", username); err != nil {
		return false
	} else {
		return i == 0
	}
}

func(s *userService) checkPhone(phone string) bool {
	if i, err := dao.User.Ctx(nil).FindCount("phone", phone); err != nil {
		return false
	} else {
		return i == 0
	}
}

func(s *userService) GetProfile(id string) (*model.UserProfile, error) {
	var user *model.User
	var profile *model.UserProfile
	err :=  dao.User.Ctx(nil).Where("id=?", id).Scan(&user)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}
	if err := gconv.Struct(user, &profile); err != nil {
		return nil, err
	}
	return profile, nil
}