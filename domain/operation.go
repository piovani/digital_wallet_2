package domain

import "time"

type Operation struct {
	ID        int64     `json:"id"`
	UserName  string    `json:"user_name"`
	Type      string    `json:"type"`
	Value     float64   `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
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

// type OperationRepository interface {
// 	GetByID(ID int64) (Operation, error)
// 	GetByUserName(name string) (Operation, error)
// 	Insert(Operation) error
// 	Update(ID int64, operation Operation) error
// 	Delete(ID int64)
// }
