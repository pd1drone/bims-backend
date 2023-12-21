package rest

import (
	"fmt"
	"strconv"
	"time"

	"github.com/signintech/gopdf"
)

func CreateIncidentsPDF(ID int64, CompliantFullName string, Respondent string, IncidentDateTime string, IncidentLocation string, IncidentNarration string, IssuingOfficer string) error {

	// ID := 1
	// CompliantFullName := "Shane Joshoua Melo"
	// Respondent := "Shane Joshoua Melo"
	// IncidentDateTime := "12/01/1981"
	// IncidentLocation := "Barangay Batis Barangay Hall"
	// IncidentNarration := "On January 15, 2023, at approximately 3:30 PM, I was shopping at XYZ Mall when an altercation occurred in the parking lot near the main entrance. The incident involved two individuals, later identified as [Person A and Person B], who were engaged in a heated argument over a parking space. The argument quickly escalated into a physical confrontation, creating an unsafe environment for shoppers and passersby."
	// IssuingOfficer := "admin"

	parsedDate, err := time.Parse("01/02/2006", IncidentDateTime)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	Dt := parsedDate.Format("January 2, 2006")

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4}) //595.28, 841.89 = A4

	err = pdf.AddTTFFont("times", "/root/bims-backend/pdf/fonts/times.ttf")
	if err != nil {
		return err
	}
	err = pdf.AddTTFFont("timesbold", "/root/bims-backend/pdf/fonts/times_bold.TTF")
	if err != nil {
		return err
	}
	err = pdf.AddTTFFont("timesitalic", "/root/bims-backend/pdf/fonts/times_italic.ttf")
	if err != nil {
		return err
	}

	err = pdf.SetFont("times", "", 15)
	if err != nil {
		return err
	}

	pdf.AddPage()

	// Import page 1
	tpl1 := pdf.ImportPage("/root/bims-backend/pdf/reference_pdf/incident-report.pdf", 1, "/MediaBox")

	// Draw pdf onto page
	pdf.UseImportedTemplate(tpl1, 0, 0, 0, 0)

	pdf.SetFont("timesbold", "", 15)
	pdf.SetXY(205, 145)
	pdf.Text(strconv.Itoa(int(ID)))

	pdf.SetXY(225, 185)
	pdf.Text(CompliantFullName)

	pdf.SetXY(225, 210)
	pdf.Text(Respondent)

	pdf.SetXY(225, 233)
	pdf.Text(Dt)

	pdf.SetXY(225, 258)
	pdf.Text(IncidentLocation)

	reasonRunes := []rune(IncidentNarration)
	purposeyData := 320.0
	var charBuffer string
	var lineNumber = 0
	for i := 0; i < len(reasonRunes); i++ {
		character := string(reasonRunes[i])
		charBuffer += character
		if len(charBuffer) >= 60 && character == " " {
			pdf.SetXY(80, purposeyData)
			pdf.Text(charBuffer)
			purposeyData += 15
			charBuffer = ""
			lineNumber++
		}

	}

	if len(charBuffer) > 0 {
		pdf.SetXY(80, purposeyData)
		pdf.Text(charBuffer)
	}

	pdf.SetXY(205, 710)
	pdf.Text(IssuingOfficer)

	err = createDirectoryIfNotExist("/root/bims-backend/files/")
	if err != nil {
		return err
	}
	err = createDirectoryIfNotExist("/root/bims-backend/files/incidents/")
	if err != nil {
		return err
	}

	err = createDirectoryIfNotExist("/root/bims-backend/files/incidents/" + strconv.Itoa(int(ID)))
	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("/root/bims-backend/files/incidents/%s/incidents_%s.pdf", strconv.Itoa(int(ID)), strconv.Itoa(int(ID)))

	err = pdf.WritePdf(filePath)
	if err != nil {
		return err
	}

	return nil
}
