package handlers

import (
	"Skool_Saver/src/entities"
	"Skool_Saver/src/logger"
	"Skool_Saver/src/models"
	"encoding/json"
	"log"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var data entities.LoginEntity
	jsonError := data.FromJSON(r.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, msg, err := models.LoginModel(&data)
		logger.Log.Println(msg)
		ThrowLoginResponse(success, msg, w, err)
	}
}
func ThrowLoginResponse(success bool, msg string, w http.ResponseWriter, err error) {
	var response entities.LoginResponse
	response.Success = success
	response.Message = msg
	response.Error = err

	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
