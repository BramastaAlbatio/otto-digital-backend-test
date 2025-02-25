package dao

import (
	"context"
	"database/sql"

	"otto-digital-backend-test/pkg/dao"
)

type DBTransaction interface {
	dao.DBTransaction

	GetCustomerDAO() CustomerDAO
}

type dbTransaction struct {
	dao.DBTransaction

	customerDAO CustomerDAO
}

func NewTransaction(ctx context.Context, sqlDB *sql.DB) DBTransaction {
	dbTrx := &dbTransaction{
		DBTransaction: dao.NewTransaction(ctx, sqlDB),
	}
	dbTrx.customerDAO = MakeCustomerDAO(dbTrx)
	return dbTrx
}

func (dbTrx *dbTransaction) GetCustomerDAO() CustomerDAO {
	return dbTrx.customerDAO
}
