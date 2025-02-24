package dao

import (
	"context"
	"fmt"
	"time"

	entity "otto-digital-backend-test/pkg/entity"

	"otto-digital-backend-test/pkg/util"

	"gitlab.com/threetopia/sqlgo/v2"
)

type BrandDAO interface {
	Search(ctx context.Context, query entity.BrandQuery) (entity.Brands, error)
	Insert(ctx context.Context, brands entity.Brands) error
	Update(ctx context.Context, brands entity.Brands) error
	Delete(ctx context.Context, id string) error
}

type brandDAO struct {
	dbTrx DBTransaction
}

func MakeBrandDAO(dbTrx DBTransaction) BrandDAO {
	return brandDAO{
		dbTrx: dbTrx,
	}
}

func (d brandDAO) Search(ctx context.Context, query entity.BrandQuery) (entity.Brands, error) {
	sqlSelect := sqlgo.NewSQLGoSelect().
		SetSQLSelect("u.id", "id").
		SetSQLSelect("u.name", "name").
		SetSQLSelect("u.created_at", "created_at").
		SetSQLSelect("u.updated_at", "updated_at")

	sqlFrom := sqlgo.NewSQLGoFrom().
		SetSQLFrom(`"brand"`, "b")

	sqlWhere := sqlgo.NewSQLGoWhere()
	if len(query.IDs) > 0 {
		sqlWhere.SetSQLWhere("AND", "b.id", "IN", query.IDs)
	}
	if len(query.Names) > 0 {
		sqlWhere.SetSQLWhere("AND", "b.name", "IN", query.Names)
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

	var Brands entity.Brands
	for rows.Next() {
		var brand entity.Brand
		if err := rows.Scan(
			&brand.ID,
			&brand.Name,
			&brand.CreatedAt,
			&brand.UpdatedAt,
		); err != nil {
			return nil, err
		}
		Brands = append(Brands, brand)
	}
	return Brands, nil
}

func (d brandDAO) Insert(ctx context.Context, brands entity.Brands) error {
	if len(brands) < 1 {
		return fmt.Errorf("empty brand(s) data")
	}

	sqlInsert := sqlgo.NewSQLGoInsert()
	sqlInsert.SetSQLInsert("brand")
	sqlInsert.SetSQLInsertColumn("id", "name", "created_at")
	for i, user := range brands {
		user.ID = util.MakeUUIDv4()
		user.CreatedAt = time.Now()
		sqlInsert.SetSQLInsertValue(user.ID, user.Name, user.CreatedAt)
		brands[i] = user
	}
	sql := sqlgo.NewSQLGo().
		SetSQLSchema("public").
		SetSQLGoInsert(sqlInsert)

	sqlStr := sql.BuildSQL()
	sqlParams := sql.GetSQLGoParameter().GetSQLParameter()
	_, err := d.dbTrx.GetSqlTx().ExecContext(ctx, sqlStr, sqlParams...)
	fmt.Println("=====================", sqlStr, sqlParams)
	if err != nil {
		return err
	}
	return nil
}

func (d brandDAO) Update(ctx context.Context, brands entity.Brands) error {
	if len(brands) < 1 {
		return fmt.Errorf("empty user(s) data")
	}

	for i, brand := range brands {
		updatedAt := time.Now()
		brand.UpdatedAt = &updatedAt
		sql := sqlgo.NewSQLGo().
			SetSQLSchema("public").
			SetSQLUpdate("brand").
			SetSQLUpdateValue("name", brand.Name).
			SetSQLUpdateValue("updated_at", brand.UpdatedAt).
			SetSQLWhere("AND", "id", "=", brand.ID)
		sqlStr := sql.BuildSQL()
		sqlParams := sql.GetSQLGoParameter().GetSQLParameter()
		_, err := d.dbTrx.GetSqlTx().ExecContext(ctx, sqlStr, sqlParams...)
		if err != nil {
			return err
		}

		brands[i] = brand
	}

	return nil
}

func (d brandDAO) Delete(ctx context.Context, id string) error {
	sql := sqlgo.NewSQLGo()
	sql.SetSQLSchema("public")
	sql.SetSQLDelete("brand")
	sql.SetSQLWhere("AND", "id", "IN", id)
	sqlStr := sql.BuildSQL()
	sqlParams := sql.GetSQLGoParameter().GetSQLParameter()

	if _, err := d.dbTrx.GetSqlTx().ExecContext(ctx, sqlStr, sqlParams...); err != nil {
		return err
	}

	return nil
}
