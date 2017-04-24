package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"time"

	"github.com/koushik-shetty/number-algo-aggregator/constants"
	"github.com/koushik-shetty/number-algo-aggregator/logger"
	"github.com/koushik-shetty/number-algo-aggregator/models"
	"github.com/koushik-shetty/number-algo-aggregator/urlparser"
)

//Numbers is the handler for the /numbers api
func Numbers(log logger.Logger) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		up, err := urlparser.New(r.URL.String(), constants.QUERYPARAM)
		if err != nil {
			log.Errorf("error parsing the url: %v", err)
			// http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		numbers := models.Numbers(r.Context(), log, up)

		numbersResponse, err := json.Marshal(numbers)
		if err != nil {
			log.Errorf("Failed to marshal response: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(numbersResponse)
		fmt.Printf("Latency: %v\n", time.Now().Sub(start))
	})
}
