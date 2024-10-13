package dao

import (
	"Skool_Saver/src/entities"
	"Skool_Saver/src/logger"
	"errors"
)

func (dbc DbConn) FetchGetInTouchDao(limit, offset int64) ([]entities.GetInTouchEntity, error) {
	var response []entities.GetInTouchEntity
	query := "SELECT user_id, user_name,user_email,user_mobileno,user_message FROM userdetails where deleteflag=0 limit ? offset ?"
	rows, err := dbc.DB.Query(query, limit, offset)
	if err != nil {
		logger.Log.Println("FetchGetInTouchDao Query Execution error :", err)
		return response, errors.New("fetchGetInTouchDao Query Execution error")
	}
	defer rows.Close()
	for rows.Next() {
		var data entities.GetInTouchEntity
		Err := rows.Scan(&data.Id, &data.UserName, &data.UserEmail, &data.MobileNumber, &data.Message)
		if Err != nil {
			logger.Log.Println("fetchGetInTouchDao Row Scan Error", err)
			return nil, err
		}
		response = append(response, data)
	}
	return response, nil

}
func (dbc DbConn) FetchJobAppliedCandidatesDao(limit, offset int64) ([]entities.JobCandidatesEntity, error) {
	var response []entities.JobCandidatesEntity
	query := "SELECT a.id,a.candidatename,a.email,a.mobno,a.current_emplyeement,a.location,a.skills,a.yearofexp,a.filename,a.applieddateandtime,b.designation,b.location FROM jobappliedcandidates a,jobs b where a.deleteflag=0 and b.deleteflag=0 and a.jobid=b.id limit ? offset ?"
	rows, err := dbc.DB.Query(query, limit, offset)
	if err != nil {
		logger.Log.Println("FetchJobAppliedCandidatesDao Query Execution error :", err)
		return response, errors.New("FetchJobAppliedCandidatesDao Query Execution error")
	}
	defer rows.Close()
	for rows.Next() {
		var data entities.JobCandidatesEntity
		Err := rows.Scan(&data.Id, &data.CandidateName, &data.Email, &data.MobileNumber, &data.CurreantEmployeement, &data.Location, &data.Skills, &data.CandidateExperienece, &data.UploadedFileName, &data.AppliedDateAndTime, &data.AppliedDesignation, &data.AppliedLocation)
		if Err != nil {
			logger.Log.Println("FetchJobAppliedCandidatesDao Row Scan Error", err)
			return nil, err
		}
		response = append(response, data)
	}
	return response, nil

}
func (dbc DbConn) FetchConsultationsDao(limit, offset int64) ([]entities.ConsultationEntity, error) {
	var response []entities.ConsultationEntity
	query := "SELECT id,fullname,email,mobilenumber,service,pricerange,location,aboutproject,upserteddateandtime FROM consultations where deleteflag=0 limit ? offset ?"
	rows, err := dbc.DB.Query(query, limit, offset)
	if err != nil {
		logger.Log.Println("FetchConsultationsDao Query Execution error :", err)
		return response, errors.New("FetchConsultationsDao Query Execution error")
	}
	defer rows.Close()
	for rows.Next() {
		var data entities.ConsultationEntity
		Err := rows.Scan(&data.Id, &data.FullName, &data.Email, &data.MobileNumber, &data.Service, &data.Pricerange, &data.Location, &data.AboutProject, &data.DateAndTime)
		if Err != nil {
			logger.Log.Println("FetchConsultationsDao Row Scan Error", err)
			return nil, err
		}
		response = append(response, data)
	}
	return response, nil

}
func (dbc DbConn) FetchJobsDao(limit, offset int64) ([]entities.FetchJobEntity, error) {
	var response []entities.FetchJobEntity
	query := "SELECT id,designation,experience,location FROM jobs where deleteflag=0 limit ? offset ?;"
	rows, err := dbc.DB.Query(query, limit, offset)
	if err != nil {
		logger.Log.Println("FetchJobsDao Query Execution error :", err)
		return response, errors.New("FetchJobsDao Query Execution error")
	}
	defer rows.Close()
	for rows.Next() {
		var data entities.FetchJobEntity
		Err := rows.Scan(&data.Id, &data.Designation, &data.Experience, &data.Location)
		if Err != nil {
			logger.Log.Println("FetchJobsDao Row Scan Error", err)
			return nil, err
		}
		response = append(response, data)
	}
	return response, nil

}
