package main

import (
	"github.com/platon-p/kpodz1/application"
	"github.com/platon-p/kpodz1/cmd"
	"github.com/platon-p/kpodz1/cmd/account"
	"github.com/platon-p/kpodz1/cmd/category"
	"github.com/platon-p/kpodz1/cmd/operation"
	"github.com/platon-p/kpodz1/domain"
	"github.com/platon-p/kpodz1/infra"
)

type Dic struct {
}

func (d *Dic) Create() cmd.Command {
	accRepo := infra.NewInMemoryAccountRepository([]domain.BankAccount{})
	operationRepo := infra.NewInMemoryOperationRepository([]domain.Operation{})
	categoryRepo := infra.NewInMemoryCategoryRepository([]domain.Category{})
	accountService := application.NewAccountsService(accRepo, operationRepo)

	operationService := application.NewOperationService(operationRepo, accRepo, categoryRepo)
	categoryService := application.NewCategoryService(categoryRepo)

	accountCmd := &cmd.GroupCmd{Commands: []cmd.NamedCommand{
		cmd.Named(&account.CreateAccountCmd{Service: accountService}, "Создать счёт"),
		cmd.Named(&account.GetAllAccountsCmd{Service: accountService}, "Показать все счета"),
		cmd.Named(&account.EditAccountNameCmd{Service: accountService}, "Изменить имя счёта"),
		cmd.Named(&account.DeleteAccountCmd{Service: accountService}, "Удалить счёт"),
	}}
	_ = accountCmd
	operationCmd := &cmd.GroupCmd{Commands: []cmd.NamedCommand{
		cmd.Named(&operation.GetAllOperationsCmd{Service: operationService}, "Показать все операции"),
		cmd.Named(&operation.CreateOperationCmd{Service: operationService}, "Создать операцию"),
		cmd.Named(&operation.EditOperationAmountCmd{Service: operationService}, "Изменить сумму операции"),
		cmd.Named(&operation.DeleteOperationCmd{Service: operationService}, "Удалить операцию"),
	}}
	_ = operationCmd
	categoryCmd := &cmd.GroupCmd{Commands: []cmd.NamedCommand{
		cmd.Named(&category.CreateCategoryCmd{Service: categoryService}, "Создать категорию"),
		cmd.Named(&category.GetAllCategoriesCmd{Service: categoryService}, "Показать все категории"),
		cmd.Named(&category.EditCategoryNameCmd{Service: categoryService}, "Изменить название категории"),
		cmd.Named(&category.DeleteCategoryCmd{Service: categoryService}, "Удалить категорию"),
	}}
	mainGroup := &cmd.GroupCmd{Commands: []cmd.NamedCommand{
		cmd.Named(cmd.Wrap(accountCmd), "Счета"),
		cmd.Named(cmd.Wrap(operationCmd), "Операции"),
		cmd.Named(cmd.Wrap(categoryCmd), "Категории"),
		cmd.Named(cmd.NewSimpleCmd(func() error {
			return nil
		}), "Выйти"),
	}}
	mainCmd := cmd.LoopCmd{Parent: mainGroup}
	return &mainCmd
}
