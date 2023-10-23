package restaurantlikestorage

import (
	"context"
	"learn/common"
)

func (s *sqlStore) Delete(ctx context.Context, userId, restaurantId int) error {
	db := s.db

	if err := db.Where("user_id = ? AND restaurant_id = ?", userId, restaurantId).
		Delete(nil).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
