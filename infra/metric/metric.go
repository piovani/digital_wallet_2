package metric

import "time"

type CLI struct {
	Name       string
	StartedAt  time.Time
	FinishedAt time.Time
	Duration   float64
}

func NewCLI(name string) *CLI {
	return &CLI{
		Name: name,
	}
}

func (c *CLI) Started() {
	c.StartedAt = time.Now()
}

func (c *CLI) Finished() {
	c.FinishedAt = time.Now()
	c.Duration = time.Since(c.StartedAt).Seconds()
}

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

type UseCase interface {
	SaveCli(c *CLI) error
	SaveService(h *HTTP)
}
