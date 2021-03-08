package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewPingHandler(t *testing.T) {

	req , _ := http.NewRequest("GET","/ping", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Ping)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"code":"200","data":"Ping Pong"}`
	assert.Equal(t, rr.Body.String(),expected)
}