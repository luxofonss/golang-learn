package reataurantlikebiz

import "context"

type UserUnlikeRestaurantRepo interface {
	UnlikeRestaurant(ctx context.Context, userId int, restaurantId int) error
}

type userUnlikeRestaurantBiz struct {
	repo UserUnlikeRestaurantRepo
}

func NewUserUnlikeRestaurantBiz(repo UserUnlikeRestaurantRepo) *userUnlikeRestaurantBiz {
	return &userUnlikeRestaurantBiz{repo: repo}
}

func (repo *userUnlikeRestaurantBiz) UnlikeRestaurant(
	ctx context.Context,
	userId int,
	restaurantId int,
) error {
	err := repo.repo.UnlikeRestaurant(ctx, userId, restaurantId)

	if err != nil {
		return err
	}

	return nil
}
