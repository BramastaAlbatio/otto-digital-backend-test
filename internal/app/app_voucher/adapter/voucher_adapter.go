package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"

	voucherService "otto-digital-backend-test/internal/app/app_voucher/service"
	entity "otto-digital-backend-test/pkg/entity"
)

type VoucherAdapter struct {
	voucherService voucherService.VoucherService
}

func MakeVoucherAdapter(voucherService voucherService.VoucherService) VoucherAdapter {
	return VoucherAdapter{
		voucherService: voucherService,
	}
}

func (h *VoucherAdapter) SearchVoucher(c echo.Context) error {
	var query entity.VoucherQuery
	if err := c.Bind(&query); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	vouchers, err := h.voucherService.Search(c.Request().Context(), query)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, vouchers)
}

func (h *VoucherAdapter) InsertVouchers(c echo.Context) error {
	var voucher entity.Vouchers
	if err := c.Bind(&voucher); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.voucherService.Insert(c.Request().Context(), voucher); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Vouchers inserted successfully"})
}

func (h *VoucherAdapter) InsertVoucher(c echo.Context) error {
	var voucher entity.Voucher
	if err := c.Bind(&voucher); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.voucherService.Insert(c.Request().Context(), entity.Vouchers{voucher}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Voucher inserted successfully"})
}

func (h *VoucherAdapter) UpdateVouchers(c echo.Context) error {
	var vouchers entity.Vouchers
	if err := c.Bind(&vouchers); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.voucherService.Update(c.Request().Context(), vouchers); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Vouchers updated successfully"})
}

func (h *VoucherAdapter) UpdateVoucher(c echo.Context) error {
	var voucher entity.Voucher
	if err := c.Bind(&voucher); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.voucherService.Update(c.Request().Context(), entity.Vouchers{voucher}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Voucher updated successfully"})
}

func (h *VoucherAdapter) DeleteVoucher(c echo.Context) error {
	if err := h.voucherService.Delete(c.Request().Context(), c.Param("id")); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Voucher deleted successfully"})
}
