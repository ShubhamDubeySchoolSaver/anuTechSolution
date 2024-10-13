package handlers

import (
	"Skool_Saver/src/entities"
	"Skool_Saver/src/logger"
	"Skool_Saver/src/models"
	"encoding/json"
	"log"
	"net/http"
)

func InsertUserDetailsHandler(w http.ResponseWriter, r *http.Request) {
	logger.Log.Println("Inside InsertUserDetailsHandler")
	var data entities.UserEntity
	jsonError := data.FromJSON(r.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		msg, responseData, _, success := models.InsertUserDetailsModel(&data)
		logger.Log.Println(msg)
		ThrowInsertFollowUpUserResponse(msg, responseData, w, success)
	}
}
func ThrowInsertFollowUpUserResponse(msg string, responsedata int64, w http.ResponseWriter, success bool) {
	var response entities.UserEntityResponse
	response.Message = msg
	response.Success = success
	response.Details = responsedata

	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

}
func UploadJobHandlers(w http.ResponseWriter, r *http.Request) {
	logger.Log.Println("Inside UploadJobHandlers")
	var data entities.JobEntity
	jsonError := data.FromJSON(r.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		msg, responseData, _, success := models.UploadJobModel(&data)
		logger.Log.Println(msg)
		ThrowUploadJobResponse(msg, responseData, w, success)
	}
}
func ThrowUploadJobResponse(msg string, responsedata int64, w http.ResponseWriter, success bool) {
	var response entities.JobEntityResponse
	response.Message = msg
	response.Success = success
	response.Details = responsedata

	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

}
