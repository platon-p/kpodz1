package domain

import "time"

type OperationType string

const (
	IncomeOperatioType   OperationType = "income"
	OutcomeOperationType OperationType = "outcome"
)

type Operation struct {
	Id            uint32
	OperationType OperationType
	BankAccountId uint32
	Amount        float64
	Date          time.Time
	CategoryId    uint32
	Description   string
}
