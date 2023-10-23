package restaurantbiz

import (
	"context"
	"learn/common"
	restaurantmodel "learn/module/restaurant/model"
)

type RestaurantCreateRepo interface {
	CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBiz struct {
	repo RestaurantCreateRepo
}

func NewCreateRestaurantBiz(repo RestaurantCreateRepo) *createRestaurantBiz {
	return &createRestaurantBiz{repo: repo}
}

func (biz *createRestaurantBiz) CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	if err := biz.repo.CreateRestaurant(context, data); err != nil {
		return common.ErrCannotCreateEntity(restaurantmodel.EntityName, err)
	}

	return nil
}
