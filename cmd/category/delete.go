package category

import (
	"fmt"
	"github.com/platon-p/kpodz1/application"
)

type DeleteCategoryCmd struct {
	Service *application.CategoryService
}

func (c *DeleteCategoryCmd) Execute() error {
	var id uint32
	fmt.Printf("Введите id категории: ")
	if _, err := fmt.Scan(&id); err != nil {
		return err
	}
	return c.Service.Delete(id)
}
