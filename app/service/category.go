package service  

import (
	"context"
	"sfmall/app/dao"
	"sfmall/app/model"
)

var Category = categoryService{}

type categoryService struct{}

func (s *categoryService) GetCategory() ([]*model.Category, error) {
	categories, err := dao.Category.Ctx(context.TODO()).All()
	if err != nil {
		g.Log().Error(err)
		return nil, err
	}
	var categorylist []*model.Category
	if err = categories.Structs(&categorylist); err != nil {
		g.Log().Error(err)
		return nil, err
	}
	return categorylist, nil
}
