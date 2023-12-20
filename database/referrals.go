package database

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Referrals struct {
	ID                 int64      `json:"ID"`
	ResidentID         int64      `json:"ResidentID"`
	DateCreated        string     `json:"DateCreated"`
	DateUpdated        string     `json:"DateUpdated"`
	HCGGGNumber        string     `json:"HCGGGNumber"`
	PhilHealthID       string     `json:"PhilHealthID"`
	PhilHealthCategory string     `json:"PhilHealthCategory"`
	ReasonForReferral  string     `json:"ReasonForReferral"`
	ValidUntil         string     `json:"ValidUntil"`
	IssuingOfficer     string     `json:"IssuingOfficer"`
	Remarks            string     `json:"Remarks"`
	DocumentStatus     string     `json:"DocumentStatus"`
	Resident           *Residents `json:"ResidentData"`
}

type ReferralsXL struct {
	ID                    int64  `json:"ID"`
	ResidentID            int64  `json:"ResidentID"`
	DateCreated           string `json:"DateCreated"`
	DateUpdated           string `json:"DateUpdated"`
	HCGGGNumber           string `json:"HCGGGNumber"`
	PhilHealthID          string `json:"PhilHealthID"`
	PhilHealthCategory    string `json:"PhilHealthCategory"`
	ReasonForReferral     string `json:"ReasonForReferral"`
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

func CreateReferrals(db sqlx.Ext, ResidentID int64, HCGGGNumber string, PhilHealthID string, PhilHealthCategory string,
	ReasonForReferral string, ValidUntil string, IssuingOfficer string, Remarks string) (int64, error) {

	currentTime := time.Now()
	// Format the time as "YYYY-MM-DD 03:04 PM"
	formattedTime := currentTime.Format("2006-01-02 03:04 PM")

	query, err := db.Exec(`INSERT INTO Referrals (
		ResidentID,
		DateCreated,
		DateUpdated,
		HCGGGNumber,
		PhilHealthID,
		PhilHealthCategory,
		ReasonForReferral,
		ValidUntil,
		IssuingOfficer,
		Remarks,
		DocumentStatus
	)
	Values(?,?,?,?,?,?,?,?,?,?,?)`,
		ResidentID,
		formattedTime,
		formattedTime,
		HCGGGNumber,
		PhilHealthID,
		PhilHealthCategory,
		ReasonForReferral,
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

func UpdateReferrals(db sqlx.Ext, ID int64, HCGGGNumber string, PhilHealthID string, PhilHealthCategory string,
	ReasonForReferral string, Remarks string, DocumentStatus string) error {

	currentTime := time.Now()
	// Format the time as "YYYY-MM-DD 03:04 PM"
	formattedTime := currentTime.Format("2006-01-02 03:04 PM")

	_, err := db.Exec(`UPDATE Referrals SET 
		DateUpdated = ?,
		HCGGGNumber = ?,
		PhilHealthID = ?,
		PhilHealthCategory = ?,
		ReasonForReferral = ?,
		Remarks = ? ,
		DocumentStatus = ? WHERE ID= ?`,
		formattedTime,
		HCGGGNumber,
		PhilHealthID,
		PhilHealthCategory,
		ReasonForReferral,
		Remarks,
		DocumentStatus,
		ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func DeleteReferrals(db sqlx.Ext, ID int64) error {

	_, err := db.Exec(`DELETE FROM Referrals WHERE ID = ? `, ID)

	if err != nil {
		return err
	}

	return nil
}

func ReadReferrals(db sqlx.Ext) ([]*Referrals, error) {

	referralsArray := make([]*Referrals, 0)
	var ID int64
	var ResidentID int64
	var DateCreated string
	var DateUpdated string
	var HCGGGNumber string
	var PhilHealthID string
	var PhilHealthCategory string
	var ReasonForReferral string
	var ValidUntil string
	var IssuingOfficer string
	var Remarks string
	var DocumentStatus string

	rows, err := db.Queryx(`SELECT ID,
				ResidentID,
				DateCreated,
				DateUpdated,
				HCGGGNumber,
				PhilHealthID,
				PhilHealthCategory,
				ReasonForReferral,
				ValidUntil,
				IssuingOfficer,
				Remarks,
				DocumentStatus FROM Referrals`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ID, &ResidentID, &DateCreated, &DateUpdated, &HCGGGNumber, &PhilHealthID, &PhilHealthCategory, &ReasonForReferral, &ValidUntil,
			&IssuingOfficer, &Remarks, &DocumentStatus)
		if err != nil {
			return nil, err
		}
		residentData, err := ReadResidentData(db, ResidentID)
		if err != nil {
			return nil, err
		}
		referralsArray = append(referralsArray, &Referrals{
			ID:                 ID,
			ResidentID:         ResidentID,
			DateCreated:        DateCreated,
			DateUpdated:        DateUpdated,
			HCGGGNumber:        HCGGGNumber,
			PhilHealthID:       PhilHealthID,
			PhilHealthCategory: PhilHealthCategory,
			ReasonForReferral:  ReasonForReferral,
			ValidUntil:         ValidUntil,
			IssuingOfficer:     IssuingOfficer,
			Remarks:            Remarks,
			DocumentStatus:     DocumentStatus,
			Resident:           residentData,
		})
	}
	return referralsArray, nil
}

func ReadReferralsData(db sqlx.Ext, DocumentID int64) (*Referrals, error) {

	var ID int64
	var ResidentID int64
	var DateCreated string
	var DateUpdated string
	var HCGGGNumber string
	var PhilHealthID string
	var PhilHealthCategory string
	var ReasonForReferral string
	var ValidUntil string
	var IssuingOfficer string
	var Remarks string
	var DocumentStatus string

	rows, err := db.Queryx(`SELECT ID,
	ResidentID,
	DateCreated,
	DateUpdated,
	HCGGGNumber,
	PhilHealthID,
	PhilHealthCategory,
	ReasonForReferral,
	ValidUntil,
	IssuingOfficer,
	Remarks,
	DocumentStatus FROM Referrals WHERE ID = ?`, DocumentID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ID, &ResidentID, &DateCreated, &DateUpdated, &HCGGGNumber, &PhilHealthID, &PhilHealthCategory, &ReasonForReferral,
			&ValidUntil, &IssuingOfficer, &Remarks, &DocumentStatus)
		if err != nil {
			return nil, err
		}
	}
	return &Referrals{
		ID:                 ID,
		ResidentID:         ResidentID,
		DateCreated:        DateCreated,
		DateUpdated:        DateUpdated,
		HCGGGNumber:        HCGGGNumber,
		PhilHealthID:       PhilHealthID,
		PhilHealthCategory: PhilHealthCategory,
		ReasonForReferral:  ReasonForReferral,
		ValidUntil:         ValidUntil,
		IssuingOfficer:     IssuingOfficer,
		Remarks:            Remarks,
		DocumentStatus:     DocumentStatus,
	}, nil
}

func ReadReferralsXL(db sqlx.Ext) ([]*ReferralsXL, error) {

	referralsArray := make([]*ReferralsXL, 0)
	var ID int64
	var ResidentID int64
	var DateCreated string
	var DateUpdated string
	var HCGGGNumber string
	var PhilHealthID string
	var PhilHealthCategory string
	var ReasonForReferral string
	var ValidUntil string
	var IssuingOfficer string
	var Remarks string
	var DocumentStatus string

	rows, err := db.Queryx(`SELECT ID,
				ResidentID,
				DateCreated,
				DateUpdated,
				HCGGGNumber,
				PhilHealthID,
				PhilHealthCategory,
				ReasonForReferral,
				ValidUntil,
				IssuingOfficer,
				Remarks,
				DocumentStatus FROM Referrals`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ID, &ResidentID, &DateCreated, &DateUpdated, &HCGGGNumber, &PhilHealthID, &PhilHealthCategory, &ReasonForReferral, &ValidUntil,
			&IssuingOfficer, &Remarks, &DocumentStatus)
		if err != nil {
			return nil, err
		}
		residentData, err := ReadResidentData(db, ResidentID)
		if err != nil {
			return nil, err
		}
		referralsArray = append(referralsArray, &ReferralsXL{
			ID:                    ID,
			ResidentID:            ResidentID,
			DateCreated:           DateCreated,
			DateUpdated:           DateUpdated,
			HCGGGNumber:           HCGGGNumber,
			PhilHealthID:          PhilHealthID,
			PhilHealthCategory:    PhilHealthCategory,
			ReasonForReferral:     ReasonForReferral,
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
	return referralsArray, nil
}
