package golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHTML(w http.ResponseWriter, r *http.Request) {
	templateText := `<html><bod>{{.}}</bod></html>`
	t := template.Must(template.New("TESTING").Parse(templateText))
	// t, err := template.New("TESTING").Parse(templateText)
	// if err != nil {
	// 	panic(err)
	// }

	t.ExecuteTemplate(w, "TESTING", "Ini Cuma Buat Pembelajaran Golang")
}

func TestSimpleHTML(t *testing.T) {
	rt := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	rc := httptest.NewRecorder()

	SimpleHTML(rc, rt)

	rs := rc.Result()
	body, _ := io.ReadAll(rs.Body)
	fmt.Println(string(body))
}

func SimpleHTMLFile(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/testing1.gohtml"))

	t.ExecuteTemplate(w, "testing1.gohtml", "Testing kedua menggunakan html file")

}

func TestSimpleHTMLFile(t *testing.T) {
	rt := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	rc := httptest.NewRecorder()

	SimpleHTMLFile(rc, rt)

	rs := rc.Result()
	body, _ := io.ReadAll(rs.Body)
	fmt.Println(string(body))
}

func TemplateDirectory(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))

	t.ExecuteTemplate(w, "testing1.gohtml", "Percobaan ketiga HTML")
}

func TestSimpleHTMLDirectory(t *testing.T) {
	rt := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	rc := httptest.NewRecorder()

	TemplateDirectory(rc, rt)

	rs := rc.Result()
	body, _ := io.ReadAll(rs.Body)
	fmt.Println(string(body))
}

//go:embed templates/*.gohtml
var templates embed.FS

func TemplateEmbed(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))

	t.ExecuteTemplate(w, "testing1.gohtml", "Percobaan Keempat Embed")
}

func TestHTTPTemplateEmbed(t *testing.T) {
	rq := httptest.NewRequest("GET", "http://localhost:8080", nil)
	rc := httptest.NewRecorder()

	TemplateEmbed(rc, rq)

	rp := rc.Result()
	body, _ := io.ReadAll(rp.Body)
	fmt.Println(string(body))
}

func TestTemplateEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//localhost:8080/templates/*.gohtml
