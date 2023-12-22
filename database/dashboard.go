package database

import (
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

type TotalClearanceMonthly struct {
	ID          int64  `json:"ID"`
	DateCreated string `json:"DateCreated"`
}

type TotalReferralsMonthly struct {
	ID          int64  `json:"ID"`
	DateCreated string `json:"DateCreated"`
}

type TotalIndigenciesMonthly struct {
	ID          int64  `json:"ID"`
	DateCreated string `json:"DateCreated"`
}

type TotalBDRRMCMonthly struct {
	ID          int64  `json:"ID"`
	DateCreated string `json:"DateCreated"`
}

type TotalIcidentsMonthly struct {
	ID          int64  `json:"ID"`
	DateCreated string `json:"DateCreated"`
}

func GetTotalClearancePerMonth(db sqlx.Ext) ([]*TotalClearanceMonthly, error) {

	clearanceData := make([]*TotalClearanceMonthly, 0)
	var ID int64
	var DateCreated string
	rows, err := db.Queryx(`SELECT DISTINCT(ID),DateCreated FROM Clearance`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ID, &DateCreated)
		if err != nil {
			return nil, err
		}
		DateOnly := strings.Split(DateCreated, " ")
		clearanceData = append(clearanceData, &TotalClearanceMonthly{
			ID:          ID,
			DateCreated: DateOnly[0],
		})
	}

	// Get the current month and year
	currentTime := time.Now()
	currentYear, currentMonth, _ := currentTime.Date()

	// Create a new slice to store clearance data for the current month
	newClearanceData := make([]*TotalClearanceMonthly, 0)

	for _, item := range clearanceData {
		// Parse the DateCreated string into a time.Time object
		dateCreated, err := time.Parse("2006-01-02", item.DateCreated)
		if err != nil {
			fmt.Println("Error parsing DateCreated:", err)
			continue // Skip this item if there's an error parsing the date
		}

		// Check if the DateCreated is in the current month and year
		if dateCreated.Year() == currentYear && dateCreated.Month() == currentMonth {
			newClearanceData = append(newClearanceData, item)
		}
	}

	return newClearanceData, nil

}

func GetTotalReferralsPerMonth(db sqlx.Ext) ([]*TotalReferralsMonthly, error) {

	referralsData := make([]*TotalReferralsMonthly, 0)
	var ID int64
	var DateCreated string
	rows, err := db.Queryx(`SELECT DISTINCT(ID),DateCreated FROM Referrals`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ID, &DateCreated)
		if err != nil {
			return nil, err
		}
		DateOnly := strings.Split(DateCreated, " ")
		referralsData = append(referralsData, &TotalReferralsMonthly{
			ID:          ID,
			DateCreated: DateOnly[0],
		})
	}

	// Get the current month and year
	currentTime := time.Now()
	currentYear, currentMonth, _ := currentTime.Date()

	// Create a new slice to store clearance data for the current month
	newReferralsData := make([]*TotalReferralsMonthly, 0)

	for _, item := range referralsData {
		// Parse the DateCreated string into a time.Time object
		dateCreated, err := time.Parse("2006-01-02", item.DateCreated)
		if err != nil {
			fmt.Println("Error parsing DateCreated:", err)
			continue // Skip this item if there's an error parsing the date
		}

		// Check if the DateCreated is in the current month and year
		if dateCreated.Year() == currentYear && dateCreated.Month() == currentMonth {
			newReferralsData = append(newReferralsData, item)
		}
	}

	return newReferralsData, nil
}

func GetTotalIndigenciesPerMonth(db sqlx.Ext) ([]*TotalIndigenciesMonthly, error) {

	indigenciesData := make([]*TotalIndigenciesMonthly, 0)
	var ID int64
	var DateCreated string
	rows, err := db.Queryx(`SELECT DISTINCT(ID),DateCreated FROM Indigencies`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ID, &DateCreated)
		if err != nil {
			return nil, err
		}
		DateOnly := strings.Split(DateCreated, " ")
		indigenciesData = append(indigenciesData, &TotalIndigenciesMonthly{
			ID:          ID,
			DateCreated: DateOnly[0],
		})
	}

	// Get the current month and year
	currentTime := time.Now()
	currentYear, currentMonth, _ := currentTime.Date()

	// Create a new slice to store clearance data for the current month
	newIndigenciesData := make([]*TotalIndigenciesMonthly, 0)

	for _, item := range indigenciesData {
		// Parse the DateCreated string into a time.Time object
		dateCreated, err := time.Parse("2006-01-02", item.DateCreated)
		if err != nil {
			fmt.Println("Error parsing DateCreated:", err)
			continue // Skip this item if there's an error parsing the date
		}

		// Check if the DateCreated is in the current month and year
		if dateCreated.Year() == currentYear && dateCreated.Month() == currentMonth {
			newIndigenciesData = append(newIndigenciesData, item)
		}
	}

	return newIndigenciesData, nil
}

func GetTotalResidentsPerMonth(db sqlx.Ext) ([]*TotalClearanceMonthly, error) {

	clearanceData := make([]*TotalClearanceMonthly, 0)
	var ID int64
	var DateCreated string
	rows, err := db.Queryx(`SELECT DISTINCT(ID),DateCreated FROM Residents`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ID, &DateCreated)
		if err != nil {
			return nil, err
		}
		DateOnly := strings.Split(DateCreated, " ")
		clearanceData = append(clearanceData, &TotalClearanceMonthly{
			ID:          ID,
			DateCreated: DateOnly[0],
		})
	}

	// Get the current month and year
	currentTime := time.Now()
	currentYear, currentMonth, _ := currentTime.Date()

	// Create a new slice to store clearance data for the current month
	newClearanceData := make([]*TotalClearanceMonthly, 0)

	for _, item := range clearanceData {
		// Parse the DateCreated string into a time.Time object
		dateCreated, err := time.Parse("2006-01-02", item.DateCreated)
		if err != nil {
			fmt.Println("Error parsing DateCreated:", err)
			continue // Skip this item if there's an error parsing the date
		}

		// Check if the DateCreated is in the current month and year
		if dateCreated.Year() == currentYear && dateCreated.Month() == currentMonth {
			newClearanceData = append(newClearanceData, item)
		}
	}

	return newClearanceData, nil

}

func GetTotalBDRRMCPerMonth(db sqlx.Ext) ([]*TotalBDRRMCMonthly, error) {

	bdrrmcData := make([]*TotalBDRRMCMonthly, 0)
	var ID int64
	var DateCreated string
	rows, err := db.Queryx(`SELECT DISTINCT(ID),DateCreated FROM BDRRMC`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ID, &DateCreated)
		if err != nil {
			return nil, err
		}
		DateOnly := strings.Split(DateCreated, " ")
		bdrrmcData = append(bdrrmcData, &TotalBDRRMCMonthly{
			ID:          ID,
			DateCreated: DateOnly[0],
		})
	}

	// Get the current month and year
	currentTime := time.Now()
	currentYear, currentMonth, _ := currentTime.Date()

	// Create a new slice to store clearance data for the current month
	newBDRRMCData := make([]*TotalBDRRMCMonthly, 0)

	for _, item := range bdrrmcData {
		// Parse the DateCreated string into a time.Time object
		dateCreated, err := time.Parse("2006-01-02", item.DateCreated)
		if err != nil {
			fmt.Println("Error parsing DateCreated:", err)
			continue // Skip this item if there's an error parsing the date
		}

		// Check if the DateCreated is in the current month and year
		if dateCreated.Year() == currentYear && dateCreated.Month() == currentMonth {
			newBDRRMCData = append(newBDRRMCData, item)
		}
	}

	return newBDRRMCData, nil
}

func GetTotalIncidentsPerMonth(db sqlx.Ext) ([]*TotalIndigenciesMonthly, error) {

	incidentsData := make([]*TotalIndigenciesMonthly, 0)
	var ID int64
	var DateCreated string
	rows, err := db.Queryx(`SELECT DISTINCT(ID),DateCreated FROM Incidents`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ID, &DateCreated)
		if err != nil {
			return nil, err
		}
		DateOnly := strings.Split(DateCreated, " ")
		incidentsData = append(incidentsData, &TotalIndigenciesMonthly{
			ID:          ID,
			DateCreated: DateOnly[0],
		})
	}

	// Get the current month and year
	currentTime := time.Now()
	currentYear, currentMonth, _ := currentTime.Date()

	// Create a new slice to store clearance data for the current month
	newIncidentsData := make([]*TotalIndigenciesMonthly, 0)

	for _, item := range incidentsData {
		// Parse the DateCreated string into a time.Time object
		dateCreated, err := time.Parse("2006-01-02", item.DateCreated)
		if err != nil {
			fmt.Println("Error parsing DateCreated:", err)
			continue // Skip this item if there's an error parsing the date
		}

		// Check if the DateCreated is in the current month and year
		if dateCreated.Year() == currentYear && dateCreated.Month() == currentMonth {
			newIncidentsData = append(newIncidentsData, item)
		}
	}

	return newIncidentsData, nil
}
