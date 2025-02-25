package dao

import (
	"context"
	"database/sql"

	"otto-digital-backend-test/pkg/dao"
)

type DBTransaction interface {
	dao.DBTransaction

	GetVoucherDAO() VoucherDAO
}

type dbTransaction struct {
	dao.DBTransaction

	voucherDAO VoucherDAO
}

func NewTransaction(ctx context.Context, sqlDB *sql.DB) DBTransaction {
	dbTrx := &dbTransaction{
		DBTransaction: dao.NewTransaction(ctx, sqlDB),
	}
	dbTrx.voucherDAO = MakeVoucherDAO(dbTrx)
	return dbTrx
}

func (dbTrx *dbTransaction) GetVoucherDAO() VoucherDAO {
	return dbTrx.voucherDAO
}
