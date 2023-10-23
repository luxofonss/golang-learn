package restaurantlikerepository

import (
	"context"
	restaurantlikemodel "learn/module/restaurant_like/model"
)

type UserLikeRestaurantStorage interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

type userLikeRestaurantRepo struct {
	store UserLikeRestaurantStorage
}

func NewUserLikeRestaurantRepo(store UserLikeRestaurantStorage) *userLikeRestaurantRepo {
	return &userLikeRestaurantRepo{store: store}
}

func (repo *userLikeRestaurantRepo) LikeRestaurant(ctx context.Context, data *restaurantlikemodel.Like) error {
	return repo.store.Create(ctx, data)
}
