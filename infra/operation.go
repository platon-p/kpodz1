package infra

import (
	"errors"
	"github.com/platon-p/kpodz1/domain"
	"github.com/platon-p/kpodz1/services"
	"slices"
)

var (
	ErrOperationNotFound = errors.New("operation not found")
)

var _ services.OperationsRepository = (*InMemoryOperationRepository)(nil)

type InMemoryOperationRepository struct {
	data []domain.Operation
}

func NewInMemoryOperationRepository(initialData []domain.Operation) *InMemoryOperationRepository {
	return &InMemoryOperationRepository{data: initialData}
}

func (r *InMemoryOperationRepository) nextId() uint32 {
	if len(r.data) == 0 {
		return 1
	}
	return r.data[len(r.data)-1].Id + 1
}

func (r *InMemoryOperationRepository) find(id uint32) (*domain.Operation, error) {
	idx := slices.IndexFunc(r.data, func(operation domain.Operation) bool {
		return operation.Id == id
	})
	if idx < 0 {
		return nil, ErrOperationNotFound
	}
	return &r.data[idx], ErrOperationNotFound
}

func (r *InMemoryOperationRepository) Create(operation domain.Operation) (domain.Operation, error) {
	operation.Id = r.nextId()
	r.data = append(r.data, operation)
	return r.data[len(r.data)-1], nil
}

func (r *InMemoryOperationRepository) GetAll() ([]domain.Operation, error) {
	return r.data, nil
}

func (r *InMemoryOperationRepository) EditAmount(id uint32, newAmount float64) (domain.Operation, error) {
	operation, err := r.find(id)
	if err != nil {
		return domain.Operation{}, err
	}
	operation.Amount = newAmount
	return *operation, nil
}

func (r *InMemoryOperationRepository) Delete(id uint32) error {
	idx := slices.IndexFunc(r.data, func(operation domain.Operation) bool {
		return operation.Id == id
	})
	if idx < 0 {
		return ErrOperationNotFound
	}
	r.data = append(r.data[:idx], r.data[idx+1:]...)
	return nil
}

func (r *InMemoryOperationRepository) DeleteOperationsByAccount(id uint32) error {
	for i := range r.data {
		if r.data[i].BankAccountId == id {
			_ = r.Delete(r.data[i].Id)
		}
	}
	return nil
}
