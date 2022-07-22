package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateDataMap(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/t_data.gohtml"))

	t.ExecuteTemplate(w, "t_data.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Raditya Bagus Putra",
		"Address": map[string]interface{}{
			"Street": "Not Yet",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	rq := httptest.NewRequest("GET", "http://localhost:8080", nil)
	rc := httptest.NewRecorder()

	TemplateDataMap(rc, rq)

	rs := rc.Result()
	body, _ := io.ReadAll(rs.Body)
	fmt.Println(string(body))
}

type Address struct {
	Street string
}

type Page struct {
	Title   string
	Name    string
	Address Address
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/t_data.gohtml"))

	t.ExecuteTemplate(w, "t_data.gohtml", Page{
		Title: "Template Data Struct",
		Name:  "John Foster",
		Address: Address{
			Street: "Not Yet",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateDataStruct),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
