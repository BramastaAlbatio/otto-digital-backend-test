package router

import (
	brandAdapter "otto-digital-backend-test/internal/app/app_brand/adapter"
	brandService "otto-digital-backend-test/internal/app/app_brand/service"
	customerAdapter "otto-digital-backend-test/internal/app/app_customer/adapter"
	customerService "otto-digital-backend-test/internal/app/app_customer/service"
	transactionAdapter "otto-digital-backend-test/internal/app/app_transaction/adapter"
	transactionService "otto-digital-backend-test/internal/app/app_transaction/service"
	transactionVoucherAdapter "otto-digital-backend-test/internal/app/app_transaction_voucher/adapter"
	transactionVoucherService "otto-digital-backend-test/internal/app/app_transaction_voucher/service"
	voucherAdapter "otto-digital-backend-test/internal/app/app_voucher/adapter"
	voucherService "otto-digital-backend-test/internal/app/app_voucher/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	brandService              brandService.BrandService
	brandAdapter              brandAdapter.BrandAdapter
	customerService           customerService.CustomerService
	customerAdapter           customerAdapter.CustomerAdapter
	voucherService            voucherService.VoucherService
	voucherAdapter            voucherAdapter.VoucherAdapter
	transactionService        transactionService.TransactionService
	transactionAdapter        transactionAdapter.TransactionAdapter
	transactionVoucherService transactionVoucherService.TransactionVoucherService
	transactionVoucherAdapter transactionVoucherAdapter.TransactionVoucherAdapter
}

func MakeRouter(
	brandService brandService.BrandService,
	customerService customerService.CustomerService,
	voucherService voucherService.VoucherService,
	transactionService transactionService.TransactionService,
	transactionVoucherService transactionVoucherService.TransactionVoucherService) Router {
	return Router{
		brandService:              brandService,
		brandAdapter:              brandAdapter.MakeBrandAdapter(brandService),
		customerService:           customerService,
		customerAdapter:           customerAdapter.MakeCustomerAdapter(customerService),
		voucherService:            voucherService,
		voucherAdapter:            voucherAdapter.MakeVoucherAdapter(voucherService),
		transactionService:        transactionService,
		transactionAdapter:        transactionAdapter.MakeTransactionAdapter(transactionService),
		transactionVoucherService: transactionVoucherService,
		transactionVoucherAdapter: transactionVoucherAdapter.MakeTransactionVoucherAdapter(transactionVoucherService),
	}
}

func (r Router) InitRouter() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS Middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	// Brand Handler
	e.GET("/brands", r.brandAdapter.SearchBrand)
	e.POST("/brands", r.brandAdapter.InsertBrands)
	e.POST("/brand", r.brandAdapter.InsertBrand)
	e.PUT("/brands", r.brandAdapter.UpdateBrands)
	e.PUT("/brand", r.brandAdapter.UpdateBrand)
	e.DELETE("/brands", r.brandAdapter.DeleteBrands)

	// Customer Handler
	e.GET("/customer", r.customerAdapter.SearchCustomer)
	e.POST("/customers", r.customerAdapter.InsertCustomers)
	e.POST("/customer", r.customerAdapter.InsertCustomer)
	e.PUT("/customers", r.customerAdapter.UpdateCustomers)
	e.PUT("/customer", r.customerAdapter.UpdateCustomer)
	e.DELETE("/customers", r.customerAdapter.DeleteCustomer)

	// Voucher Handler
	e.GET("/voucher", r.voucherAdapter.SearchVoucher)
	e.POST("/vouchers", r.voucherAdapter.InsertVouchers)
	e.POST("/voucher", r.voucherAdapter.InsertVoucher)
	e.PUT("/vouchers", r.voucherAdapter.UpdateVouchers)
	e.PUT("/voucher", r.voucherAdapter.UpdateVoucher)
	e.DELETE("/vouchers", r.voucherAdapter.DeleteVoucher)

	// Transaction Handler
	e.GET("/transaction", r.transactionAdapter.SearchTransaction)
	e.POST("/transactions", r.transactionAdapter.InsertTransactions)
	e.POST("/transaction", r.transactionAdapter.InsertTransaction)
	e.PUT("/transactions", r.transactionAdapter.UpdateTransactions)
	e.PUT("/transaction", r.transactionAdapter.UpdateTransaction)
	e.DELETE("/transactions", r.transactionAdapter.DeleteTransaction)

	// Transaction Voucher Handler
	e.GET("/transaction-voucher", r.transactionVoucherAdapter.SearchTransactionVoucher)
	e.POST("/transaction-vouchers", r.transactionVoucherAdapter.InsertTransactionVouchers)
	e.POST("/transaction-voucher", r.transactionVoucherAdapter.InsertTransactionVoucher)
	e.PUT("/transaction-vouchers", r.transactionVoucherAdapter.UpdateTransactionVouchers)
	e.PUT("/transaction-voucher", r.transactionVoucherAdapter.UpdateTransactionVoucher)
	e.DELETE("/transaction-vouchers", r.transactionVoucherAdapter.DeleteTransactionVoucher)

	return e
}
