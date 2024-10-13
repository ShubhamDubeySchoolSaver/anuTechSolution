package entities

import (
	"encoding/json"
	"io"
)

type QuotationEntity struct {
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	MobileNumber string `json:"mob_no"`
	Service      string `json:"service"`
	Pricerange   string `json:"price_range"`
	Location     string `json:"location"`
	AboutProject string `json:"about_project"`
}
type QuotationResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *QuotationEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
