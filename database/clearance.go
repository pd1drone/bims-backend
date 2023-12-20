package database

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Clearance struct {
	ID                 int64      `json:"ID"`
	ResidentID         int64      `json:"ResidentID"`
	DateCreated        string     `json:"DateCreated"`
	DateUpdated        string     `json:"DateUpdated"`
	ValidUntil         string     `json:"ValidUntil"`
	IssuingOfficer     string     `json:"IssuingOfficer"`
	Remarks            string     `json:"Remarks"`
	ResidentLastName   string     `json:"ResidentLastName"`
	ResidentFirstName  string     `json:"ResidentFirstName"`
	ResidentMiddleName string     `json:"ResidentMiddleName"`
	Purpose            string     `json:"Purpose"`
	CedulaNo           string     `json:"cedulaNo"`
	PrecintNo          string     `json:"precintNo"`
	DocumentStatus     string     `json:"DocumentStatus"`
	Resident           *Residents `json:"ResidentData"`
}

type ClearanceXL struct {
	ID                    int64  `json:"ID"`
	ResidentID            int64  `json:"ResidentID"`
	DateCreated           string `json:"DateCreated"`
	DateUpdated           string `json:"DateUpdated"`
	ValidUntil            string `json:"ValidUntil"`
	IssuingOfficer        string `json:"IssuingOfficer"`
	Remarks               string `json:"Remarks"`
	ResidentLastName      string `json:"ResidentLastName"`
	ResidentFirstName     string `json:"ResidentFirstName"`
	ResidentMiddleName    string `json:"ResidentMiddleName"`
	Purpose               string `json:"Purpose"`
	CedulaNo              string `json:"cedulaNo"`
	PrecintNo             string `json:"precintNo"`
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

func CreateClearance(db sqlx.Ext, ResidentID int64, ValidUntil string, IssuingOfficer string, Remarks string, ResidentLastName string,
	ResidentFirstName string, ResidentMiddleName string, Purpose string, cedulaNo string, precintNo string) (int64, error) {

	currentTime := time.Now()
	// Format the time as "YYYY-MM-DD 03:04 PM"
	formattedTime := currentTime.Format("2006-01-02 03:04 PM")

	query, err := db.Exec(`INSERT INTO Clearance (
		ResidentID,
		DateCreated,
		DateUpdated,
		ValidUntil,
		IssuingOfficer,
		Remarks,
		ResidentLastName,
		ResidentFirstName,
		ResidentMiddleName,
		Purpose,
		CedulaNo,
		PrecintNo,
		DocumentStatus
	)
	Values(?,?,?,?,?,?,?,?,?,?,?,?,?)`,
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
		cedulaNo,
		precintNo,
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

func DeleteClearance(db sqlx.Ext, ID int64) error {

	_, err := db.Exec(`DELETE FROM Clearance WHERE ID = ? `, ID)

	if err != nil {
		return err
	}

	return nil
}

func UpdateClearance(db sqlx.Ext, ID int64, Remarks string, ResidentLastName string,
	ResidentFirstName string, ResidentMiddleName string, Purpose string, cedulaNo string, precintNo string, documentStatus string) error {

	currentTime := time.Now()
	// Format the time as "YYYY-MM-DD 03:04 PM"
	formattedTime := currentTime.Format("2006-01-02 03:04 PM")

	_, err := db.Exec(`UPDATE Clearance SET 
		DateUpdated = ?,
		Remarks = ?,
		ResidentLastName = ?,
		ResidentFirstName = ?,
		ResidentMiddleName = ?,
		Purpose = ?,
		CedulaNo = ?,
		PrecintNo = ?,
		DocumentStatus = ? WHERE ID= ?`,
		formattedTime,
		Remarks,
		ResidentLastName,
		ResidentFirstName,
		ResidentMiddleName,
		Purpose,
		cedulaNo,
		precintNo,
		documentStatus,
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
	var CedulaNo string
	var PrecintNo string
	var DocumentStatus string

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
				Purpose,
				CedulaNo,
				PrecintNo,
				DocumentStatus FROM Clearance`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ID, &ResidentID, &DateCreated, &DateUpdated, &ValidUntil, &IssuingOfficer,
			&Remarks, &ResidentLastName, &ResidentFirstName, &ResidentMiddleName, &Purpose, &CedulaNo, &PrecintNo, &DocumentStatus)
		if err != nil {
			return nil, err
		}
		residentData, err := ReadResidentData(db, ResidentID)
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
			CedulaNo:           CedulaNo,
			PrecintNo:          PrecintNo,
			DocumentStatus:     DocumentStatus,
			Resident:           residentData,
		})
	}
	return clearanceArray, nil
}

func ReadClearanceData(db sqlx.Ext, DocumentID int64) (*Clearance, error) {

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
	var CedulaNo string
	var PrecintNo string
	var DocumentStatus string

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
				Purpose,
				CedulaNo,
				PrecintNo,
				DocumentStatus FROM Clearance WHERE ID = ?`, DocumentID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ID, &ResidentID, &DateCreated, &DateUpdated, &ValidUntil, &IssuingOfficer, &Remarks,
			&ResidentLastName, &ResidentFirstName, &ResidentMiddleName, &Purpose, &CedulaNo, &PrecintNo, &DocumentStatus)
		if err != nil {
			return nil, err
		}
	}
	return &Clearance{
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
		CedulaNo:           CedulaNo,
		PrecintNo:          PrecintNo,
		DocumentStatus:     DocumentStatus,
	}, nil
}

func ReadClearanceXL(db sqlx.Ext) ([]*ClearanceXL, error) {

	clearanceArray := make([]*ClearanceXL, 0)
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
	var CedulaNo string
	var PrecintNo string
	var DocumentStatus string

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
				Purpose,
				CedulaNo,
				PrecintNo,
				DocumentStatus FROM Clearance`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ID, &ResidentID, &DateCreated, &DateUpdated, &ValidUntil, &IssuingOfficer,
			&Remarks, &ResidentLastName, &ResidentFirstName, &ResidentMiddleName, &Purpose, &CedulaNo, &PrecintNo, &DocumentStatus)
		if err != nil {
			return nil, err
		}
		residentData, err := ReadResidentData(db, ResidentID)
		if err != nil {
			return nil, err
		}
		clearanceArray = append(clearanceArray, &ClearanceXL{
			ID:                    ID,
			ResidentID:            ResidentID,
			DateCreated:           DateCreated,
			DateUpdated:           DateUpdated,
			ValidUntil:            ValidUntil,
			IssuingOfficer:        IssuingOfficer,
			Remarks:               Remarks,
			ResidentLastName:      ResidentLastName,
			ResidentFirstName:     ResidentFirstName,
			ResidentMiddleName:    ResidentMiddleName,
			Purpose:               Purpose,
			CedulaNo:              CedulaNo,
			PrecintNo:             PrecintNo,
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
	return clearanceArray, nil
}
