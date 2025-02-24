package dao

import (
	"context"
	"database/sql"
)

type DBTransaction interface {
	GetSqlTx() *sql.Tx
	GetSqlDB() *sql.DB

	Commit(ctx context.Context) error
}

type dbTransaction struct {
	sqlDB *sql.DB
	sqlTx *sql.Tx
}

func NewTransaction(ctx context.Context, sqlDB *sql.DB) DBTransaction {
	sqlTx, _ := sqlDB.BeginTx(ctx, nil)
	dbTrx := &dbTransaction{
		sqlTx: sqlTx,
		sqlDB: sqlDB,
	}
	return dbTrx
}

func (dbTrx *dbTransaction) GetSqlTx() *sql.Tx {
	return dbTrx.sqlTx
}

func (dbTrx *dbTransaction) GetSqlDB() *sql.DB {
	return dbTrx.sqlDB
}

func (dbTrx *dbTransaction) Commit(ctx context.Context) error {
	return dbTrx.sqlTx.Commit()
}
