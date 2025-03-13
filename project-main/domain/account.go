package domain

import (
	appexecption "project-common/exception"
	"project-main/model"
)

const dbTSLayout = "2006-01-02 15:04:05"

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

func (a Account) ToNewAccountResponseDto() *model.NewAccountResponse {
	return &model.NewAccountResponse{AccountId: a.AccountId}
}

//go:generate mockgen -destination=../mocks/domain/mockAccountRepository.go -package=domain github.com/ashishjuyal/banking/domain AccountRepository
type AccountRepository interface {
	Save(account Account) (*Account, *appexecption.AppError)
	SaveTransaction(transaction Transaction) (*Transaction, *appexecption.AppError)
	FindBy(accountId string) (*Account, *appexecption.AppError)
}

func (a Account) CanWithdraw(amount float64) bool {
	if a.Amount < amount {
		return false
	}
	return true
}

func NewAccount(customerId, accountType string, amount float64) Account {
	return Account{
		CustomerId:  customerId,
		OpeningDate: dbTSLayout,
		AccountType: accountType,
		Amount:      amount,
		Status:      "1",
	}
}
