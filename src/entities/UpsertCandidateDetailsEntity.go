package entities

import (
	"encoding/json"
	"io"
)

type CandidateEntity struct {
	CandidateName        string `json:"candidatename"`
	Email                string `json:"email"`
	MobileNumber         string `json:"mobno"`
	CurreantEmployeement string `json:"current_emplyeement"`
	Location             string `json:"location"`
	Skills               string `json:"skills"`
	CandidateExperienece string `json:"yoe"`
	JobId                int64  `json:"jobid"`
}
type CandidateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (w *CandidateEntity) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
