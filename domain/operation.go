package domain

import "time"

type Operation struct {
	ID        int64
	UserName  string
	Type      string
	Coin      string
	Value     float64
	CreatedAt time.Time
	UpdatedAt time.Time
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

type OperationRepository interface {
	Insert(opt *Operation) error
}
