package metric

import "time"

type HTTP struct {
	Handler    string
	Method     string
	StatusCode string
	StartedAt  time.Time
	FinishedAt time.Time
	Duration   float64
}

func NewHTTP(handler, method string) *HTTP {
	return &HTTP{
		Handler: handler,
		Method:  method,
	}
}

func (h *HTTP) Started() {
	h.StartedAt = time.Now()
}

func (h *HTTP) Finished() {
	h.FinishedAt = time.Now()
	h.Duration = time.Since(h.StartedAt).Seconds()
}
