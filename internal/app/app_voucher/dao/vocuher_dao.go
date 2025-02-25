package dao

import (
	"context"
	"fmt"
	"time"

	entity "otto-digital-backend-test/pkg/entity"

	"otto-digital-backend-test/pkg/util"

	"gitlab.com/threetopia/sqlgo/v2"
)

type VoucherDAO interface {
	Search(ctx context.Context, query entity.VoucherQuery) (entity.Vouchers, error)
	Insert(ctx context.Context, vouchers entity.Vouchers) error
	Update(ctx context.Context, vouchers entity.Vouchers) error
	Delete(ctx context.Context, id string) error
}

type voucherDAO struct {
	dbTrx DBTransaction
}

func MakeVoucherDAO(dbTrx DBTransaction) VoucherDAO {
	return voucherDAO{
		dbTrx: dbTrx,
	}
}

func (d voucherDAO) Search(ctx context.Context, query entity.VoucherQuery) (entity.Vouchers, error) {
	sqlSelect := sqlgo.NewSQLGoSelect().
		SetSQLSelect("v.id", "id").
		SetSQLSelect("v.brand_id", "brand_id").
		SetSQLSelect("v.name", "name").
		SetSQLSelect("v.cost_in_point", "cost_in_point").
		SetSQLSelect("v.created_at", "created_at").
		SetSQLSelect("v.updated_at", "updated_at")

	sqlFrom := sqlgo.NewSQLGoFrom().
		SetSQLFrom(`"voucher"`, "v")

	sqlWhere := sqlgo.NewSQLGoWhere()
	if len(query.IDs) > 0 {
		sqlWhere.SetSQLWhere("AND", "v.id", "IN", query.IDs)
	}
	if len(query.Names) > 0 {
		sqlWhere.SetSQLWhere("AND", "v.name", "IN", query.Names)
	}
	if len(query.BrandIDs) > 0 {
		sqlWhere.SetSQLWhere("AND", "v.brand_id", "IN", query.BrandIDs)
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

	var Vouchers entity.Vouchers
	for rows.Next() {
		var Voucher entity.Voucher
		if err := rows.Scan(
			&Voucher.ID,
			&Voucher.BrandID,
			&Voucher.Name,
			&Voucher.CostInPoint,
			&Voucher.CreatedAt,
			&Voucher.UpdatedAt,
		); err != nil {
			return nil, err
		}
		Vouchers = append(Vouchers, Voucher)
	}
	return Vouchers, nil
}

func (d voucherDAO) Insert(ctx context.Context, vouchers entity.Vouchers) error {
	if len(vouchers) < 1 {
		return fmt.Errorf("empty Customer(s) data")
	}

	sqlInsert := sqlgo.NewSQLGoInsert()
	sqlInsert.SetSQLInsert("voucher")
	sqlInsert.SetSQLInsertColumn("id", "brand_id", "name", "cost_in_point", "created_at")
	for i, voucher := range vouchers {
		voucher.ID = util.MakeUUIDv4()
		voucher.CreatedAt = time.Now()
		sqlInsert.SetSQLInsertValue(voucher.ID, voucher.BrandID, voucher.Name, voucher.CostInPoint, voucher.CreatedAt)
		vouchers[i] = voucher
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

func (d voucherDAO) Update(ctx context.Context, vouchers entity.Vouchers) error {
	if len(vouchers) < 1 {
		return fmt.Errorf("empty user(s) data")
	}

	for i, voucher := range vouchers {
		updatedAt := time.Now()
		voucher.UpdatedAt = &updatedAt
		sql := sqlgo.NewSQLGo().
			SetSQLSchema("public").
			SetSQLUpdate("voucher").
			SetSQLUpdateValue("name", voucher.Name).
			SetSQLUpdateValue("cost_in_point", voucher.CostInPoint).
			SetSQLUpdateValue("updated_at", voucher.UpdatedAt).
			SetSQLWhere("AND", "id", "=", voucher.ID)
		sqlStr := sql.BuildSQL()
		sqlParams := sql.GetSQLGoParameter().GetSQLParameter()
		_, err := d.dbTrx.GetSqlTx().ExecContext(ctx, sqlStr, sqlParams...)
		if err != nil {
			return err
		}

		vouchers[i] = voucher
	}

	return nil
}

func (d voucherDAO) Delete(ctx context.Context, id string) error {
	sql := sqlgo.NewSQLGo()
	sql.SetSQLSchema("public")
	sql.SetSQLDelete("voucher")
	sql.SetSQLWhere("AND", "id", "IN", id)
	sqlStr := sql.BuildSQL()
	sqlParams := sql.GetSQLGoParameter().GetSQLParameter()

	if _, err := d.dbTrx.GetSqlTx().ExecContext(ctx, sqlStr, sqlParams...); err != nil {
		return err
	}

	return nil
}
