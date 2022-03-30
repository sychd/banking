package domain

import (
	"github.com/dsych/banking/errs"
	"github.com/dsych/banking/logger"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (db AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"
	result, err := db.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)

	if err != nil {
		logger.Error("Error while creating an account, " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error while account creation")
	}
	id, err := result.LastInsertId()

	if err != nil {
		logger.Error("Error while getting created account id, " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error while account creation")
	}
	a.AccountId = strconv.FormatInt(id, 10)

	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
