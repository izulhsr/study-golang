package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recoder := httptest.NewRecorder()

	HelloHandler(recoder, request)

	resonse := recoder.Result()
	body, _ := io.ReadAll(resonse.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}
