package healthRouter



import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)


func TestPingRoute(t *testing.T) {
	router := gin.Default()
	Init(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestInfoRoute(t *testing.T) {
	router := gin.Default()
	Init(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/info", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"state\":\"UP\",\"version\":\"master\"}", w.Body.String())

}

func TestHealthRoute(t *testing.T) {
	router := gin.Default()
	Init(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"status\":\"OK\"}", w.Body.String())

}