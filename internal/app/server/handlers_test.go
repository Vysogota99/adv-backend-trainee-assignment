package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Vysogota99/adv-backend-trainee-assignment/internal/app/store/mock"
	"github.com/stretchr/testify/assert"
)

const (
	serverPort = ":8081"
)

func TestSaveStatHandler(t *testing.T) {
	type testCase struct {
		name   string
		params map[string]interface{}
		code   int
	}

	tCases := []testCase{
		testCase{
			name: "correct data",
			params: map[string]interface{}{
				"date":   "2006-10-21",
				"views":  1,
				"clicks": 19,
				"cost":   0.99,
			},
			code: 200,
		},
		testCase{
			name: "invalid date",
			params: map[string]interface{}{
				"date":   "2006-01-21",
				"views":  1,
				"clicks": 19,
				"cost":   0.99,
			},
			code: 400,
		},
		testCase{
			name: "invalid cost",
			params: map[string]interface{}{
				"date":   "2006-1-21",
				"views":  1,
				"clicks": 19,
				"cost":   -0.99,
			},
			code: 400,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(t *testing.T) {
			store := mock.New()
			r := NewRouter(serverPort, store)
			w := httptest.NewRecorder()
			body, err := json.Marshal(tc.params)
			assert.NoError(t, err)

			req, err := http.NewRequest("POST", "/stat", bytes.NewBuffer(body))
			assert.NoError(t, err)

			router, err := r.Setup()
			assert.NoError(t, err)

			router.ServeHTTP(w, req)
			assert.Equal(t, tc.code, w.Result().StatusCode)
		})
	}

}

func TestDeleteHandler(t *testing.T) {
	store := mock.New()
	r := NewRouter(serverPort, store)
	w := httptest.NewRecorder()

	req, err := http.NewRequest("DELETE", "/stat", nil)
	assert.NoError(t, err)

	router, err := r.Setup()
	assert.NoError(t, err)

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Result().StatusCode)
}

func TestGetStatus(t *testing.T) {
	type testCase struct {
		name   string
		params map[string]string
		code   int
	}

	tCases := []testCase{
		testCase{
			name: "correct data",
			params: map[string]string{
				"from": "2006-10-21",
				"to":   "2009-02-12",
			},
			code: 200,
		},
		testCase{
			name: "correct data",
			params: map[string]string{
				"from": "2010-10-21",
				"to":   "2009-02-12",
			},
			code: 400,
		},
		testCase{
			name: "correct data",
			params: map[string]string{
				"from": "asd-10-21",
				"to":   "2009-02-12",
			},
			code: 400,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(t *testing.T) {
			store := mock.New()
			r := NewRouter(serverPort, store)
			w := httptest.NewRecorder()

			url := fmt.Sprintf("/stat?from=%s&to=%s", tc.params["from"], tc.params["to"])
			req, err := http.NewRequest("GET", url, nil)
			assert.NoError(t, err)

			router, err := r.Setup()
			assert.NoError(t, err)

			router.ServeHTTP(w, req)
			assert.Equal(t, tc.code, w.Result().StatusCode)
		})
	}

}
