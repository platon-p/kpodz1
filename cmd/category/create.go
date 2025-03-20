package category

import (
	"bufio"
	"fmt"
	"github.com/platon-p/kpodz1/application"
	"github.com/platon-p/kpodz1/domain"
	"os"
	"strings"
)

type CreateCategoryCmd struct {
	Service *application.CategoryService
}

func (c *CreateCategoryCmd) Execute() error {
	fmt.Printf("Введите тип категории (%s/%s): ", domain.IncomeCategoryType, domain.OutcomeCategoryType)
	rd := bufio.NewReader(os.Stdin)
	categoryTypeStr, err := rd.ReadString('\n')
	if err != nil {
		return err
	}
	categoryTypeStr = strings.TrimSuffix(categoryTypeStr, "\n")
	fmt.Printf("Введите название категории: ")
	name, err := rd.ReadString('\n')
	if err != nil {
		return err
	}
	name = strings.TrimSuffix(name, "\n")
	res, err := c.Service.Create(domain.CategoryType(categoryTypeStr), name)
	if err != nil {
		return err
	}
	fmt.Printf("Категория создана: %v\n", res)
	return nil
}
