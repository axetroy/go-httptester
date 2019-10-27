package httptester

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
)

type Tester struct {
	responseRecorder *httptest.ResponseRecorder
	testing          assert.TestingT
	valid            bool
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (t *Tester) Status(status string) *Tester {
	t.setValid(assert.Equal(t.testing, status, t.responseRecorder.Result().Status))
	return t
}

func (t *Tester) StatusCode(statusCode int) *Tester {
	t.setValid(assert.Equal(t.testing, statusCode, t.responseRecorder.Code))
	return t
}

func (t *Tester) ResponseType(responseType string) *Tester {
	headers := t.responseRecorder.Header()

	if values, ok := headers["Content-Type"]; !ok {
		t.setValid(assert.Fail(t.testing, "expect contain header 'Content-Type' but got nil"))
	} else {
		t.setValid(assert.True(t.testing, contains(values, responseType), fmt.Sprintf("expect response type is '%s'", responseType)))
	}
	return t
}

func (t *Tester) ContentLength(length int) *Tester {
	t.setValid(assert.Equal(t.testing, length, t.responseRecorder.Result().ContentLength))
	return t
}

func (t *Tester) ContainHeader(key string, value string) *Tester {
	headers := t.responseRecorder.Header()

	if values, ok := headers[key]; !ok {
		t.setValid(assert.Failf(t.testing, "expect contain header '%s' but got nil", key))
		t.setValid(assert.Equal(t.testing, key, nil, fmt.Sprintf("expect contain header '%s' but got nil", key)))
	} else {
		t.setValid(assert.True(t.testing, contains(values, value), fmt.Sprintf("expect contain header with '%s=%s'", key, value)))
	}

	return t
}

func (t *Tester) Body(body []byte) *Tester {
	response, err := ioutil.ReadAll(t.responseRecorder.Body)

	t.setValid(assert.Nil(t.testing, err, "response error should be nil"))
	t.setValid(assert.Equal(t.testing, body, response))

	return t
}

func (t *Tester) BodyStruct(actuallyPointer interface{}, expectedPointer interface{}) *Tester {
	response, err := ioutil.ReadAll(t.responseRecorder.Body)

	t.setValid(assert.Nil(t.testing, err, "response error should be nil"))

	t.setValid(assert.Nil(t.testing, json.Unmarshal(response, actuallyPointer)))

	t.setValid(assert.Equal(t.testing, expectedPointer, actuallyPointer))
	return t
}

func (t *Tester) setValid(isValid bool) {
	if t.valid == false {
		return
	} else if isValid == false {
		t.valid = false
	}
}

func (t *Tester) IsValid() bool {
	return t.valid
}
