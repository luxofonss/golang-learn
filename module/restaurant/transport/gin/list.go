package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"learn/common"
	"learn/component/appctx"
	restaurantbiz "learn/module/restaurant/biz"
	restaurantmodel "learn/module/restaurant/model"
	restaurantrepository "learn/module/restaurant/repository"
	restaurantstorage "learn/module/restaurant/storage"
	"learn/module/restaurant_like/repository"
	restaurantlikestorage "learn/module/restaurant_like/storage"
	"net/http"
)

func ListRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restaurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var pagingData common.Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.Fulfill()

		db := appCtx.GetMainDBConnection()
		listStore := restaurantstorage.NewSQLStore(db)
		likeStore := restaurantlikestorage.NewSQLStore(db)
		listRepo := restaurantrepository.NewListRestaurantRepo(listStore)
		likeRepo := restaurantlikerepository.NewLikeRestaurantRepo(likeStore)
		biz := restaurantbiz.NewListRestaurantBiz(listRepo, likeRepo)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &pagingData)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}
