package database

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Indigencies struct {
	ID             int64      `json:"ID"`
	ResidentID     int64      `json:"ResidentID"`
	DateCreated    string     `json:"DateCreated"`
	DateUpdated    string     `json:"DateUpdated"`
	Reason         string     `json:"Reason"`
	ValidUntil     string     `json:"ValidUntil"`
	IssuingOfficer string     `json:"IssuingOfficer"`
	Remarks        string     `json:"Remarks"`
	DocumentStatus string     `json:"DocumentStatus"`
	Resident       *Residents `json:"ResidentData"`
}
type IndigenciesXL struct {
	ID                    int64  `json:"ID"`
	ResidentID            int64  `json:"ResidentID"`
	DateCreated           string `json:"DateCreated"`
	DateUpdated           string `json:"DateUpdated"`
	Reason                string `json:"Reason"`
	ValidUntil            string `json:"ValidUntil"`
	IssuingOfficer        string `json:"IssuingOfficer"`
	Remarks               string `json:"Remarks"`
	DocumentStatus        string `json:"DocumentStatus"`
	ResidentDateCreated   string `json:"ResidentDateCreated"`
	ResidentDateUpdated   string `json:"ResidentDateUpdated"`
	Address               string `json:"Address"`
	BirthDate             string `json:"BirthDate"`
	BirthPlace            string `json:"BirthPlace"`
	Gender                string `json:"Gender"`
	CivilStatus           string `json:"CivilStatus"`
	ContactNumber         string `json:"ContactNumber"`
	GuardianName          string `json:"GuardianName"`
	GurdianContactNumbers string `json:"GurdianContactNumbers"`
	DocumentType          string `json:"DocumentType"`
	DocumentID            int64  `json:"DocumentID"`
}

func CreateIndigencies(db sqlx.Ext, ResidentID int64, Reason string, ValidUntil string, IssuingOfficer string, Remarks string) (int64, error) {

	currentTime := time.Now()
	// Format the time as "YYYY-MM-DD 03:04 PM"
	formattedTime := currentTime.Format("2006-01-02 03:04 PM")

	query, err := db.Exec(`INSERT INTO Indigencies (
		ResidentID,
		DateCreated,
		DateUpdated,
		Reason,
		ValidUntil,
		IssuingOfficer,
		Remarks,
		DocumentStatus
	)
	Values(?,?,?,?,?,?,?,?)`,
		ResidentID,
		formattedTime,
		formattedTime,
		Reason,
		ValidUntil,
		IssuingOfficer,
		Remarks,
		"For Printing",
	)

	if err != nil {
		return 0, err
	}
	documentID, err := query.LastInsertId()
	if err != nil {
		return 0, err
	}

	return documentID, nil
}

func DeleteIndigencies(db sqlx.Ext, ID int64) error {

	_, err := db.Exec(`DELETE FROM Indigencies WHERE ID = ? `, ID)

	if err != nil {
		return err
	}

	return nil
}

func UpdateIndigencies(db sqlx.Ext, ID int64, Reason string, Remarks string, DocumentStatus string) error {

	currentTime := time.Now()
	// Format the time as "YYYY-MM-DD 03:04 PM"
	formattedTime := currentTime.Format("2006-01-02 03:04 PM")

	_, err := db.Exec(`UPDATE Indigencies SET 
		DateUpdated = ?,
		Reason = ?,
		Remarks = ?,
		DocumentStatus = ? WHERE ID= ?`,
		formattedTime,
		Reason,
		Remarks,
		DocumentStatus,
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
	var DocumentStatus string

	rows, err := db.Queryx(`SELECT ID,
				ResidentID,
				DateCreated,
				DateUpdated,
				Reason,
				ValidUntil,
				IssuingOfficer,
				Remarks,
				DocumentStatus FROM Indigencies`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ID, &ResidentID, &DateCreated, &DateUpdated, &Reason, &ValidUntil, &IssuingOfficer, &Remarks, &DocumentStatus)
		if err != nil {
			return nil, err
		}
		residentData, err := ReadResidentData(db, ResidentID)
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
			DocumentStatus: DocumentStatus,
			Resident:       residentData,
		})
	}
	return indigenciesArray, nil
}

func ReadIndigencyData(db sqlx.Ext, DocumentID int64) (*Indigencies, error) {

	var ID int64
	var ResidentID int64
	var DateCreated string
	var DateUpdated string
	var Reason string
	var ValidUntil string
	var IssuingOfficer string
	var Remarks string
	var DocumentStatus string

	rows, err := db.Queryx(`SELECT ID,
	ResidentID,
	DateCreated,
	DateUpdated,
	Reason,
	ValidUntil,
	IssuingOfficer,
	Remarks,
	DocumentStatus FROM Indigencies WHERE ID = ?`, DocumentID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ID, &ResidentID, &DateCreated, &DateUpdated, &Reason, &ValidUntil, &IssuingOfficer, &Remarks, &DocumentStatus)
		if err != nil {
			return nil, err
		}
	}
	return &Indigencies{
		ID:             ID,
		ResidentID:     ResidentID,
		DateCreated:    DateCreated,
		DateUpdated:    DateUpdated,
		Reason:         Reason,
		ValidUntil:     ValidUntil,
		IssuingOfficer: IssuingOfficer,
		Remarks:        Remarks,
		DocumentStatus: DocumentStatus,
	}, nil
}

func ReadIndigenciesXL(db sqlx.Ext) ([]*IndigenciesXL, error) {

	indigenciesArray := make([]*IndigenciesXL, 0)
	var ID int64
	var ResidentID int64
	var DateCreated string
	var DateUpdated string
	var Reason string
	var ValidUntil string
	var IssuingOfficer string
	var Remarks string
	var DocumentStatus string

	rows, err := db.Queryx(`SELECT ID,
				ResidentID,
				DateCreated,
				DateUpdated,
				Reason,
				ValidUntil,
				IssuingOfficer,
				Remarks,
				DocumentStatus FROM Indigencies`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ID, &ResidentID, &DateCreated, &DateUpdated, &Reason, &ValidUntil, &IssuingOfficer, &Remarks, &DocumentStatus)
		if err != nil {
			return nil, err
		}
		residentData, err := ReadResidentData(db, ResidentID)
		if err != nil {
			return nil, err
		}
		indigenciesArray = append(indigenciesArray, &IndigenciesXL{
			ID:                    ID,
			ResidentID:            ResidentID,
			DateCreated:           DateCreated,
			DateUpdated:           DateUpdated,
			Reason:                Reason,
			ValidUntil:            ValidUntil,
			IssuingOfficer:        IssuingOfficer,
			Remarks:               Remarks,
			DocumentStatus:        DocumentStatus,
			ResidentDateCreated:   residentData.DateCreated,
			ResidentDateUpdated:   residentData.DateUpdated,
			Address:               residentData.Address,
			BirthDate:             residentData.BirthDate,
			BirthPlace:            residentData.BirthPlace,
			Gender:                residentData.Gender,
			CivilStatus:           residentData.CivilStatus,
			ContactNumber:         residentData.ContactNumber,
			GuardianName:          residentData.GuardianName,
			GurdianContactNumbers: residentData.GurdianContactNumbers,
			DocumentType:          residentData.DocumentType,
			DocumentID:            residentData.DocumentID,
		})
	}
	return indigenciesArray, nil
}
