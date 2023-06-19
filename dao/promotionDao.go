package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"test/simpleApi/model"
)

type PromotionDao struct  {
	c *mongo.Collection 
}

func NewPromotionDao(collection * mongo.Collection) (PromotionDao) {
	return PromotionDao {	
		c: collection,
	}
}

func(dao *PromotionDao) GetById(id string) model.Promotion {
	result :=  dao.c.FindOne(context.TODO(), bson.M{"_id": id})
	var promotion model.Promotion
	result.Decode(&promotion)
	return promotion
}

func (dao *PromotionDao) BatchInsert(promotions []model.Promotion) []interface{} { 
	result, err := dao.c.InsertMany(context.TODO(), mapToInterface(promotions))
	if err != nil {
		return nil
	} 
	return result.InsertedIDs
}

func (dao *PromotionDao) DeleteAllValuesFromCollection() {
	dao.c.DeleteMany(context.TODO(), bson.D{})
}

func mapToInterface(promotions []model.Promotion) []interface{} {
	var result []interface{}
	for _, p := range promotions {
		s := bson.D{{"_id", p.Id},{"price", p.Price}, {"expirationDate", p.ExpirationDate}}
		result = append(result, s)
	}
	return result
}
