package entities

import (
	"encoding/json"
	"io"
)

type LoginEntity struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (w *LoginEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}

type LoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   error  `json:"errormessage"`
}
type LoginEntityReq struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
