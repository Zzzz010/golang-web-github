package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Single Query Parameter

func SayPerson(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "Hello.., Kamu siapa ?")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestQuery(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?name=raung", nil)
	recorder := httptest.NewRecorder()

	SayPerson(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(recorder.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}

// Multipe Query Parameter

func MultipleParameters(w http.ResponseWriter, r *http.Request) {
	firstname := r.URL.Query().Get("firstname")
	lastname := r.URL.Query().Get("lastname")
	fmt.Fprintf(w, "%s %s", firstname, lastname)
}

func TestMultipleQuery(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?firstname=Raung&lastname=Kawijayan", nil)
	recorder := httptest.NewRecorder()

	MultipleParameters(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(recorder.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}

// Multiple Query Parameter Value

func MultipleValues(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names := query["name"]
	fmt.Fprint(w, strings.Join(names, " "))
}

func TestMultipleValues(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?name=John&name=Foster", nil)
	recorder := httptest.NewRecorder()

	MultipleValues(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(recorder.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}
