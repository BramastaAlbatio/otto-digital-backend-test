package service

import (
	"context"
	"database/sql"

	"otto-digital-backend-test/internal/app/app_brand/dao"
	entity "otto-digital-backend-test/pkg/entity"
)

type BrandService interface {
	Search(ctx context.Context, query entity.BrandQuery) (entity.Brands, error)
	Insert(ctx context.Context, brands entity.Brands) error
	Update(ctx context.Context, brands entity.Brands) error
	Delete(ctx context.Context, id string) error
}

type brandService struct {
	sqlDB *sql.DB
}

func MakeBrandService(sqlDB *sql.DB) BrandService {
	return brandService{
		sqlDB: sqlDB,
	}
}

func (s brandService) Search(ctx context.Context, query entity.BrandQuery) (entity.Brands, error) {
	dbTrx := dao.NewTransaction(ctx, s.sqlDB)
	defer dbTrx.GetSqlTx().Rollback()
	users, err := dbTrx.GetBrandDAO().Search(ctx, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s brandService) Insert(ctx context.Context, brands entity.Brands) error {
	dbTrx := dao.NewTransaction(ctx, s.sqlDB)
	defer dbTrx.GetSqlTx().Rollback()

	err := dbTrx.GetBrandDAO().Insert(ctx, brands)
	if err != nil {
		return err
	}

	if err := dbTrx.GetSqlTx().Commit(); err != nil {
		return err
	}

	return nil
}

func (s brandService) Update(ctx context.Context, brands entity.Brands) error {
	dbTrx := dao.NewTransaction(ctx, s.sqlDB)
	defer dbTrx.GetSqlTx().Rollback()

	err := dbTrx.GetBrandDAO().Update(ctx, brands)
	if err != nil {
		return err
	}

	if err := dbTrx.GetSqlTx().Commit(); err != nil {
		return err
	}

	return nil
}

func (s brandService) Delete(ctx context.Context, id string) error {
	dbTrx := dao.NewTransaction(ctx, s.sqlDB)
	defer dbTrx.GetSqlTx().Rollback()

	if err := dbTrx.GetBrandDAO().Delete(ctx, id); err != nil {
		return err
	} else if err := dbTrx.GetSqlTx().Commit(); err != nil {
		return err
	}

	return nil
}
