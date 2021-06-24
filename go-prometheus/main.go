package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-prometheus/prom"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()

	engine.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "world")
	})

	engine.GET("/counter", func(c *gin.Context) {
		purl, _ := url.Parse(c.Request.RequestURI)
		prom.AccessCounter.With(prometheus.Labels{
			"method": c.Request.Method,
			"path":   purl.Path,
		}).Add(1)
	})

	engine.GET("/queue", func(c *gin.Context) {
		num := c.Query("num")
		numF, _ := strconv.ParseFloat(num, 32)
		prom.QueueGauge.With(prometheus.Labels{"name": "queue_gauge"}).Set(numF)
	})

	engine.GET("/histogram", func(c *gin.Context) {
		purl, _ := url.Parse(c.Request.RequestURI)
		prom.HttpDurationsHistogram.With(prometheus.Labels{"path": purl.Path}).Observe(float64(rand.Intn(30)))
	})

	engine.GET("/summary", func(c *gin.Context) {
		purl, _ := url.Parse(c.Request.RequestURI)
		prom.HttpDurations.With(prometheus.Labels{"path": purl.Path}).Observe(float64(rand.Intn(30)))
	})

	engine.GET("/metrics", gin.WrapH(promhttp.Handler()))

	_ = engine.Run(":8000")
}
