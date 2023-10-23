package restaurantlikerepository

import "context"

type LikeRestaurantStore interface {
	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
}

type likeRestaurantRepo struct {
	store LikeRestaurantStore
}

func NewLikeRestaurantRepo(store LikeRestaurantStore) *likeRestaurantRepo {
	return &likeRestaurantRepo{store: store}
}

func (repo *likeRestaurantRepo) GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error) {
	return repo.store.GetRestaurantLikes(ctx, ids)
}
