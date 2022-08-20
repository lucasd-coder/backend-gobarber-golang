package middlewares

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	responseStatus = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "response_status",
			Help: "Status of HTTP response",
		},
		[]string{"status"},
	)

	httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_response_time_seconds",
		Help: "Duration of HTTP requests.",
	}, []string{"path"})

	totalRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of get requests.",
		},
		[]string{"path"},
	)
)

func init() {
	prometheus.Register(totalRequests)
	prometheus.Register(responseStatus)
	prometheus.Register(httpDuration)
}

func PrometheusHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.String()

		if strings.Contains(path, "/metrics") {
			c.Next()
			return
		}

		c.Next()
		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))

		responseStatus.WithLabelValues(strconv.Itoa(c.Writer.Status())).Inc()
		totalRequests.WithLabelValues(path).Inc()

		timer.ObserveDuration()
	}
}
