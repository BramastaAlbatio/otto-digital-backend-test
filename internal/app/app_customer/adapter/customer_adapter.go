package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"

	customerService "otto-digital-backend-test/internal/app/app_customer/service"
	entity "otto-digital-backend-test/pkg/entity"
)

type CustomerAdapter struct {
	customerService customerService.CustomerService
}

func MakeCustomerAdapter(customerService customerService.CustomerService) CustomerAdapter {
	return CustomerAdapter{
		customerService: customerService,
	}
}

func (h *CustomerAdapter) SearchCustomer(c echo.Context) error {
	var query entity.CustomerQuery
	if err := c.Bind(&query); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	brands, err := h.customerService.Search(c.Request().Context(), query)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, brands)
}

func (h *CustomerAdapter) InsertCustomers(c echo.Context) error {
	var customer entity.Customers
	if err := c.Bind(&customer); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.customerService.Insert(c.Request().Context(), customer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Customers inserted successfully"})
}

func (h *CustomerAdapter) InsertCustomer(c echo.Context) error {
	var customer entity.Customer
	if err := c.Bind(&customer); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.customerService.Insert(c.Request().Context(), entity.Customers{customer}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Customer inserted successfully"})
}

func (h *CustomerAdapter) UpdateCustomers(c echo.Context) error {
	var customers entity.Customers
	if err := c.Bind(&customers); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.customerService.Update(c.Request().Context(), customers); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Customers updated successfully"})
}

func (h *CustomerAdapter) UpdateCustomer(c echo.Context) error {
	var customer entity.Customer
	if err := c.Bind(&customer); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.customerService.Update(c.Request().Context(), entity.Customers{customer}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Customer updated successfully"})
}

func (h *CustomerAdapter) DeleteCustomer(c echo.Context) error {
	if err := h.customerService.Delete(c.Request().Context(), c.Param("id")); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Customer deleted successfully"})
}
