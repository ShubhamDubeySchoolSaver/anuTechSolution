package dao

import (
	"Skool_Saver/src/entities"
	"log"
)

var insertdetails = "INSERT INTO userdetails (user_name, user_email, user_mobileno,user_message) VALUES (?,?,?,?)"

func (dbc DbConn) InsertUserDetailsDao(tz *entities.UserEntity) (int64, error) {

	stmt, err := dbc.DB.Prepare(insertdetails)
	if err != nil {
		log.Print("InsertFollowUpUserDao Prepare Statement  Error", err)
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(tz.UserName, tz.UserEmail, tz.MobileNumber, tz.Message)
	if err != nil {
		log.Print("InsertFollowUpUserDao Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, _ := res.LastInsertId()

	return lastInsertedId, nil
}

var jobinsert = "INSERT INTO jobs (designation, experience, location) VALUES (?,?,?)"

func (dbc DbConn) UploadJobDao(tz *entities.JobEntity) (int64, error) {

	stmt, err := dbc.DB.Prepare(jobinsert)
	if err != nil {
		log.Print("JobEntity Prepare Statement  Error", err)
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(tz.Designation, tz.Experience, tz.JobLocation)
	if err != nil {
		log.Print("JobEntity Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, _ := res.LastInsertId()
	return lastInsertedId, nil
}
