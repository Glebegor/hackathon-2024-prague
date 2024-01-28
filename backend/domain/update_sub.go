package domain

type UpdateSub struct {
	ChangePerson string `json:"change_person`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Description  string `json:"description"`
	Images       string `json:"images"`
	Tags         string `json:"tags"`
	Wallet       string `json:"wallet"`
}
