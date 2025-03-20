package account

import (
	"bufio"
	"fmt"
	"github.com/platon-p/kpodz1/services"
	"github.com/platon-p/kpodz1/utils"
	"os"
	"strings"
)

type CreateAccountCmd struct {
	Service *services.AccountsService
}

func (c *CreateAccountCmd) Execute() error {
	fmt.Print("Введите имя счёта: ")
	name, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return err
	}
	name = strings.TrimSuffix(name, "\n")
	created, err := c.Service.Create(name)
	if err != nil {
		return err
	}
	fmt.Println("Счёт создан: \n", utils.PrettyBankAccount(created))
	return nil
}
