package httptester

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
)

type Mocker struct {
	router  http.Handler
	testing assert.TestingT
}

type Header map[string]string

func New(router http.Handler, testing assert.TestingT) *Mocker {
	return &Mocker{
		router:  router,
		testing: testing,
	}
}

func (c *Mocker) Request(method string, path string, body []byte, header *Header) *Tester {
	reader := bytes.NewReader(body)
	req, _ := http.NewRequest(method, path, reader)

	if header != nil {
		for key, value := range *header {
			req.Header.Set(key, value)
		}
	}

	w := httptest.NewRecorder()
	c.router.ServeHTTP(w, req)

	return &Tester{
		responseRecorder: w,
		testing:          c.testing,
	}
}

func (c *Mocker) Head(path string, body []byte, header *Header) *Tester {
	return c.Request(http.MethodHead, path, body, header)
}

func (c *Mocker) Options(path string, body []byte, header *Header) *Tester {
	return c.Request(http.MethodOptions, path, body, header)
}

func (c *Mocker) Get(path string, body []byte, header *Header) *Tester {
	return c.Request(http.MethodGet, path, body, header)
}

func (c *Mocker) Put(path string, body []byte, header *Header) *Tester {
	return c.Request(http.MethodPut, path, body, header)
}

func (c *Mocker) Post(path string, body []byte, header *Header) *Tester {
	return c.Request(http.MethodPost, path, body, header)
}

func (c *Mocker) Delete(path string, body []byte, header *Header) *Tester {
	return c.Request(http.MethodDelete, path, body, header)
}

func (c *Mocker) Patch(path string, body []byte, header *Header) *Tester {
	return c.Request(http.MethodPatch, path, body, header)
}

func (c *Mocker) Trace(path string, body []byte, header *Header) *Tester {
	return c.Request(http.MethodTrace, path, body, header)
}
