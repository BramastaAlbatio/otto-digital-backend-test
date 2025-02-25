package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"otto-digital-backend-test/pkg/client"
	"otto-digital-backend-test/pkg/entity"
	"syscall"
	"time"

	"github.com/joho/godotenv"

	brandService "otto-digital-backend-test/internal/app/app_brand/service"

	customerService "otto-digital-backend-test/internal/app/app_customer/service"

	transactionService "otto-digital-backend-test/internal/app/app_transaction/service"
	voucherService "otto-digital-backend-test/internal/app/app_voucher/service"
	"otto-digital-backend-test/internal/app/router"

	"gitlab.com/threetopia/envgo"
)

func main() {
	godotenv.Load("../../.env")

	port := envgo.GetString("PORT", "6060")
	dsn := entity.DSNEntity{
		Host:     envgo.GetString("DB_HOST", "localhost"),
		User:     envgo.GetString("DB_USER", "bramasta"),
		Password: envgo.GetString("DB_PASS", "bramasta"),
		Port:     envgo.GetInt("DB_PORT", 5432),
		SSLMode:  envgo.GetBool("DB_SSL_MODE", false),
		Database: envgo.GetString("DB_DATABASE", "otto-digital-test"),
		Schema:   envgo.GetString("DB_SCHEMA", "public"),
		TimeZone: envgo.GetString("DB_TZ", "Asia/Jakarta"),
	}

	psqlDB := client.MakePostgreSQLClient(dsn)
	psqlDB.Migration()

	brandSrv := brandService.MakeBrandService(psqlDB.GetSQLDB())
	customerSrv := customerService.MakeCustomerService(psqlDB.GetSQLDB())
	voucherSrv := voucherService.MakeVoucherService(psqlDB.GetSQLDB())
	transactionSrv := transactionService.MakeTransactionService(psqlDB.GetSQLDB())

	router := router.MakeRouter(brandSrv, customerSrv, voucherSrv, transactionSrv)

	newRouter := router.InitRouter()

	server := &http.Server{
		Addr:    ":" + port,
		Handler: newRouter,
	}
	log.Printf("API server listening on %s", server.Addr)

	go func() {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("API server closed: err: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("got shutdown signal. shutting down server...")

	localCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(localCtx); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}

	log.Println("server shutdown complete")
}
