package models

import (
	"Skool_Saver/src/config"
	"Skool_Saver/src/dao"
	"Skool_Saver/src/entities"
	"Skool_Saver/src/logger"
	"encoding/base64"
	"errors"
)

func LoginModel(tz *entities.LoginEntity) (bool, string, error) {
	var success bool
	var message string
	var errr error
	logger.Log.Println("Inside Login Model")
	db, err := config.ConnectMySqlDbSlaveSingleton()
	if err != nil {
		logger.Log.Println("DbConnection Error : ", err)
		return false, "DbConnection Failed", err
	}
	dataAccess := dao.DbConn{DB: db}
	values, err := dataAccess.LoginDao(tz)
	if err != nil {
		logger.Log.Println("Fetching LoginDetails Error :", err)
		return false, "Fetching Login Details Error", err
	}
	logger.Log.Println("Login Details : ", values)
	if len(values) > 0 {
		if values[0].UserName == tz.UserName {
			encodedPassword := base64.StdEncoding.EncodeToString([]byte(tz.Password))
			if values[0].Password == encodedPassword {
				updted, err := dataAccess.UpdateLoginTime(tz.UserName, encodedPassword)
				if err != nil {
					logger.Log.Println("UpdateLoginTime Inside Model Error ", err)
					return false, "Update Login Time Error", errors.New("update Login Time Error")
				}
				if updted {
					success = updted
					message = "User LoggedIn Succesfully"
					errr = nil
				}
			}
		}
	}
	return success, message, errr
}
