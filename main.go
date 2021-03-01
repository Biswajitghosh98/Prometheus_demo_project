package main

import (
	"log"
	"net/http"

	"github.com/biswajitghosh98/demo_project/pkg"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//metrics to collect :
//1) no. of requests handled
//2) 200 rquest count
//3) non-200 request count
//4) no. of unique no. returned

var (
	NumRequestHandled = promauto.NewCounter(prometheus.CounterOpts{
		Name: "Num_request_handled_total",
		Help: "The total Number of processed requests",
	})
	Num200Requests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "Num_request_handled_total_200",
		Help: "The total Number of successful requests",
	})
	NumNon200Requests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "Num_request_handled_total_non_200",
		Help: "The total Number of unsuccessful requests",
	})
	NumUniqueRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "Num_unique_request_ID_handled",
		Help: "The total Number of unique ID requests",
	})
)

func main() {

	log.Printf("starting the server")
	srv := &pkg.Server{
		NumRequestHandled: NumRequestHandled,
		Num200Requests:    Num200Requests,
		NumNon200Requests: NumNon200Requests,
		NumUniqueRequests: NumUniqueRequests,
	}
	http.HandleFunc("/hello", srv.ReturnUniqueNumber)
	http.HandleFunc("/oddeven", srv.ReturnIfOdd)
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("some error happened : %v", err)
	}
	log.Printf("started listening at localhost:8080")
}
