package db

import (
	"context"
	"errors"
	"go-rest-api-example/internal/models"
	"go-rest-api-example/pkg/log"
	"go-rest-api-example/pkg/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type OrdersCrudService struct {
	Collection *mongo.Collection
}

func (orderCrudService *OrdersCrudService) Prepare(database string, collection string) {
	log.Logger.Debug("In Prepare")
	orderCrudService.Collection = GetDBClient().Database(database).Collection(collection)
}

func (orderCrudService *OrdersCrudService) CreateOrder(purchaseOrder *models.Order) (interface{}, error) {
	err := validate(orderCrudService.Collection)
	if err != nil {
		return nil, err
	}

	if !purchaseOrder.ID.IsZero() {
		return nil, errors.New("invalid request")
	}
	purchaseOrder.LastUpdatedAt = util.CurrentISOTime()

	result, err := orderCrudService.Collection.InsertOne(context.TODO(), purchaseOrder)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateOrder - Create and Update can be merged using upsert, but this is to demonstrate CRUD rest API so ...
func (orderCrudService *OrdersCrudService) UpdateOrder(purchaseOrder *models.Order) (int64, error) {
	err := validate(orderCrudService.Collection)
	if err != nil {
		return 0, err
	}

	if primitive.ObjectID.IsZero(purchaseOrder.ID) {
		return 0, errors.New("invalid request")
	}

	purchaseOrder.LastUpdatedAt = util.CurrentISOTime()

	opts := options.Update().SetUpsert(true)
	filter := bson.D{primitive.E{Key: "_id", Value: purchaseOrder.ID}}
	update := bson.D{primitive.E{Key: "$set", Value: purchaseOrder}}
	result, err := orderCrudService.Collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Logger.Error("Error occurred while updating order", zap.Error(err))
	}
	if result.MatchedCount != 0 {
		log.Logger.Info("matched and replaced an existing document", zap.Any("OrderID", result.UpsertedID))
		return result.MatchedCount, nil
	}
	if result.UpsertedCount != 0 {
		log.Logger.Warn("inserted a new order with ID:", zap.Any("OrderID", result.UpsertedID))
	}

	return result.MatchedCount, nil
}

func (orderCrudService *OrdersCrudService) GetAllOrders() ([]models.Order, error) {
	err := validate(orderCrudService.Collection)
	if err != nil {
		return nil, err
	}

	filter := bson.D{}

	cursor, err := orderCrudService.Collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []models.Order
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results, nil
}

func (orderCrudService *OrdersCrudService) GetOrderByID(id string) (interface{}, error) {
	err := validate(orderCrudService.Collection)
	if err != nil {
		return nil, err
	}

	isValidId := primitive.IsValidObjectID(id)
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil || !isValidId {
		return nil, errors.New("bad request")
	}
	filter := bson.D{primitive.E{Key: "_id", Value: docID}}

	var result models.Order
	error := orderCrudService.Collection.FindOne(context.TODO(), filter).Decode(&result)
	if error != nil {
		if error == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, error
	}

	return result, nil
}

func (orderCrudService *OrdersCrudService) DeleteOrderByID(id string) (int64, error) {
	err := validate(orderCrudService.Collection)
	if err != nil {
		return 0, err
	}

	isValidId := primitive.IsValidObjectID(id)
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil || !isValidId {
		return 0, errors.New("bad request")
	}
	filter := bson.D{primitive.E{Key: "_id", Value: docID}}

	res, error := orderCrudService.Collection.DeleteOne(context.TODO(), filter)
	if error != nil {
		return 0, error
	}

	return res.DeletedCount, nil
}

func validate(collection *mongo.Collection) error {
	if collection == nil {
		log.Logger.Error("Collection is not defined, Please override Prepare and Bind collection")
		return errors.New("collection is not defined")
	}
	return nil
}
