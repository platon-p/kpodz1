package domain

func PerformOperation(operation Operation) {
	// pull account and category
	// validate: category matches operation and account has sufficient funds
}

func CreateCategory()   {}
func DeleteCategory()   {}
func EditCategoryName() {}

func CreateAccount()      {}
func EditAccountBalance() {}
func EditAccountName()    {}
func DeleteAccount()      {}

func CreateOperation()     {}
func DeleteOperation()     {}
func EditOperationAmount() {}

func RecalculateAmount() {}

func DumpJson() {}
func LoadJson() {}
