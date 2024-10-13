package entities

import (
	"encoding/json"
	"io"
)

type UserEntity struct {
	UserName     string `json:"username"`
	UserEmail    string `json:"email"`
	MobileNumber string `json:"mobno"`
	Message      string `json:"message"`
}
type UserEntityResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *UserEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}

type JobEntity struct {
	Designation string `json:"designation"`
	Experience  int64  `json:"experience"`
	JobLocation string `json:"joblocation"`
}

func (w *JobEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}

type JobEntityResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}
