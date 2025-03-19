package domain

import "time"

type OperationType string

const (
	IncomeOperatioType   OperationType = "Доход"
	OutcomeOperationType OperationType = "Расход"
)

type Operation struct {
	id            uint32
	operationType OperationType
	bankAccountId uint32
	amount        float64
	date          time.Time
	categoryId    uint32
	description   string
}
