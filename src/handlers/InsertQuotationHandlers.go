package handlers

import (
	"Skool_Saver/src/entities"
	"Skool_Saver/src/logger"
	"Skool_Saver/src/models"
	"encoding/json"
	"log"
	"net/http"
)

func InsertQuotationHandler(w http.ResponseWriter, r *http.Request) {
	var data entities.QuotationEntity
	jsonError := data.FromJSON(r.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		msg, responseData, _, success := models.InsertQuotationModel(&data)
		logger.Log.Println(msg)
		ThrowInsertQuotationResponse(msg, responseData, w, success)
	}

}
func ThrowInsertQuotationResponse(message string, data int64, w http.ResponseWriter, success bool) {
	var response entities.QuotationResponse
	response.Success = success
	response.Message = message
	response.Details = data

	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
