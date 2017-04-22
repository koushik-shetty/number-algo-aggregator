package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/koushik-shetty/number-algo-aggregator/testhelper"
)

func TestNumbersShouldRespondWithEmptyNumbersIfQueryParametersAreAbsent(t *testing.T) {
	assert := testhelper.Assert{t}

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/numbers", nil)

	Numbers(w, r)

	responseBody, err := ioutil.ReadAll(w.Body)
	assert.NoError(err, "Expected response body to be readable", err)

	expectedReponse, err := json.Marshal(&NumbersResponse{})

	assert.NoError(err, "Expected marshal to succeed", err)
	assert.Equal(responseBody, expectedReponse, "The response is not correct")
}

func TestNumbersShouldRespondWithEmptyNumbersIfQueryParametersAreNotPresent(t *testing.T) {
	assert := testhelper.Assert{t}

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/numbers?u='randomvalue'", nil)

	Numbers(w, r)

	responseBody, err := ioutil.ReadAll(w.Body)
	assert.NoError(err, "Expected response body to be readable", err)

	expectedReponse, err := json.Marshal(&NumbersResponse{})

	assert.NoError(err, "Expected marshal to succeed", err)
	assert.Equal(responseBody, expectedReponse, "The response is not correct")
}

func TestNumbersShouldRespondWithEmptyNumbersIfQueryParametersAreNotValidURLs(t *testing.T) {
	assert := testhelper.Assert{t}

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/numbers?u='randomvalue'", nil)

	Numbers(w, r)

	responseBody, err := ioutil.ReadAll(w.Body)
	assert.NoError(err, "Expected response body to be readable", err)

	expectedReponse, err := json.Marshal(&NumbersResponse{})

	assert.NoError(err, "Expected marshal to succeed", err)
	assert.Equal(responseBody, expectedReponse, "The response is not correct")
}
