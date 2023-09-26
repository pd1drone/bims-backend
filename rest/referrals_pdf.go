package rest

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/signintech/gopdf"
)

func CreateReferralsPDF(residentID int64, documentID int64, LastName string, MiddleName string, FirstName string, Address string,
	MobileNumber string, ParentName string, ParentNumber string, Reason string, HCGGGnumber string, PhilHealthNumber string,
	PhilHealthCategory string, Gender string, BirthDate string, CivilStatus string, Religion string, Occupation string, BirthPlace string) error {

	// LastName := "Lim"
	// MiddleName := "Navarro"
	// FirstName := "Karl Angelo"
	// Address := "16-A Avelino Mejico Sr. Santolan Pasig City"
	// MobileNumber := "09164385846"
	// ParentName := "Myhr Lim"
	// ParentNumber := "09164418014"
	// Reason := "Lorem Ipsum is simply dummy text of the printing and typesetting"
	// HCGGGnumber := "123123123123"
	// PhilHealthNumber := "1456745645667"
	// PhilHealthCategory := "Dependent"
	// Gender := "Male"
	// BirthDate := "09/10/1997"
	// CivilStatus := "Married"
	// Religion := "Roman Catholic"
	// Occupation := "Employee"
	// BirthPlace := "Paranaque City"

	currentTime := time.Now()
	formattedTime := currentTime.Format("01/02/2006")

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4

	err := pdf.AddTTFFont("times", "/root/bims-backend/pdf/fonts/times.ttf")
	if err != nil {
		return err
	}

	err = pdf.SetFont("times", "", 20)
	if err != nil {
		return err
	}

	pdf.AddPage()

	// Import page 1
	tpl1 := pdf.ImportPage("/root/bims-backend/pdf/reference_pdf/new_referral.pdf", 1, "/MediaBox")

	// Draw pdf onto page
	pdf.UseImportedTemplate(tpl1, 0, 0, 0, 0)

	pdf.SetFontSize(12)
	pdf.SetXY(125, 177)
	pdf.Text(LastName)
	pdf.SetXY(188, 177)
	pdf.Text(MiddleName)
	pdf.SetXY(240, 177)
	pdf.Text(FirstName)

	words := strings.Fields(Address)
	yData := 189.0
	for i := 0; i < len(words); i += 4 {
		end := i + 4
		if end > len(words) {
			end = len(words)
		}
		chunk := strings.Join(words[i:end], " ")
		pdf.SetXY(120, yData)
		pdf.Text(chunk)
		yData = yData + 12
	}

	pdf.SetXY(120, 225)
	pdf.Text(MobileNumber)

	pdf.SetXY(150, 237)
	pdf.Text(ParentName)

	pdf.SetXY(150, 249)
	pdf.Text(ParentNumber)

	pdf.SetXY(130, 424)
	pdf.Text(Reason)

	pdf.SetXY(400, 165)
	pdf.Text(formattedTime)
	pdf.SetXY(400, 177)
	pdf.Text(HCGGGnumber)
	pdf.SetXY(400, 189)
	pdf.Text(PhilHealthNumber)
	pdf.SetXY(400, 201)
	pdf.Text(PhilHealthCategory)
	pdf.SetXY(400, 213)
	pdf.Text(Gender)
	pdf.SetXY(400, 225)
	pdf.Text(BirthDate)
	pdf.SetXY(400, 237)
	pdf.Text(BirthPlace)
	pdf.SetXY(400, 249)
	pdf.Text(CivilStatus)
	pdf.SetXY(400, 261)
	pdf.Text(Religion)
	pdf.SetXY(400, 273)
	pdf.Text(Occupation)

	age, err := GetAge(BirthDate)

	pdf.SetXY(520, 225)
	pdf.Text(age)

	err = createDirectoryIfNotExist("/root/bims-backend/files/")
	if err != nil {
		return err
	}
	err = createDirectoryIfNotExist("/root/bims-backend/files/referrals/")
	if err != nil {
		return err
	}
	err = createDirectoryIfNotExist("/root/bims-backend/files/referrals/" + strconv.Itoa(int(residentID)))
	if err != nil {
		return err
	}
	err = createDirectoryIfNotExist("/root/bims-backend/files/referrals/" + strconv.Itoa(int(residentID)) + "/" + strconv.Itoa(int(documentID)))
	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("/root/bims-backend/files/referrals/%s/%s/referrals_%s_%s.pdf",
		strconv.Itoa(int(residentID)), strconv.Itoa(int(documentID)), strconv.Itoa(int(residentID)), strconv.Itoa(int(documentID)))

	err = pdf.WritePdf(filePath)
	if err != nil {
		return err
	}

	return err

}

func GetAge(birthDateStr string) (string, error) {
	// Parse the birthdate string into a time.Time object
	birthDate, err := time.Parse("01/02/2006", birthDateStr)
	if err != nil {
		return "", err
	}

	// Get the current date
	currentDate := time.Now()

	// Calculate the user's age
	age := currentDate.Year() - birthDate.Year()

	// Check if the user's birthday has occurred this year
	if currentDate.YearDay() < birthDate.YearDay() {
		age--
	}

	return strconv.Itoa(age), nil
}
