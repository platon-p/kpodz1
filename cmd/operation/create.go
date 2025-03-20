package operation

import (
	"bufio"
	"fmt"
	"github.com/platon-p/kpodz1/domain"
	"github.com/platon-p/kpodz1/services"
	"os"
	"strings"
	"time"
)

type CreateOperationCmd struct {
	Service *services.OperationService
}

func (c *CreateOperationCmd) Execute() error {
	const layout = time.DateTime
	op, err := scanOperation(layout)
	_, err = c.Service.Perform(op)
	return err
}

func scanOperation(layout string) (op domain.Operation, err error) {
	fmt.Print("Введите ID счёта: ")
	if _, err = fmt.Scan(&op.BankAccountId); err != nil {
		return
	}
	fmt.Print("Введите ID категории: ")
	if _, err = fmt.Scan(&op.CategoryId); err != nil {
		return
	}
	fmt.Printf("Введите тип операции (%s/%s): ", domain.IncomeOperatioType, domain.OutcomeCategoryType)
	if _, err = fmt.Scan(&op.OperationType); err != nil {
		return
	}
	fmt.Printf("Введите время операции (в формате %s, пустая строка = now): ", layout)
	rd := bufio.NewReader(os.Stdin)
	dateStr, err := rd.ReadString('\n')
	if err != nil {
		return
	}
	dateStr = strings.TrimSuffix(dateStr, "\n")
	if dateStr == "" {
		op.Date = time.Now()
	} else if op.Date, err = time.Parse(layout, dateStr); err != nil {
		return
	}
	fmt.Print("Введите сумму операции (дробное число): ")
	if _, err = fmt.Scan(&op.Amount); err != nil {
		return
	}
	fmt.Print("Введите описание: ")
	descriptionStr, err := rd.ReadString('\n')
	if err != nil {
		return
	}
	op.Description = strings.TrimSuffix(descriptionStr, "\n")
	return
}
