package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"

	brandService "otto-digital-backend-test/internal/app/app_brand/service"
	entity "otto-digital-backend-test/pkg/entity"
)

type BrandAdapter struct {
	brandService brandService.BrandService
}

func MakeBrandAdapter(brandService brandService.BrandService) BrandAdapter {
	return BrandAdapter{
		brandService: brandService,
	}
}

func (h *BrandAdapter) SearchBrand(c echo.Context) error {
	var query entity.BrandQuery
	if err := c.Bind(&query); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	brands, err := h.brandService.Search(c.Request().Context(), query)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, brands)
}

func (h *BrandAdapter) InsertBrands(c echo.Context) error {
	var brands entity.Brands
	if err := c.Bind(&brands); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.brandService.Insert(c.Request().Context(), brands); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Brands inserted successfully"})
}

func (h *BrandAdapter) InsertBrand(c echo.Context) error {
	var brand entity.Brand
	if err := c.Bind(&brand); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.brandService.Insert(c.Request().Context(), entity.Brands{brand}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Brands inserted successfully"})
}

func (h *BrandAdapter) UpdateBrands(c echo.Context) error {
	var brands entity.Brands
	if err := c.Bind(&brands); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.brandService.Update(c.Request().Context(), brands); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Brands updated successfully"})
}

func (h *BrandAdapter) UpdateBrand(c echo.Context) error {
	var brand entity.Brand
	if err := c.Bind(&brand); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.brandService.Update(c.Request().Context(), entity.Brands{brand}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Brands updated successfully"})
}

func (h *BrandAdapter) DeleteBrands(c echo.Context) error {
	if err := h.brandService.Delete(c.Request().Context(), c.Param("id")); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Rents deleted successfully"})
}
