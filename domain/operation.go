package domain

import "time"

type Operation struct {
	ID        int64     `gorm:"primary_key;autoIncrement"`
	UserName  string    `gorm:"type:varchar(255);column:user_name"`
	Type      string    `gorm:"type:varchar(255);column:type"`
	Coin      string    `gorm:"type:varchar(255);column:coin"`
	Value     float64   `gorm:"type:double;column:value"`
	CreatedAt time.Time `gorm:<-:create"type:timestamp;column:created_at"`
	UpdatedAt time.Time `gorm:<-:update"type:timestamp;column:updated_at"`
}

func NewOperation(name, t, coin string, value float64) *Operation {
	return &Operation{
		UserName:  name,
		Type:      t,
		Coin:      coin,
		Value:     value,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type OperationRepository interface {
	Insert(opt *Operation) error
}
