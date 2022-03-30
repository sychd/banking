package domain

import (
	"database/sql"
	"github.com/dsych/banking/errs"
	"github.com/dsych/banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (db CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	var args []interface{}

	if status == "active" || status == "inactive" {
		findAllSql = strings.Join([]string{findAllSql, " where status = ?"}, "")
		args = append(args, CustomerStatusDict[status])
	}

	customers := make([]Customer, 0)
	err := db.client.Select(&customers, findAllSql, args...)

	if err != nil {
		logger.Error("Error while querying customer table, " + err.Error())
		return nil, errs.NewUnexpectedError("Error while querying customer table")
	}

	return customers, nil
}

func (db CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	var c Customer
	err := db.client.Get(&c, customerSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			logger.Error("Error while reading customer" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}

	return &c, nil
}

func NewCustomerRepositoryDb(dbUrl string) CustomerRepositoryDb {
	client, err := sqlx.Open("mysql", dbUrl)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}
