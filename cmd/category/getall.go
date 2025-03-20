package category

import (
	"fmt"
	"github.com/platon-p/kpodz1/services"
	"github.com/platon-p/kpodz1/utils"
)

type GetAllCategoriesCmd struct {
	Service *services.CategoryService
}

func (c *GetAllCategoriesCmd) Execute() error {
	categories, err := c.Service.GetAll()
	if err != nil {
		return err
	}
	if len(categories) == 0 {
		fmt.Println("Пустой список")
		return nil
	}
	fmt.Println(utils.CategoryTitle())
	for _, category := range categories {
		fmt.Println(utils.PrettyCategory(category))
	}
	return nil
}
