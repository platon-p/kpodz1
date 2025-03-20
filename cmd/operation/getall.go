package operation

import (
	"fmt"
	"github.com/platon-p/kpodz1/services"
	"github.com/platon-p/kpodz1/utils"
)

type GetAllOperationsCmd struct {
	Service *services.OperationService
}

func (c *GetAllOperationsCmd) Execute() error {
	res, err := c.Service.GetAll()
	if err != nil {
		return err
	}
	if len(res) == 0 {
		fmt.Println("Пустой список")
		return nil
	}
	fmt.Println(utils.OperationTitle())
	for i := range res {
		fmt.Println(utils.PrettyOperation(res[i]))
	}
	return nil
}
