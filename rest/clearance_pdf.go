package rest

import (
	"fmt"
	"strconv"

	"github.com/signintech/gopdf"
)

func CreateClearancePDF(residentID int64, documentID int64, formattedTime string, Birthday string, birthPlace string, fullName string,
	Address string, CivilStatus string, Purpose string) error {

	// currentTime := time.Now()
	// formattedTime := currentTime.Format("January 2, 2006")

	// dob := "12/01/1981"
	// parsedDate, err := time.Parse("01/02/2006", dob)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// Birthday := parsedDate.Format("January 2, 2006")

	// fullName := "ARLIN G. BAYLON"
	// Address := "NO.4 F. MANALO STREET Barangay Batis, City of San Juan"

	// birthPlace := "Surigao"
	// CivilStatus := "Married"
	// Purpose := "RESIDENCY/MANILA WATER APPLICATION"

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4Landscape}) //595.28, 841.89 = A4

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
		pdf.Image("/root/bims-backend/pdf/img/san_juan_seal_resize.jpg", 70, 20, nil)
		pdf.Image("/root/bims-backend/pdf/img/Batis_Seal.jpg", 690, 20, nil)

		pdf.SetFont("timesbold", "", 12)
		pdf.SetXY(365, 20)
		pdf.Cell(nil, "Republic of The Philippines")

		pdf.SetXY(380, 32)
		pdf.Cell(nil, "CITY OF SAN JUAN")

		pdf.SetXY(330, 44)
		pdf.SetFont("times", "", 12)
		pdf.Cell(nil, "420 F. Manalo Street, Barangay Batis Hall")

		pdf.SetXY(380, 56)
		pdf.SetFont("timesbold", "", 12)
		pdf.Cell(nil, "BARANGAY BATIS")

		pdf.SetXY(360, 68)
		pdf.SetFont("times", "", 12)
		pdf.Cell(nil, "Telephone Number. 7744-0737")
	})

	pdf.AddPage()
	pdf.Image("/root/bims-backend/pdf/img/batis_blur.png", 330, 200, nil)

	pdf.SetFont("timesbold", "", 20)
	pdf.SetXY(220, 100)
	pdf.Text("OFFICE OF THE BARANGAY CHAIRPERSON")

	pdf.SetXY(310, 125)
	pdf.Text("BARANGAY CLEARANCE")

	pdf.SetFont("timesbold", "", 10)
	pdf.SetXY(40, 150)
	pdf.Text("DINO C. GENESTON")
	pdf.SetXY(40, 160)
	pdf.SetFont("times", "", 10)
	pdf.Text("Barangay Chairperson")

	pdf.SetFont("timesbold", "", 10)
	pdf.SetXY(40, 190)
	pdf.Text("KAGAWAD:")

	pdf.SetXY(40, 200)
	pdf.Text("NOEL B. ASTRERO")

	pdf.SetXY(40, 230)
	pdf.Text("ROSALIE A. DE OCA")

	pdf.SetXY(40, 260)
	pdf.Text("MELINDA M. LAGUMEN")

	pdf.SetXY(40, 290)
	pdf.Text("ARIEL C. TORRES")

	pdf.SetXY(40, 320)
	pdf.Text("CAMILLE MIKAELA M. MALAPITAN")

	pdf.SetXY(40, 350)
	pdf.Text("NICHOLAS G. DUMPIT II")

	pdf.SetXY(40, 380)
	pdf.Text("FRANCIS B. YABAO")

	pdf.SetXY(40, 410)
	pdf.Text("JUAN PAULO R. MARTIN")
	pdf.SetFont("timesitalic", "", 10)
	pdf.SetXY(40, 420)
	pdf.Text("SK Chairperson")

	pdf.SetXY(40, 450)
	pdf.SetFont("timesbold", "", 10)
	pdf.Text("JERICA P. PAMPLONA")
	pdf.SetFont("times", "", 10)
	pdf.SetXY(40, 460)
	pdf.Text("Barangay Treasurer")

	pdf.SetXY(40, 490)
	pdf.SetFont("timesbold", "", 10)
	pdf.Text("RAMIL B. ASTERO")
	pdf.SetFont("times", "", 10)
	pdf.SetXY(40, 500)
	pdf.Text("Barangay Secretary")

	pdf.SetXY(40, 530)
	pdf.SetFont("times", "", 10)
	pdf.Text("Note: This certification is not valid")
	pdf.SetXY(40, 540)
	pdf.Text("without the Barangay Batis Offical")
	pdf.SetXY(40, 550)
	pdf.Text("Seal")

	pdf.SetStrokeColorCMYK(0, 24, 53, 2)
	pdf.SetLineWidth(2)
	pdf.SetFillColorCMYK(0, 24, 53, 2)
	err = pdf.Rectangle(25, 130, 230, 570, "D", 3, 5)
	if err != nil {
		fmt.Println(err)
	}

	pdf.SetStrokeColorCMYK(0, 24, 53, 2)
	pdf.SetLineWidth(2)
	pdf.SetFillColorCMYK(0, 24, 53, 2)
	err = pdf.Rectangle(240, 130, 780, 570, "D", 3, 5)
	if err != nil {
		fmt.Println(err)
	}

	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("timesbold", "", 12)
	pdf.SetXY(260, 150)
	pdf.Text("TO WHOM IT MAY CONCERN:")

	pdf.SetFont("times", "", 12)
	pdf.SetXY(640, 150)
	pdf.Text("Date: " + formattedTime)

	pdf.SetXY(250, 180)
	pdf.Text("THIS IS TO CERTIFY that according to the records available in this office, that name appeared below is a")
	pdf.SetXY(250, 192)
	pdf.Text("resident of Barangay Batis, City of San Juan, Metro Manila, whoes picture, thumb marks and signature")
	pdf.SetXY(250, 204)
	pdf.Text("were affix on this certificate")

	pdf.SetXY(250, 230)
	pdf.Text("Name:")
	pdf.SetFont("timesbold", "", 12)
	pdf.SetXY(430, 230)
	pdf.Text(":  " + fullName)

	pdf.SetFont("times", "", 12)
	pdf.SetXY(250, 260)
	pdf.Text("Address:")
	pdf.SetFont("timesbold", "", 12)
	pdf.SetXY(430, 260)
	pdf.Text(":  " + Address)

	pdf.SetFont("times", "", 12)
	pdf.SetXY(250, 272)
	pdf.Text("Date Of Birth:")
	pdf.SetFont("timesbold", "", 12)
	pdf.SetXY(430, 272)
	pdf.Text(":  " + Birthday)

	pdf.SetFont("times", "", 12)
	pdf.SetXY(250, 284)
	pdf.Text("Place of Birth:")
	pdf.SetFont("timesbold", "", 12)
	pdf.SetXY(430, 284)
	pdf.Text(":  " + birthPlace)

	pdf.SetFont("times", "", 12)
	pdf.SetXY(250, 296)
	pdf.Text("Civil Status:")
	pdf.SetFont("timesbold", "", 12)
	pdf.SetXY(430, 296)
	pdf.Text(":  " + CivilStatus)

	pdf.SetFont("times", "", 12)
	pdf.SetXY(250, 326)
	pdf.Text("Purpose:")
	pdf.SetFont("timesbold", "", 12)
	pdf.SetXY(430, 326)
	pdf.Text(":  " + Purpose)

	pdf.SetFont("times", "", 12)
	pdf.SetXY(250, 356)
	pdf.Text("CTC No.:")
	pdf.SetFont("timesbold", "", 12)
	pdf.SetXY(430, 356)
	pdf.Text(":  ")

	pdf.SetFont("times", "", 12)
	pdf.SetXY(250, 368)
	pdf.Text("Issued at:")
	pdf.SetFont("timesbold", "", 12)
	pdf.SetXY(430, 368)
	pdf.Text(":  ")

	pdf.SetFont("times", "", 12)
	pdf.SetXY(250, 380)
	pdf.Text("Date Issued: ")
	pdf.SetFont("timesbold", "", 12)
	pdf.SetXY(430, 380)
	pdf.Text(":  ")

	pdf.SetFont("times", "", 12)
	pdf.SetXY(250, 392)
	pdf.Text("Precint No.:")
	pdf.SetFont("timesbold", "", 12)
	pdf.SetXY(430, 392)
	pdf.Text(":  ")

	pdf.SetFont("times", "", 12)
	pdf.SetXY(250, 404)
	pdf.Text("Expiry Date:")
	pdf.SetFont("timesbold", "", 12)
	pdf.SetXY(430, 404)
	pdf.Text(":  ")

	pdf.SetFont("times", "", 12)
	pdf.SetXY(250, 416)
	pdf.Text("R.B.I No.:")
	pdf.SetFont("timesbold", "", 12)
	pdf.SetXY(430, 416)
	pdf.Text(":  ")

	pdf.SetFont("times", "", 12)
	pdf.SetXY(250, 450)
	pdf.Text("Verified & Issued by:")
	pdf.SetXY(450, 450)
	pdf.Text("Noted by:")

	pdf.SetFont("times", "", 12)
	pdf.SetXY(630, 380)
	pdf.Text("_____________________")

	pdf.SetFont("times", "", 12)
	pdf.SetXY(640, 392)
	pdf.Text("Signature of Applicant")

	pdf.SetStrokeColorCMYK(0, 24, 53, 2)
	pdf.SetLineWidth(2)
	pdf.SetFillColorCMYK(0, 24, 53, 2)
	pdf.Oval(620, 430, 760, 530)

	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("times", "", 12)
	pdf.SetXY(640, 550)
	pdf.Text("Right Thumb Mark")

	pdf.SetFont("timesbold", "", 14)
	pdf.SetXY(250, 538)
	pdf.Text("NOEL B. ASTRERO")
	pdf.SetFont("times", "", 12)
	pdf.SetXY(250, 550)
	pdf.Text("Barangay Kagawad")

	pdf.SetFont("timesbold", "", 16)
	pdf.SetXY(450, 538)
	pdf.Text("DINO C. GENESTON")
	pdf.SetFont("times", "", 12)
	pdf.SetXY(450, 550)
	pdf.Text("Barangay Chairperson")

	err = createDirectoryIfNotExist("/root/bims-backend/files/")
	if err != nil {
		return err
	}
	err = createDirectoryIfNotExist("/root/bims-backend/files/clearances/")
	if err != nil {
		return err
	}
	err = createDirectoryIfNotExist("/root/bims-backend/files/clearances/" + strconv.Itoa(int(residentID)))
	if err != nil {
		return err
	}
	err = createDirectoryIfNotExist("/root/bims-backend/files/clearances/" + strconv.Itoa(int(residentID)) + "/" + strconv.Itoa(int(documentID)))
	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("/root/bims-backend/files/clearances/%s/%s/clearances_%s_%s.pdf",
		strconv.Itoa(int(residentID)), strconv.Itoa(int(documentID)), strconv.Itoa(int(residentID)), strconv.Itoa(int(documentID)))

	err = pdf.WritePdf(filePath)
	if err != nil {
		return err
	}

	return nil
}
