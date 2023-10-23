package restaurantlikestorage

import (
	"context"
	"learn/common"
	restaurantlikemodel "learn/module/restaurant_like/model"
)

func (s *sqlStore) UnLikeRestaurant(ctx context.Context, userId int, restaurantId int) error {
	db := s.db

	if err := db.Where("user_id = ? AND restaurant_id = ?", userId, restaurantId).Delete(&restaurantlikemodel.Like{}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
