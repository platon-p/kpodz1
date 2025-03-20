package main

import (
	"fmt"
	"github.com/platon-p/kpodz1/application"
	"github.com/platon-p/kpodz1/cmd/account"
	"github.com/platon-p/kpodz1/domain"
	"github.com/platon-p/kpodz1/infra"
)

func main() {
	accRepo := infra.NewInMemoryAccountRepository([]domain.BankAccount{})
	operationRepo := infra.NewInMemoryOperationRepository([]domain.Operation{})
	//categoryRepo := infra.NewInMemoryCategoryRepository([]domain.Category{})
	accountService := application.NewAccountsService(accRepo, operationRepo)

	//operationService := application.NewOperationService(operationRepo, accRepo, categoryRepo)

	accountCmd := &GroupCmd{Commands: []NamedCommand{
		Named(&account.CreateAccountCmd{Service: accountService}, "Создать счёт"),
		Named(&account.GetAllAccountsCmd{Service: accountService}, "Показать все счета"),
		Named(&account.EditAccountNameCmd{Service: accountService}, "Изменить имя счёта"),
		Named(&account.DeleteAccountCmd{Service: accountService}, "Удалить счёт"),
	}}
	cmd := &GroupCmd{Commands: []NamedCommand{
		Named(accountCmd, "Работа со счетами"),
	}}
	for {
		fmt.Println(cmd.Execute())
		fmt.Println()
	}
}

func Wrap(cmd Command) Command {
	return &StatusWrapperCmd{Parent: cmd}
}

type StatusWrapperCmd struct {
	Parent Command
}

func (s *StatusWrapperCmd) Execute() error {
	if err := s.Parent.Execute(); err != nil {
		fmt.Printf("Команда завершилась ошибкой: %s", err)
		return err
	}
	fmt.Println("Успешно")
	return nil
}

type Command interface {
	Execute() error
}
