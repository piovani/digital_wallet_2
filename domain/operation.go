package domain

import "time"

type Operation struct {
	ID        int64     `json:"id"`
	UserName  string    `json:"user_name"`
	Type      string    `json:"type"`
	Value     float64   `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewOperation(name, t string, value float64) *Operation {
	return &Operation{
		UserName:  name,
		Type:      t,
		Value:     value,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
