package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sfmall/app/dao"
	"sfmall/app/model"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/guuid"
)

// 中间件管理服务
var User = userService{}

type userService struct{}

func (s *userService) GetUser(req *model.AuthServiceLoginReq) (g.Map, error) {
	var user *model.User
	err := dao.User.Ctx(context.TODO()).Where("phone=? and password=?", req.Phone, req.Password).Scan(&user)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("手机号或密码错误")
	}
	return g.Map{
		"id":       user.Id,
		"username": user.Username,
		"roleID":   user.RoleId,
	}, nil
}

func (s *userService) SignUp(req *model.UserServiceSignUpReq) error {
	if !s.checkUsername(req.Username) {
		return fmt.Errorf("账号 %s 已经存在", req.Username)
	}
	if !s.checkPhone(req.Phone) {
		return fmt.Errorf("手机号 %s 已经注册", req.Phone)
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
	if _, err := dao.User.Ctx(context.TODO()).Save(user); err != nil {
		return err
	}
	return nil
}

func (s *userService) checkUsername(username string) bool {
	if i, err := dao.User.Ctx(context.TODO()).FindCount("username", username); err != nil {
		return false
	} else {
		return i == 0
	}
}

func (s *userService) checkPhone(phone string) bool {
	if i, err := dao.User.Ctx(context.TODO()).FindCount("phone", phone); err != nil {
		return false
	} else {
		return i == 0
	}
}

func (s *userService) GetProfile(id string) (*model.UserProfile, error) {
	var user *model.User
	var profile *model.UserProfile
	err := dao.User.Ctx(context.TODO()).WherePri(id).Scan(&user)
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


func (s *userService) AddBalance(id string, money uint) error {
	err := dao.User.Ctx(context.TODO()).Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) error {
		var user *model.User
		err := dao.User.Ctx(ctx).WherePri(id).Scan(&user)
		if err != nil {
			return err
		}
		user.Balance += uint64(money)
		
		_, err = dao.User.Ctx(ctx).WherePri(id).Update(user)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) BuyProduct(id string, req *model.BuyProduct) error {
	err := dao.User.Ctx(context.TODO()).Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) error {
		price := 0
		for _, productId := range req.ProductIds {
			var product *model.Product
			err := dao.Product.Ctx(ctx).WherePri(productId).Scan(&product)
			if err != nil {
				g.Log().Error(err)
				return err
			}
			if product == nil {
				return errors.New("不存在此商品")
			}
			if product.Amount < 1 {
				return errors.New("商品数量不足")
			}
			product.Amount -= 1
			product.Sales += 1
			price += int(product.Price)
		}
		
		var user *model.User
		err := dao.User.Ctx(ctx).WherePri(id).Scan(&user)
		if err != nil || user == nil {
			return errors.New("用户不存在")
		}
		if user.Balance < uint64(price) {
			return errors.New("余额不足")
		}
		user.Balance -= gconv.Uint64(price)
		var address *model.Address
		err = dao.Address.Ctx(ctx).WherePri(req.AddressId).Scan(&address)
		if err != nil {
			g.Log().Error(err)
			return err
		}
		if address == nil || address.UserId != id{
			return errors.New("地址不存在")
		}
		
		var order model.Order
		order.AddressId = int64(req.AddressId)
		var productItem []string 
		for _, value := range req.ProductIds {
			productItem = append(productItem, gconv.String(value))
		}
		order.ProductItem = strings.Join(productItem, ",")
		order.NickName = user.Username
		order.AddressId = int64(req.AddressId)
		order.Status = "配送中"
		order.TotalPrice = float64(price)
		order.UserId = id
		_, err = dao.Order.Ctx(context.TODO()).Save(order)
		if err != nil {
			g.Log().Error(err)
			return errors.New("下订单失败")
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
