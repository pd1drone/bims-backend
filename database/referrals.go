package database

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Referrals struct {
	ID                 int64  `json:"ID"`
	ResidentID         int64  `json:"ResidentID"`
	DateCreated        string `json:"DateCreated"`
	DateUpdated        string `json:"DateUpdated"`
	HCGGGNumber        string `json:"HCGGGNumber"`
	PhilHealthID       string `json:"PhilHealthID"`
	PhilHealthCategory string `json:"PhilHealthCategory"`
	ReasonForReferral  string `json:"ReasonForReferral"`
	ValidUntil         string `json:"ValidUntil"`
	IssuingOfficer     string `json:"IssuingOfficer"`
	Remarks            string `json:"Remarks"`
}

func CreateReferrals(db sqlx.Ext, ResidentID int64, HCGGGNumber string, PhilHealthID string, PhilHealthCategory string,
	ReasonForReferral string, ValidUntil string, IssuingOfficer string, Remarks string) error {

	currentTime := time.Now()
	// Format the time as "YYYY-MM-DD 03:04 PM"
	formattedTime := currentTime.Format("2006-01-02 03:04 PM")

	_, err := db.Exec(`INSERT INTO Referrals (
		ResidentID,
		DateCreated,
		DateUpdated,
		HCGGGNumber,
		PhilHealthID,
		PhilHealthCategory,
		ReasonForReferral,
		ValidUntil,
		IssuingOfficer,
		Remarks
	)
	Values(?,?,?,?,?,?,?,?,?,?)`,
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
	)

	if err != nil {
		return err
	}

	return nil
}

func UpdateReferrals(db sqlx.Ext, ID int64, HCGGGNumber string, PhilHealthID string, PhilHealthCategory string,
	ReasonForReferral string, Remarks string) error {

	currentTime := time.Now()
	// Format the time as "YYYY-MM-DD 03:04 PM"
	formattedTime := currentTime.Format("2006-01-02 03:04 PM")

	_, err := db.Exec(`UPDATE Referrals SET 
		DateUpdated = ?,
		HCGGGNumber = ?,
		PhilHealthID = ?,
		PhilHealthCategory = ?,
		ReasonForReferral = ?,
		Remarks = ? WHERE ID= ?`,
		formattedTime,
		HCGGGNumber,
		PhilHealthID,
		PhilHealthCategory,
		ReasonForReferral,
		Remarks,
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
				Remarks FROM Referrals`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ID, &ResidentID, &DateCreated, &DateUpdated, &HCGGGNumber, &PhilHealthID, &PhilHealthCategory, &ReasonForReferral, &ValidUntil,
			&IssuingOfficer, &Remarks)
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
		})
	}
	return referralsArray, nil
}
