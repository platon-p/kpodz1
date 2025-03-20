package account

import (
	"fmt"
	"github.com/platon-p/kpodz1/application"
)

type DeleteAccountCmd struct {
	Service *application.AccountsService
}

func (c *DeleteAccountCmd) Execute() error {
	fmt.Print("Введите ID счёта: ")
	var id uint32
	if _, err := fmt.Scan(&id); err != nil {
		return err
	}
	if err := c.Service.Delete(id); err != nil {
		return err
	}
	return nil
}
