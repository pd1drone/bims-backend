package rest

import (
	"bims/database"
	"net/http"
)

type PieChartData struct {
	Labels   []string           `json:"labels"`
	Datasets []*PiechartDataSet `json:"datasets"`
}
type PiechartDataSet struct {
	Label           string   `json:"label"`
	Data            []int64  `json:"data"`
	BackgroundColor []string `json:"backgroundColor"`
}

type PieChartCounter struct {
	Printting int64 `json:"Printting"`
	Printted  int64 `json:"Printted"`
	Claimed   int64 `json:"Claimed"`
}

func (b *BimsConfiguration) GetPieChartData(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	ClearancePrinted, err := database.GetTotalClearancePerMonthPrintted(b.BIMSdb)
	if err != nil {
		respondJSON(w, 400, nil)
	}

	ClearancePrinting, err := database.GetTotalClearancePerMonthPrinting(b.BIMSdb)
	if err != nil {
		respondJSON(w, 400, nil)
	}

	ClearanceClaimed, err := database.GetTotalClearancePerMonthClaimed(b.BIMSdb)
	if err != nil {
		respondJSON(w, 400, nil)
	}

	IndigenciesPrinted, err := database.GetTotalIndigenciesPerMonthPrintted(b.BIMSdb)
	if err != nil {
		respondJSON(w, 400, nil)
	}

	IndigenciesPrinting, err := database.GetTotalIndigenciesPerMonthPrinting(b.BIMSdb)
	if err != nil {
		respondJSON(w, 400, nil)
	}

	IndigenciesClaimed, err := database.GetTotalIndigenciesPerMonthClaimed(b.BIMSdb)
	if err != nil {
		respondJSON(w, 400, nil)
	}

	ReferralsPrinted, err := database.GetTotalReferralsPerMonthPrintted(b.BIMSdb)
	if err != nil {
		respondJSON(w, 400, nil)
	}
	ReferralsPrinting, err := database.GetTotalReferralsPerMonthPrinting(b.BIMSdb)
	if err != nil {
		respondJSON(w, 400, nil)
	}
	ReferralsClaimed, err := database.GetTotalReferralsPerMonthClaimed(b.BIMSdb)
	if err != nil {
		respondJSON(w, 400, nil)
	}

	printing := len(ClearancePrinted) + len(IndigenciesPrinted) + len(ReferralsPrinted)
	printted := len(ClearancePrinting) + len(IndigenciesPrinting) + len(ReferralsPrinting)
	claimed := len(ClearanceClaimed) + len(IndigenciesClaimed) + len(ReferralsClaimed)

	// labels: ['For print records', 'Printed Records', 'Claimed records'],
	// datasets: [
	//   {
	// 	label: 'Records',
	// 	data: [1,2,3],
	// 	backgroundColor: ['#EF4444','#F59E0B','#22C55E'],
	//   },
	// ]

	labels := make([]string, 0)
	labels = append(labels, "For print records")
	labels = append(labels, "Printed Records")
	labels = append(labels, "Claimed records")

	datasets := make([]*PiechartDataSet, 0)
	piechartData := make([]int64, 0)
	piechartData = append(piechartData, int64(printing))
	piechartData = append(piechartData, int64(printted))
	piechartData = append(piechartData, int64(claimed))

	bgColor := make([]string, 0)
	bgColor = append(bgColor, "#EF4444")
	bgColor = append(bgColor, "#F59E0B")
	bgColor = append(bgColor, "#22C55E")

	datasets = append(datasets, &PiechartDataSet{
		Label:           "Records",
		Data:            piechartData,
		BackgroundColor: bgColor,
	})

	resp := &PieChartData{
		Labels:   labels,
		Datasets: datasets,
	}

	respondJSON(w, 200, resp)
}

func (b *BimsConfiguration) GetPieChartDataCounter(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	ClearancePrinted, err := database.GetTotalClearancePerMonthPrintted(b.BIMSdb)
	if err != nil {
		respondJSON(w, 400, nil)
	}

	ClearancePrinting, err := database.GetTotalClearancePerMonthPrinting(b.BIMSdb)
	if err != nil {
		respondJSON(w, 400, nil)
	}

	ClearanceClaimed, err := database.GetTotalClearancePerMonthClaimed(b.BIMSdb)
	if err != nil {
		respondJSON(w, 400, nil)
	}

	IndigenciesPrinted, err := database.GetTotalIndigenciesPerMonthPrintted(b.BIMSdb)
	if err != nil {
		respondJSON(w, 400, nil)
	}

	IndigenciesPrinting, err := database.GetTotalIndigenciesPerMonthPrinting(b.BIMSdb)
	if err != nil {
		respondJSON(w, 400, nil)
	}

	IndigenciesClaimed, err := database.GetTotalIndigenciesPerMonthClaimed(b.BIMSdb)
	if err != nil {
		respondJSON(w, 400, nil)
	}

	ReferralsPrinted, err := database.GetTotalReferralsPerMonthPrintted(b.BIMSdb)
	if err != nil {
		respondJSON(w, 400, nil)
	}
	ReferralsPrinting, err := database.GetTotalReferralsPerMonthPrinting(b.BIMSdb)
	if err != nil {
		respondJSON(w, 400, nil)
	}
	ReferralsClaimed, err := database.GetTotalReferralsPerMonthClaimed(b.BIMSdb)
	if err != nil {
		respondJSON(w, 400, nil)
	}

	printing := len(ClearancePrinting) + len(IndigenciesPrinting) + len(ReferralsPrinting)
	printted := len(ClearancePrinted) + len(IndigenciesPrinted) + len(ReferralsPrinted)
	claimed := len(ClearanceClaimed) + len(IndigenciesClaimed) + len(ReferralsClaimed)

	resp := &PieChartCounter{
		Printting: int64(printing),
		Printted:  int64(printted),
		Claimed:   int64(claimed),
	}

	respondJSON(w, 200, resp)
}
