package operation

import (
	"fmt"
	"github.com/platon-p/kpodz1/services"
	"github.com/platon-p/kpodz1/utils"
)

type EditOperationAmountCmd struct {
	Service *services.OperationService
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
	fmt.Println("Операция изменена: ", utils.PrettyOperation(op))
	return nil
}
