package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/waydevs/sections-api/internal/designpatters"
)

type designPatternServiceMock struct{}

func (s *designPatternServiceMock) GetByID(ctx context.Context, id string) (designpatters.DesignPattern, error) {
	switch id {
	case "ok":
		return designpatters.DesignPattern{
			Title: "Design Pattern",
		}, nil

	case "not_found":
		return designpatters.DesignPattern{}, designpatters.ErrDesignPatternNotFound

	default:
		return designpatters.DesignPattern{}, errors.New("unexpected error")
	}
}

func (s *designPatternServiceMock) Create(ctx context.Context, designPattern designpatters.DesignPattern) (designpatters.DesignPattern, error) {
	switch designPattern.Title {
	case "ok":
		return designpatters.DesignPattern{
			Title: "Design Pattern",
		}, nil
	default:
		return designpatters.DesignPattern{}, errors.New("unexpected error")
	}
}

func (s *designPatternServiceMock) Delete(ctx context.Context, id string) error {
	switch id {
	case "ok":
		return nil
	case "not_found":
		return designpatters.ErrDesignPatternNotFound
	default:
		return errors.New("unexpected error")
	}
}

func (s *designPatternServiceMock) Update(ctx context.Context, designPattern designpatters.DesignPattern) (designpatters.DesignPattern, error) {
	switch designPattern.Title {
	case "ok":
		return designpatters.DesignPattern{
			Title: "Design Pattern",
		}, nil
	case "not_found":
		return designpatters.DesignPattern{}, designpatters.ErrDesignPatternNotFound
	default:
		return designpatters.DesignPattern{}, errors.New("unexpected error")
	}
}

func TestNewDesignPatternsHandler(t *testing.T) {
	service := &designPatternServiceMock{}
	handler := NewDesignPatternsHandler(service)

	assert.NotNil(t, handler)
}

func TestDesignPatternsHandler_GetPatternByID(t *testing.T) {
	tests := []struct {
		name             string
		id               string
		service          DesignPatternService
		expectedStatus   int
		expectedResponse string
	}{
		{
			name:             "Ok - Get Design Pattern by ID",
			id:               "ok",
			service:          &designPatternServiceMock{},
			expectedStatus:   200,
			expectedResponse: "{\"status\":200,\"message\":\"\",\"data\":{\"id\":\"\",\"title\":\"Design Pattern\",\"subtitle\":\"\",\"contentData\":null}}",
		},
		{
			name:             "Not Found - Get Design Pattern by ID",
			id:               "not_found",
			service:          &designPatternServiceMock{},
			expectedStatus:   404,
			expectedResponse: "{\"status\":404,\"message\":\"Design Pattern not found\",\"data\":null}",
		},
		{
			name:             "Internal Server Error - Get Design Pattern by ID",
			id:               "unexpected_error",
			service:          &designPatternServiceMock{},
			expectedStatus:   500,
			expectedResponse: "{\"status\":500,\"message\":\"unexpected error\",\"data\":null}",
		},
	}

	for _, tt := range tests {
		test := tt
		t.Run(tt.name, func(t *testing.T) {
			app := gin.Default()
			app = DesignPatternRoutes(app, tt.service)

			r, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/%s/%s", designPattersGroup, tt.id), nil)
			require.NoError(t, err)
			rr := httptest.NewRecorder()
			app.ServeHTTP(rr, r)

			resp := rr.Result()
			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err)

			require.Equal(t, test.expectedStatus, resp.StatusCode)
			require.Equal(t, tt.expectedResponse, string(body))

			err = resp.Body.Close()
			require.NoError(t, err)
		})
	}
}

func TestDesignPatternsHandler_CreatePattern(t *testing.T) {
	tests := []struct {
		name             string
		service          DesignPatternService
		bodyPost         designpatters.DesignPattern
		expectedStatus   int
		expectedResponse string
	}{
		{
			name:             "Ok - Create Design Pattern",
			service:          &designPatternServiceMock{},
			bodyPost:         designpatters.DesignPattern{Title: "ok"},
			expectedStatus:   201,
			expectedResponse: "{\"status\":201,\"message\":\"\",\"data\":{\"id\":\"\",\"title\":\"Design Pattern\",\"subtitle\":\"\",\"contentData\":null}}",
		},
		{
			name:             "Internal Server Error - Create Design Pattern",
			service:          &designPatternServiceMock{},
			bodyPost:         designpatters.DesignPattern{Title: "unexpected_error"},
			expectedStatus:   500,
			expectedResponse: "{\"status\":500,\"message\":\"unexpected error\",\"data\":null}",
		},
	}

	for _, tt := range tests {
		test := tt
		t.Run(tt.name, func(t *testing.T) {
			app := gin.Default()
			app = DesignPatternRoutes(app, tt.service)

			bodyPost, err := json.Marshal(tt.bodyPost)
			require.NoError(t, err)

			r, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/%s", designPattersGroup), bytes.NewReader(bodyPost))
			require.NoError(t, err)
			rr := httptest.NewRecorder()
			app.ServeHTTP(rr, r)

			resp := rr.Result()
			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err)

			require.Equal(t, test.expectedStatus, resp.StatusCode)
			require.Equal(t, tt.expectedResponse, string(body))

			err = resp.Body.Close()
			require.NoError(t, err)
		})
	}
}

func TestDesignPatternsHandler_DeletePattern(t *testing.T) {
	tests := []struct {
		name             string
		id               string
		service          DesignPatternService
		expectedStatus   int
		expectedResponse string
	}{
		{
			name:             "Ok - Delete Design Pattern",
			id:               "ok",
			service:          &designPatternServiceMock{},
			expectedStatus:   200,
			expectedResponse: "{\"status\":200,\"message\":\"Design pattern deleted successfully\",\"data\":null}",
		},
		{
			name:             "Not Found - Delete Design Pattern",
			id:               "not_found",
			service:          &designPatternServiceMock{},
			expectedStatus:   404,
			expectedResponse: "{\"status\":404,\"message\":\"Design Pattern not found\",\"data\":null}",
		},
		{
			name:             "Internal Server Error - Delete Design Pattern",
			id:               "unexpected_error",
			service:          &designPatternServiceMock{},
			expectedStatus:   500,
			expectedResponse: "{\"status\":500,\"message\":\"unexpected error\",\"data\":null}",
		},
	}

	for _, tt := range tests {
		test := tt
		t.Run(tt.name, func(t *testing.T) {
			app := gin.Default()
			app = DesignPatternRoutes(app, tt.service)

			r, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/%s/%s", designPattersGroup, tt.id), nil)
			require.NoError(t, err)
			rr := httptest.NewRecorder()
			app.ServeHTTP(rr, r)

			resp := rr.Result()
			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err)

			require.Equal(t, test.expectedStatus, resp.StatusCode)
			require.Equal(t, tt.expectedResponse, string(body))

			err = resp.Body.Close()
			require.NoError(t, err)
		})
	}
}

func TestDesignPatternsHandler_UpdatePattern(t *testing.T) {
	tests := []struct {
		name             string
		id               string
		service          DesignPatternService
		bodyPost         designpatters.DesignPattern
		expectedStatus   int
		expectedResponse string
	}{
		{
			name:             "Ok - Update Design Pattern",
			id:               "ok",
			service:          &designPatternServiceMock{},
			bodyPost:         designpatters.DesignPattern{Title: "ok"},
			expectedStatus:   200,
			expectedResponse: "{\"status\":200,\"message\":\"\",\"data\":{\"id\":\"\",\"title\":\"Design Pattern\",\"subtitle\":\"\",\"contentData\":null}}",
		},
		{
			name:             "Not Found - Update Design Pattern",
			id:               "not_found",
			service:          &designPatternServiceMock{},
			bodyPost:         designpatters.DesignPattern{Title: "not_found"},
			expectedStatus:   404,
			expectedResponse: "{\"status\":404,\"message\":\"Design Pattern not found\",\"data\":null}",
		},
		{
			name:             "Internal Server Error - Update Design Pattern",
			id:               "unexpected_error",
			service:          &designPatternServiceMock{},
			bodyPost:         designpatters.DesignPattern{Title: "unexpected_error"},
			expectedStatus:   500,
			expectedResponse: "{\"status\":500,\"message\":\"unexpected error\",\"data\":null}",
		},
	}

	for _, tt := range tests {
		test := tt
		t.Run(tt.name, func(t *testing.T) {
			app := gin.Default()
			app = DesignPatternRoutes(app, tt.service)

			bodyPost, err := json.Marshal(tt.bodyPost)
			require.NoError(t, err)

			r, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/%s", designPattersGroup), bytes.NewReader(bodyPost))
			require.NoError(t, err)
			rr := httptest.NewRecorder()
			app.ServeHTTP(rr, r)

			resp := rr.Result()
			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err)

			require.Equal(t, test.expectedStatus, resp.StatusCode)
			require.Equal(t, tt.expectedResponse, string(body))

			err = resp.Body.Close()
			require.NoError(t, err)
		})
	}
}
