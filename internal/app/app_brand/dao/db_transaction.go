package dao

import (
	"context"
	"database/sql"

	"otto-digital-backend-test/pkg/dao"
)

type DBTransaction interface {
	dao.DBTransaction

	GetBrandDAO() BrandDAO
}

type dbTransaction struct {
	dao.DBTransaction

	brandDAO BrandDAO
}

func NewTransaction(ctx context.Context, sqlDB *sql.DB) DBTransaction {
	dbTrx := &dbTransaction{
		DBTransaction: dao.NewTransaction(ctx, sqlDB),
	}
	dbTrx.brandDAO = MakeBrandDAO(dbTrx)
	return dbTrx
}

func (dbTrx *dbTransaction) GetBrandDAO() BrandDAO {
	return dbTrx.brandDAO
}
