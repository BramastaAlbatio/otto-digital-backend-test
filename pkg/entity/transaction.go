package entity

import (
	"fmt"
	"time"
)

type (
	TransactionQuery struct {
		IDs []string `query:"id"`
	}

	Transaction struct {
		ID          string     `json:"id"`
		CustomerID  string     `json:"customerId"`
		TotalPoints int        `json:"totalPoint"`
		CreatedAt   time.Time  `json:"createdAt"`
		UpdatedAt   *time.Time `json:"updatedAt"`

		TransactionVoucher TransactionVoucher `json:"transactionVouchers"`
	}

	Transactions []Transaction
)

func (transactions Transactions) GetIDs() []string {
	var ids []string

	for _, transaction := range transactions {
		ids = append(ids, transaction.ID)
	}
	return ids
}

func (transactions Transactions) GetTransactionVouchers() TransactionVouchers {
	var transactionVouchers TransactionVouchers
	for _, transaction := range transactions {
		transaction.TransactionVoucher.TransactionID = transaction.ID
		transactionVouchers = append(transactionVouchers, transaction.TransactionVoucher)
	}

	return transactionVouchers
}

func (t *Transaction) CalculateTotalPoints(vouchers Vouchers) error {
	var totalPoints int
	voucherMap := make(map[string]int)

	// Buat map voucher agar lookup lebih cepat
	for _, v := range vouchers {
		voucherMap[v.ID] = v.CostInPoint
	}

	// Pastikan TransactionVoucher memiliki voucher ID yang valid
	point, exists := voucherMap[t.TransactionVoucher.VoucherID]
	if !exists {
		return fmt.Errorf("voucher ID %s not found", t.TransactionVoucher.VoucherID)
	}

	// Hitung subtotal dan total points
	t.TransactionVoucher.SubtotalPoint = t.TransactionVoucher.Quantity * point
	totalPoints += t.TransactionVoucher.SubtotalPoint

	t.TotalPoints = totalPoints
	return nil
}
