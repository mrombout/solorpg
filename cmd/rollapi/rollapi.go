package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/mrombout/solorpg/backend/rollsvc"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	fieldKeys := []string{"method", "error"}
	rollCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "roll_service",
		Name:      "roll_count",
		Help:      " Number of requests to roll.",
	}, fieldKeys)
	rollLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "roll_service",
		Name:      "roll_latency",
		Help:      "The latency on rolls.",
	}, fieldKeys)

	var service rollsvc.RollService
	service = rollsvc.RollServiceImpl{}
	service = rollsvc.LoggingMiddleware{
		Logger: logger,
		Next:   service,
	}
	service = rollsvc.InstrumentingMiddleware{
		RollCount:   rollCount,
		RollLatency: rollLatency,
		Next:        service,
	}

	rollHandler := httptransport.NewServer(
		rollsvc.MakeRollEndpoint(service),
		rollsvc.DecodeRollRequest,
		rollsvc.EncodeResponse,
	)

	http.Handle("/roll", rollHandler)
	http.Handle("/metrics", promhttp.Handler())
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}
