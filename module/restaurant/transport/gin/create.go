package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"learn/common"
	"learn/component/appctx"
	restaurantbiz "learn/module/restaurant/biz"
	restaurantmodel "learn/module/restaurant/model"
	restaurantrepository "learn/module/restaurant/repository"
	restaurantstorage "learn/module/restaurant/storage"
	"net/http"
)

func CreateRestaurant(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		db := ctx.GetMainDBConnection()
		store := restaurantstorage.NewSQLStore(db)
		repo := restaurantrepository.NewRestaurantRepo(store)
		biz := restaurantbiz.NewCreateRestaurantBiz(repo)

		data.UserId = requester.GetUserId()

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
