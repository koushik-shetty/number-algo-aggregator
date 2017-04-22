package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

const (
	//TIMEOUT - Timeout for the responses.
	TIMEOUT = time.Millisecond * 500
)

//NumbersResponse represents the response to the /numbers endpoint
type NumbersResponse struct {
	Numbers []int `json:"numbers"`
}

//Numbers is the handler for the /numbers endpoint
func Numbers(w http.ResponseWriter, r *http.Request) {
	numbers := &NumbersResponse{}
	numbersResponse, err := json.Marshal(numbers)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	w.Write(numbersResponse)
}
