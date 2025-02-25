package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"

	transactionVoucherService "otto-digital-backend-test/internal/app/app_transaction_voucher/service"
	entity "otto-digital-backend-test/pkg/entity"
)

type TransactionVoucherAdapter struct {
	transactionVoucherService transactionVoucherService.TransactionVoucherService
}

func MakeTransactionVoucherAdapter(transactionVoucherService transactionVoucherService.TransactionVoucherService) TransactionVoucherAdapter {
	return TransactionVoucherAdapter{
		transactionVoucherService: transactionVoucherService,
	}
}

func (h *TransactionVoucherAdapter) SearchTransactionVoucher(c echo.Context) error {
	var query entity.TransactionVoucherQuery
	if err := c.Bind(&query); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	transactionVouchers, err := h.transactionVoucherService.Search(c.Request().Context(), query)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, transactionVouchers)
}

func (h *TransactionVoucherAdapter) InsertTransactionVouchers(c echo.Context) error {
	var transactionVouchers entity.TransactionVouchers
	if err := c.Bind(&transactionVouchers); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.transactionVoucherService.Insert(c.Request().Context(), transactionVouchers); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Transaction Vouchers inserted successfully"})
}

func (h *TransactionVoucherAdapter) InsertTransactionVoucher(c echo.Context) error {
	var transactionVoucher entity.TransactionVoucher
	if err := c.Bind(&transactionVoucher); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.transactionVoucherService.Insert(c.Request().Context(), entity.TransactionVouchers{transactionVoucher}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "TransactionVoucher inserted successfully"})
}

func (h *TransactionVoucherAdapter) UpdateTransactionVouchers(c echo.Context) error {
	var transactionVouchers entity.TransactionVouchers
	if err := c.Bind(&transactionVouchers); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.transactionVoucherService.Update(c.Request().Context(), transactionVouchers); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Transaction Vouchers updated successfully"})
}

func (h *TransactionVoucherAdapter) UpdateTransactionVoucher(c echo.Context) error {
	var transactionVoucher entity.TransactionVoucher
	if err := c.Bind(&transactionVoucher); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.transactionVoucherService.Update(c.Request().Context(), entity.TransactionVouchers{transactionVoucher}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Transaction Voucher updated successfully"})
}

func (h *TransactionVoucherAdapter) DeleteTransactionVoucher(c echo.Context) error {
	if err := h.transactionVoucherService.Delete(c.Request().Context(), c.Param("id")); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Transaction Voucher deleted successfully"})
}
