package operation

import (
	"fmt"
	"github.com/platon-p/kpodz1/services"
)

type DeleteOperationCmd struct {
	Service *services.OperationService
}

func (c *DeleteOperationCmd) Execute() error {
	var id uint32
	fmt.Print("Введите ID операции: ")
	if _, err := fmt.Scan(&id); err != nil {
		return err
	}
	return c.Service.Delete(id)
}
