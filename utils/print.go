package utils

import (
	"fmt"
	"github.com/platon-p/kpodz1/domain"
	"time"
)

func CategoryTitle() string {
	return fmt.Sprintf("%-10s %-7s %s", "ID", "Type", "Name")
}

func PrettyCategory(category domain.Category) string {
	return fmt.Sprintf("%-10d %-7s %s", category.Id, category.CategoryType, category.Name)
}

func OperationTitle() string {
	return fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\t%s", "ID", "AccountID", "CategoryID", "Amount", "Type", "Date", "Description")
}

func PrettyOperation(op domain.Operation) string {
	return fmt.Sprintf(
		"%v\t%v\t%v\t%v\t%s\t%s\t%s",
		op.Id,
		op.BankAccountId,
		op.CategoryId,
		op.Amount,
		op.OperationType,
		op.Date.Format(time.DateTime),
		op.Description,
	)
}

func BankAccountTitle() string {
	return fmt.Sprintf("%-10s %s\t%s", "ID", "Name", "Balance")
}

func PrettyBankAccount(account domain.BankAccount) string {
	return fmt.Sprintf("%-10d %s\t%v", account.Id, account.Name, account.Balance)
}
