package dao

import (
	"Skool_Saver/src/entities"
	"Skool_Saver/src/logger"
	"errors"
	"log"
)

var logindetails = "SELECT username,password FROM skoolsaverusers WHERE username=? AND deleteflag=0"

func (dbc DbConn) LoginDao(tz *entities.LoginEntity) ([]entities.LoginEntityReq, error) {
	log.Println("In side dao")
	values := []entities.LoginEntityReq{}
	rows, err := dbc.DB.Query(logindetails, tz.UserName)
	if err != nil {
		logger.Log.Print("Login Get Statement Prepare Error", err)
		log.Print("Login Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.LoginEntityReq{}
		err = rows.Scan(&value.UserName, &value.Password)
		if err != nil {
			logger.Log.Print("Login Scan Error", err)
			log.Print("Login Scan Error", err)
			return values, err
		}
		values = append(values, value)
	}
	return values, nil
}

var updatetimequery = "UPDATE skoolsaverusers SET logintime=CURRENT_TIMESTAMP WHERE username=? AND password=? AND deleteflag=0"

func (dbc DbConn) UpdateLoginTime(username, pass string) (bool, error) {
	logger.Log.Println("Inside UpdateLoginTime Dao")
	var res bool
	smt, err := dbc.DB.Prepare(updatetimequery)
	if err != nil {
		logger.Log.Println("UpdateLoginTime Prepared Statement Error ", err)
		return false, errors.New("updateLoginTime Prepared Statement Error")
	}
	defer smt.Close()

	result, Err := smt.Exec(username, pass)
	if Err != nil {
		logger.Log.Println("UpdateLoginTime Execution Error ", Err)
		return false, errors.New("updateLoginTime Execution Error")
	}
	row, _ := result.RowsAffected()
	if (row) > 0 {
		res = true
	}
	return res, nil
}
