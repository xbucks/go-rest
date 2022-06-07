package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rameshsunkara/go-rest-api-example/internal/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type MockMongoDBClient struct{}

var (
	pingFunc func(ctx context.Context, rp *readpref.ReadPref) error
	svcInfo  = &models.ServiceInfo{
		Name:        "test-api-service",
		Version:     "rams-fav",
		UpTime:      time.Now(),
		Environment: "test",
	}
	sc = NewStatusHandler(svcInfo, &MockMongoDBClient{})
)

func (m *MockMongoDBClient) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	return pingFunc(ctx, rp)
}

func UnMarshalStatusResponse(resp *http.Response) (StatusResponse, error) {
	body, _ := io.ReadAll(resp.Body)
	var statusResponse StatusResponse
	err := json.Unmarshal(body, &statusResponse)
	return statusResponse, err
}

func TestStatusSuccess(t *testing.T) {
	// Test Setup
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	pingFunc = func(ctx context.Context, rp *readpref.ReadPref) error {
		return nil
	}

	// Call actual function
	sc.CheckStatus(c)

	// Check results
	resp := w.Result()
	statusResponse, err := UnMarshalStatusResponse(resp)
	if err != nil {
		t.Fail()
	}
	assert.EqualValues(t, http.StatusOK, resp.StatusCode)
	assert.EqualValues(t, "test", statusResponse.Environment)
}

func TestStatusDown(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	pingFunc = func(ctx context.Context, rp *readpref.ReadPref) error {
		return errors.New("DB Connection Failed")
	}

	sc.CheckStatus(c)

	resp := w.Result()
	statusResponse, err := UnMarshalStatusResponse(resp)
	if err != nil {
		t.Fail()
	}

	assert.EqualValues(t, http.StatusFailedDependency, resp.StatusCode)
	assert.EqualValues(t, "rams-fav", statusResponse.Version)
}
