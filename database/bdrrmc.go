package database

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// `ID` int PRIMARY KEY AUTO_INCREMENT,
// `TypeOfRecord` varchar(255),
// `PartiesInvolved` varchar(255),
// `DateTime` varchar(255),
// `Location` varchar(255),
// `RecordDetails` longtext,
// `IssuingOfficer` varchar(255)
type BDRRMC struct {
	ID              int64  `json:"ID"`
	DateCreated     string `json:"DateCreated"`
	DateUpdated     string `json:"DateUpdated"`
	TypeOfRecord    string `json:"TypeOfRecord"`
	PartiesInvolved string `json:"PartiesInvolved"`
	DateTime        string `json:"DateTime"`
	Location        string `json:"Location"`
	RecordDetails   string `json:"RecordDetails"`
	IssuingOfficer  string `json:"IssuingOfficer"`
}

func CreateBDRRMC(db sqlx.Ext, typeOfRecord string, partiesInvolved string, dateTime string, location string, recordDetails string, issuingOfficer string) (int64, error) {
	currentTime := time.Now()
	// Format the time as "YYYY-MM-DD 03:04 PM"
	formattedTime := currentTime.Format("2006-01-02 03:04 PM")
	query, err := db.Exec(`INSERT INTO BDRRMC (
		DateCreated,
		DateUpdated,
		TypeOfRecord,
		PartiesInvolved,
		DateTime,
		Location,
		RecordDetails,
		IssuingOfficer
	)
	Values(?,?,?,?,?,?,?,?)`,
		formattedTime,
		formattedTime,
		typeOfRecord,
		partiesInvolved,
		dateTime,
		location,
		recordDetails,
		issuingOfficer,
	)

	if err != nil {
		return 0, err
	}

	docID, err := query.LastInsertId()
	if err != nil {
		return 0, err
	}

	return docID, nil
}

func DeleteBDRRMC(db sqlx.Ext, ID int64) error {

	_, err := db.Exec(`DELETE FROM BDRRMC WHERE ID = ? `, ID)

	if err != nil {
		return err
	}

	return nil
}

func UpdateBDRRMC(db sqlx.Ext, ID int64, typeOfRecord string, partiesInvolved string, dateTime string, location string, recordDetails string, issuingOfficer string) error {
	currentTime := time.Now()
	// Format the time as "YYYY-MM-DD 03:04 PM"
	formattedTime := currentTime.Format("2006-01-02 03:04 PM")

	_, err := db.Exec(`UPDATE BDRRMC SET 
		DateUpdated = ?,
		TypeOfRecord = ?,
		PartiesInvolved = ?,
		DateTime = ?,
		Location = ?,
		RecordDetails = ?,
		IssuingOfficer =? WHERE ID= ?`,
		formattedTime,
		typeOfRecord,
		partiesInvolved,
		dateTime,
		location,
		recordDetails,
		issuingOfficer,
		ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func ReadBDRRMC(db sqlx.Ext) ([]*BDRRMC, error) {

	bdrrmcArray := make([]*BDRRMC, 0)
	var ID int64
	var DateCreated string
	var DateUpdated string
	var TypeOfRecord string
	var PartiesInvolved string
	var DateTime string
	var Location string
	var RecordDetails string
	var IssuingOfficer string

	rows, err := db.Queryx(`SELECT ID,
				DateCreated,
				DateUpdated,				
				TypeOfRecord,
				PartiesInvolved,
				DateTime,
				Location,
				RecordDetails,
				IssuingOfficer FROM BDRRMC`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ID, &DateCreated, &DateUpdated, &TypeOfRecord, &PartiesInvolved, &DateTime, &Location, &RecordDetails, &IssuingOfficer)
		if err != nil {
			return nil, err
		}

		bdrrmcArray = append(bdrrmcArray, &BDRRMC{
			ID:              ID,
			DateCreated:     DateCreated,
			DateUpdated:     DateUpdated,
			TypeOfRecord:    TypeOfRecord,
			PartiesInvolved: PartiesInvolved,
			DateTime:        DateTime,
			Location:        Location,
			RecordDetails:   RecordDetails,
			IssuingOfficer:  IssuingOfficer,
		})

	}
	return bdrrmcArray, nil
}
