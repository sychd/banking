package domain

import (
	"database/sql"
	"github.com/ashishjuyal/banking-lib/errs"
	"github.com/ashishjuyal/banking-lib/logger"
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

func (db AccountRepositoryDb) ById(id string) (*Account, *errs.AppError) {
	accountSql := "select account_id, customer_id, opening_date, account_type, amount, status from accounts where account_id = ?"
	var a Account
	err := db.client.Get(&a, accountSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("account not found")
		} else {
			logger.Error("Error while reading account, " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}

	return &a, nil
}

func (db AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	tx, err := db.client.Begin()
	if err != nil {
		logger.Error("Error while transaction start, " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error while transaction creation")
	}

	res, _ := tx.Exec(`insert into transactions (account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?)`,
		t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	if t.IsWithdrawal() {
		_, err = tx.Exec(`update accounts set amount = amount - ? where account_id = ?`, t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec(`update accounts set amount = amount + ? where account_id = ?`, t.Amount, t.AccountId)
	}

	if err != nil {
		tx.Rollback()
		logger.Error("Error while transaction save, " + err.Error())
		return nil, errs.NewUnexpectedError("Error while saving transaction, " + err.Error())
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while transaction save, " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	transactionId, err := res.LastInsertId()
	if err != nil {
		logger.Error("Error while getting inserted id, " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	account, appErr := db.ById(t.AccountId)
	if appErr != nil {
		logger.Error("Error while getting inserted id, " + err.Error())
		return nil, appErr
	}

	t.TransactionId = strconv.FormatInt(transactionId, 10)
	t.Amount = account.Amount

	return &t, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
