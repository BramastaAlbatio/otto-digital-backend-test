package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"

	transactionService "otto-digital-backend-test/internal/app/app_transaction/service"
	entity "otto-digital-backend-test/pkg/entity"
)

type TransactionAdapter struct {
	transactionService transactionService.TransactionService
}

func MakeTransactionAdapter(transactionService transactionService.TransactionService) TransactionAdapter {
	return TransactionAdapter{
		transactionService: transactionService,
	}
}

func (h *TransactionAdapter) SearchTransaction(c echo.Context) error {
	var query entity.TransactionQuery
	if err := c.Bind(&query); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	transactions, err := h.transactionService.Search(c.Request().Context(), query)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, transactions)
}

func (h *TransactionAdapter) InsertTransactions(c echo.Context) error {
	var transactions entity.Transactions
	if err := c.Bind(&transactions); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.transactionService.Insert(c.Request().Context(), transactions); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Transactions inserted successfully"})
}

func (h *TransactionAdapter) InsertTransaction(c echo.Context) error {
	var transaction entity.Transaction
	if err := c.Bind(&transaction); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.transactionService.Insert(c.Request().Context(), entity.Transactions{transaction}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Transaction inserted successfully"})
}

func (h *TransactionAdapter) UpdateTransactions(c echo.Context) error {
	var transactions entity.Transactions
	if err := c.Bind(&transactions); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.transactionService.Update(c.Request().Context(), transactions); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Transactions updated successfully"})
}

func (h *TransactionAdapter) UpdateTransaction(c echo.Context) error {
	var transaction entity.Transaction
	if err := c.Bind(&transaction); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.transactionService.Update(c.Request().Context(), entity.Transactions{transaction}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Transaction updated successfully"})
}

func (h *TransactionAdapter) DeleteTransaction(c echo.Context) error {
	if err := h.transactionService.Delete(c.Request().Context(), c.Param("id")); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Transaction deleted successfully"})
}
