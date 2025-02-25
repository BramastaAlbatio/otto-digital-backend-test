package service

import (
	"context"
	"database/sql"

	"otto-digital-backend-test/internal/app/app_transaction_voucher/dao"
	entity "otto-digital-backend-test/pkg/entity"
)

type TransactionVoucherService interface {
	Search(ctx context.Context, query entity.TransactionVoucherQuery) (entity.TransactionVouchers, error)
	Insert(ctx context.Context, transactionVouchers entity.TransactionVouchers) error
	Update(ctx context.Context, transactionVouchers entity.TransactionVouchers) error
	Delete(ctx context.Context, id string) error
}

type transactionVoucherService struct {
	sqlDB *sql.DB
}

func MakeTransactionVoucherService(sqlDB *sql.DB) TransactionVoucherService {
	return transactionVoucherService{
		sqlDB: sqlDB,
	}
}

func (s transactionVoucherService) Search(ctx context.Context, query entity.TransactionVoucherQuery) (entity.TransactionVouchers, error) {
	dbTrx := dao.NewTransaction(ctx, s.sqlDB)
	defer dbTrx.GetSqlTx().Rollback()
	transactionVouchers, err := dbTrx.GetTransactionVoucherDAO().Search(ctx, query)
	if err != nil {
		return nil, err
	}
	return transactionVouchers, nil
}

func (s transactionVoucherService) Insert(ctx context.Context, transactionVouchers entity.TransactionVouchers) error {
	dbTrx := dao.NewTransaction(ctx, s.sqlDB)
	defer dbTrx.GetSqlTx().Rollback()

	err := dbTrx.GetTransactionVoucherDAO().Insert(ctx, transactionVouchers)
	if err != nil {
		return err
	}

	if err := dbTrx.GetSqlTx().Commit(); err != nil {
		return err
	}

	return nil
}

func (s transactionVoucherService) Update(ctx context.Context, transactionVouchers entity.TransactionVouchers) error {
	dbTrx := dao.NewTransaction(ctx, s.sqlDB)
	defer dbTrx.GetSqlTx().Rollback()

	err := dbTrx.GetTransactionVoucherDAO().Update(ctx, transactionVouchers)
	if err != nil {
		return err
	}

	if err := dbTrx.GetSqlTx().Commit(); err != nil {
		return err
	}

	return nil
}

func (s transactionVoucherService) Delete(ctx context.Context, id string) error {
	dbTrx := dao.NewTransaction(ctx, s.sqlDB)
	defer dbTrx.GetSqlTx().Rollback()

	if err := dbTrx.GetTransactionVoucherDAO().Delete(ctx, id); err != nil {
		return err
	} else if err := dbTrx.GetSqlTx().Commit(); err != nil {
		return err
	}

	return nil
}
