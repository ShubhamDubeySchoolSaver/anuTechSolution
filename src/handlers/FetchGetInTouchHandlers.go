package handlers

import (
	"Skool_Saver/src/entities"
	"Skool_Saver/src/logger"
	"Skool_Saver/src/models"
	"encoding/json"
	"net/http"
	"strconv"
)

func FetchGetInTouchHandlers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	var limit, err = strconv.ParseInt(query.Get("limit"), 10, 64)
	if err != nil {
		logger.Log.Println("Error parsing limit")
	}
	var offset, errors = strconv.ParseInt(query.Get("offset"), 10, 64)
	if errors != nil {
		logger.Log.Println("Error parsing offset")
	}
	data, success := models.FetchGetInTouchModels(limit, offset)
	if success {
		ThrowFetchGetInTouchResponse(data, success, w, "Data Fetched Successfully")
	}

}
func ThrowFetchGetInTouchResponse(data []entities.GetInTouchEntity, success bool, w http.ResponseWriter, msg string) {
	var response entities.GetInTouchResponse
	response.Success = success
	response.Message = msg
	response.Details = data

	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
func FetchJobAppliedCandidatesHandlers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	var limit, err = strconv.ParseInt(query.Get("limit"), 10, 64)
	if err != nil {
		logger.Log.Println("Error parsing limit")
	}
	var offset, errors = strconv.ParseInt(query.Get("offset"), 10, 64)
	if errors != nil {
		logger.Log.Println("Error parsing offset")
	}
	data, success := models.FetchJobAppliedCandidatesModels(limit, offset)
	if success {
		ThrowFetchJobAppliedCandidatesResponse(data, success, w, "Data Fetched Successfully")
	}

}
func ThrowFetchJobAppliedCandidatesResponse(data []entities.JobCandidatesEntity, success bool, w http.ResponseWriter, msg string) {
	var response entities.JobCandidatesResponse
	response.Success = success
	response.Message = msg
	response.Details = data

	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
func FetchConsultationsHandlers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	var limit, err = strconv.ParseInt(query.Get("limit"), 10, 64)
	if err != nil {
		logger.Log.Println("Error parsing limit")
	}
	var offset, errors = strconv.ParseInt(query.Get("offset"), 10, 64)
	if errors != nil {
		logger.Log.Println("Error parsing offset")
	}
	data, success := models.FetchConsultationsModels(limit, offset)
	if success {
		ThrowFetchConsultationsResponse(data, success, w, "Data Fetched Successfully")
	}

}
func ThrowFetchConsultationsResponse(data []entities.ConsultationEntity, success bool, w http.ResponseWriter, msg string) {
	var response entities.ConsultationResponse
	response.Success = success
	response.Message = msg
	response.Details = data

	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
func FetchJobsHandlers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	var limit, err = strconv.ParseInt(query.Get("limit"), 10, 64)
	if err != nil {
		logger.Log.Println("Error parsing limit")
	}
	var offset, errors = strconv.ParseInt(query.Get("offset"), 10, 64)
	if errors != nil {
		logger.Log.Println("Error parsing offset")
	}
	data, success := models.FetchJobsModels(limit, offset)
	if success {
		ThrowFetchJobsResponse(data, success, w, "Data Fetched Successfully")
	}

}
func ThrowFetchJobsResponse(data []entities.FetchJobEntity, success bool, w http.ResponseWriter, msg string) {
	var response entities.FetchJobEntityResponse
	response.Success = success
	response.Message = msg
	response.Details = data

	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
