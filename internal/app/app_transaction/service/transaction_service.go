package service

import (
	"context"
	"database/sql"
	"errors"

	"otto-digital-backend-test/internal/app/app_transaction/dao"
	entity "otto-digital-backend-test/pkg/entity"
)

type TransactionService interface {
	Search(ctx context.Context, query entity.TransactionQuery) (entity.Transactions, error)
	Insert(ctx context.Context, transactions entity.Transactions) error
	Update(ctx context.Context, transactions entity.Transactions) error
	Delete(ctx context.Context, id string) error
}

type transactionService struct {
	sqlDB *sql.DB
}

func MakeTransactionService(sqlDB *sql.DB) TransactionService {
	return transactionService{
		sqlDB: sqlDB,
	}
}

func (s transactionService) Search(ctx context.Context, query entity.TransactionQuery) (entity.Transactions, error) {
	dbTrx := dao.NewTransaction(ctx, s.sqlDB)
	defer dbTrx.GetSqlTx().Rollback()
	transactions, err := dbTrx.GetTransactionDAO().Search(ctx, query)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (s transactionService) Insert(ctx context.Context, transactions entity.Transactions) error {
	dbTrx := dao.NewTransaction(ctx, s.sqlDB)
	defer dbTrx.GetSqlTx().Rollback()

	// Loop setiap transaksi
	for _, transaction := range transactions {
		// Ambil data voucher sekali saja
		vouchers, err := dbTrx.GetVoucherDAO().Search(ctx, entity.VoucherQuery{})
		if err != nil || len(vouchers) == 0 {
			return errors.New("vouchers not found")
		}

		// Hitung total poin transaksi
		if err := transaction.CalculateTotalPoints(vouchers); err != nil {
			return err
		}

		// Simpan transaksi utama terlebih dahulu
		if err := dbTrx.GetTransactionDAO().Insert(ctx, entity.Transactions{transaction}); err != nil {
			return err
		}

		transactionVouchers := transaction.GetTransactionVouchers()

		// Simpan detail voucher transaksi
		if err := dbTrx.GetTransactionVoucherDAO().Insert(ctx, transactionVouchers); err != nil {
			return err
		}
	}

	if err := dbTrx.GetSqlTx().Commit(); err != nil {
		return err
	}

	return nil
}

func (s transactionService) Update(ctx context.Context, transactions entity.Transactions) error {
	dbTrx := dao.NewTransaction(ctx, s.sqlDB)
	defer dbTrx.GetSqlTx().Rollback()

	err := dbTrx.GetTransactionDAO().Update(ctx, transactions)
	if err != nil {
		return err
	}

	if err := dbTrx.GetSqlTx().Commit(); err != nil {
		return err
	}

	return nil
}

func (s transactionService) Delete(ctx context.Context, id string) error {
	dbTrx := dao.NewTransaction(ctx, s.sqlDB)
	defer dbTrx.GetSqlTx().Rollback()

	if err := dbTrx.GetTransactionDAO().Delete(ctx, id); err != nil {
		return err
	} else if err := dbTrx.GetSqlTx().Commit(); err != nil {
		return err
	}

	return nil
}
