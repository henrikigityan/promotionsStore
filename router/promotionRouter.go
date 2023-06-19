package routers

import (
	"test/simpleApi/dao"
	"test/simpleApi/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var PromotionDao dao.PromotionDao 

func LoadCsv(c *gin.Context) {
	PromotionDao.DeleteAllValuesFromCollection()
	service.HandleUpload(c.Request, PromotionDao)
}

func GetPromotion(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, PromotionDao.GetById(c.Param("id")))
}