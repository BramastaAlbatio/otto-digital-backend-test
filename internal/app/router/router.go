package router

import (
	brandAdapter "otto-digital-backend-test/internal/app/app_brand/adapter"
	brandService "otto-digital-backend-test/internal/app/app_brand/service"
	customerAdapter "otto-digital-backend-test/internal/app/app_customer/adapter"
	customerService "otto-digital-backend-test/internal/app/app_customer/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	brandService    brandService.BrandService
	brandAdapter    brandAdapter.BrandAdapter
	customerService customerService.CustomerService
	customerAdapter customerAdapter.CustomerAdapter
}

func MakeRouter(
	brandService brandService.BrandService,
	customerService customerService.CustomerService) Router {
	return Router{
		brandService:    brandService,
		brandAdapter:    brandAdapter.MakeBrandAdapter(brandService),
		customerService: customerService,
		customerAdapter: customerAdapter.MakeCustomerAdapter(customerService),
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

	return e
}
