package dao

import (
	"context"
	"fmt"
	"time"

	entity "otto-digital-backend-test/pkg/entity"

	"otto-digital-backend-test/pkg/util"

	"gitlab.com/threetopia/sqlgo/v2"
)

type TransactionVoucherDAO interface {
	Search(ctx context.Context, query entity.TransactionVoucherQuery) (entity.TransactionVouchers, error)
	Insert(ctx context.Context, transactionVouchers entity.TransactionVouchers) error
	Update(ctx context.Context, transactionVouchers entity.TransactionVouchers) error
	Delete(ctx context.Context, id string) error
}

type transactionVoucherDAO struct {
	dbTrx DBTransaction
}

func MakeTransactionVoucherDAO(dbTrx DBTransaction) TransactionVoucherDAO {
	return transactionVoucherDAO{
		dbTrx: dbTrx,
	}
}

func (d transactionVoucherDAO) Search(ctx context.Context, query entity.TransactionVoucherQuery) (entity.TransactionVouchers, error) {
	sqlSelect := sqlgo.NewSQLGoSelect().
		SetSQLSelect("tv.id", "id").
		SetSQLSelect("tv.transaction_id", "transaction_id").
		SetSQLSelect("tv.voucher_id", "voucher_id").
		SetSQLSelect("tv.quantity", "quantity").
		SetSQLSelect("tv.subtotal_points", "subtotal_points").
		SetSQLSelect("tv.created_at", "created_at").
		SetSQLSelect("tv.updated_at", "updated_at")

	sqlFrom := sqlgo.NewSQLGoFrom().
		SetSQLFrom(`"transaction_voucher"`, "tv")

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

	var TransactionVouchers entity.TransactionVouchers
	for rows.Next() {
		var TransactionVoucher entity.TransactionVoucher
		if err := rows.Scan(
			&TransactionVoucher.ID,
			&TransactionVoucher.TransactionID,
			&TransactionVoucher.VoucherID,
			&TransactionVoucher.Quantity,
			&TransactionVoucher.SubtotalPoints,
			&TransactionVoucher.CreatedAt,
			&TransactionVoucher.UpdatedAt,
		); err != nil {
			return nil, err
		}
		TransactionVouchers = append(TransactionVouchers, TransactionVoucher)
	}
	return TransactionVouchers, nil
}

func (d transactionVoucherDAO) Insert(ctx context.Context, transactionVouchers entity.TransactionVouchers) error {
	if len(transactionVouchers) < 1 {
		return fmt.Errorf("empty transaction voucher(s) data")
	}

	sqlInsert := sqlgo.NewSQLGoInsert()
	sqlInsert.SetSQLInsert("transaction_voucher")
	sqlInsert.SetSQLInsertColumn("id", "transaction_id", "voucher_id", "quantity", "subtotal_points", "created_at")
	for i, transactionVoucher := range transactionVouchers {
		transactionVoucher.ID = util.MakeUUIDv4()
		transactionVoucher.CreatedAt = time.Now()
		sqlInsert.SetSQLInsertValue(transactionVoucher.ID, transactionVoucher.TransactionID, transactionVoucher.VoucherID, transactionVoucher.Quantity, transactionVoucher.SubtotalPoints, transactionVoucher.CreatedAt)
		transactionVouchers[i] = transactionVoucher
	}
	sql := sqlgo.NewSQLGo().
		SetSQLSchema("public").
		SetSQLGoInsert(sqlInsert)

	sqlStr := sql.BuildSQL()
	sqlParams := sql.GetSQLGoParameter().GetSQLParameter()
	_, err := d.dbTrx.GetSqlTx().ExecContext(ctx, sqlStr, sqlParams...)
	if err != nil {
		return err
	}
	return nil
}

func (d transactionVoucherDAO) Update(ctx context.Context, transactionVouchers entity.TransactionVouchers) error {
	if len(transactionVouchers) < 1 {
		return fmt.Errorf("empty transaction voucher(s) data")
	}

	for i, transactionVoucher := range transactionVouchers {
		updatedAt := time.Now()
		transactionVoucher.UpdatedAt = &updatedAt
		sql := sqlgo.NewSQLGo().
			SetSQLSchema("public").
			SetSQLUpdate("transaction_voucher").
			SetSQLUpdateValue("quantity", transactionVoucher.Quantity).
			SetSQLUpdateValue("subtotal_points", transactionVoucher.SubtotalPoints).
			SetSQLUpdateValue("updated_at", transactionVoucher.UpdatedAt).
			SetSQLWhere("AND", "id", "=", transactionVoucher.ID)
		sqlStr := sql.BuildSQL()
		sqlParams := sql.GetSQLGoParameter().GetSQLParameter()
		_, err := d.dbTrx.GetSqlTx().ExecContext(ctx, sqlStr, sqlParams...)
		if err != nil {
			return err
		}

		transactionVouchers[i] = transactionVoucher
	}

	return nil
}

func (d transactionVoucherDAO) Delete(ctx context.Context, id string) error {
	sql := sqlgo.NewSQLGo()
	sql.SetSQLSchema("public")
	sql.SetSQLDelete("transaction_voucher")
	sql.SetSQLWhere("AND", "id", "IN", id)
	sqlStr := sql.BuildSQL()
	sqlParams := sql.GetSQLGoParameter().GetSQLParameter()

	if _, err := d.dbTrx.GetSqlTx().ExecContext(ctx, sqlStr, sqlParams...); err != nil {
		return err
	}

	return nil
}
