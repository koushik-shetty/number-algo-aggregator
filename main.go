package main

import (
	"os"

	"net/http"

	"github.com/koushik-shetty/number-algo-aggregator/handlers"
	"github.com/koushik-shetty/number-algo-aggregator/logger"
)

func main() {
	log := logger.New(os.Stdout)

	r := http.NewServeMux()
	r.Handle("/numbers", handlers.Numbers(log))
	log.Fatalf("Connection broke: %v", http.ListenAndServe(":8080", r))
}
