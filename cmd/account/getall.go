package account

import (
	"fmt"
	"github.com/platon-p/kpodz1/application"
)

type GetAllAccountsCmd struct {
	Service *application.AccountsService
}

func (c *GetAllAccountsCmd) Execute() error {
	res, err := c.Service.GetAll()
	if err != nil {
		return err
	}
	fmt.Println(res, len(res))
	return nil
}
