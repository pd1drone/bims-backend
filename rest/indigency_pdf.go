package rest

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/signintech/gopdf"
)

func CreateIndigencyPDF(residentID int64, documentID int64, fullname string, address string, Purpose string) error {

	// Get the current time
	currentTime := time.Now()

	// Format the day with the appropriate suffix
	day := currentTime.Day()
	dayString := dayWithSuffix(day)

	// Format the time in the desired string format
	formattedTime := currentTime.Format("day of January 2006")

	DateToday := fmt.Sprintf("%s %s", dayString, formattedTime)

	// fullname := "JECELL BELTRAN"
	// address := "No. 2 F. MANALO ST. Barangay Batis, San Juan City"
	// Purpose := "MOTHER - MARCELA BELTRAN (Deceased)"

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4}) //595.28, 841.89 = A4

	err := pdf.AddTTFFont("times", "/root/bims-backend/pdf/fonts/times.ttf")
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

	err = pdf.SetFont("times", "", 12)
	if err != nil {
		return err
	}

	pdf.AddHeader(func() {
		pdf.SetY(20)
		pdf.Image("/root/bims-backend/pdf/img/san_juan_seal_resize.jpg", 20, 20, nil)
		pdf.Image("/root/bims-backend/pdf/img/Batis_Seal.jpg", 490, 20, nil)

		pdf.SetXY(235, 20)
		pdf.Cell(nil, "Republic of The Philippines")

		pdf.SetXY(250, 32)
		pdf.SetFont("timesbold", "", 12)
		pdf.Cell(nil, "CITY OF SAN JUAN")

		pdf.SetXY(200, 44)
		pdf.SetFont("times", "", 12)
		pdf.Cell(nil, "420 F. Manalo Street, Barangay Batis Hall")

		pdf.SetXY(250, 56)
		pdf.SetFont("timesbold", "", 12)
		pdf.Cell(nil, "BARANGAY BATIS")

		pdf.SetXY(230, 68)
		pdf.SetFont("times", "", 12)
		pdf.Cell(nil, "Telephone Number. 7744-0737")

		pdf.SetXY(210, 80)
		pdf.SetFont("timesitalic", "", 12)
		pdf.SetTextColor(0, 0, 205)
		pdf.Cell(nil, "Email add: batis.sanjuan@gmail.com")
	})

	pdf.AddPage()
	pdf.Image("/root/bims-backend/pdf/img/batis_blur.png", 180, 275, nil)

	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("timesbold", "", 20)
	pdf.SetXY(135, 150)
	pdf.Text("OFFICE OF THE PUNONG BARANGAY")

	pdf.SetFont("timesbold", "", 12)
	pdf.SetXY(20, 190)
	pdf.Text("SANGGINIANG BARANGAY")

	pdf.SetFontSize(20)
	pdf.SetXY(270, 195)
	pdf.SetFont("timesbold", "", 20)
	pdf.Text("BARANGAY RESIDENCY AND")

	pdf.SetXY(350, 215)
	pdf.SetFont("timesbold", "", 20)
	pdf.Text("INDIGENCY")

	pdf.SetFont("timesbold", "", 10)
	pdf.SetXY(20, 230)
	pdf.Text("DINO C. GENESTON")
	pdf.SetXY(20, 240)
	pdf.SetFont("times", "", 10)
	pdf.Text("Punong Barangay")
	pdf.SetXY(20, 250)
	pdf.SetFont("timesitalic", "", 10)
	pdf.Text("Commitee on Peace & Order/BADAC")

	pdf.SetXY(20, 270)
	pdf.SetFont("timesbold", "", 10)
	pdf.Text("Barangay Kagawad")

	pdf.SetXY(20, 290)
	pdf.Text("NOEL B. ASTRERO")
	pdf.SetFont("timesitalic", "", 10)
	pdf.SetXY(20, 300)
	pdf.Text("Committee on Bids and Awards")
	pdf.SetXY(20, 310)
	pdf.Text("Commitee on Appropriations, Ways and Means")
	pdf.SetXY(20, 320)
	pdf.Text("Commitee on Infrastructure and Public Works")

	pdf.SetXY(20, 340)
	pdf.SetFont("timesbold", "", 10)
	pdf.Text("ROSALIE A. DE OCA")
	pdf.SetFont("timesitalic", "", 10)
	pdf.SetXY(20, 350)
	pdf.Text("Committee on Health, Nutrition and Women")
	pdf.SetXY(20, 360)
	pdf.Text("Children's Welfare")
	pdf.SetXY(20, 370)
	pdf.Text("Commitee on Clean and Green Solid Waste")
	pdf.SetXY(20, 380)
	pdf.Text("Management")

	pdf.SetXY(20, 400)
	pdf.SetFont("timesbold", "", 10)
	pdf.Text("MELINDA M. LAGUMEN")
	pdf.SetFont("timesitalic", "", 10)
	pdf.SetXY(20, 410)
	pdf.Text("Committee on Senior Citizen's Affairs")

	pdf.SetXY(20, 430)
	pdf.SetFont("timesbold", "", 10)
	pdf.Text("ARIEL C. TORRES")
	pdf.SetFont("timesitalic", "", 10)
	pdf.SetXY(20, 440)
	pdf.Text("Committee on Barangay Disaster Risk Reduction")
	pdf.SetXY(20, 450)
	pdf.Text("Management, Red Cros 143 and")
	pdf.SetXY(20, 460)
	pdf.Text("Barangay Citizens Program")

	pdf.SetXY(20, 480)
	pdf.SetFont("timesbold", "", 10)
	pdf.Text("CAMILLE MIKAELA M. MALAPITAN")
	pdf.SetFont("timesitalic", "", 10)
	pdf.SetXY(20, 490)
	pdf.Text("Committee on Person With Disability")

	pdf.SetXY(20, 510)
	pdf.SetFont("timesbold", "", 10)
	pdf.Text("NICHOLAS G. DUMPIT II")
	pdf.SetFont("timesitalic", "", 10)
	pdf.SetXY(20, 520)
	pdf.Text("Committee on Education, Public Information &")
	pdf.SetXY(20, 530)
	pdf.Text("Cultural")
	pdf.SetXY(20, 540)
	pdf.Text("Committee on Housing and Land Utilization")
	pdf.SetXY(20, 550)
	pdf.Text("Committee on Traffic and Parking Management")

	pdf.SetXY(20, 570)
	pdf.SetFont("timesbold", "", 10)
	pdf.Text("FRANCIS B. YABAO")
	pdf.SetFont("timesitalic", "", 10)
	pdf.SetXY(20, 580)
	pdf.Text("Committee on Livelihood/Entrepreneurship")
	pdf.SetXY(20, 590)
	pdf.Text("Committee on Public Utility")

	pdf.SetXY(20, 610)
	pdf.SetFont("timesbold", "", 10)
	pdf.Text("JUAN PAULO R. MARTIN")
	pdf.SetFont("timesitalic", "", 10)
	pdf.SetXY(20, 620)
	pdf.Text("SK Chairperson")
	pdf.SetXY(20, 630)
	pdf.Text("Committee on Youth and Sports Development")

	pdf.SetXY(20, 650)
	pdf.SetFont("timesbold", "", 10)
	pdf.Text("RAMIL B. ASTERO")
	pdf.SetFont("times", "", 10)
	pdf.SetXY(20, 660)
	pdf.Text("Barangay Secretary")

	pdf.SetXY(20, 680)
	pdf.SetFont("timesbold", "", 10)
	pdf.Text("JERICA P. PAMPLONA")
	pdf.SetFont("times", "", 10)
	pdf.SetXY(20, 690)
	pdf.Text("Barangay Treasurer")

	pdf.SetXY(360, 270)
	pdf.SetFont("times", "", 12)
	pdf.Text("This is to certify that")

	fullNameArray := strings.Split(fullname, " ")
	if len(fullNameArray) == 2 {
		pdf.SetXY(310, 300)
	} else if len(fullNameArray) == 3 {
		pdf.SetXY(310, 300)
	} else if len(fullNameArray) == 4 {
		pdf.SetXY(310, 300)
	}
	pdf.SetFont("timesbold", "", 20)
	pdf.Cell(nil, fullname)

	pdf.SetFont("times", "", 12)
	pdf.SetXY(315, 340)
	pdf.Text("Is a resident of our Barangay and with")
	pdf.SetXY(365, 352)
	pdf.Text("postal address at")

	pdf.SetFont("timesbold", "", 12)
	pdf.SetXY(290, 382)
	pdf.Text(address)

	pdf.SetFont("times", "", 12)
	pdf.SetXY(300, 412)
	pdf.Text("This certification is issued upon the request of")
	pdf.SetXY(300, 424)
	pdf.Text("the above-mentioned person for the purpose of:")
	pdf.SetXY(290, 446)
	pdf.SetFont("timesbold", "", 12)
	pdf.Text(Purpose)

	pdf.SetFont("times", "", 12)
	pdf.SetXY(320, 476)
	pdf.Text("Issued this " + DateToday + "")
	pdf.SetXY(335, 488)
	pdf.Text("at Barangay Batis,San Juan City")

	pdf.SetFont("timesbold", "", 18)
	pdf.SetTextColor(0, 0, 205)
	pdf.SetXY(380, 650)
	pdf.Text("DINO C. GENESTON")
	pdf.SetFont("timesitalic", "", 14)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetXY(415, 670)
	pdf.Text("Punong Barangay")

	pdf.SetXY(195, 750)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("timesitalic", "", 12)
	pdf.Cell(nil, "Not Valid without Official Barangay Seal")

	err = createDirectoryIfNotExist("/root/bims-backend/files/")
	if err != nil {
		return err
	}
	err = createDirectoryIfNotExist("/root/bims-backend/files/indigencies/")
	if err != nil {
		return err
	}
	err = createDirectoryIfNotExist("/root/bims-backend/files/indigencies/" + strconv.Itoa(int(residentID)))
	if err != nil {
		return err
	}
	err = createDirectoryIfNotExist("/root/bims-backend/files/indigencies/" + strconv.Itoa(int(residentID)) + "/" + strconv.Itoa(int(documentID)))
	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("/root/bims-backend/files/indigencies/%s/%s/indigencies_%s_%s.pdf",
		strconv.Itoa(int(residentID)), strconv.Itoa(int(documentID)), strconv.Itoa(int(residentID)), strconv.Itoa(int(documentID)))

	err = pdf.WritePdf(filePath)
	if err != nil {
		return err
	}

	return nil
}

func dayWithSuffix(day int) string {
	switch day {
	case 1, 21, 31:
		return fmt.Sprintf("%dst", day)
	case 2, 22:
		return fmt.Sprintf("%dnd", day)
	case 3, 23:
		return fmt.Sprintf("%drd", day)
	default:
		return fmt.Sprintf("%dth", day)
	}
}
