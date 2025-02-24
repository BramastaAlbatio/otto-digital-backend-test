package router

import (
	BrandAdapter "otto-digital-backend-test/internal/app/app_brand/adapter"
	BrandService "otto-digital-backend-test/internal/app/app_brand/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	brandService BrandService.BrandService
	brandAdapter BrandAdapter.BrandAdapter
}

func MakeRouter(
	brandService BrandService.BrandService) Router {
	return Router{
		brandService: brandService,
		brandAdapter: BrandAdapter.MakeBrandAdapter(brandService),
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

	// User Handler
	e.GET("/brands", r.brandAdapter.SearchBrand)
	e.POST("/brands", r.brandAdapter.InsertBrands)
	e.POST("/brand", r.brandAdapter.InsertBrand)
	e.PUT("/brands", r.brandAdapter.UpdateBrands)
	e.PUT("/brand", r.brandAdapter.UpdateBrand)
	e.DELETE("/brands", r.brandAdapter.DeleteBrands)

	return e
}
