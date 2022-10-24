package products

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

var s = createServer()

func createServer() *gin.Engine {

	repo := NewRepository()
	service := NewService(repo)
	p := NewHandler(service)

	r := gin.Default()

	r.GET("/api/v1/products", p.GetProducts)

	return r
}
func createRequestTest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {

	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	q := req.URL.Query()
	q.Add("seller_id", "1")
	req.URL.RawQuery = q.Encode()

	return req, httptest.NewRecorder()
}

func TestHandler(t *testing.T) {

	req, rr := createRequestTest("GET", "/api/v1/products", "")

	s.ServeHTTP(rr, req)
	//var resp []Product
	assert.Equal(t, http.StatusOK, rr.Code)

}
