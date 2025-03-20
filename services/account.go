package services

import (
	"github.com/platon-p/kpodz1/domain"
)

type AccountsService struct {
	accountRepository    AccountRepository
	operationsRepository OperationsRepository
}

func NewAccountsService(accountRepository AccountRepository, operationsRepository OperationsRepository) *AccountsService {
	return &AccountsService{
		accountRepository:    accountRepository,
		operationsRepository: operationsRepository,
	}
}

func (s *AccountsService) Create(name string) (domain.BankAccount, error) {
	return s.accountRepository.Create(name)
}

func (s *AccountsService) GetAll() ([]domain.BankAccount, error) {
	return s.accountRepository.GetAll()
}

func (s *AccountsService) Rename(id uint32, newName string) (domain.BankAccount, error) {
	return s.accountRepository.EditName(id, newName)
}

func (s *AccountsService) Delete(id uint32) error {
	if err := s.operationsRepository.DeleteOperationsByAccount(id); err != nil {
		return err
	}
	return s.accountRepository.Delete(id)
}

type AccountRepository interface {
	Create(name string) (domain.BankAccount, error)
	Find(id uint32) (domain.BankAccount, error)
	GetAll() ([]domain.BankAccount, error)
	EditName(id uint32, newName string) (domain.BankAccount, error)
	EditBalance(id uint32, newBalance float64) (domain.BankAccount, error)
	Delete(id uint32) error
}
