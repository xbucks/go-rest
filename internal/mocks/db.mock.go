package mocks

import (
	"context"
	"github.com/rameshsunkara/go-rest-api-example/internal/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	PingFunc       func() error
	CreateFunc     func(ctx context.Context, purchaseOrder interface{}) (*mongo.InsertOneResult, error)
	UpdateFunc     func(ctx context.Context, purchaseOrder interface{}) (int64, error)
	GetAllFunc     func(ctx context.Context) (interface{}, error)
	GetByIdFunc    func(ctx context.Context, id string) (interface{}, error)
	DeleteByIdFunc func(ctx context.Context, id string) (int64, error)
)

type MockMongoMgr struct{}

func (m *MockMongoMgr) Ping() error {
	return PingFunc()
}

func (m *MockMongoMgr) Database() (db.MongoDatabase, error) {
	return &MockMongoDataBase{}, nil
}

func (m *MockMongoMgr) Disconnect() error {
	return nil
}

type MockMongoDataBase struct{}

func (m *MockMongoDataBase) Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection {
	return nil
}

type MockDataService struct{}

func (m *MockDataService) Create(ctx context.Context, purchaseOrder interface{}) (*mongo.InsertOneResult, error) {
	return CreateFunc(ctx, purchaseOrder)
}

func (m *MockDataService) Update(ctx context.Context, purchaseOrder interface{}) (int64, error) {
	return UpdateFunc(ctx, purchaseOrder)
}

func (m *MockDataService) GetAll(ctx context.Context) (interface{}, error) {
	return GetAllFunc(ctx)
}

func (m *MockDataService) GetById(ctx context.Context, id string) (interface{}, error) {
	return GetByIdFunc(ctx, id)
}

func (m *MockDataService) DeleteById(ctx context.Context, id string) (int64, error) {
	return DeleteByIdFunc(ctx, id)
}
