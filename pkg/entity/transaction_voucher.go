package entity

import "time"

type (
	TransactionVoucherQuery struct {
		IDs []string `query:"id"`
	}

	TransactionVoucher struct {
		ID            string     `json:"id"`
		TransactionID string     `json:"transactionId"`
		VoucherID     string     `json:"voucherId"`
		Quantity      int        `json:"quantity"`
		SubtotalPoint int        `json:"subtotalPoint"`
		CreatedAt     time.Time  `json:"createdAt"`
		UpdatedAt     *time.Time `json:"updatedAt"`
	}

	TransactionVouchers []TransactionVoucher
)

func (transactionVoucers TransactionVouchers) GetIDs() []string {
	var ids []string

	for _, trtransactionVoucer := range transactionVoucers {
		ids = append(ids, trtransactionVoucer.ID)
	}
	return ids
}

func (transactionVouchers TransactionVouchers) GetTransactionIDS() []string {
	var ids []string

	for _, transactionVoucher := range transactionVouchers {
		ids = append(ids, transactionVoucher.TransactionID)
	}

	return ids
}
