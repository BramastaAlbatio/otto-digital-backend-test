package entity

import "time"

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
