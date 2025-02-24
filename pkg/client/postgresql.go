package client

import (
	"database/sql"
	"fmt"
	"log"
	"otto-digital-backend-test/pkg/entity"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"gitlab.com/threetopia/envgo"
)

type PostgreSQLClient interface {
	GetSQLDB() *sql.DB
	Migration() error
}

type postgreSQLClient struct {
	sqlDB *sql.DB
	dsn   entity.DSNEntity
}

func MakePostgreSQLClient(dsn entity.DSNEntity) PostgreSQLClient {
	return postgreSQLClient{
		dsn:   dsn,
		sqlDB: dbConn(dsn),
	}
}

func dbConn(dsn entity.DSNEntity) *sql.DB {
	sqlDB, err := sql.Open("postgres", dsn.GetPostgresParam())
	if err != nil {
		panic(err.Error())
	}
	return sqlDB
}

func (c postgreSQLClient) GetSQLDB() *sql.DB {
	return c.sqlDB
}

func (c postgreSQLClient) Migration() error {
	log.Printf("DSN %s", c.dsn.GetPostgresParam())

	driver, err := postgres.WithInstance(c.sqlDB, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", envgo.GetString("MIGRATION_PATH", "/home/bramasta/Project/go/otto-digital-backend-test/pkg/scripts")),
		"postgres", driver)

	if err != nil {
		log.Fatalf("Error %s", err.Error())
		return err
	}

	m.Up()
	return nil
}
