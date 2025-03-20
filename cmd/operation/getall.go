package operation

import (
	"fmt"
	"github.com/platon-p/kpodz1/application"
)

type GetAllOperationsCmd struct {
	Service *application.OperationService
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
	for i := range res {
		fmt.Println(res[i])
	}
	return nil
}
