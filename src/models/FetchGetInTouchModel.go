package models

import (
	"Skool_Saver/src/config"
	"Skool_Saver/src/dao"
	"Skool_Saver/src/entities"
	"Skool_Saver/src/logger"
)

func FetchGetInTouchModels(limit, offset int64) ([]entities.GetInTouchEntity, bool) {
	logger.Log.Println("Inside FetchGetInTouchModels")
	var res []entities.GetInTouchEntity
	db, err := config.ConnectMySqlDbSlaveSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return res, false
	}
	dataAccess := dao.DbConn{DB: db}
	data, err := dataAccess.FetchGetInTouchDao(limit, offset)
	if err != nil {
		logger.Log.Println("FetchGetInTouchDao Details Fetch error :", err)
		return res, false
	}
	res = data

	return res, true

}
func FetchJobAppliedCandidatesModels(limit, offset int64) ([]entities.JobCandidatesEntity, bool) {
	logger.Log.Println("Inside FetchJobAppliedCandidatesModels")
	var res []entities.JobCandidatesEntity
	db, err := config.ConnectMySqlDbSlaveSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return res, false
	}
	dataAccess := dao.DbConn{DB: db}
	data, err := dataAccess.FetchJobAppliedCandidatesDao(limit, offset)
	if err != nil {
		logger.Log.Println("FetchJobAppliedCandidatesModels Details Fetch error :", err)
		return res, false
	}
	res = data

	return res, true

}
func FetchConsultationsModels(limit, offset int64) ([]entities.ConsultationEntity, bool) {
	logger.Log.Println("Inside FetchConsultationsModels")
	var res []entities.ConsultationEntity
	db, err := config.ConnectMySqlDbSlaveSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return res, false
	}
	dataAccess := dao.DbConn{DB: db}
	data, err := dataAccess.FetchConsultationsDao(limit, offset)
	if err != nil {
		logger.Log.Println("FetchConsultationsModels Details Fetch error :", err)
		return res, false
	}
	res = data

	return res, true

}
func FetchJobsModels(limit, offset int64) ([]entities.FetchJobEntity, bool) {
	logger.Log.Println("Inside FetchJobsModels")
	var res []entities.FetchJobEntity
	db, err := config.ConnectMySqlDbSlaveSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return res, false
	}
	dataAccess := dao.DbConn{DB: db}
	data, err := dataAccess.FetchJobsDao(limit, offset)
	if err != nil {
		logger.Log.Println("FetchJobsModels Details Fetch error :", err)
		return res, false
	}
	res = data

	return res, true

}
