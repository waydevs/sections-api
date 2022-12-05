package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestNewDesignPatterns(t *testing.T) {
	db := &databaseHelperMock{}
	designPatterns := NewDesignPatterns(db)

	assert.NotNil(t, designPatterns)
}

func TestDesignPatterns_GetByID(t *testing.T) {
	tt := []struct {
		name           string
		id             string
		database       DatabaseHelper
		expectedResult DesignPattern
		expectedError  error
	}{
		{
			name:     "Ok - GetByID",
			id:       "5f9f1c5b9b9b9b9b9b9b9b9b",
			database: &databaseHelperMock{},
			expectedResult: DesignPattern{
				Title: "Some Design Pattern",
			},
			expectedError: nil,
		},
		{
			name:           "Error - GetByID",
			id:             "5f9f1c5b9b9b9b9b9b9b9b9b",
			database:       &databaseHelperErrorMock{},
			expectedResult: DesignPattern{},
			expectedError:  errors.New("some-error"),
		},
		{
			name:           "Error - Erroneous ID",
			id:             "aaaa",
			database:       &databaseHelperMock{},
			expectedResult: DesignPattern{},
			expectedError:  errors.New("the provided hex string is not a valid ObjectID"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			designPatterns := NewDesignPatterns(tc.database)

			result, err := designPatterns.GetByID(context.Background(), tc.id)

			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestDesignPatterns_Create(t *testing.T) {
	id, _ := primitive.ObjectIDFromHex(someId)

	tt := []struct {
		name           string
		designPattern  DesignPattern
		database       DatabaseHelper
		expectedResult DesignPattern
		expectedError  error
	}{
		{
			name: "Ok - Create",
			designPattern: DesignPattern{
				Title: "Some Design Pattern",
			},
			database: &databaseHelperMock{},
			expectedResult: DesignPattern{
				MongoID: id,
				Title:   "Some Design Pattern",
			},
			expectedError: nil,
		},
		{
			name: "Error - Create",
			designPattern: DesignPattern{
				Title: "Some Design Pattern",
			},
			database:       &databaseHelperErrorMock{},
			expectedResult: DesignPattern{},
			expectedError:  errors.New("some-error"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			designPatterns := NewDesignPatterns(tc.database)

			result, err := designPatterns.Create(context.Background(), tc.designPattern)

			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestDesignPatterns_Update(t *testing.T) {
	tt := []struct {
		name           string
		designPattern  DesignPattern
		database       DatabaseHelper
		expectedResult DesignPattern
		expectedError  error
	}{
		{
			name: "Ok - Update",
			designPattern: DesignPattern{
				Title: "Some Design Pattern",
			},
			database: &databaseHelperMock{},
			expectedResult: DesignPattern{
				Title: "Some Design Pattern",
			},
			expectedError: nil,
		},
		{
			name: "Error - Update",
			designPattern: DesignPattern{
				Title: "Some Design Pattern",
			},
			database:       &databaseHelperErrorMock{},
			expectedResult: DesignPattern{},
			expectedError:  errors.New("some-error"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			designPatterns := NewDesignPatterns(tc.database)

			result, err := designPatterns.Update(context.Background(), tc.designPattern)

			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestDesignPatterns_Delete(t *testing.T) {
	tt := []struct {
		name          string
		id            string
		database      DatabaseHelper
		expectedError error
	}{
		{
			name:          "Ok - Delete",
			id:            "5f9f1c5b9b9b9b9b9b9b9b9b",
			database:      &databaseHelperMock{},
			expectedError: nil,
		},
		{
			name:          "Error - Delete",
			id:            "5f9f1c5b9b9b9b9b9b9b9b9b",
			database:      &databaseHelperErrorMock{},
			expectedError: errors.New("some-error"),
		},
		{
			name:          "Error - Erroneous ID",
			id:            "aaaa",
			database:      &databaseHelperMock{},
			expectedError: errors.New("the provided hex string is not a valid ObjectID"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			designPatterns := NewDesignPatterns(tc.database)

			err := designPatterns.Delete(context.Background(), tc.id)

			assert.Equal(t, tc.expectedError, err)
		})
	}
}
