package service

import (
	"context"
	"database/sql"

	"otto-digital-backend-test/internal/app/app_customer/dao"
	entity "otto-digital-backend-test/pkg/entity"
)

type CustomerService interface {
	Search(ctx context.Context, query entity.CustomerQuery) (entity.Customers, error)
	Insert(ctx context.Context, customers entity.Customers) error
	Update(ctx context.Context, customers entity.Customers) error
	Delete(ctx context.Context, id string) error
}

type customerService struct {
	sqlDB *sql.DB
}

func MakeCustomerService(sqlDB *sql.DB) CustomerService {
	return customerService{
		sqlDB: sqlDB,
	}
}

func (s customerService) Search(ctx context.Context, query entity.CustomerQuery) (entity.Customers, error) {
	dbTrx := dao.NewTransaction(ctx, s.sqlDB)
	defer dbTrx.GetSqlTx().Rollback()
	customers, err := dbTrx.GetCustomerDAO().Search(ctx, query)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (s customerService) Insert(ctx context.Context, customers entity.Customers) error {
	dbTrx := dao.NewTransaction(ctx, s.sqlDB)
	defer dbTrx.GetSqlTx().Rollback()

	err := dbTrx.GetCustomerDAO().Insert(ctx, customers)
	if err != nil {
		return err
	}

	if err := dbTrx.GetSqlTx().Commit(); err != nil {
		return err
	}

	return nil
}

func (s customerService) Update(ctx context.Context, customers entity.Customers) error {
	dbTrx := dao.NewTransaction(ctx, s.sqlDB)
	defer dbTrx.GetSqlTx().Rollback()

	err := dbTrx.GetCustomerDAO().Update(ctx, customers)
	if err != nil {
		return err
	}

	if err := dbTrx.GetSqlTx().Commit(); err != nil {
		return err
	}

	return nil
}

func (s customerService) Delete(ctx context.Context, id string) error {
	dbTrx := dao.NewTransaction(ctx, s.sqlDB)
	defer dbTrx.GetSqlTx().Rollback()

	if err := dbTrx.GetCustomerDAO().Delete(ctx, id); err != nil {
		return err
	} else if err := dbTrx.GetSqlTx().Commit(); err != nil {
		return err
	}

	return nil
}
