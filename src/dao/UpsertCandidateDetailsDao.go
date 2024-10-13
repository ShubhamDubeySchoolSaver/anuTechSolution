package dao

import (
	"Skool_Saver/src/entities"
	"Skool_Saver/src/logger"
)

var uploadFile = "INSERT INTO jobappliedcandidates (candidatename,email,mobno,current_emplyeement,location,skills,yearofexp,blobname,ext,filename,jobid) VALUES (?,?,?,?,?,?,?,?,?,?,?)"

func (dbc DbConn) UpsertCandidateDetailsDao(tz *entities.CandidateEntity, blobName, ext, fileName string) (int64, bool, error) {
	logger.Log.Println("Inside Upload File Dao function")
	stmt, err := dbc.DB.Prepare(uploadFile)
	if err != nil {
		logger.Log.Print("UploadFile Prepare Statement Error", err)
		return 0, false, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(tz.CandidateName, tz.Email, tz.MobileNumber, tz.CurreantEmployeement, tz.Location, tz.Skills, tz.CandidateExperienece, blobName, ext, fileName, tz.JobId)
	if err != nil {
		logger.Log.Print("UploadFile Execute Statement Error", err)
		return 0, false, err
	}

	lastInsertedID, _ := res.LastInsertId()
	return lastInsertedID, true, nil
}
