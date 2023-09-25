package database

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Clearance struct {
	ID                 int64  `json:"ID"`
	ResidentID         int64  `json:"ResidentID"`
	DateCreated        string `json:"DateCreated"`
	DateUpdated        string `json:"DateUpdated"`
	ValidUntil         string `json:"ValidUntil"`
	IssuingOfficer     string `json:"IssuingOfficer"`
	Remarks            string `json:"Remarks"`
	ResidentLastName   string `json:"ResidentLastName"`
	ResidentFirstName  string `json:"ResidentFirstName"`
	ResidentMiddleName string `json:"ResidentMiddleName"`
	Purpose            string `json:"Purpose"`
}

func CreateClearance(db sqlx.Ext, ResidentID int64, ValidUntil string, IssuingOfficer string, Remarks string, ResidentLastName string,
	ResidentFirstName string, ResidentMiddleName string, Purpose string) error {

	currentTime := time.Now()
	// Format the time as "YYYY-MM-DD 03:04 PM"
	formattedTime := currentTime.Format("2006-01-02 03:04 PM")

	_, err := db.Exec(`INSERT INTO Clearance (
		ResidentID,
		DateCreated,
		DateUpdated,
		ValidUntil,
		IssuingOfficer,
		Remarks,
		ResidentLastName,
		ResidentFirstName,
		ResidentMiddleName,
		Purpose
	)
	Values(?,?,?,?,?,?,?,?,?,?)`,
		ResidentID,
		formattedTime,
		formattedTime,
		ValidUntil,
		IssuingOfficer,
		Remarks,
		ResidentLastName,
		ResidentFirstName,
		ResidentMiddleName,
		Purpose,
	)

	if err != nil {
		return err
	}

	return nil
}

func DeleteClearance(db sqlx.Ext, ID int64) error {

	_, err := db.Exec(`DELETE FROM Clearance WHERE ID = ? `, ID)

	if err != nil {
		return err
	}

	return nil
}

func UpdateClearance(db sqlx.Ext, ID int64, Remarks string, ResidentLastName string,
	ResidentFirstName string, ResidentMiddleName string, Purpose string) error {

	currentTime := time.Now()
	// Format the time as "YYYY-MM-DD 03:04 PM"
	formattedTime := currentTime.Format("2006-01-02 03:04 PM")

	_, err := db.Exec(`UPDATE Clearance SET 
		DateUpdated = ?,
		Remarks = ?,
		ResidentLastName = ?,
		ResidentFirstName = ?,
		ResidentMiddleName = ?,
		Purpose = ? WHERE ID= ?`,
		formattedTime,
		Remarks,
		ResidentLastName,
		ResidentFirstName,
		ResidentMiddleName,
		Purpose,
		ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func ReadClearance(db sqlx.Ext) ([]*Clearance, error) {

	clearanceArray := make([]*Clearance, 0)
	var ID int64
	var ResidentID int64
	var DateCreated string
	var DateUpdated string
	var ValidUntil string
	var IssuingOfficer string
	var Remarks string
	var ResidentLastName string
	var ResidentFirstName string
	var ResidentMiddleName string
	var Purpose string

	rows, err := db.Queryx(`SELECT ID,
				ResidentID,
				DateCreated,
				DateUpdated,
				ValidUntil,
				IssuingOfficer,
				Remarks,
				ResidentLastName,
				ResidentFirstName,
				ResidentMiddleName,
				Purpose FROM Clearance`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ID, &ResidentID, &DateCreated, &DateUpdated, &ValidUntil, &IssuingOfficer,
			&Remarks, &ResidentLastName, &ResidentFirstName, &ResidentMiddleName, &Purpose)
		if err != nil {
			return nil, err
		}
		clearanceArray = append(clearanceArray, &Clearance{
			ID:                 ID,
			ResidentID:         ResidentID,
			DateCreated:        DateCreated,
			DateUpdated:        DateUpdated,
			ValidUntil:         ValidUntil,
			IssuingOfficer:     IssuingOfficer,
			Remarks:            Remarks,
			ResidentLastName:   ResidentLastName,
			ResidentFirstName:  ResidentFirstName,
			ResidentMiddleName: ResidentMiddleName,
			Purpose:            Purpose,
		})
	}
	return clearanceArray, nil
}
