package domain

type CategoryType string

const (
	IncomeCategoryType  CategoryType = "income"
	OutcomeCategoryType CategoryType = "outcome"
)

type Category struct {
	Id           uint32
	CategoryType CategoryType
	Name         string
}
