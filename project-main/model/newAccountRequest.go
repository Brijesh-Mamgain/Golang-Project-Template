package model

import (
	appexception "project-common/exception"
	"strings"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *appexception.AppError {
	if r.Amount < 5000 {
		return appexception.ValidationError("To open a new account you need to deposit atleast 5000.00")
	}
	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return appexception.ValidationError("Account type should be checking or saving")
	}
	return nil
}
