package restaurantrepository

import (
	"context"
	"learn/common"
	restaurantmodel "learn/module/restaurant/model"
)

type ListRestaurantStore interface {
	ListDataWithCondition(context context.Context, filter *restaurantmodel.Filter, paging *common.Paging, moreKeys ...string) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantRepo struct {
	store ListRestaurantStore
}

func NewListRestaurantRepo(store ListRestaurantStore) *listRestaurantRepo {
	return &listRestaurantRepo{store: store}
}

func (repo *listRestaurantRepo) ListRestaurant(context context.Context, filter *restaurantmodel.Filter, paging *common.Paging, moreKey ...string) ([]restaurantmodel.Restaurant, error) {
	result, err := repo.store.ListDataWithCondition(context, filter, paging, moreKey...)

	if err != nil {
		return nil, err
	}

	return result, nil
}
