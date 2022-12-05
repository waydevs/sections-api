package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	someId = "5f9f1c5b9b9b9b9b9b9b9b9b"
)

type databaseHelperMock struct {
}

func (d *databaseHelperMock) Collection(name string) CollectionHelper {
	return &collectionHelperMock{}
}

func (d *databaseHelperMock) Client() ClientHelper {
	return &clientHelperMock{}
}

type collectionHelperMock struct {
}

func (c *collectionHelperMock) FindOne(ctx context.Context, filter interface{}) SingleResultHelper {
	id := filter.(map[string]primitive.ObjectID)["_id"]
	switch id.Hex() {
	case "5f9f1c5b9b9b9b9b9b9b9b9b":
		return &singleResultHelperMock{
			designPattern: DesignPattern{
				Title: "Some Design Pattern",
			},
		}

	default:
		return &singleResultHelperMock{}

	}
}

func (c *collectionHelperMock) InsertOne(ctx context.Context, field interface{}) (interface{}, error) {
	id, _ := primitive.ObjectIDFromHex(someId)
	return id, nil
}

func (c *collectionHelperMock) DeleteOne(ctx context.Context, filter interface{}) (int64, error) {
	return 0, nil
}

func (c *collectionHelperMock) ReplaceOne(ctx context.Context, filter interface{}, update interface{}) (int64, error) {
	return 0, nil
}

type singleResultHelperMock struct {
	designPattern DesignPattern
}

func (s *singleResultHelperMock) Decode(v interface{}) error {
	designPattern := v.(*DesignPattern)
	*designPattern = s.designPattern

	return nil
}

type clientHelperMock struct {
}

func (c *clientHelperMock) Database(string) DatabaseHelper {
	return &databaseHelperMock{}
}

func (c *clientHelperMock) Connect() error {
	return nil
}

func (c *clientHelperMock) Close() error {
	return nil
}

// Generate mocks with errors

type databaseHelperErrorMock struct {
}

func (d *databaseHelperErrorMock) Collection(name string) CollectionHelper {
	return &collectionHelperErrorMock{}
}

func (d *databaseHelperErrorMock) Client() ClientHelper {
	return &clientHelperErrorMock{}
}

type collectionHelperErrorMock struct {
}

func (c *collectionHelperErrorMock) FindOne(ctx context.Context, filter interface{}) SingleResultHelper {
	return &singleResultHelperErrorMock{}
}

func (c *collectionHelperErrorMock) InsertOne(ctx context.Context, designPattern interface{}) (interface{}, error) {
	return nil, errors.New("some-error")
}

func (c *collectionHelperErrorMock) DeleteOne(ctx context.Context, filter interface{}) (int64, error) {
	return 0, errors.New("some-error")
}

func (c *collectionHelperErrorMock) ReplaceOne(ctx context.Context, filter interface{}, update interface{}) (int64, error) {
	return 0, errors.New("some-error")
}

type singleResultHelperErrorMock struct {
}

func (s *singleResultHelperErrorMock) Decode(v interface{}) error {
	return errors.New("some-error")
}

type clientHelperErrorMock struct {
}

func (c *clientHelperErrorMock) Database(string) DatabaseHelper {
	return &databaseHelperErrorMock{}
}

func (c *clientHelperErrorMock) Connect() error {
	return errors.New("some-error")
}

func (c *clientHelperErrorMock) Close() error {
	return errors.New("some-error")
}
