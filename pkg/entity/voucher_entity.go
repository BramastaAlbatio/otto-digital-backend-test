package entity

import "time"

type (
	VoucherQuery struct {
		IDs      []string `query:"id"`
		Names    []string `query:"name"`
		BrandIDs []string `query:"brand_id"`
	}

	Voucher struct {
		ID          string     `json:"id"`
		BrandID     string     `json:"brandId"`
		Name        string     `json:"name"`
		CostInPoint int        `json:"costInPoint"`
		CreatedAt   time.Time  `json:"createdAt"`
		UpdatedAt   *time.Time `json:"updatedAt"`
	}

	Vouchers []Voucher
)

func (vouchers Vouchers) GetIDs() []string {
	var ids []string

	for _, voucher := range vouchers {
		ids = append(ids, voucher.ID)
	}

	return ids
}
