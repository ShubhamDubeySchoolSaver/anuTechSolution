package dao

import (
	"Skool_Saver/src/entities"
	"log"
)

var insertquotation = "INSERT INTO consultations (fullname, email, mobilenumber,service,pricerange,location,aboutproject) VALUES (?,?,?,?,?,?,?)"

func (dbc DbConn) InsertQuotationDao(tz *entities.QuotationEntity) (int64, error) {

	stmt, err := dbc.DB.Prepare(insertquotation)
	if err != nil {
		log.Print("InsertQuotationDao Prepare Statement  Error", err)
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(tz.FullName, tz.Email, tz.MobileNumber, tz.Service, tz.Pricerange, tz.Location, tz.AboutProject)
	if err != nil {
		log.Print("InsertQuotationDao Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, _ := res.LastInsertId()
	return lastInsertedId, nil
}
