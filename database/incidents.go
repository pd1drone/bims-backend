package database

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// `ID` int PRIMARY KEY AUTO_INCREMENT,
// `DateCreated` varchar(255),
// `DateUpdated` varchar(255),
// `CompliantFullName` varchar(255),
// `Respondent` varchar(255),
// `IncidentStatus` varchar(255),
// `IncidentDateTime` varchar(255),
// `IncidentLocation` varchar(255),
// `IncidentNarration` longtext,
// `IssuingOfficer` varchar(255)
type Incidents struct {
	ID                int64  `json:"ID"`
	DateCreated       string `json:"DateCreated"`
	DateUpdated       string `json:"DateUpdated"`
	CompliantFullName string `json:"CompliantFullName"`
	Respondent        string `json:"Respondent"`
	IncidentStatus    string `json:"IncidentStatus"`
	IncidentDateTime  string `json:"IncidentDateTime"`
	IncidentLocation  string `json:"IncidentLocation"`
	IncidentNarration string `json:"IncidentNarration"`
	IssuingOfficer    string `json:"IssuingOfficer"`
}

func CreateIncidents(db sqlx.Ext, CompliantFullName string, Respondent string, IncidentStatus string, IncidentDateTime string, IncidentLocation string, IncidentNarration string, issuingOfficer string) (int64, error) {
	currentTime := time.Now()
	// Format the time as "YYYY-MM-DD 03:04 PM"
	formattedTime := currentTime.Format("2006-01-02 03:04 PM")
	query, err := db.Exec(`INSERT INTO Incidents (
		DateCreated,
		DateUpdated,
		CompliantFullName,
		Respondent,
		IncidentStatus,
		IncidentDateTime,
		IncidentLocation,
		IncidentNarration,
		IssuingOfficer
	)
	Values(?,?,?,?,?,?,?,?,?)`,
		formattedTime,
		formattedTime,
		CompliantFullName,
		Respondent,
		IncidentStatus,
		IncidentDateTime,
		IncidentLocation,
		IncidentNarration,
		issuingOfficer,
	)

	if err != nil {
		return 0, err
	}
	incidentID, err := query.LastInsertId()

	if err != nil {
		return 0, err
	}

	return incidentID, nil
}

func DeleteIncidents(db sqlx.Ext, ID int64) error {

	_, err := db.Exec(`DELETE FROM Incidents WHERE ID = ? `, ID)

	if err != nil {
		return err
	}

	return nil
}

func UpdateIncidents(db sqlx.Ext, ID int64, CompliantFullName string, Respondent string, IncidentStatus string, IncidentDateTime string, IncidentLocation string, IncidentNarration string, issuingOfficer string) error {
	currentTime := time.Now()
	// Format the time as "YYYY-MM-DD 03:04 PM"
	formattedTime := currentTime.Format("2006-01-02 03:04 PM")

	_, err := db.Exec(`UPDATE Incidents SET 
		DateUpdated = ?,
		CompliantFullName = ?,
		Respondent = ?,
		IncidentStatus = ?,
		IncidentDateTime = ?,
		IncidentLocation = ?,
		IncidentNarration = ?,
		IssuingOfficer =? WHERE ID= ?`,
		formattedTime,
		CompliantFullName,
		Respondent,
		IncidentStatus,
		IncidentDateTime,
		IncidentLocation,
		IncidentNarration,
		issuingOfficer,
		ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func ReadIncidents(db sqlx.Ext) ([]*Incidents, error) {

	incidentArray := make([]*Incidents, 0)
	var ID int64
	var DateCreated string
	var DateUpdated string
	var CompliantFullName string
	var Respondent string
	var IncidentStatus string
	var IncidentDateTime string
	var IncidentLocation string
	var IncidentNarration string
	var IssuingOfficer string

	rows, err := db.Queryx(`SELECT ID,
				DateCreated,
				DateUpdated,				
				CompliantFullName,
				Respondent,
				IncidentStatus,
				IncidentDateTime,
				IncidentLocation,
				IncidentNarration,
				IssuingOfficer FROM Incidents`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ID, &DateCreated, &DateUpdated, &CompliantFullName, &Respondent, &IncidentStatus, &IncidentDateTime, &IncidentLocation, &IncidentNarration, &IssuingOfficer)
		if err != nil {
			return nil, err
		}

		incidentArray = append(incidentArray, &Incidents{
			ID:                ID,
			DateCreated:       DateCreated,
			DateUpdated:       DateUpdated,
			CompliantFullName: CompliantFullName,
			Respondent:        Respondent,
			IncidentStatus:    IncidentStatus,
			IncidentDateTime:  IncidentDateTime,
			IncidentLocation:  IncidentLocation,
			IncidentNarration: IncidentNarration,
			IssuingOfficer:    IssuingOfficer,
		})

	}
	return incidentArray, nil
}
