package entity

import "time"

type (
	BrandQuery struct {
		IDs   []string `json:"ids"`
		Names []string `json:"name"`
	}

	Brand struct {
		ID        string     `json:"id"`
		Name      string     `json:"name"`
		CreatedAt time.Time  `json:"createdAt"`
		UpdatedAt *time.Time `json:"updatedAt"`
	}
	Brands []Brand
)

func (brands Brands) GetIds() []string {
	var ids []string
	for _, brand := range brands {
		ids = append(ids, brand.ID)
	}
	return ids
}
