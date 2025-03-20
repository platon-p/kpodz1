package application

import (
	"errors"

	"github.com/platon-p/kpodz1/domain"
)

var (
	ErrCategoryOperationMatch = errors.New("category and operation types does not match")
	ErrBadOperationType       = errors.New("bad operation type")
	ErrBadAmount              = errors.New("non-positive amount")
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

func NewOperationService(operationsRepository OperationsRepository, accountRepository AccountRepository, categoryRepository CategoryRepository) *OperationService {
	return &OperationService{
		operationsRepository: operationsRepository,
		accountRepository:    accountRepository,
		categoryRepository:   categoryRepository,
	}
}

func (s *OperationService) Perform(operation domain.Operation) (*domain.Operation, error) {
	if err := validateAmount(operation.Amount); err != nil {
		return nil, err
	}
	acc, err := s.accountRepository.Find(operation.BankAccountId)
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

func (s *OperationService) GetAll() ([]domain.Operation, error) {
	return s.operationsRepository.GetAll()
}

func (s *OperationService) EditAmount(id uint32, newAmount float64) (domain.Operation, error) {
	if err := validateAmount(newAmount); err != nil {
		return domain.Operation{}, err
	}
	return s.operationsRepository.EditAmount(id, newAmount)
}

func (s *OperationService) Delete(id uint32) error {
	return s.operationsRepository.Delete(id)
}

func validateAmount(amount float64) error {
	if amount <= 0 {
		return ErrBadAmount
	}
	return nil
}
