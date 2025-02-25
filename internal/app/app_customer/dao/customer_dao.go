package dao

import (
	"context"
	"fmt"
	"time"

	entity "otto-digital-backend-test/pkg/entity"

	"otto-digital-backend-test/pkg/util"

	"gitlab.com/threetopia/sqlgo/v2"
)

type CustomerDAO interface {
	Search(ctx context.Context, query entity.CustomerQuery) (entity.Customers, error)
	Insert(ctx context.Context, customers entity.Customers) error
	Update(ctx context.Context, customers entity.Customers) error
	Delete(ctx context.Context, id string) error
}

type customerDAO struct {
	dbTrx DBTransaction
}

func MakeCustomerDAO(dbTrx DBTransaction) CustomerDAO {
	return customerDAO{
		dbTrx: dbTrx,
	}
}

func (d customerDAO) Search(ctx context.Context, query entity.CustomerQuery) (entity.Customers, error) {
	sqlSelect := sqlgo.NewSQLGoSelect().
		SetSQLSelect("c.id", "id").
		SetSQLSelect("c.name", "name").
		SetSQLSelect("c.email", "email").
		SetSQLSelect("c.created_at", "created_at").
		SetSQLSelect("c.updated_at", "updated_at")

	sqlFrom := sqlgo.NewSQLGoFrom().
		SetSQLFrom(`"customer"`, "c")

	sqlWhere := sqlgo.NewSQLGoWhere()
	if len(query.IDs) > 0 {
		sqlWhere.SetSQLWhere("AND", "c.id", "IN", query.IDs)
	}
	if len(query.Names) > 0 {
		sqlWhere.SetSQLWhere("AND", "c.name", "IN", query.Names)
	}

	sql := sqlgo.NewSQLGo().
		SetSQLSchema("public").
		SetSQLGoSelect(sqlSelect).
		SetSQLGoFrom(sqlFrom).
		SetSQLGoWhere(sqlWhere)

	sqlStr := sql.BuildSQL()
	sqlParams := sql.GetSQLGoParameter().GetSQLParameter()
	rows, err := d.dbTrx.GetSqlDB().QueryContext(ctx, sqlStr, sqlParams...)
	if err != nil {
		return nil, err
	}

	var Customers entity.Customers
	for rows.Next() {
		var Customer entity.Customer
		if err := rows.Scan(
			&Customer.ID,
			&Customer.Name,
			&Customer.Email,
			&Customer.CreatedAt,
			&Customer.UpdatedAt,
		); err != nil {
			return nil, err
		}
		Customers = append(Customers, Customer)
	}
	return Customers, nil
}

func (d customerDAO) Insert(ctx context.Context, customers entity.Customers) error {
	if len(customers) < 1 {
		return fmt.Errorf("empty Customer(s) data")
	}

	sqlInsert := sqlgo.NewSQLGoInsert()
	sqlInsert.SetSQLInsert("customer")
	sqlInsert.SetSQLInsertColumn("id", "name", "email", "created_at")
	for i, customer := range customers {
		customer.ID = util.MakeUUIDv4()
		customer.CreatedAt = time.Now()
		sqlInsert.SetSQLInsertValue(customer.ID, customer.Name, customer.Email, customer.CreatedAt)
		customers[i] = customer
	}
	sql := sqlgo.NewSQLGo().
		SetSQLSchema("public").
		SetSQLGoInsert(sqlInsert)

	sqlStr := sql.BuildSQL()
	sqlParams := sql.GetSQLGoParameter().GetSQLParameter()
	_, err := d.dbTrx.GetSqlTx().ExecContext(ctx, sqlStr, sqlParams...)
	if err != nil {
		return err
	}
	return nil
}

func (d customerDAO) Update(ctx context.Context, customers entity.Customers) error {
	if len(customers) < 1 {
		return fmt.Errorf("empty user(s) data")
	}

	for i, Customer := range customers {
		updatedAt := time.Now()
		Customer.UpdatedAt = &updatedAt
		sql := sqlgo.NewSQLGo().
			SetSQLSchema("public").
			SetSQLUpdate("customer").
			SetSQLUpdateValue("name", Customer.Name).
			SetSQLUpdateValue("email", Customer.Email).
			SetSQLUpdateValue("updated_at", Customer.UpdatedAt).
			SetSQLWhere("AND", "id", "=", Customer.ID)
		sqlStr := sql.BuildSQL()
		sqlParams := sql.GetSQLGoParameter().GetSQLParameter()
		_, err := d.dbTrx.GetSqlTx().ExecContext(ctx, sqlStr, sqlParams...)
		if err != nil {
			return err
		}

		customers[i] = Customer
	}

	return nil
}

func (d customerDAO) Delete(ctx context.Context, id string) error {
	sql := sqlgo.NewSQLGo()
	sql.SetSQLSchema("public")
	sql.SetSQLDelete("customer")
	sql.SetSQLWhere("AND", "id", "IN", id)
	sqlStr := sql.BuildSQL()
	sqlParams := sql.GetSQLGoParameter().GetSQLParameter()

	if _, err := d.dbTrx.GetSqlTx().ExecContext(ctx, sqlStr, sqlParams...); err != nil {
		return err
	}

	return nil
}
