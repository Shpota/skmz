package cors

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDisable(t *testing.T) {
	writer := httptest.NewRecorder()
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "test")
	}
	handler = Disable(handler)

	handler(writer, nil)

	if got := writer.Header().Get("Access-Control-Allow-Origin"); got != "*" {
		t.Errorf("Expected Access-Control-Allow-Origin = *, got = %s", got)
	}
	if got := writer.Header().Get("Access-Control-Allow-Methods"); got != "*" {
		t.Errorf("Expected Access-Control-Allow-Methods = *, got = %s", got)
	}
	if got := writer.Header().Get("Access-Control-Allow-Headers"); got != "*" {
		t.Errorf("Expected Access-Control-Allow-Headers = *, got = %s", got)
	}
	if body, _ := ioutil.ReadAll(writer.Result().Body); string(body) != "test" {
		t.Errorf("Original handler has not been invoked")
	}
}
