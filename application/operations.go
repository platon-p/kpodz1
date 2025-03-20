package application

import (
	"errors"

	"github.com/platon-p/kpodz1/domain"
)

var (
	ErrCategoryOperationMatch = errors.New("category and operation types does not match")
	ErrBadOperationType       = errors.New("bad operation type")
)

type OperationsRepository interface {
	Create(operation domain.Operation) (domain.Operation, error)
	GetAll() ([]domain.Operation, error)
	EditAmount(id uint32, newAmount float64) (domain.Operation, error)
	Delete(id uint32) error
	DeleteOperationsByAccount(id uint32) error
}

type OperationService struct {
	operationsRepository OperationsRepository
	accountRepository    AccountRepository
	categoryRepository   CategoryRepository
}

func (s *OperationService) Perform(operation domain.Operation) (*domain.Operation, error) {
	acc, err := s.accountRepository.Find(operation.Id)
	if err != nil {
		return nil, err
	}
	category, err := s.categoryRepository.Find(operation.CategoryId)
	if err != nil {
		return nil, err
	}
	if !validateTypeMatch(category.CategoryType, operation.OperationType) {
		return nil, ErrCategoryOperationMatch
	}
	instance, err := s.operationsRepository.Create(operation)
	if err != nil {
		return nil, err
	}
	newBalance := acc.Balance
	switch operation.OperationType {
	case domain.IncomeOperatioType:
		newBalance += operation.Amount
	case domain.OutcomeOperationType:
		newBalance -= operation.Amount
	default:
		return nil, ErrBadOperationType
	}
	if _, err := s.accountRepository.EditBalance(acc.Id, newBalance); err != nil {
		return nil, err
	}
	return &instance, nil
}

func validateTypeMatch(category domain.CategoryType, operation domain.OperationType) bool {
	if category == domain.IncomeCategoryType && operation == domain.IncomeOperatioType {
		return true
	}
	if category == domain.OutcomeCategoryType && operation == domain.OutcomeOperationType {
		return true
	}
	return false
}
