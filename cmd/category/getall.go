package category

import (
	"fmt"
	"github.com/platon-p/kpodz1/application"
)

type GetAllCategoriesCmd struct {
	Service *application.CategoryService
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
	for _, category := range categories {
		fmt.Println(category)
	}
	return nil
}
