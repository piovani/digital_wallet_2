package metric

import (
	"github.com/piovani/digital_wallet_2/infra/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

type Metric struct {
	pHistogram           *prometheus.HistogramVec
	httpRequestHistogram *prometheus.HistogramVec
}

func NewMetric() (*Metric, error) {
	cli := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "pushgateway",
		Name:      "cmd_duration_seconds",
		Help:      "CLI application execution in seconds",
		Buckets:   prometheus.DefBuckets,
	}, []string{"name"})

	http := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "http",
		Name:      "request_duration_seconds",
		Help:      "The latency of the HTTP requests.",
		Buckets:   prometheus.DefBuckets,
	}, []string{"handler", "method", "code"})

	m := &Metric{
		pHistogram:           cli,
		httpRequestHistogram: http,
	}

	err := prometheus.Register(m.pHistogram)
	if err != nil && err.Error() != "duplicate metrics collector registration attempted" {
		return nil, err
	}

	err = prometheus.Register(m.httpRequestHistogram)
	if err != nil && err.Error() != "duplicate metrics collector registration attempted" {
		return nil, err
	}

	return m, nil
}

func (m *Metric) SaveCMD(c *CMD) error {
	gatewayURL := config.Env.Prometheuspushgateway

	m.pHistogram.WithLabelValues(c.Name).Observe(c.Duration)

	return push.New(gatewayURL, "cmd_job").Collector(m.pHistogram).Push()
}

func (m *Metric) SaveHTTP(h *HTTP) {
	m.httpRequestHistogram.WithLabelValues(h.Handler, h.Method, h.StatusCode).Observe(h.Duration)
}

type UseCase interface {
	SaveCli(c *CMD) error
	SaveService(h *HTTP)
}
