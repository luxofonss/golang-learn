package ginrestaurantlike

import (
	"github.com/gin-gonic/gin"
	"learn/common"
	"learn/component/appctx"
	reataurantlikebiz "learn/module/restaurant_like/biz"
	restaurantlikemodel "learn/module/restaurant_like/model"
	restaurantlikerepository "learn/module/restaurant_like/repository"
	restaurantlikestorage "learn/module/restaurant_like/storage"
	"net/http"
)

func LikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalId()),
			UserId:       requester.GetUserId(),
		}

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		repo := restaurantlikerepository.NewUserLikeRestaurantRepo(store)

		biz := reataurantlikebiz.NewUserLikeRestaurantBiz(repo)

		if err := biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
