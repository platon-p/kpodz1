package domain

type CategoryType string

const (
	IncomeCategoryType  CategoryType = "Доход"
	OutcomeCategoryType CategoryType = "Расход"
)

type Category struct {
	id           uint32
	categoryType CategoryType
	name         string
}
