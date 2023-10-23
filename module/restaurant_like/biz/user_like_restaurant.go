package reataurantlikebiz

import (
	"context"
	restaurantlikemodel "learn/module/restaurant_like/model"
)

type UserLikeRestaurantRepo interface {
	LikeRestaurant(ctx context.Context, data *restaurantlikemodel.Like) error
}

type userLikeRestaurantBiz struct {
	repo UserLikeRestaurantRepo
}

func NewUserLikeRestaurantBiz(repo UserLikeRestaurantRepo) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{repo: repo}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {
	err := biz.repo.LikeRestaurant(ctx, data)

	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	return nil
}
