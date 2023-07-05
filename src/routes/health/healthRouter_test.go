package healthRouter

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine
func TestMain(m *testing.M) {
	router = gin.Default()
	Init(router)
    os.Exit(m.Run())
}

func fireGet(path string) *httptest.ResponseRecorder{
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", path, nil)
	router.ServeHTTP(w, req)
	return w
}

func TestPingRoute(t *testing.T) {
	w := fireGet("/ping")
	assert.Equal(t, 200, w.Code)
}

func TestInfoRoute(t *testing.T) {
	w := fireGet("/info")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"state\":\"UP\",\"version\":\"master\"}", w.Body.String())
}

func TestHealthRoute(t *testing.T) {
	w := fireGet("/health")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"status\":\"OK\"}", w.Body.String())
}