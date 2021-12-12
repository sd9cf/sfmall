package dao

import (
	"context"
	"sfmall/app/model"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	var ctx context.Context
	var userinfo model.User
	var userID  = "oUT385ZLmRr6R_a9xKSfSW9SekYI"
	User.Ctx(ctx).Fields(userinfo).Where(User.Columns.Id, userID).Scan(&userinfo)
	t.Log(userinfo.Phone)
}
