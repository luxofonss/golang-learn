package ginrestaurantlike

import (
	"github.com/gin-gonic/gin"
	"learn/common"
	"learn/component/appctx"
	reataurantlikebiz "learn/module/restaurant_like/biz"
	restaurantlikestorage "learn/module/restaurant_like/storage"
)

func UserDislikeRestaurant(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := restaurantlikestorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := reataurantlikebiz.NewUserUnlikeRestaurantBiz(store)
	}
}
