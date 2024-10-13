package handlers

import (
	"Skool_Saver/src/entities"
	"Skool_Saver/src/logger"
	"Skool_Saver/src/models"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func UpsertCandidateDetails(w http.ResponseWriter, r *http.Request) {
	logger.Log.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		logger.Log.Println("Error Retrieving the File")
		logger.Log.Println(err)
		return
	}
	defer file.Close()
	logger.Log.Printf("Uploaded File: %+v\n", handler.Filename)
	logger.Log.Printf("File Size: %+v\n", handler.Size)
	logger.Log.Printf("MIME Header: %+v\n", handler.Header)

	fileBytes, err := io.ReadAll(file)

	CanidateName := r.FormValue("candidate_name")
	CanidateEmail := r.FormValue("email")
	CanidateExperience := r.FormValue("yoe")
	CanidateLocation := r.FormValue("candidate_location")
	CanidateMobNo := r.FormValue("candidate_mobno")
	CanidateSkills := r.FormValue("skills")
	CanidateCurrentEmployement := r.FormValue("candidate_currentemployeement")
	JobId := r.FormValue("jobid")

	var data = entities.CandidateEntity{}
	data.CandidateName = CanidateName
	data.Email = CanidateEmail
	data.CandidateExperienece = CanidateExperience
	data.Location = CanidateLocation
	data.MobileNumber = CanidateMobNo
	data.Skills = CanidateSkills
	data.CurreantEmployeement = CanidateCurrentEmployement
	jobid, _ := strconv.ParseInt(JobId, 10, 64)
	data.JobId = jobid

	success, msg, _ := models.UpsertCandidateDetailsModel(&data, fileBytes, handler.Filename)
	if success {
		ThrowUpsertCandidateDetailsResponse(msg, w, success)
	}
}

func ThrowUpsertCandidateDetailsResponse(successMessage string, w http.ResponseWriter, success bool) {
	var response = entities.CandidateResponse{}
	response.Success = success
	response.Message = successMessage
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
