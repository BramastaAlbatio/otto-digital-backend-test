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

		TransactionVouchers TransactionVouchers `json:"transactionVouchers"`
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

func (transaction Transaction) GetTransactionVouchers() TransactionVouchers {
	var transactionVouchers TransactionVouchers
	for _, tv := range transaction.TransactionVouchers {
		tv.TransactionID = transaction.ID

		transactionVouchers = append(transactionVouchers, tv)
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

	for i, tv := range t.TransactionVouchers {
		point, exists := voucherMap[tv.VoucherID]
		if !exists {
			return fmt.Errorf("voucher ID %s not found", tv.VoucherID)
		}
		t.TransactionVouchers[i].SubtotalPoint = tv.Quantity * point
		totalPoints += t.TransactionVouchers[i].SubtotalPoint
	}

	t.TotalPoints = totalPoints
	return nil
}
