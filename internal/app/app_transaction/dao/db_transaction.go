package dao

import (
	"context"
	"database/sql"

	"otto-digital-backend-test/pkg/dao"
)

type DBTransaction interface {
	dao.DBTransaction

	GetTransactionDAO() TransactionDAO
	GetTransactionVoucherDAO() TransactionVoucherDAO
	GetVoucherDAO() VoucherDAO
}

type dbTransaction struct {
	dao.DBTransaction

	transactionDAO        TransactionDAO
	transactionVoucherDAO TransactionVoucherDAO
	voucherDAO            VoucherDAO
}

func NewTransaction(ctx context.Context, sqlDB *sql.DB) DBTransaction {
	dbTrx := &dbTransaction{
		DBTransaction: dao.NewTransaction(ctx, sqlDB),
	}
	dbTrx.transactionDAO = MakeTransactionDAO(dbTrx)
	dbTrx.transactionVoucherDAO = MakeTransactionVoucherDAO(dbTrx)
	dbTrx.voucherDAO = MakeVoucherDAO(dbTrx)
	return dbTrx
}

func (dbTrx *dbTransaction) GetTransactionDAO() TransactionDAO {
	return dbTrx.transactionDAO
}

func (dbTrx *dbTransaction) GetTransactionVoucherDAO() TransactionVoucherDAO {
	return dbTrx.transactionVoucherDAO
}

func (dbTrx *dbTransaction) GetVoucherDAO() VoucherDAO {
	return dbTrx.voucherDAO
}
