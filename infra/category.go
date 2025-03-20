package infra

import (
	"errors"
	"github.com/platon-p/kpodz1/domain"
	"slices"
)

var (
	ErrCategoryNotFound = errors.New("category not found")
)

type InMemoryCategoryRepository struct {
	data []domain.Category
}

func NewInMemoryCategoryRepository(initialData []domain.Category) *InMemoryCategoryRepository {
	return &InMemoryCategoryRepository{data: initialData}
}

func (r *InMemoryCategoryRepository) nextId() uint32 {
	if len(r.data) == 0 {
		return 1
	}
	return r.data[len(r.data)-1].Id + 1
}

func (r *InMemoryCategoryRepository) Create(categoryType domain.CategoryType, name string) (domain.Category, error) {
	category := domain.Category{
		Id:           r.nextId(),
		CategoryType: categoryType,
		Name:         name,
	}
	r.data = append(r.data, category)
	return category, nil
}

func (r *InMemoryCategoryRepository) find(id uint32) (int, error) {
	idx := slices.IndexFunc(r.data, func(category domain.Category) bool {
		return category.Id == id
	})
	if idx < 0 {
		return 0, ErrCategoryNotFound
	}
	return idx, nil
}

func (r *InMemoryCategoryRepository) Find(id uint32) (domain.Category, error) {
	idx, err := r.find(id)
	if err != nil {
		return domain.Category{}, err
	}
	return r.data[idx], nil
}

func (r *InMemoryCategoryRepository) GetAll() ([]domain.Category, error) {
	return r.data, nil
}

func (r *InMemoryCategoryRepository) EditName(id uint32, name string) (domain.Category, error) {
	idx, err := r.find(id)
	if err != nil {
		return domain.Category{}, err
	}
	r.data[idx].Name = name
	return r.data[idx], nil
}

func (r *InMemoryCategoryRepository) Delete(id uint32) error {
	idx, err := r.find(id)
	if err != nil {
		return err
	}
	r.data = append(r.data[:idx], r.data[idx+1:]...)
	return nil
}
