package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"fmt"

	"github.com/koushik-shetty/number-algo-aggregator/models"
	"github.com/koushik-shetty/number-algo-aggregator/testhelper"
)

func TestNumbersShouldRespondWithEmptyNumbersIfQueryParametersAreAbsent(t *testing.T) {
	assert := testhelper.Assert{t}

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://localhost:0/numbers", nil)

	sl := &testhelper.StubLogger{}

	Numbers(sl)(w, r)

	responseBody, err := ioutil.ReadAll(w.Body)
	assert.NoError(err, "Expected response body to be readable", err)

	expectedReponse, err := json.Marshal(&models.NumbersResponse{
		Numbers: []int{},
	})

	assert.NoError(err, "Expected marshal to succeed", err)
	assert.Equal(responseBody, expectedReponse, "The response is not correct")
}

func TestNumbersShouldRespondWithEmptyNumbersIfQueryParametersAreInvalid(t *testing.T) {
	assert := testhelper.Assert{t}

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://localhost:0/numbers?r='randomvalue'", nil)

	sl := &testhelper.StubLogger{}
	Numbers(sl)(w, r)

	responseBody, err := ioutil.ReadAll(w.Body)
	assert.NoError(err, "Expected response body to be readable", err)

	// assert.Equal(sl.Called, true, "expected the logger to be called on error")

	expectedReponse, err := json.Marshal(&models.NumbersResponse{
		Numbers: []int{},
	})

	assert.NoError(err, "Expected marshal to succeed", err)
	assert.Equal(responseBody, expectedReponse, "The response is not correct")
}

func TestNumbersShouldRespondWithEmptyNumbersIfQueryParametersValuesAreNotValidURLs(t *testing.T) {
	assert := testhelper.Assert{t}

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://localhost:0/numbers?u='randomvalue'", nil)

	sl := &testhelper.StubLogger{}
	Numbers(sl)(w, r)

	responseBody, err := ioutil.ReadAll(w.Body)
	assert.NoError(err, "Expected response body to be readable", err)

	expectedReponse, err := json.Marshal(&models.NumbersResponse{
		Numbers: []int{},
	})

	assert.NoError(err, "Expected marshal to succeed", err)
	assert.Equal(responseBody, expectedReponse, "The response is not correct")
}

func TestNumbersShouldRespondWithNumbersIfQueryParametersValuesAreValidURLs(t *testing.T) {
	assert := testhelper.Assert{t}

	w := httptest.NewRecorder()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		numbers := &models.NumbersResponse{
			Numbers: []int{1, 1, 2, 3, 5, 8, 13, 21},
		}
		resp, err := json.Marshal(numbers)
		assert.NoError(err, "Expected no marshal error in test")
		w.Write(resp)
	}))
	defer server.Close()
	requestURL := fmt.Sprintf("http://localhost:0/numbers?u=%v/fibo", server.URL)
	r := httptest.NewRequest("GET", requestURL, nil)

	sl := &testhelper.StubLogger{}
	Numbers(sl)(w, r)

	responseBody, err := ioutil.ReadAll(w.Body)
	assert.NoError(err, "Expected response body to be readable", err)

	expectedReponse, err := json.Marshal(&models.NumbersResponse{
		Numbers: []int{1, 2, 3, 5, 8, 13, 21},
	})

	assert.NoError(err, "Expected marshal to succeed", err)
	assert.Equal(responseBody, expectedReponse, "The response is not correct")
}

func TestNumbersShouldRespondWithNumbersAggregatedFromMultipleValidURLs(t *testing.T) {
	assert := testhelper.Assert{t}

	w := httptest.NewRecorder()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var numbers *models.NumbersResponse
		if r.URL.String() == "/fibo" {
			numbers = &models.NumbersResponse{
				Numbers: []int{1, 1, 2, 3, 5, 8, 13, 21},
			}
		} else {
			numbers = &models.NumbersResponse{
				Numbers: []int{1, 3, 5, 7, 9, 11, 13},
			}
		}
		resp, err := json.Marshal(numbers)
		assert.NoError(err, "Expected no marshal error in test")
		w.Write(resp)
	}))
	defer server.Close()
	requestURL := fmt.Sprintf("http://localhost:0/numbers?u=%v/fibo&u=%v/odd", server.URL, server.URL)
	r := httptest.NewRequest("GET", requestURL, nil)

	sl := &testhelper.StubLogger{}
	Numbers(sl)(w, r)

	responseBody, err := ioutil.ReadAll(w.Body)
	assert.NoError(err, "Expected response body to be readable", err)

	expectedReponse, err := json.Marshal(&models.NumbersResponse{
		Numbers: []int{1, 2, 3, 5, 7, 8, 9, 11, 13, 21},
	})

	assert.NoError(err, "Expected marshal to succeed", err)
	assert.Equal(responseBody, expectedReponse, "The response is not correct")
}

func TestNumbersShouldRespondWithNumbersAggregatedFromMultipleValidURLsAndIgnoreInvalidAndUnavailableServices(t *testing.T) {
	assert := testhelper.Assert{t}

	w := httptest.NewRecorder()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var numbers *models.NumbersResponse
		if r.URL.String() == "/fibo" {
			numbers = &models.NumbersResponse{
				Numbers: []int{1, 1, 2, 3, 5, 8, 13, 21},
			}
			resp, err := json.Marshal(numbers)
			assert.NoError(err, "Expected no marshal error in test")
			w.Write(resp)
		} else {
			numbers = &models.NumbersResponse{
				Numbers: []int{1, 3, 5, 7, 9, 11, 13},
			}
			http.Error(w, "service unavailable", http.StatusServiceUnavailable)
		}
	}))
	defer server.Close()

	requestURL := fmt.Sprintf("http://localhost:0/numbers?u=%v/fibo&u=%v/abc", server.URL, server.URL)
	r := httptest.NewRequest("GET", requestURL, nil)

	sl := &testhelper.StubLogger{}
	Numbers(sl)(w, r)

	responseBody, err := ioutil.ReadAll(w.Body)
	assert.NoError(err, "Expected response body to be readable", err)

	expectedReponse, err := json.Marshal(&models.NumbersResponse{
		Numbers: []int{1, 2, 3, 5, 8, 13, 21},
	})

	assert.NoError(err, "Expected marshal to succeed", err)
	assert.Equal(responseBody, expectedReponse, "The response is not correct")
}

func TestNumbersShouldRespondWithNumbersWithinTimeoutIgnoringSlowServices(t *testing.T) {
	assert := testhelper.Assert{t}

	w := httptest.NewRecorder()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var numbers *models.NumbersResponse
		if r.URL.String() == "/fibo" {
			numbers = &models.NumbersResponse{
				Numbers: []int{1, 1, 2, 3, 5, 8, 13, 21},
			}
			resp, err := json.Marshal(numbers)
			assert.NoError(err, "Expected no marshal error in test")
			w.Write(resp)
		} else if r.URL.String() == "/odd" {
			numbers = &models.NumbersResponse{
				Numbers: []int{1, 3, 5, 7, 9, 11, 13},
			}
			resp, err := json.Marshal(numbers)
			assert.NoError(err, "Expected no marshal error in test")
			time.Sleep(time.Millisecond * 1000)
			w.Write(resp)
		}
	}))
	defer server.Close()

	requestURL := fmt.Sprintf("http://localhost:0/numbers?u=%v/fibo&u=%v/odd", server.URL, server.URL)
	r := httptest.NewRequest("GET", requestURL, nil)

	start := time.Now()
	sl := &testhelper.StubLogger{}
	Numbers(sl)(w, r)
	elapsedTime := time.Now().Sub(start)
	if elapsedTime > (time.Millisecond * 500) {
		t.Error("expected the response within the timeout", elapsedTime)
	}

	responseBody, err := ioutil.ReadAll(w.Body)
	assert.NoError(err, "Expected response body to be readable", err)

	expectedReponse, err := json.Marshal(&models.NumbersResponse{
		Numbers: []int{1, 2, 3, 5, 8, 13, 21},
	})

	assert.NoError(err, "Expected marshal to succeed", err)
	assert.Equal(responseBody, expectedReponse, "The response is not correct")
}
