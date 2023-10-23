package restaurantbiz

import (
	"context"
	"learn/common"
	restaurantmodel "learn/module/restaurant/model"
)

type RestaurantListRepo interface {
	ListRestaurant(context context.Context, filter *restaurantmodel.Filter, paging *common.Paging, moreKeys ...string) ([]restaurantmodel.Restaurant, error)
}

type RestaurantLikeRepo interface {
	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
}

type listRestaurantBiz struct {
	listRepo RestaurantListRepo
	likeRepo RestaurantLikeRepo
}

func NewListRestaurantBiz(listRepo RestaurantListRepo, likeRepo RestaurantLikeRepo) *listRestaurantBiz {
	return &listRestaurantBiz{listRepo: listRepo, likeRepo: likeRepo}
}

func (biz *listRestaurantBiz) ListRestaurant(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.listRepo.ListRestaurant(context, filter, paging, "User")

	if err != nil {
		return nil, err
	}

	ids := make([]int, len(result))

	for i := range ids {
		ids[i] = result[i].Id
	}

	likeMap, err := biz.likeRepo.GetRestaurantLikes(context, ids)

	if err != nil {
		return result, nil
	}

	for i := range result {
		result[i].LikeCount = likeMap[result[i].Id]
	}

	return result, nil
}
