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

	pdf.SetXY(20, 270)
	pdf.SetFont("timesbold", "", 10)
	pdf.Text("BARANGAY KAGAWAD")

	pdf.SetXY(20, 310)
	pdf.Text("JONEL V. FABROA")

	pdf.SetXY(20, 350)
	pdf.SetFont("timesbold", "", 10)
	pdf.Text("ROSALIE A. DE OCA")

	pdf.SetXY(20, 390)
	pdf.SetFont("timesbold", "", 10)
	pdf.Text("MELINDA M. LAGUMEN")

	pdf.SetXY(20, 430)
	pdf.SetFont("timesbold", "", 10)
	pdf.Text("RICARDO M. TUBOG")

	pdf.SetXY(20, 470)
	pdf.SetFont("timesbold", "", 10)
	pdf.Text("JUAN PAULO R. MARTIN")

	pdf.SetXY(20, 510)
	pdf.SetFont("timesbold", "", 10)
	pdf.Text("ARIEL C. TORRES")

	pdf.SetXY(20, 550)
	pdf.SetFont("timesbold", "", 10)
	pdf.Text("CAMILLE MIKAELA M. MALAPITAN")

	pdf.SetXY(20, 590)
	pdf.SetFont("timesbold", "", 10)
	pdf.Text("SHEREENA B. TAN")
	pdf.SetFont("timesitalic", "", 10)
	pdf.SetXY(20, 600)
	pdf.Text("SK Chairperson")

	pdf.SetXY(20, 640)
	pdf.SetFont("timesbold", "", 10)
	pdf.Text("RAMIL B. ASTERO")
	pdf.SetFont("times", "", 10)
	pdf.SetXY(20, 650)
	pdf.Text("Barangay Secretary")

	pdf.SetXY(20, 690)
	pdf.SetFont("timesbold", "", 10)
	pdf.Text("JERICA P. PAMPLONA")
	pdf.SetFont("times", "", 10)
	pdf.SetXY(20, 700)
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
	pdf.SetXY(270, 382)
	pdf.Text(address)

	pdf.SetFont("times", "", 12)
	pdf.SetXY(300, 412)
	pdf.Text("This certification is issued upon the request of the above")
	pdf.SetXY(360, 424)
	pdf.Text("mentioned person for his/her ")
	pdf.SetXY(290, 446)
	pdf.SetFont("timesbold", "", 12)
	pdf.Text(Purpose)
	pdf.SetFont("times", "", 12)
	pdf.SetXY(360, 460)
	pdf.Text("for obtaining/securing:")

	pdf.SetFont("timesbold", "", 12)
	pdf.SetXY(330, 480)
	pdf.Text("______ Proof of Residency")
	pdf.SetXY(330, 494)
	pdf.Text("______ Medical Assistance")
	pdf.SetXY(330, 508)
	pdf.Text("______ Belongs to one of the Indigent Families")
	pdf.SetXY(390, 520)
	pdf.Text("of the Barangay")

	pdf.SetFont("times", "", 12)
	pdf.SetXY(320, 550)
	pdf.Text("Done this " + DateToday + "")
	pdf.SetXY(335, 562)
	pdf.Text("at Barangay Batis,San Juan City")

	pdf.SetFont("timesbold", "", 18)
	pdf.SetTextColor(0, 0, 205)
	pdf.SetXY(380, 630)
	pdf.Text("DINO C. GENESTON")
	pdf.SetFont("timesitalic", "", 14)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetXY(415, 650)
	pdf.Text("Punong Barangay")

	pdf.SetXY(250, 670)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("timesitalic", "", 12)
	pdf.Cell(nil, "Issued by:")

	pdf.SetXY(250, 720)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("timesbold", "", 12)
	pdf.Cell(nil, "JONEL V. FABROA")
	pdf.SetXY(250, 732)
	pdf.SetFont("timesitalic", "", 12)
	pdf.Cell(nil, "Barangay Kagawad")

	pdf.SetXY(195, 780)
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
