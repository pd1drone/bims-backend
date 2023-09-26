package rest

import (
	"bims/database"
	"net/http"
	"time"
)

type ClearanceMonthlyGraphData struct {
	Labels   []string    `json:"Labels"`
	DataSets []*DataSets `json:"DataSets"`
}

type DataSets struct {
	Label string `json:"Label"`
	Data  []int  `json:"Data"`
}

type TotalNumberOfCreatedDocumentsPerMonth struct {
	Indigencies int `json:"Indigencies"`
	Referrals   int `json:"Referrals"`
	Clearance   int `json:"Clearance"`
	Residents   int `json:"Residents"`
}

func (b *BimsConfiguration) GetTotalNumberOfCreatedDocumentsPerMonth(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	Clearance, err := database.GetTotalClearancePerMonth(b.BIMSdb)
	if err != nil {
		respondJSON(w, 400, nil)
	}

	Indigencies, err := database.GetTotalIndigenciesPerMonth(b.BIMSdb)
	if err != nil {
		respondJSON(w, 400, nil)
	}

	Referrals, err := database.GetTotalReferralsPerMonth(b.BIMSdb)
	if err != nil {
		respondJSON(w, 400, nil)
	}

	Residents, err := database.GetTotalResidentsPerMonth(b.BIMSdb)
	if err != nil {
		respondJSON(w, 400, nil)
	}

	resp := &TotalNumberOfCreatedDocumentsPerMonth{
		Indigencies: len(Indigencies),
		Referrals:   len(Referrals),
		Clearance:   len(Clearance),
		Residents:   len(Residents),
	}

	respondJSON(w, 200, resp)
}

func (b *BimsConfiguration) ReadMonthlyTotalGraph(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	resp := &ClearanceMonthlyGraphData{}
	labels, IndigencyCount, err := b.MonthlyIndigencies()
	if err != nil {
		respondJSON(w, 400, nil)
	}
	resp.Labels = labels

	Dataset := make([]*DataSets, 0)

	Dataset = append(Dataset, &DataSets{
		Label: "Indigencies",
		Data:  IndigencyCount,
	})

	_, ClearanceCount, err := b.MonthlyClearance()
	if err != nil {
		respondJSON(w, 400, nil)
	}

	Dataset = append(Dataset, &DataSets{
		Label: "Clearance",
		Data:  ClearanceCount,
	})

	_, ReferralsCount, err := b.MonthlyReferrals()
	if err != nil {
		respondJSON(w, 400, nil)
	}

	Dataset = append(Dataset, &DataSets{
		Label: "Referrals",
		Data:  ReferralsCount,
	})

	resp.DataSets = Dataset

	respondJSON(w, 200, resp)
}

func (b *BimsConfiguration) MonthlyIndigencies() ([]string, []int, error) {
	IndigenciesData, err := database.GetTotalIndigenciesPerMonth(b.BIMSdb)
	if err != nil {
		return nil, nil, err
	}

	// Initialize arrays for labels and count data
	var Labels []string
	var countData []int

	// Calculate the date range from "2023-09-01" to today
	startDate, _ := time.Parse("2006-01-02", "2023-09-01")
	endDate := time.Now()

	// Loop through the date range
	currentDate := startDate
	for currentDate.Before(endDate) || currentDate.Equal(endDate) {
		// Format the date to "Sep 1" style
		formattedDate := currentDate.Format("Jan 2")

		// Count the entries for the current date
		count := 0
		for _, entry := range IndigenciesData {
			if entry.DateCreated == currentDate.Format("2006-01-02") {
				count++
			}
		}

		// Append the formatted date to Labels and the count to countData
		Labels = append(Labels, formattedDate)
		countData = append(countData, count)

		// Move to the next day
		currentDate = currentDate.AddDate(0, 0, 1)
	}

	return Labels, countData, nil
}

func (b *BimsConfiguration) MonthlyClearance() ([]string, []int, error) {
	ClearanceData, err := database.GetTotalClearancePerMonth(b.BIMSdb)
	if err != nil {
		return nil, nil, err
	}

	// Initialize arrays for labels and count data
	var Labels []string
	var countData []int

	// Calculate the date range from "2023-09-01" to today
	startDate, _ := time.Parse("2006-01-02", "2023-09-01")
	endDate := time.Now()

	// Loop through the date range
	currentDate := startDate
	for currentDate.Before(endDate) || currentDate.Equal(endDate) {
		// Format the date to "Sep 1" style
		formattedDate := currentDate.Format("Jan 2")

		// Count the entries for the current date
		count := 0
		for _, entry := range ClearanceData {
			if entry.DateCreated == currentDate.Format("2006-01-02") {
				count++
			}
		}

		// Append the formatted date to Labels and the count to countData
		Labels = append(Labels, formattedDate)
		countData = append(countData, count)

		// Move to the next day
		currentDate = currentDate.AddDate(0, 0, 1)
	}

	return Labels, countData, nil
}

func (b *BimsConfiguration) MonthlyReferrals() ([]string, []int, error) {
	ReferralData, err := database.GetTotalReferralsPerMonth(b.BIMSdb)
	if err != nil {
		return nil, nil, err
	}

	// Initialize arrays for labels and count data
	var Labels []string
	var countData []int

	// Calculate the date range from "2023-09-01" to today
	startDate, _ := time.Parse("2006-01-02", "2023-09-01")
	endDate := time.Now()

	// Loop through the date range
	currentDate := startDate
	for currentDate.Before(endDate) || currentDate.Equal(endDate) {
		// Format the date to "Sep 1" style
		formattedDate := currentDate.Format("Jan 2")

		// Count the entries for the current date
		count := 0
		for _, entry := range ReferralData {
			if entry.DateCreated == currentDate.Format("2006-01-02") {
				count++
			}
		}

		// Append the formatted date to Labels and the count to countData
		Labels = append(Labels, formattedDate)
		countData = append(countData, count)

		// Move to the next day
		currentDate = currentDate.AddDate(0, 0, 1)
	}

	return Labels, countData, nil
}
