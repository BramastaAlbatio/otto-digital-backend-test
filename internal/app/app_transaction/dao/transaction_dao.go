package dao

import (
	"context"
	"fmt"
	"time"

	entity "otto-digital-backend-test/pkg/entity"

	"otto-digital-backend-test/pkg/util"

	"gitlab.com/threetopia/sqlgo/v2"
)

type TransactionDAO interface {
	Search(ctx context.Context, query entity.TransactionQuery) (entity.Transactions, error)
	Insert(ctx context.Context, Transactions entity.Transactions) error
	Update(ctx context.Context, Transactions entity.Transactions) error
	Delete(ctx context.Context, id string) error
}

type transactionDAO struct {
	dbTrx DBTransaction
}

func MakeTransactionDAO(dbTrx DBTransaction) TransactionDAO {
	return transactionDAO{
		dbTrx: dbTrx,
	}
}

func (d transactionDAO) Search(ctx context.Context, query entity.TransactionQuery) (entity.Transactions, error) {
	sqlSelect := sqlgo.NewSQLGoSelect().
		SetSQLSelect("t.id", "id").
		SetSQLSelect("t.customer_id", "customer_id").
		SetSQLSelect("t.total_point", "total_point").
		SetSQLSelect("t.created_at", "created_at").
		SetSQLSelect("t.updated_at", "updated_at")

	sqlFrom := sqlgo.NewSQLGoFrom().
		SetSQLFrom(`"transaction"`, "t")

	sqlWhere := sqlgo.NewSQLGoWhere()
	if len(query.IDs) > 0 {
		sqlWhere.SetSQLWhere("AND", "t.id", "IN", query.IDs)
	}

	sql := sqlgo.NewSQLGo().
		SetSQLSchema("public").
		SetSQLGoSelect(sqlSelect).
		SetSQLGoFrom(sqlFrom).
		SetSQLGoWhere(sqlWhere)

	sqlStr := sql.BuildSQL()
	sqlParams := sql.GetSQLGoParameter().GetSQLParameter()
	rows, err := d.dbTrx.GetSqlDB().QueryContext(ctx, sqlStr, sqlParams...)
	if err != nil {
		return nil, err
	}

	var Transactions entity.Transactions
	for rows.Next() {
		var Transaction entity.Transaction
		if err := rows.Scan(
			&Transaction.ID,
			&Transaction.CustomerID,
			&Transaction.TotalPoints,
			&Transaction.CreatedAt,
			&Transaction.UpdatedAt,
		); err != nil {
			return nil, err
		}
		Transactions = append(Transactions, Transaction)
	}
	return Transactions, nil
}

func (d transactionDAO) Insert(ctx context.Context, transactions entity.Transactions) error {
	if len(transactions) < 1 {
		return fmt.Errorf("empty transaction(s) data")
	}

	sqlInsert := sqlgo.NewSQLGoInsert()
	sqlInsert.SetSQLInsert("transaction")
	sqlInsert.SetSQLInsertColumn("id", "customer_id", "total_point", "created_at")
	for i, transaction := range transactions {
		transaction.ID = util.MakeUUIDv4()
		transaction.CreatedAt = time.Now()
		sqlInsert.SetSQLInsertValue(transaction.ID, transaction.CustomerID, transaction.TotalPoints, transaction.CreatedAt)
		transactions[i] = transaction
	}
	sql := sqlgo.NewSQLGo().
		SetSQLSchema("public").
		SetSQLGoInsert(sqlInsert)

	sqlStr := sql.BuildSQL()
	sqlParams := sql.GetSQLGoParameter().GetSQLParameter()
	_, err := d.dbTrx.GetSqlTx().ExecContext(ctx, sqlStr, sqlParams...)
	fmt.Println("=======================================", sqlStr, sqlParams)
	if err != nil {
		return err
	}
	return nil
}

func (d transactionDAO) Update(ctx context.Context, transactions entity.Transactions) error {
	if len(transactions) < 1 {
		return fmt.Errorf("empty transaction(s) data")
	}

	for i, transaction := range transactions {
		updatedAt := time.Now()
		transaction.UpdatedAt = &updatedAt
		sql := sqlgo.NewSQLGo().
			SetSQLSchema("public").
			SetSQLUpdate("transaction").
			SetSQLUpdateValue("total_point", transaction.TotalPoints).
			SetSQLUpdateValue("updated_at", transaction.UpdatedAt).
			SetSQLWhere("AND", "id", "=", transaction.ID)
		sqlStr := sql.BuildSQL()
		sqlParams := sql.GetSQLGoParameter().GetSQLParameter()
		_, err := d.dbTrx.GetSqlTx().ExecContext(ctx, sqlStr, sqlParams...)
		if err != nil {
			return err
		}

		transactions[i] = transaction
	}

	return nil
}

func (d transactionDAO) Delete(ctx context.Context, id string) error {
	sql := sqlgo.NewSQLGo()
	sql.SetSQLSchema("public")
	sql.SetSQLDelete("transaction")
	sql.SetSQLWhere("AND", "id", "IN", id)
	sqlStr := sql.BuildSQL()
	sqlParams := sql.GetSQLGoParameter().GetSQLParameter()

	if _, err := d.dbTrx.GetSqlTx().ExecContext(ctx, sqlStr, sqlParams...); err != nil {
		return err
	}

	return nil
}
