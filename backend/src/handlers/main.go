package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	muxHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	Articles   []Article
	appVersion string
	version    = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "version",
		Help: "Version information about this binary",
		ConstLabels: map[string]string{
			"version": appVersion,
		},
	})

	httpRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Count of all HTTP requests",
	}, []string{"code", "method"})

	httpSuccessTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "http_success_total",
		Help: "Count of all HTTP 200 requests",
	})

	httpRequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_request_duration_seconds",
		Help: "Duration of all HTTP requests",
	}, []string{"code", "handler", "method"})
)

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	httpSuccessTotal.Inc()
	json.NewEncoder(w).Encode(Articles)
}

func registerMetrics(r *prometheus.Registry) {
	r.MustRegister(httpRequestsTotal)
	r.MustRegister(httpSuccessTotal)
	r.MustRegister(httpRequestDuration)
	r.MustRegister(version)
}

func HandleRequests() {
	Articles = GetArticlesData()
	promRegistry := prometheus.NewRegistry()

	router := mux.NewRouter().StrictSlash(true)

	registerMetrics(promRegistry)

	cors := muxHandlers.CORS(
		muxHandlers.AllowedHeaders([]string{"content-type"}),
		muxHandlers.AllowedOrigins([]string{"*"}),
		muxHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		muxHandlers.AllowCredentials(),
	)

	router.Use(cors)

	router.HandleFunc("/articles", getAllArticles).Methods("GET")
	router.Handle("/metrics", promhttp.HandlerFor(promRegistry, promhttp.HandlerOpts{}))

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Listening on port 8000")
	log.Fatal(srv.ListenAndServe())
}
