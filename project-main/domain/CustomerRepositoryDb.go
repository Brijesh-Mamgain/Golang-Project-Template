package domain

import (
	"database/sql"
	appexception "project-common/exception"
	"project-common/logger"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *appexception.AppError) {
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = $1"
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customers table " + err.Error())
		return nil, appexception.UnexpectedError("Unexpected database error")
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *appexception.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = $1"

	var c Customer
	err := d.client.Get(&c, customerSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, appexception.NotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, appexception.UnexpectedError("Unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
}
