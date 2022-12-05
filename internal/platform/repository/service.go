package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	database   = "sections-db"
	ctxTimeout = 10 * time.Second
)

type DatabaseHelper interface {
	Collection(name string) CollectionHelper
	Client() ClientHelper
}

type CollectionHelper interface {
	FindOne(context.Context, interface{}) SingleResultHelper
	InsertOne(context.Context, interface{}) (interface{}, error)
	DeleteOne(ctx context.Context, filter interface{}) (int64, error)
	ReplaceOne(ctx context.Context, filter interface{}, update interface{}) (int64, error)
}

type SingleResultHelper interface {
	Decode(v interface{}) error
}

type ClientHelper interface {
	Database(string) DatabaseHelper
	Connect() error
	Close() error
}

type mongoClient struct {
	cl *mongo.Client
}

type mongoDatabase struct {
	db *mongo.Database
}

type mongoCollection struct {
	coll *mongo.Collection
}

type mongoSingleResult struct {
	sr *mongo.SingleResult
}

// NewClient returns a new mongo client
func NewClient(mongoDBURI string) (ClientHelper, error) {
	c, err := mongo.NewClient(options.Client().ApplyURI(mongoDBURI))

	return &mongoClient{cl: c}, err
}

// NewDatabase returns a new mongo database
func NewDatabase(client ClientHelper) DatabaseHelper {
	client.Connect()
	return client.Database(database)
}

func (mc *mongoClient) Database(dbName string) DatabaseHelper {
	db := mc.cl.Database(dbName)
	return &mongoDatabase{db: db}
}

func (mc *mongoClient) Close() error {
	return mc.cl.Disconnect(context.Background())
}

func (mc *mongoClient) Connect() error {
	return mc.cl.Connect(context.Background())
}

func (md *mongoDatabase) Collection(colName string) CollectionHelper {
	collection := md.db.Collection(colName)
	return &mongoCollection{coll: collection}
}

func (md *mongoDatabase) Client() ClientHelper {
	client := md.db.Client()
	return &mongoClient{cl: client}
}

func (mc *mongoCollection) FindOne(ctx context.Context, filter interface{}) SingleResultHelper {
	singleResult := mc.coll.FindOne(ctx, filter)
	return &mongoSingleResult{sr: singleResult}
}

func (mc *mongoCollection) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	id, err := mc.coll.InsertOne(ctx, document)
	return id.InsertedID, err
}

func (mc *mongoCollection) DeleteOne(ctx context.Context, filter interface{}) (int64, error) {
	count, err := mc.coll.DeleteOne(ctx, filter)
	return count.DeletedCount, err
}

func (mc *mongoCollection) ReplaceOne(ctx context.Context, filter interface{}, update interface{}) (int64, error) {
	count, err := mc.coll.ReplaceOne(ctx, filter, update)
	return count.ModifiedCount, err
}

func (sr *mongoSingleResult) Decode(v interface{}) error {
	return sr.sr.Decode(v)
}
