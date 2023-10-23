package restaurantrepository

import (
	"context"
	restaurantmodel "learn/module/restaurant/model"
)

type RestaurantStore interface {
	Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantRepo struct {
	store RestaurantStore
}

func NewRestaurantRepo(store RestaurantStore) *createRestaurantRepo {
	return &createRestaurantRepo{store: store}
}

func (repo *createRestaurantRepo) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := repo.store.Create(ctx, data); err != nil {
		return err
	}

	return nil
}
