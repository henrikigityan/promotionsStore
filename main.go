package main

import (
	"log"
	"net/http"
	"test/simpleApi/configuration"
	"test/simpleApi/dao"
	"test/simpleApi/router"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
var collection *mongo.Collection

func main() {	
	collection = dbconfig.GetCollection(dbconfig.DB, "promotion", "testing")
	var promotionDao dao.PromotionDao = dao.NewPromotionDao(collection)
	routers.PromotionDao = promotionDao
    router := gin.Default()
	router.POST("promotions/csv", routers.LoadCsv)
	router.GET("promotions/:id", routers.GetPromotion)
	log.Fatal(http.ListenAndServe(":8000", router))
}

