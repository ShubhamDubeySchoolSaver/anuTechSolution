package entities

import (
	"encoding/json"
	"io"
)

type GetInTouchEntity struct {
	Id           int64  `json:"id"`
	UserName     string `json:"username"`
	UserEmail    string `json:"email"`
	MobileNumber string `json:"mobno"`
	Message      string `json:"message"`
}

func (w *GetInTouchEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}

type GetInTouchResponse struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Details []GetInTouchEntity `json:"details"`
}
type JobCandidatesEntity struct {
	Id                   string `json:"id"`
	CandidateName        string `json:"candidatename"`
	Email                string `json:"email"`
	MobileNumber         string `json:"mobno"`
	CurreantEmployeement string `json:"current_emplyeement"`
	Location             string `json:"location"`
	Skills               string `json:"skills"`
	CandidateExperienece string `json:"yoe"`
	UploadedFileName     string `json:"filename"`
	AppliedDateAndTime   string `json:"applieddateandtime"`
	AppliedDesignation   string `json:"applieddesignation"`
	AppliedLocation      string `json:"appliedlocation"`
}
type JobCandidatesResponse struct {
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	Details []JobCandidatesEntity `json:"details"`
}
type ConsultationEntity struct {
	Id           int64  `json:"id"`
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	MobileNumber string `json:"mob_no"`
	Service      string `json:"service"`
	Pricerange   string `json:"price_range"`
	Location     string `json:"location"`
	AboutProject string `json:"about_project"`
	DateAndTime  string `json:"dateandtime"`
}
type ConsultationResponse struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Details []ConsultationEntity `json:"details"`
}
type FetchJobEntity struct {
	Id          int64  `json:"id"`
	Designation string `json:"designation"`
	Experience  int64  `json:"experience"`
	Location    string `json:"location"`
}
type FetchJobEntityResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Details []FetchJobEntity `json:"details"`
}
