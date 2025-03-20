package category

import (
	"bufio"
	"fmt"
	"github.com/platon-p/kpodz1/application"
	"os"
	"strings"
)

type EditCategoryNameCmd struct {
	Service *application.CategoryService
}

func (c *EditCategoryNameCmd) Execute() error {
	var id uint32
	fmt.Print("Введите id категории: ")
	if _, err := fmt.Scan(&id); err != nil {
		return err
	}
	fmt.Print("Введите новое имя категории: ")
	nameStr, err := bufio.NewReader(bufio.NewReader(os.Stdin)).ReadString('\n')
	if err != nil {
		return err
	}
	nameStr = strings.TrimSuffix(nameStr, "\n")
	_, err = c.Service.EditName(id, nameStr)
	return err
}
