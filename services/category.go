package services

import (
	"errors"
	"github.com/platon-p/kpodz1/domain"
)

var (
	ErrInvalidCategoryType = errors.New("invalid category type")
)

type CategoryRepository interface {
	Create(categoryType domain.CategoryType, name string) (domain.Category, error)
	Find(id uint32) (domain.Category, error)
	GetAll() ([]domain.Category, error)
	EditName(id uint32, name string) (domain.Category, error)
	Delete(id uint32) error
}

type CategoryService struct {
	repo CategoryRepository
}

func NewCategoryService(repo CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) Create(categoryType domain.CategoryType, name string) (domain.Category, error) {
	if err := validateCategoryType(categoryType); err != nil {
		return domain.Category{}, err
	}
	return s.repo.Create(categoryType, name)
}

func (s *CategoryService) GetAll() ([]domain.Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) EditName(id uint32, name string) (domain.Category, error) {
	return s.repo.EditName(id, name)
}

func (s *CategoryService) Delete(id uint32) error {
	return s.repo.Delete(id)
}

func validateCategoryType(categoryType domain.CategoryType) error {
	if categoryType != domain.IncomeCategoryType && categoryType != domain.OutcomeCategoryType {
		return ErrInvalidCategoryType
	}
	return nil
}
