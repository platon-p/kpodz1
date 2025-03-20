package infra

import (
	"errors"
	"github.com/platon-p/kpodz1/domain"
	"slices"
)

var (
	ErrAccountNotFound = errors.New("account not found")
)

type InMemoryAccountRepository struct {
	accounts []domain.BankAccount
}

func NewInMemoryAccountRepository(initialData []domain.BankAccount) *InMemoryAccountRepository {
	return &InMemoryAccountRepository{
		accounts: initialData,
	}
}

func (i *InMemoryAccountRepository) Create(name string) (domain.BankAccount, error) {
	acc := domain.BankAccount{
		Id:      i.nextId(),
		Name:    name,
		Balance: 0,
	}
	i.accounts = append(i.accounts, acc)
	return acc, nil
}

func (i *InMemoryAccountRepository) nextId() uint32 {
	if len(i.accounts) == 0 {
		return 1
	}
	return i.accounts[len(i.accounts)-1].Id + 1
}

func (i *InMemoryAccountRepository) Find(id uint32) (domain.BankAccount, error) {
	accRef, err := i.find(id)
	if err != nil {
		return domain.BankAccount{}, err
	}
	return *accRef, nil
}

func (i *InMemoryAccountRepository) find(id uint32) (*domain.BankAccount, error) {
	idx := slices.IndexFunc(i.accounts, func(account domain.BankAccount) bool {
		return account.Id == id
	})
	if idx < 0 {
		return nil, ErrAccountNotFound
	}
	return &i.accounts[idx], nil
}

func (i *InMemoryAccountRepository) GetAll() ([]domain.BankAccount, error) {
	accs := make([]domain.BankAccount, 0, len(i.accounts))
	for _, acc := range i.accounts {
		accs = append(accs, acc)
	}
	return accs, nil
}

func (i *InMemoryAccountRepository) EditName(id uint32, newName string) (domain.BankAccount, error) {
	acc, err := i.find(id)
	if err != nil {
		return domain.BankAccount{}, err
	}
	acc.Name = newName
	return *acc, nil
}

func (i *InMemoryAccountRepository) EditBalance(id uint32, newBalance float64) (domain.BankAccount, error) {
	acc, err := i.find(id)
	if err != nil {
		return domain.BankAccount{}, err
	}
	acc.Balance = newBalance
	return *acc, nil
}

func (i *InMemoryAccountRepository) Delete(id uint32) error {
	idx := slices.IndexFunc(i.accounts, func(account domain.BankAccount) bool {
		return account.Id == id
	})
	if idx < 0 {
		return ErrAccountNotFound
	}
	i.accounts = append(i.accounts[:idx], i.accounts[idx+1:]...)
	return nil
}
