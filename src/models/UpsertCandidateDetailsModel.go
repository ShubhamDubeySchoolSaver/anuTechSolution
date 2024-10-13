package models

import (
	"Skool_Saver/src/config"
	"Skool_Saver/src/dao"
	"Skool_Saver/src/entities"
	"Skool_Saver/src/logger"
	"Skool_Saver/src/utility"
	"errors"
	"os"
	"strings"
)

func UpsertCandidateDetailsModel(tz *entities.CandidateEntity, fileBytes []byte, fileName string) (bool, string, error) {
	splitedFileName := strings.Split(fileName, ".")
	ext := splitedFileName[len(splitedFileName)-1]
	ext = strings.ToLower(ext)

	if ext != "pdf" && ext != "docx" {
		logger.Log.Println("Unsupported file type:", ext)
		return false, "File type not supported", errors.New("file type not supported")
	}
	db, err := config.ConnectMySqlDbSlaveSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, "Something Went Wrong", err
	}

	blobName := utility.GetBlobName(fileName)

	// // Directory where files will be saved
	// directoryPath := "../src/public/"

	// // Create directory if it doesn't exist
	// err = os.MkdirAll(directoryPath, os.ModePerm)
	// if err != nil {
	// 	logger.Log.Println("Failed to create directory", err)
	// 	return false, "Failed to create directory", err
	// }

	// // Define the full file path
	// fullFilePath := filepath.Join(directoryPath, blobName)

	contextpath, _ := os.Getwd()
	//filePath := config.VolumePath + blobName
	filePath := contextpath + "/src/public/" + blobName
	file, err := os.Create(filePath)

	// Write the file to the directory
	file.Write(fileBytes)
	// err = os.WriteFile(fullFilePath, fileBytes, 0644)
	// if err != nil {
	// 	logger.Log.Println("Failed to save the file", err)
	// 	return false, "Failed to save the file", err
	// }
	logger.Log.Println("File successfully saved at:", filePath)

	dataAccess := dao.DbConn{DB: db}
	_, uploadStatus, err := dataAccess.UpsertCandidateDetailsDao(tz, blobName, ext, splitedFileName[0])
	if err != nil {
		logger.Log.Println("Error while uploading the document", err)
		return false, "Something went wrong", err
	}
	msg := "Document is not Uploaded"
	if uploadStatus {
		msg = "Document is Uploaded"
	}

	return true, msg, nil
}
