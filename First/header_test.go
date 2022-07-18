package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Request

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	fmt.Fprint(w, contentType)
}

func TestRequestHeaders(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

// Response

func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Di Sponsori Oleh", "Tokyo Metro desu")
	fmt.Fprint(w, "OK")
}

func TestResponseHeaders(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	sponsor := recorder.Header().Get("di sponsori oleh")
	fmt.Println(sponsor)
}
