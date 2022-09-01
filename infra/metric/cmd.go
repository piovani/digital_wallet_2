package metric

import "time"

type CMD struct {
	Name       string
	StartedAt  time.Time
	FinishedAt time.Time
	Duration   float64
}

func NewCMD(name string) *CMD {
	return &CMD{
		Name: name,
	}
}

func (c *CMD) Started() {
	c.StartedAt = time.Now()
}

func (c *CMD) Finished() {
	c.FinishedAt = time.Now()
	c.Duration = time.Since(c.StartedAt).Seconds()
}
