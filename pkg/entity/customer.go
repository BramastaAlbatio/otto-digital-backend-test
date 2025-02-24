package entity

import "time"

type (
	CustomerQuery struct {
		IDs   []string `query:"id"`
		Names []string `query:"name"`
	}

	Customer struct {
		ID        string     `json:"id"`
		Name      string     `json:"name"`
		Email     string     `json:"email"`
		CreatedAt time.Time  `json:"createdAt"`
		UpdatedAt *time.Time `json:"updatedAt"`
	}

	Customers []Customer
)

func (brands Brands) GetIDs() []string {
	var ids []string
	for _, user := range brands {
		ids = append(ids, user.ID)
	}
	return ids
}
