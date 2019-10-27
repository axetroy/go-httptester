package httptester_test

import (
	"encoding/json"
	"github.com/axetroy/go-httptester"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

type serverHelloWorld struct{}

func (h *serverHelloWorld) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Foo", "BAR")
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write([]byte("Hello world"))
}

type serverJSON struct{}

type addr struct {
	Country string `json:"country"`
	Code    int    `json:"code"`
}

type serverJSONResponse struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address addr
}

func (h *serverJSON) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Foo", "BAR")
	w.WriteHeader(http.StatusBadRequest)

	b, _ := json.Marshal(&serverJSONResponse{
		Name: "axetroy",
		Age:  18,
		Address: addr{
			Country: "China",
			Code:    10000,
		},
	})

	_, _ = w.Write(b)
}

type serverMultipleHeader struct{}

func (h *serverMultipleHeader) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Foo", "FOO")
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

func TestNewJSON(t *testing.T) {
	// test json response
	{
		h := &serverJSON{}

		r := httptester.New(h, t)

		assert.Equal(t, nil, nil)

		r.Get("/json", nil, nil).
			StatusCode(http.StatusBadRequest).
			BodyStruct(new(serverJSONResponse), &serverJSONResponse{
				Name: "axetroy",
				Age:  18,
				Address: addr{
					Country: "China",
					Code:    10000,
				},
			}).
			ContainHeader("X-Foo", "BAR")
	}
}

func TestNewMultipleHeader(t *testing.T) {
	{
		h := &serverMultipleHeader{}

		r := httptester.New(h, t)

		assert.Equal(t, nil, nil)

		r.Get("/header", nil, nil).
			StatusCode(http.StatusBadRequest).
			ContainHeader("X-Foo", "BAR").
			ContainHeader("X-Foo", "FOO")
	}
}
