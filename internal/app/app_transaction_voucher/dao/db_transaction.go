package dao

import (
	"context"
	"database/sql"

	"otto-digital-backend-test/pkg/dao"
)

type DBTransaction interface {
	dao.DBTransaction

	GetTransactionVoucherDAO() TransactionVoucherDAO
}

type dbTransaction struct {
	dao.DBTransaction

	transactionVoucherDAO TransactionVoucherDAO
}

func NewTransaction(ctx context.Context, sqlDB *sql.DB) DBTransaction {
	dbTrx := &dbTransaction{
		DBTransaction: dao.NewTransaction(ctx, sqlDB),
	}
	dbTrx.transactionVoucherDAO = MakeTransactionVoucherDAO(dbTrx)
	return dbTrx
}

func (dbTrx *dbTransaction) GetTransactionVoucherDAO() TransactionVoucherDAO {
	return dbTrx.transactionVoucherDAO
}
