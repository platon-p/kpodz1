package operation

import (
	"fmt"
	"github.com/platon-p/kpodz1/application"
)

type EditOperationAmountCmd struct {
	Service *application.OperationService
}

func (c *EditOperationAmountCmd) Execute() error {
	var id uint32
	fmt.Print("Введите ID операции: ")
	if _, err := fmt.Scan(&id); err != nil {
		return err
	}
	var newAmount float64
	fmt.Print("Введите новую сумму операции: ")
	if _, err := fmt.Scan(&newAmount); err != nil {
		return err
	}
	op, err := c.Service.EditAmount(id, newAmount)
	if err != nil {
		return err
	}
	fmt.Println(op)
	return nil
}
