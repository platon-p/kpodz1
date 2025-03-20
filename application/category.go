package application

import "github.com/platon-p/kpodz1/domain"

type CategoryRepository interface {
	Create(categoryType domain.CategoryType, name string) (domain.Category, error)
	Find(id uint32) (domain.Category, error)
	GetAll() ([]domain.Category, error)
	EditName(id uint32, name string) (domain.Category, error)
	Delete(id uint32) error
}
