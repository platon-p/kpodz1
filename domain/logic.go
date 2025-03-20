package domain

import "io"

func PerformOperation(operation Operation) {
	// pull account and category
	// validate: category matches operation and account has sufficient funds
}

func CreateCategory(entity Category)          {}
func DeleteCategory(id uint32)                {}
func EditCategoryName(id uint32, name string) {}

func CreateAccount(account BankAccount)                {}
func EditAccountBalance(id uint32, newBalance float64) {}
func EditAccountName(id uint32, newName string)        {}
func DeleteAccount(id uint32)                          {}

func CreateOperation(operation Operation)              {}
func DeleteOperation(id uint32)                        {}
func EditOperationAmount(id uint32, newAmount float64) {}

func RecalculateAmount(accountId uint32) {}

type GlobalState struct {
	Accounts   []BankAccount
	Operations []Operation
	Categories []Category
}

func Dump(state GlobalState, writer io.Writer)
func Load(reader io.Reader) GlobalState
