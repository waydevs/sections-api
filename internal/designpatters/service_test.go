package designpatters

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/waydevs/sections-api/internal/platform/repository"
)

type designPatternRepositoryMock struct{}

func (d designPatternRepositoryMock) GetByID(_ context.Context, id string) (repository.DesignPattern, error) {
	switch id {
	case "ok":
		return repository.DesignPattern{
			Title: "ok",
		}, nil

	case "not-found":
		return repository.DesignPattern{}, repository.ErrNotFound

	default:
		return repository.DesignPattern{}, errors.New("some-error")
	}
}

func (d designPatternRepositoryMock) Create(_ context.Context, designPattern repository.DesignPattern) (repository.DesignPattern, error) {
	switch designPattern.Title {
	case "ok":
		return designPattern, nil

	default:
		return repository.DesignPattern{}, errors.New("some-error")
	}
}

func (d designPatternRepositoryMock) Delete(_ context.Context, id string) error {
	switch id {
	case "ok":
		return nil

	default:
		return errors.New("some-error")
	}
}

func (d designPatternRepositoryMock) Update(_ context.Context, designPattern repository.DesignPattern) (repository.DesignPattern, error) {
	switch designPattern.Title {
	case "ok":
		return designPattern, nil

	default:
		return repository.DesignPattern{}, errors.New("some-error")
	}
}

func TestNewService(t *testing.T) {
	db := designPatternRepositoryMock{}
	service := NewService(db)

	require.NotNil(t, service)
}

func TestService_GetByID(t *testing.T) {
	tt := []struct {
		name             string
		id               string
		expectedResponse DesignPattern
		expectedError    error
	}{
		{
			name: "ok",
			id:   "ok",
			expectedResponse: DesignPattern{
				// Is an empty ObjectID
				ID:    "000000000000000000000000",
				Title: "ok",
			},
			expectedError: nil,
		},
		{
			name:             "error not found",
			id:               "not-found",
			expectedResponse: DesignPattern{},
			expectedError:    ErrDesignPatternNotFound,
		},
		{
			name:             "error",
			id:               "error",
			expectedResponse: DesignPattern{},
			expectedError:    ErrSomethingWentWrong,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			db := designPatternRepositoryMock{}
			service := NewService(db)

			response, err := service.GetByID(context.Background(), tc.id)

			require.Equal(t, tc.expectedResponse, response)
			require.Equal(t, tc.expectedError, err)
		})
	}
}

func TestService_Create(t *testing.T) {
	tt := []struct {
		name             string
		designPattern    DesignPattern
		expectedResponse DesignPattern
		expectedError    error
	}{
		{
			name: "ok",
			designPattern: DesignPattern{
				ID:    "638d568a507b6e07cd39de82",
				Title: "ok",
			},
			expectedResponse: DesignPattern{
				// Is an empty ObjectID
				ID:    "000000000000000000000000",
				Title: "ok",
			},
			expectedError: nil,
		},
		{
			name:             "error",
			designPattern:    DesignPattern{},
			expectedResponse: DesignPattern{},
			expectedError:    ErrSomethingWentWrong,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			db := designPatternRepositoryMock{}
			service := NewService(db)

			response, err := service.Create(context.Background(), tc.designPattern)

			require.Equal(t, tc.expectedResponse, response)
			require.Equal(t, tc.expectedError, err)
		})
	}
}

func TestService_Delete(t *testing.T) {
	tt := []struct {
		name          string
		id            string
		expectedError error
	}{
		{
			name:          "ok",
			id:            "ok",
			expectedError: nil,
		},
		{
			name:          "error",
			id:            "error",
			expectedError: ErrSomethingWentWrong,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			db := designPatternRepositoryMock{}
			service := NewService(db)

			err := service.Delete(context.Background(), tc.id)

			require.Equal(t, tc.expectedError, err)
		})
	}
}

func TestService_Update(t *testing.T) {
	tt := []struct {
		name             string
		designPattern    DesignPattern
		expectedResponse DesignPattern
		expectedError    error
	}{
		{
			name: "ok",
			designPattern: DesignPattern{
				ID:    "638d568a507b6e07cd39de82",
				Title: "ok",
			},
			expectedResponse: DesignPattern{
				ID:    "638d568a507b6e07cd39de82",
				Title: "ok",
			},
			expectedError: nil,
		},
		{
			name:             "error",
			designPattern:    DesignPattern{},
			expectedResponse: DesignPattern{},
			expectedError:    ErrSomethingWentWrong,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			db := designPatternRepositoryMock{}
			service := NewService(db)

			response, err := service.Update(context.Background(), tc.designPattern)

			require.Equal(t, tc.expectedResponse, response)
			require.Equal(t, tc.expectedError, err)
		})
	}
}
