package infra

import (
	"errors"
	"github.com/platon-p/kpodz1/application"
	"github.com/platon-p/kpodz1/domain"
	"slices"
)

var (
	ErrAccountNotFound = errors.New("account not found")
	ErrAccountExists   = errors.New("account already exists")
)

var _ application.AccountRepository = (*InMemoryAccountRepository)(nil)

type InMemoryAccountRepository struct {
	data []domain.BankAccount
}

func NewInMemoryAccountRepository(initialData []domain.BankAccount) *InMemoryAccountRepository {
	return &InMemoryAccountRepository{
		data: initialData,
	}
}

func (r *InMemoryAccountRepository) Create(name string) (domain.BankAccount, error) {
	if _, err := r.findByName(name); err == nil {
		return domain.BankAccount{}, ErrAccountExists
	}
	acc := domain.BankAccount{
		Id:      r.nextId(),
		Name:    name,
		Balance: 0,
	}
	r.data = append(r.data, acc)
	return acc, nil
}

func (r *InMemoryAccountRepository) Find(id uint32) (domain.BankAccount, error) {
	idx, err := r.find(id)
	if err != nil {
		return domain.BankAccount{}, err
	}
	return r.data[idx], nil
}

func (r *InMemoryAccountRepository) GetAll() ([]domain.BankAccount, error) {
	accs := make([]domain.BankAccount, 0, len(r.data))
	for _, acc := range r.data {
		accs = append(accs, acc)
	}
	return accs, nil
}

func (r *InMemoryAccountRepository) EditName(id uint32, newName string) (domain.BankAccount, error) {
	idx, err := r.find(id)
	if err != nil {
		return domain.BankAccount{}, err
	}
	if _, err := r.findByName(newName); err == nil {
		return domain.BankAccount{}, ErrAccountExists
	}
	r.data[idx].Name = newName
	return r.data[idx], nil
}

func (r *InMemoryAccountRepository) EditBalance(id uint32, newBalance float64) (domain.BankAccount, error) {
	idx, err := r.find(id)
	if err != nil {
		return domain.BankAccount{}, err
	}
	r.data[idx].Balance = newBalance
	return r.data[idx], nil
}

func (r *InMemoryAccountRepository) Delete(id uint32) error {
	idx, err := r.find(id)
	if err != nil {
		return err
	}
	r.data = append(r.data[:idx], r.data[idx+1:]...)
	return nil
}

func (r *InMemoryAccountRepository) find(id uint32) (int, error) {
	idx := slices.IndexFunc(r.data, func(account domain.BankAccount) bool {
		return account.Id == id
	})
	if idx < 0 {
		return 0, ErrAccountNotFound
	}
	return idx, nil
}

func (r *InMemoryAccountRepository) findByName(name string) (int, error) {
	idx := slices.IndexFunc(r.data, func(account domain.BankAccount) bool {
		return account.Name == name
	})
	if idx < 0 {
		return 0, ErrAccountNotFound
	}
	return idx, nil
}

func (r *InMemoryAccountRepository) nextId() uint32 {
	if len(r.data) == 0 {
		return 1
	}
	return r.data[len(r.data)-1].Id + 1
}
