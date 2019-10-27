### HTTP testing util for Golang

A library for testing the http interface

```golang
package httptester_test

import (
	"github.com/axetroy/go-httptester"
	"net/http"
	"testing"
)

type serverHelloWorld struct{}

func (h *serverHelloWorld) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Foo", "BAR")
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write([]byte("Hello world"))
}

func TestNew(t *testing.T) {
	// test string response
	{
		h := &serverHelloWorld{}

		r := httptester.New(h, t)

		assert.Equal(t, nil, nil)

		r.Get("/hello_world", nil, nil).
			StatusCode(http.StatusBadRequest).
			Body([]byte("Hello world")).
			ContainHeader("X-Foo", "BAR")
	}
}
```