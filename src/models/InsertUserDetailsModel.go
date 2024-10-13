package models

import (
	"Skool_Saver/src/config"
	"Skool_Saver/src/dao"
	"Skool_Saver/src/entities"
	"Skool_Saver/src/logger"
)

func InsertUserDetailsModel(data *entities.UserEntity) (string, int64, error, bool) {
	logger.Log.Println("Inside InsertUserDetailsModel")
	var res int64
	db, err := config.ConnectMySqlDbSlaveSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return "DbConnection Error", res, err, false
	}
	dataAccess := dao.DbConn{DB: db}
	id, err := dataAccess.InsertUserDetailsDao(data)
	if err != nil {
		logger.Log.Println("InsertUserDetailsDao Insertion Error : ", err)
		return "Insertion Error", res, err, false
	}
	res = id
	return "Our Team Will Contact Soon...!", res, nil, true

}
func UploadJobModel(data *entities.JobEntity) (string, int64, error, bool) {
	logger.Log.Println("Inside UploadJobModel")
	var res int64
	db, err := config.ConnectMySqlDbSlaveSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return "DbConnection Error", res, err, false
	}
	dataAccess := dao.DbConn{DB: db}
	id, err := dataAccess.UploadJobDao(data)
	if err != nil {
		logger.Log.Println("UploadJobModel Insertion Error : ", err)
		return "Insertion Error", res, err, false
	}
	res = id
	return "Job Posted Successfully", res, nil, true

}
