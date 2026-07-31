package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Projeto struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

var httpRequests = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Quantidade total de requisições HTTP.",
	},
)

func init() {
	prometheus.MustRegister(httpRequests)
}

func projetoHandler(w http.ResponseWriter, r *http.Request) {
	httpRequests.Inc()

	resposta := Projeto{
		Nome:    "Projeto Korp",
		Horario: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resposta)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {

	http.HandleFunc("/projeto-korp", projetoHandler)

	http.HandleFunc("/health", healthHandler)

	http.Handle("/metrics", promhttp.Handler())

	log.Println("Servidor iniciado na porta :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}