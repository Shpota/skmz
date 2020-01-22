package cors

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDisable(t *testing.T) {
	var handler http.HandlerFunc
	w := httptest.NewRecorder()
	handler = func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "test")
	}
	corsHandler := Disable(handler)

	corsHandler(w, nil)

	if got := w.Header().Get("Access-Control-Allow-Origin"); got != "*" {
		t.Errorf("Expected Access-Control-Allow-Origin = *, got = %s", got)
	}
	if got := w.Header().Get("Access-Control-Allow-Methods"); got != "*" {
		t.Errorf("Expected Access-Control-Allow-Methods = *, got = %s", got)
	}
	if got := w.Header().Get("Access-Control-Allow-Headers"); got != "*" {
		t.Errorf("Expected Access-Control-Allow-Headers = *, got = %s", got)
	}
	if body, _ := ioutil.ReadAll(w.Result().Body); string(body) != "test" {
		t.Errorf("Original handler has not been invoked")
	}
}
