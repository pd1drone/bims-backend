package database

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type ResidentsXL struct {
	ID                    int64  `json:"ID"`
	DateCreated           string `json:"DateCreated"`
	DateUpdated           string `json:"DateUpdated"`
	LastName              string `json:"LastName"`
	FirstName             string `json:"FirstName"`
	MiddleName            string `json:"MiddleName"`
	Address               string `json:"Address"`
	BirthDate             string `json:"BirthDate"`
	BirthPlace            string `json:"BirthPlace"`
	Gender                string `json:"Gender"`
	CivilStatus           string `json:"CivilStatus"`
	ContactNumber         string `json:"ContactNumber"`
	GuardianName          string `json:"GuardianName"`
	GurdianContactNumbers string `json:"GurdianContactNumbers"`
	ResidentID            int64  `json:"ResidentID"`
	DocumentDateCreated   string `json:"DocumentDateCreated"`
	DocumentDateUpdated   string `json:"DocumentDateUpdated"`
	Purpose               string `json:"Purpose"`
	ValidUntil            string `json:"ValidUntil"`
	Remarks               string `json:"Remarks"`
	DocumentStatus        string `json:"DocumentStatus"`
	HCGGGNumber           string `json:"HCGGGNumber"`
	PhilHealthID          string `json:"PhilHealthID"`
	PhilHealthCategory    string `json:"PhilHealthCategory"`
	ReasonForReferral     string `json:"ReasonForReferral"`
	Reason                string `json:"Reason"`
	CedulaNo              string `json:"CedulaNo"`
	PrecintNo             string `json:"PrecintNo"`
	DocumentType          string `json:"DocumentType"`
	DocumentID            int64  `json:"DocumentID"`
	IssuingOfficer        string `json:"IssuingOfficer"`
}

type Residents struct {
	ID                    int64       `json:"ID"`
	DateCreated           string      `json:"DateCreated"`
	DateUpdated           string      `json:"DateUpdated"`
	LastName              string      `json:"LastName"`
	FirstName             string      `json:"FirstName"`
	MiddleName            string      `json:"MiddleName"`
	Address               string      `json:"Address"`
	BirthDate             string      `json:"BirthDate"`
	BirthPlace            string      `json:"BirthPlace"`
	Gender                string      `json:"Gender"`
	CivilStatus           string      `json:"CivilStatus"`
	ContactNumber         string      `json:"ContactNumber"`
	GuardianName          string      `json:"GuardianName"`
	GurdianContactNumbers string      `json:"GurdianContactNumbers"`
	IssuingOfficer        string      `json:"IssuingOfficer"`
	DocumentType          string      `json:"DocumentType"`
	DocumentID            int64       `json:"DocumentID"`
	DocumentData          interface{} `json:"DocumentData"`
}

type DocumentData struct {
}

func CreateResident(db sqlx.Ext, LastName string, FirstName string, MiddleName string, Address string, BirthDate string,
	BirthPlace string, Gender string, CivilStatus string, ContactNumber string, GuardianName string, GurdianContactNumbers string,
	IssuingOfficer string, DocumentType string) (int64, error) {

	currentTime := time.Now()
	// Format the time as "YYYY-MM-DD 03:04 PM"
	formattedTime := currentTime.Format("2006-01-02 03:04 PM")

	query, err := db.Exec(`INSERT INTO Residents (
		DateCreated,
		DateUpdated,
		LastName,
		FirstName,
		MiddleName,
		Address,
		BirthDate,
		BirthPlace,
		Gender,
		CivilStatus,
		ContactNumber,
		GuardianName,
		GurdianContactNumber,
		IssuingOfficer,
		DocumentType
	)
	Values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		formattedTime,
		formattedTime,
		LastName,
		FirstName,
		MiddleName,
		Address,
		BirthDate,
		BirthPlace,
		Gender,
		CivilStatus,
		ContactNumber,
		GuardianName,
		GurdianContactNumbers,
		IssuingOfficer,
		DocumentType,
	)

	if err != nil {
		return 0, err
	}

	residentID, err := query.LastInsertId()
	if err != nil {
		return 0, err
	}
	return residentID, nil
}

func UpdateResidents(db sqlx.Ext, ID int64, LastName string, FirstName string, MiddleName string, Address string, BirthDate string,
	BirthPlace string, Gender string, CivilStatus string, ContactNumber string, GuardianName string, GurdianContactNumbers string) error {

	currentTime := time.Now()
	// Format the time as "YYYY-MM-DD 03:04 PM"
	formattedTime := currentTime.Format("2006-01-02 03:04 PM")

	_, err := db.Exec(`UPDATE Residents SET 
		DateUpdated = ?,
		LastName = ?,
		FirstName = ?,
		MiddleName = ?,
		Address = ?,
		BirthDate = ?,
		BirthPlace = ?,
		Gender = ?,
		CivilStatus = ?,
		ContactNumber = ?,
		GuardianName = ?,
		GurdianContactNumber = ? WHERE ID= ?`,
		formattedTime,
		LastName,
		FirstName,
		MiddleName,
		Address,
		BirthDate,
		BirthPlace,
		Gender,
		CivilStatus,
		ContactNumber,
		GuardianName,
		GurdianContactNumbers,
		ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func DeleteResidents(db sqlx.Ext, ID int64) error {

	_, err := db.Exec(`DELETE FROM Residents WHERE ID = ? `, ID)

	if err != nil {
		return err
	}

	return nil
}

func ReadResidents(db sqlx.Ext) ([]*Residents, error) {

	residentsArray := make([]*Residents, 0)
	var ID int64
	var DateCreated string
	var DateUpdated string
	var LastName string
	var FirstName string
	var MiddleName string
	var Address string
	var BirthDate string
	var BirthPlace string
	var Gender string
	var CivilStatus string
	var ContactNumber string
	var GuardianName string
	var GurdianContactNumbers string
	var IssuingOfficer string
	var DocumentType string
	var DocumentID int64

	rows, err := db.Queryx(`SELECT ID,
				DateCreated,
				DateUpdated,
				LastName,
				FirstName,
				MiddleName,
				Address,
				BirthDate,
				BirthPlace,
				Gender,
				CivilStatus,
				ContactNumber,
				GuardianName,
				GurdianContactNumber,
				IssuingOfficer,
				DocumentType,
				DocumentID FROM Residents`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ID, &DateCreated, &DateUpdated, &LastName, &FirstName, &MiddleName, &Address, &BirthDate, &BirthPlace,
			&Gender, &CivilStatus, &ContactNumber, &GuardianName, &GurdianContactNumbers, &IssuingOfficer,
			&DocumentType, &DocumentID)
		if err != nil {
			return nil, err
		}
		var documentData interface{}
		if DocumentType == "Barangay Clearance" {
			clearanceData, err := ReadClearanceData(db, DocumentID)
			if err != nil {
				return nil, err
			}
			documentData = clearanceData
		} else if DocumentType == "Barangay Indigency" {
			indigencyData, err := ReadIndigencyData(db, DocumentID)
			if err != nil {
				return nil, err
			}
			documentData = indigencyData
		} else if DocumentType == "Referral Slip" {
			referralData, err := ReadReferralsData(db, DocumentID)
			if err != nil {
				return nil, err
			}
			documentData = referralData
		}

		residentsArray = append(residentsArray, &Residents{
			ID:                    ID,
			DateCreated:           DateCreated,
			DateUpdated:           DateUpdated,
			LastName:              LastName,
			FirstName:             FirstName,
			MiddleName:            MiddleName,
			Address:               Address,
			BirthDate:             BirthDate,
			BirthPlace:            BirthPlace,
			Gender:                Gender,
			CivilStatus:           CivilStatus,
			ContactNumber:         ContactNumber,
			GuardianName:          GuardianName,
			GurdianContactNumbers: GurdianContactNumbers,
			IssuingOfficer:        IssuingOfficer,
			DocumentType:          DocumentType,
			DocumentID:            DocumentID,
			DocumentData:          documentData,
		})
	}
	return residentsArray, nil
}

func ReadResidentData(db sqlx.Ext, residentID int64) (*Residents, error) {

	var ID int64
	var DateCreated string
	var DateUpdated string
	var LastName string
	var FirstName string
	var MiddleName string
	var Address string
	var BirthDate string
	var BirthPlace string
	var Gender string
	var CivilStatus string
	var ContactNumber string
	var GuardianName string
	var GurdianContactNumbers string
	var IssuingOfficer string
	var DocumentType string
	var DocumentID int64

	rows, err := db.Queryx(`SELECT ID,
				DateCreated,
				DateUpdated,
				LastName,
				FirstName,
				MiddleName,
				Address,
				BirthDate,
				BirthPlace,
				Gender,
				CivilStatus,
				ContactNumber,
				GuardianName,
				GurdianContactNumber,
				IssuingOfficer,
				DocumentType,
				DocumentID FROM Residents WHERE ID = ?`, residentID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ID, &DateCreated, &DateUpdated, &LastName, &FirstName, &MiddleName, &Address, &BirthDate, &BirthPlace,
			&Gender, &CivilStatus, &ContactNumber, &GuardianName, &GurdianContactNumbers, &IssuingOfficer,
			&DocumentType, &DocumentID)
		if err != nil {
			return nil, err
		}
	}
	return &Residents{
		ID:                    ID,
		DateCreated:           DateCreated,
		DateUpdated:           DateUpdated,
		LastName:              LastName,
		FirstName:             FirstName,
		MiddleName:            MiddleName,
		Address:               Address,
		BirthDate:             BirthDate,
		BirthPlace:            BirthPlace,
		Gender:                Gender,
		CivilStatus:           CivilStatus,
		ContactNumber:         ContactNumber,
		GuardianName:          GuardianName,
		GurdianContactNumbers: GurdianContactNumbers,
		IssuingOfficer:        IssuingOfficer,
		DocumentType:          DocumentType,
		DocumentID:            DocumentID,
	}, nil

}

func UpdateDocumentID(db sqlx.Ext, DocumentID int64, ResidentID int64) error {

	_, err := db.Exec(`UPDATE Residents SET 
		DocumentID = ? WHERE ID= ?`,
		DocumentID,
		ResidentID,
	)

	if err != nil {
		return err
	}

	return nil
}

func ReadResidentsXL(db sqlx.Ext) ([]*ResidentsXL, error) {

	residentsArray := make([]*ResidentsXL, 0)
	var ID int64
	var DateCreated string
	var DateUpdated string
	var LastName string
	var FirstName string
	var MiddleName string
	var Address string
	var BirthDate string
	var BirthPlace string
	var Gender string
	var CivilStatus string
	var ContactNumber string
	var GuardianName string
	var GurdianContactNumbers string
	var IssuingOfficer string
	var DocumentType string
	var DocumentID int64

	rows, err := db.Queryx(`SELECT ID,
				DateCreated,
				DateUpdated,
				LastName,
				FirstName,
				MiddleName,
				Address,
				BirthDate,
				BirthPlace,
				Gender,
				CivilStatus,
				ContactNumber,
				GuardianName,
				GurdianContactNumber,
				IssuingOfficer,
				DocumentType,
				DocumentID FROM Residents`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ID, &DateCreated, &DateUpdated, &LastName, &FirstName, &MiddleName, &Address, &BirthDate, &BirthPlace,
			&Gender, &CivilStatus, &ContactNumber, &GuardianName, &GurdianContactNumbers, &IssuingOfficer,
			&DocumentType, &DocumentID)
		if err != nil {
			return nil, err
		}
		if DocumentType == "Barangay Clearance" {
			clearanceData, err := ReadClearanceData(db, DocumentID)
			if err != nil {
				return nil, err
			}

			residentsArray = append(residentsArray, &ResidentsXL{
				ID:                    ID,
				DateCreated:           DateCreated,
				DateUpdated:           DateUpdated,
				LastName:              LastName,
				FirstName:             FirstName,
				MiddleName:            MiddleName,
				Address:               Address,
				BirthDate:             BirthDate,
				BirthPlace:            BirthPlace,
				Gender:                Gender,
				CivilStatus:           CivilStatus,
				ContactNumber:         ContactNumber,
				GuardianName:          GuardianName,
				GurdianContactNumbers: GurdianContactNumbers,
				IssuingOfficer:        IssuingOfficer,
				DocumentType:          DocumentType,
				DocumentID:            DocumentID,
				ResidentID:            clearanceData.ResidentID,
				DocumentDateCreated:   clearanceData.DateCreated,
				DocumentDateUpdated:   clearanceData.DateUpdated,
				Purpose:               clearanceData.Purpose,
				ValidUntil:            clearanceData.ValidUntil,
				Remarks:               clearanceData.Remarks,
				DocumentStatus:        clearanceData.DocumentStatus,
				CedulaNo:              clearanceData.CedulaNo,
				PrecintNo:             clearanceData.PrecintNo,
			})

		} else if DocumentType == "Barangay Indigency" {
			indigencyData, err := ReadIndigencyData(db, DocumentID)
			if err != nil {
				return nil, err
			}
			residentsArray = append(residentsArray, &ResidentsXL{
				ID:                    ID,
				DateCreated:           DateCreated,
				DateUpdated:           DateUpdated,
				LastName:              LastName,
				FirstName:             FirstName,
				MiddleName:            MiddleName,
				Address:               Address,
				BirthDate:             BirthDate,
				BirthPlace:            BirthPlace,
				Gender:                Gender,
				CivilStatus:           CivilStatus,
				ContactNumber:         ContactNumber,
				GuardianName:          GuardianName,
				GurdianContactNumbers: GurdianContactNumbers,
				IssuingOfficer:        IssuingOfficer,
				DocumentType:          DocumentType,
				DocumentID:            DocumentID,
				ResidentID:            indigencyData.ResidentID,
				DocumentDateCreated:   indigencyData.DateCreated,
				DocumentDateUpdated:   indigencyData.DateUpdated,
				ValidUntil:            indigencyData.ValidUntil,
				Remarks:               indigencyData.Remarks,
				DocumentStatus:        indigencyData.DocumentStatus,
				Reason:                indigencyData.Reason,
			})
		} else if DocumentType == "Referral Slip" {
			referralData, err := ReadReferralsData(db, DocumentID)
			if err != nil {
				return nil, err
			}
			residentsArray = append(residentsArray, &ResidentsXL{
				ID:                    ID,
				DateCreated:           DateCreated,
				DateUpdated:           DateUpdated,
				LastName:              LastName,
				FirstName:             FirstName,
				MiddleName:            MiddleName,
				Address:               Address,
				BirthDate:             BirthDate,
				BirthPlace:            BirthPlace,
				Gender:                Gender,
				CivilStatus:           CivilStatus,
				ContactNumber:         ContactNumber,
				GuardianName:          GuardianName,
				GurdianContactNumbers: GurdianContactNumbers,
				IssuingOfficer:        IssuingOfficer,
				DocumentType:          DocumentType,
				DocumentID:            DocumentID,
				ResidentID:            referralData.ResidentID,
				DocumentDateCreated:   referralData.DateCreated,
				DocumentDateUpdated:   referralData.DateUpdated,
				ValidUntil:            referralData.ValidUntil,
				Remarks:               referralData.Remarks,
				DocumentStatus:        referralData.DocumentStatus,
				HCGGGNumber:           referralData.HCGGGNumber,
				PhilHealthID:          referralData.PhilHealthID,
				PhilHealthCategory:    referralData.PhilHealthCategory,
				ReasonForReferral:     referralData.ReasonForReferral,
			})
		}

	}
	return residentsArray, nil
}
