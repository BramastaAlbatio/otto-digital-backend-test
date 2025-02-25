package service

import (
	"context"
	"database/sql"

	"otto-digital-backend-test/internal/app/app_voucher/dao"
	entity "otto-digital-backend-test/pkg/entity"
)

type VoucherService interface {
	Search(ctx context.Context, query entity.VoucherQuery) (entity.Vouchers, error)
	Insert(ctx context.Context, vouchers entity.Vouchers) error
	Update(ctx context.Context, vouchers entity.Vouchers) error
	Delete(ctx context.Context, id string) error
}

type voucherService struct {
	sqlDB *sql.DB
}

func MakeVoucherService(sqlDB *sql.DB) VoucherService {
	return voucherService{
		sqlDB: sqlDB,
	}
}

func (s voucherService) Search(ctx context.Context, query entity.VoucherQuery) (entity.Vouchers, error) {
	dbTrx := dao.NewTransaction(ctx, s.sqlDB)
	defer dbTrx.GetSqlTx().Rollback()
	vouchers, err := dbTrx.GetVoucherDAO().Search(ctx, query)
	if err != nil {
		return nil, err
	}
	return vouchers, nil
}

func (s voucherService) Insert(ctx context.Context, vouchers entity.Vouchers) error {
	dbTrx := dao.NewTransaction(ctx, s.sqlDB)
	defer dbTrx.GetSqlTx().Rollback()

	err := dbTrx.GetVoucherDAO().Insert(ctx, vouchers)
	if err != nil {
		return err
	}

	if err := dbTrx.GetSqlTx().Commit(); err != nil {
		return err
	}

	return nil
}

func (s voucherService) Update(ctx context.Context, vouchers entity.Vouchers) error {
	dbTrx := dao.NewTransaction(ctx, s.sqlDB)
	defer dbTrx.GetSqlTx().Rollback()

	err := dbTrx.GetVoucherDAO().Update(ctx, vouchers)
	if err != nil {
		return err
	}

	if err := dbTrx.GetSqlTx().Commit(); err != nil {
		return err
	}

	return nil
}

func (s voucherService) Delete(ctx context.Context, id string) error {
	dbTrx := dao.NewTransaction(ctx, s.sqlDB)
	defer dbTrx.GetSqlTx().Rollback()

	if err := dbTrx.GetVoucherDAO().Delete(ctx, id); err != nil {
		return err
	} else if err := dbTrx.GetSqlTx().Commit(); err != nil {
		return err
	}

	return nil
}
