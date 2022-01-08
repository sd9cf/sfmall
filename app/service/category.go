package service  

import (
	"context"
	"sfmall/app/dao"
	"sfmall/app/model"
	"sfmall/app/myerror"
	"github.com/gogf/gf/frame/g"
)

var Category = categoryService{}

type categoryService struct{}

func (s *categoryService) GetCategory() ([]*model.Category, error) {
	categories, err := dao.Category.Ctx(context.TODO()).All()
	if err != nil {
		g.Log().Errorf("数据库查找category错误:%v", err)
		return nil, myerror.DATABASEERROR
	}
	var categorylist []*model.Category
	if err = categories.Structs(&categorylist); err != nil {
		g.Log().Errorf("数据%v映射错误", categories)
		return nil, myerror.MAPPINGERROR
	}
	return categorylist, nil
}
