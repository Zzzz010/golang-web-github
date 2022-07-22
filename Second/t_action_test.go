package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Action struct {
	Title string
	Name  string
}

func TemplateIf(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))

	t.ExecuteTemplate(w, "if.gohtml", Action{
		Title: "Template Action",
		Name:  "Foster",
	})
}

func TestTemplateIf(t *testing.T) {
	rq := httptest.NewRequest("GET", "localhost:8080", nil)
	rc := httptest.NewRecorder()

	TemplateIf(rc, rq)

	rs := rc.Result()
	body, _ := io.ReadAll(rs.Body)
	fmt.Println(string(body))
}

func TemplateAppeal(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/appeal.gohtml"))

	t.ExecuteTemplate(w, "appeal.gohtml", map[string]interface{}{
		"Title":      "Template Action",
		"NilaiAkhir": 30,
	})
}

func TestTemplateAppeal(t *testing.T) {
	rq := httptest.NewRequest("GET", "localhost:8080", nil)
	rc := httptest.NewRecorder()

	TemplateAppeal(rc, rq)

	rs := rc.Result()
	body, _ := io.ReadAll(rs.Body)
	fmt.Println(string(body))
}

func TemplateRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))

	t.ExecuteTemplate(w, "range.gohtml", map[string]interface{}{
		"Mobil": []string{
			"Toyota Kijang", "GTR R32", "Subaru Impreza",
		},
	})
}

func TestTemplateRange(t *testing.T) {
	rq := httptest.NewRequest("GET", "localhost:8080", nil)
	rc := httptest.NewRecorder()

	TemplateRange(rc, rq)

	rs := rc.Result()
	body, _ := io.ReadAll(rs.Body)
	fmt.Println(string(body))
}

func TemplateWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/with.gohtml"))

	t.ExecuteTemplate(w, "with.gohtml", map[string]interface{}{
		"Title": "Template Action",
		"Name":  "John Foster",
		"Address": map[string]interface{}{
			"Street": "Jl. Sirsak Blok 2 no 1",
			"City":   "Jombang",
		},
	})
}

func TestTemplateWith(t *testing.T) {
	rq := httptest.NewRequest("GET", "localhost:8080", nil)
	rc := httptest.NewRecorder()

	TemplateWith(rc, rq)

	rs := rc.Result()
	body, _ := io.ReadAll(rs.Body)
	fmt.Println(string(body))
}

//localhost:8080/templates/with.gohtml

func TestTemplateAction(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateWith),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
