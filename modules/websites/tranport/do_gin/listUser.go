package do_gin

import (
	"food_delivery/common_function"
	"food_delivery/component"
	"food_delivery/modules/websites/WebsiteModel"
	"food_delivery/modules/websites/WebsiteStorage"
	"food_delivery/modules/websites/Websitebiz"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetList(AppCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data WebsiteModel.Filter
		if err := c.ShouldBind(&data); err != nil {
			c.String(http.StatusBadRequest, "Get account list failed")
			return
		}
		store := WebsiteStorage.NewSQL(AppCtx.GetMainDBConnection())
		biz := Websitebiz.NewGetListBiz(store)

		result, err := biz.DoGetList(c.Request.Context(), &data)
		if err != nil {
			c.String(http.StatusBadRequest, "Get accountlist failed")
			return
		}
		c.JSON(http.StatusAccepted, common_function.DataResponseOnly(result))
	}
}
