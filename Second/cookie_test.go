package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "Test_Cookie"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(w, cookie)
	fmt.Fprint(w, "Sukses Membuat Cookie")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("Test_Cookie")
	if err != nil {
		fmt.Fprint(w, "No Cookie")
	} else {
		fmt.Fprintf(w, "Hello %s", cookie.Value)
	}
}

func TestCookieLocalhost(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func TestCookieOfflineSet(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/?name=Zero", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("%s : %s\n", cookie.Name, cookie.Value)
	}
}

func TestCookieOfflineGet(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/?name=Owen", nil)
	cookie := new(http.Cookie)
	cookie.Name = "Test_Cookie"
	cookie.Value = "Zero"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
