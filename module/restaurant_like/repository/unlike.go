package restaurantlikerepository

import "context"

type UserUnlikeRestaurantStorage interface {
	UnLikeRestaurant(ctx context.Context, userId int, restaurantId int) error
}

type userUnlikeRestaurantRepo struct {
	store UserUnlikeRestaurantStorage
}

func NewUserUnlikeRestaurantRepo(store UserUnlikeRestaurantStorage) *userUnlikeRestaurantRepo {
	return &userUnlikeRestaurantRepo{store: store}
}

func (repo *userUnlikeRestaurantRepo) UnlikeRestaurant(ctx context.Context, userId int, restaurantId int) error {
	return repo.store.UnLikeRestaurant(ctx, userId, restaurantId)
}
