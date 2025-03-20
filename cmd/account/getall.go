package account

import (
	"fmt"
	"github.com/platon-p/kpodz1/services"
	"github.com/platon-p/kpodz1/utils"
)

type GetAllAccountsCmd struct {
	Service *services.AccountsService
}

func (c *GetAllAccountsCmd) Execute() error {
	res, err := c.Service.GetAll()
	if err != nil {
		return err
	}
	if len(res) == 0 {
		fmt.Println("Пустой список")
		return nil
	}
	fmt.Println(utils.BankAccountTitle())
	for i := range res {
		fmt.Println(utils.PrettyBankAccount(res[i]))
	}
	return nil
}
