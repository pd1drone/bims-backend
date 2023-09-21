package database

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Indigencies struct {
	ID             int64  `json:"ID"`
	ResidentID     int64  `json:"ResidentID"`
	DateCreated    string `json:"DateCreated"`
	DateUpdated    string `json:"DateUpdated"`
	Reason         string `json:"Reason"`
	ValidUntil     string `json:"ValidUntil"`
	IssuingOfficer string `json:"IssuingOfficer"`
	Remarks        string `json:"Remarks"`
}

func CreateIndigencies(db sqlx.Ext, ResidentID int64, Reason string, ValidUntil string, IssuingOfficer string, Remarks string) error {

	currentTime := time.Now()
	// Format the time as "YYYY-MM-DD 03:04 PM"
	formattedTime := currentTime.Format("2006-01-02 03:04 PM")

	_, err := db.Exec(`INSERT INTO Indigencies (
		ResidentID,
		DateCreated,
		DateUpdated,
		Reason,
		ValidUntil,
		IssuingOfficer,
		Remarks
	)
	Values(?,?,?,?,?,?,?)`,
		ResidentID,
		formattedTime,
		formattedTime,
		Reason,
		ValidUntil,
		IssuingOfficer,
		Remarks,
	)

	if err != nil {
		return err
	}

	return nil
}

func DeleteIndigencies(db sqlx.Ext, ID int64) error {

	_, err := db.Exec(`DELETE FROM Indigencies WHERE ID = ? `, ID)

	if err != nil {
		return err
	}

	return nil
}

func UpdateIndigencies(db sqlx.Ext, ID int64, ResidentID int64, Reason string, IssuingOfficer string, Remarks string) error {

	currentTime := time.Now()
	// Format the time as "YYYY-MM-DD 03:04 PM"
	formattedTime := currentTime.Format("2006-01-02 03:04 PM")

	_, err := db.Exec(`UPDATE Indigencies SET 
		ResidentID = ?,
		DateUpdated = ?,
		Reason = ?,
		IssuingOfficer = ?,
		Remarks = ? WHERE ID= ?`,
		ResidentID,
		formattedTime,
		Reason,
		IssuingOfficer,
		Remarks,
		ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func ReadIndigencies(db sqlx.Ext) ([]*Indigencies, error) {

	indigenciesArray := make([]*Indigencies, 0)
	var ID int64
	var ResidentID int64
	var DateCreated string
	var DateUpdated string
	var Reason string
	var ValidUntil string
	var IssuingOfficer string
	var Remarks string

	rows, err := db.Queryx(`SELECT ID,
				ID,
				ResidentID,
				DateCreated,
				DateUpdated,
				Reason,
				ValidUntil,
				IssuingOfficer,
				Remarks FROM Indigencies`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ID, &ResidentID, &DateCreated, &DateUpdated, &Reason, &ValidUntil, &IssuingOfficer, &Remarks)
		if err != nil {
			return nil, err
		}
		indigenciesArray = append(indigenciesArray, &Indigencies{
			ID:             ID,
			ResidentID:     ResidentID,
			DateCreated:    DateCreated,
			DateUpdated:    DateUpdated,
			Reason:         Reason,
			ValidUntil:     ValidUntil,
			IssuingOfficer: IssuingOfficer,
			Remarks:        Remarks,
		})
	}
	return indigenciesArray, nil
}
