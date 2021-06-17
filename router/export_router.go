package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"pdb/database"
	"pdb/model"
	"strings"
)

func GetDocumentationPDF(c *gin.Context) {
	patient := c.MustGet("patient").(model.Patient)

	c.Header("Content-Type", "application/pdf")
	pdf := GenerateDocPDF(&patient)

	if err := pdf.Output(c.Writer); err != nil {
		fmt.Println(err.Error())
	}
}

func CreatePDF(title string) *gofpdf.Fpdf {
	pdf := gofpdf.New("P", "mm", "A4", "")
	tr := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.SetFooterFunc(func() {
		// Position at 1.5 cm from bottom
		pdf.SetY(-15)
		// Arial italic 8
		pdf.SetFont("Arial", "I", 8)
		// Text color in gray
		pdf.SetTextColor(128, 128, 128)
		// Page number
		pdf.CellFormat(0, 10, fmt.Sprintf("Seite %d", pdf.PageNo()),
			"", 0, "C", false, 0, "")
	})

	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, tr(title))
	pdf.SetTitle(tr(title), false)

	return pdf
}

func GenerateDocPDF(p *model.Patient) *gofpdf.Fpdf {
	title := fmt.Sprintf("Dokumentation: %s %s", p.FirstName, p.LastName)
	pdf := CreatePDF(title)
	tr := pdf.UnicodeTranslatorFromDescriptor("")

	documentations := model.FindDocumentation(p.ID, database.GormDB)

	_, lineHt := pdf.GetFontSize()
	pdf.Ln(15)
	pdf.Line(10, 20, 200, 20)
	html := pdf.HTMLBasicNew()
	pdf.SetFont("Arial", "", 12)
	for _, documentation := range documentations {
		content := strings.ReplaceAll(documentation.Content, "&nbsp;", " ")
		content = strings.ReplaceAll(content, "</p>", "</p><br />")
		content = strings.ReplaceAll(content, "<br>", "<br><br />")

		timeFormat := documentation.Time.Format("Am 02.01.2006 um 15:04")

		html.Write(lineHt, fmt.Sprintf("<b>%s</b>", timeFormat))
		pdf.Ln(2)
		html.Write(lineHt, fmt.Sprintf("<br /><p>%s</p><br /><br />", tr(content)))
		pdf.Ln(5)
	}

	return pdf
}
