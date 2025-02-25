package dao

import (
	"context"
	"database/sql"

	"otto-digital-backend-test/pkg/dao"
)

type DBTransaction interface {
	dao.DBTransaction

	GetTransactionDAO() TransactionDAO
}

type dbTransaction struct {
	dao.DBTransaction

	transactionDAO TransactionDAO
}

func NewTransaction(ctx context.Context, sqlDB *sql.DB) DBTransaction {
	dbTrx := &dbTransaction{
		DBTransaction: dao.NewTransaction(ctx, sqlDB),
	}
	dbTrx.transactionDAO = MakeTransactionDAO(dbTrx)
	return dbTrx
}

func (dbTrx *dbTransaction) GetTransactionDAO() TransactionDAO {
	return dbTrx.transactionDAO
}
