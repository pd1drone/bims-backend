package database

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Residents struct {
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
	Religion              string `json:"Religion"`
	Occupation            string `json:"Occupation"`
	IssuingOfficer        string `json:"IssuingOfficer"`
}

func CreateResident(db sqlx.Ext, LastName string, FirstName string, MiddleName string, Address string, BirthDate string,
	BirthPlace string, Gender string, CivilStatus string, ContactNumber string, GuardianName string, GurdianContactNumbers string,
	Religion string, Occupation string, IssuingOfficer string) (int64, error) {

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
		GurdianContactNumbers,
		Religion,
		Occupation,
		IssuingOfficer
	)
	Values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
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
		Religion,
		Occupation,
		IssuingOfficer,
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
	BirthPlace string, Gender string, CivilStatus string, ContactNumber string, GuardianName string, GurdianContactNumbers string,
	Religion string, Occupation string) error {

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
		GurdianContactNumbers = ?,
		Religion = ?,
		Occupation = ? WHERE ID= ?`,
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
		Religion,
		Occupation,
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
	var Religion string
	var Occupation string
	var IssuingOfficer string

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
				GurdianContactNumbers,
				Religion,
				Occupation,
				IssuingOfficer FROM Residents`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ID, &DateCreated, &DateUpdated, &LastName, &FirstName, &MiddleName, &Address, &BirthDate, &BirthPlace,
			&Gender, &CivilStatus, &ContactNumber, &GuardianName, &GurdianContactNumbers, &Religion, &Occupation, IssuingOfficer)
		if err != nil {
			return nil, err
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
			Religion:              Religion,
			Occupation:            Occupation,
			IssuingOfficer:        IssuingOfficer,
		})
	}
	return residentsArray, nil
}
