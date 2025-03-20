package account

import (
	"bufio"
	"fmt"
	"github.com/platon-p/kpodz1/services"
	"os"
	"strings"
)

type EditAccountNameCmd struct {
	Service *services.AccountsService
}

func (c *EditAccountNameCmd) Execute() error {
	var id uint32
	fmt.Print("Введите ID счёта для редактирования: ")
	if _, err := fmt.Scan(&id); err != nil {
		return err
	}
	fmt.Print("Введите новое имя счёта: ")
	newName, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return err
	}
	newName = strings.TrimSuffix(newName, "\n")

	if _, err := c.Service.Rename(id, newName); err != nil {
		return err
	}
	return nil
}
