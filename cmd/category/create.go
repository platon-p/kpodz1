package category

import (
	"bufio"
	"fmt"
	"github.com/platon-p/kpodz1/domain"
	"github.com/platon-p/kpodz1/services"
	"github.com/platon-p/kpodz1/utils"
	"os"
	"strings"
)

type CreateCategoryCmd struct {
	Service *services.CategoryService
}

func (c *CreateCategoryCmd) Execute() error {
	category, err := scanCategory()
	if err != nil {
		return err
	}
	res, err := c.Service.Create(category.CategoryType, category.Name)
	if err != nil {
		return err
	}
	fmt.Printf("Категория создана: %s\n", utils.PrettyCategory(res))
	return nil
}

func scanCategory() (domain.Category, error) {
	rd := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите тип категории (%s/%s): ", domain.IncomeCategoryType, domain.OutcomeCategoryType)
	categoryTypeStr, err := rd.ReadString('\n')
	if err != nil {
		return domain.Category{}, err
	}
	categoryTypeStr = strings.TrimSuffix(categoryTypeStr, "\n")

	fmt.Printf("Введите название категории: ")
	name, err := rd.ReadString('\n')
	if err != nil {
		return domain.Category{}, err
	}
	name = strings.TrimSuffix(name, "\n")

	return domain.Category{CategoryType: domain.CategoryType(categoryTypeStr), Name: name}, nil
}
