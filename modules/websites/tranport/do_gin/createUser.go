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

func Register(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data WebsiteModel.User
		if err := c.ShouldBind(&data); err != nil {
			c.String(http.StatusBadRequest, "cant register")
			return
		}
		store := WebsiteStorage.NewSQL(appCtx.GetMainDBConnection())
		biz := Websitebiz.NewCreateUserBiz(store)
		if err := biz.Register(c.Request.Context(), &data); err != nil {
			c.String(http.StatusBadRequest, "Data have already existed in db")
			return
		}
		c.JSON(http.StatusAccepted, common_function.DataResponseOnly(data))
	}
}
