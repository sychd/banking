package domain

import (
	"database/sql"
	"github.com/dsych/banking/errs"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
	"time"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (db CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	var args []interface{}

	if status == "active" || status == "inactive" {
		findAllSql = strings.Join([]string{findAllSql, " where status = ?"}, "")
		args = append(args, CustomerStatusDict[status])
	}

	rows, err := db.client.Query(findAllSql, args...)

	if err != nil {
		log.Println("Error while querying customer table" + err.Error())
		return nil, errs.NewUnexpectedError("Error while querying customer table")
	}

	customers := make([]Customer, 0)

	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)

		if err != nil {
			log.Println("Error while scanning customers" + err.Error())
			return nil, errs.NewUnexpectedError("Error while querying customer customers")
		}
		customers = append(customers, c)
	}

	return customers, nil
}

func (db CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	row := db.client.QueryRow(customerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			log.Println("Error while reading customer" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}

	return &c, nil
}

func NewCustomerRepositoryDb(dbUrl string) CustomerRepositoryDb {
	client, err := sql.Open("mysql", dbUrl)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}
