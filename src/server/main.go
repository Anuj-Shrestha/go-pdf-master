package main

import (
	"fmt"
	"github.com/hhrutter/pdfcpu/pkg/api"
	"github.com/hhrutter/pdfcpu/pkg/pdfcpu"
	"github.com/jung-kurt/gofpdf"
	"log"
	"net/http"
	"strings"
)

func main()  {
	//http.HandleFunc("/", sayHello)
	//http.HandleFunc("/pdf", createPdf)
	//if err := http.ListenAndServe(":8081", nil); err != nil {
	//	panic(err)
	//}

	//creates simple pdf. includes lines, image, html string
	hello()

	//expample for merging two pdfs using pdfcpu
	exampleProcessMerge()

	//exampleProcessExtractPages()

}

func hello() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Nurse Profile")

	pdf.MoveTo(20, 25)
	pdf.LineTo(170, 25)
	pdf.ClosePath()
	pdf.SetFillColor(200, 200, 200)
	pdf.SetLineWidth(3)
	pdf.DrawPath("DF")

	imageFile(pdf)

	pdf.MoveTo(45, 30)
	pdf.LineTo(45, 65)
	//pdf.ArcTo(170, 40, 20, 20, 0, 90, 0)
	//pdf.CurveTo(190, 100, 105, 100)
	//pdf.CurveBezierCubicTo(20, 100, 105, 200, 20, 200)
	pdf.ClosePath()
	pdf.SetFillColor(200, 200, 200)
	pdf.SetLineWidth(1)
	pdf.DrawPath("DF")

	pdf.SetFont("Times", "", 14)

	pdf.SetLeftMargin(50)
	pdf.SetFontSize(14)
	_, lineHt := pdf.GetFontSize()
	pdf.SetY(30)
	pdf.Cell(40, 5, "Nurse Hellen")
	pdf.SetY(40)

	htmlStr := `You can now easily print text mixing different styles: <b>bold</b>, ` +
		`<i>italic</i>, <u>underlined</u>, or <b><i><u>all at once</u></i></b>!<br><br>` +
		`<center>You can also center text.</center>` +
		`<right>Or align it to the right.</right>` +
		`You can also insert links on text, such as ` +
		`<a href="http://www.fpdf.org">www.fpdf.org</a>, or on an image: click on the logo.`
	html := pdf.HTMLBasicNew()
	html.Write(lineHt, htmlStr)

	//pdf.CellFormat(40, 40, "Name: Nurse Hellen", "*", 0, "R", false, 0, "")

	//pdf.SetY(80)
	pdf.AddPage()

	pdf.Text(10, 10, "Driving Licence:")
	var opt gofpdf.ImageOptions

	opt.ImageType = "jpeg"
	//pdf.ImageOptions("download.jpeg", -10, 10, 30, 0, false, opt, 0, "")
	opt.AllowNegativePosition = true
	pdf.ImageOptions("driving.JPG", 10, 20, 80, 0, false, opt, 0, "")


	err := pdf.OutputFileAndClose("profile.pdf")
	log.Println(err)

}

// generate image
func imageFile(pdf *gofpdf.Fpdf) {
	var opt gofpdf.ImageOptions

	//pdf := gofpdf.New("P", "mm", "A4", "")
	//pdf.AddPage()
	pdf.SetFont("Arial", "", 11)
	pdf.SetX(60)
	opt.ImageType = "jpeg"
	//pdf.ImageOptions("download.jpeg", -10, 10, 30, 0, false, opt, 0, "")
	opt.AllowNegativePosition = true
	pdf.ImageOptions("download.jpeg", 10, 30, 30, 0, false, opt, 0, "")
	//err := pdf.OutputFileAndClose("image.pdf")
}

// basing pdf
func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello bro " + message
	w.Write([]byte(message))
}

// footer expample
func TestFooterFuncLpi() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	var (
		oldFooterFnc  = "oldFooterFnc"
		bothPages     = "bothPages"
		firstPageOnly = "firstPageOnly"
		lastPageOnly  = "lastPageOnly"
	)

	// This set just for testing, only set SetFooterFuncLpi.
	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFont("Arial", "I", 8)
		pdf.CellFormat(0, 10, oldFooterFnc,
			"", 0, "C", false, 0, "")
	})
	pdf.SetFooterFuncLpi(func(lastPage bool) {
		pdf.SetY(-15)
		pdf.SetFont("Arial", "I", 8)
		pdf.CellFormat(0, 10, bothPages, "", 0, "L", false, 0, "")
		if !lastPage {
			pdf.CellFormat(0, 10, firstPageOnly, "", 0, "C", false, 0, "")
		} else {
			pdf.CellFormat(0, 10, lastPageOnly, "", 0, "C", false, 0, "")
		}
	})
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	for j := 1; j <= 40; j++ {
		pdf.CellFormat(0, 10, fmt.Sprintf("Printing line number %d", j),
			"", 1, "", false, 0, "")
	}

	pdf.AddPage()
	pdf.MoveTo(20, 20)
	pdf.LineTo(170, 20)
	pdf.ArcTo(170, 40, 20, 20, 0, 90, 0)
	pdf.CurveTo(190, 100, 105, 100)
	pdf.CurveBezierCubicTo(20, 100, 105, 200, 20, 200)
	pdf.ClosePath()
	pdf.SetFillColor(200, 200, 200)
	pdf.SetLineWidth(3)
	pdf.DrawPath("DF")

	err := pdf.OutputFileAndClose("footer.pdf")
	log.Println(err)
}

// merge example
func exampleProcessMerge() {
	log.Println("here at merge")
	// Concatenate this sequence of PDF files:
	filenamesIn := []string{"profile.pdf", "footer.pdf", "hello.pdf"}

	_, err := api.Process(api.MergeCommand(filenamesIn, "mergeProfile.pdf", pdfcpu.NewDefaultConfiguration()))
	if err != nil {
		return
	}

}



// ***** Following are Examples of pdfcpu functions *****
func exampleProcessValidate() {

	config := pdfcpu.NewDefaultConfiguration()

	// Set optional password(s).
	//config.UserPW = "upw"
	//config.OwnerPW = "opw"

	// Set relaxed validation mode.
	config.ValidationMode = pdfcpu.ValidationRelaxed

	_, err := api.Process(api.ValidateCommand("in.pdf", config))
	if err != nil {
		return
	}

}

func exampleProcessOptimize() {

	config := pdfcpu.NewDefaultConfiguration()

	// Set optional password(s).
	//config.UserPW = "upw"
	//config.OwnerPW = "opw"

	// Generate optional stats.
	config.StatsFileName = "stats.csv"

	// Configure end of line sequence for writing.
	config.Eol = pdfcpu.EolLF

	_, err := api.Process(api.OptimizeCommand("merge.pdf", "optimize.pdf", config))
	if err != nil {
		return
	}

}

func exampleProcessSplit() {

	config := pdfcpu.NewDefaultConfiguration()

	// Set optional password(s).
	//config.UserPW = "upw"
	//config.OwnerPW = "opw"

	// Split into single-page PDFs.

	_, err := api.Process(api.SplitCommand("in.pdf", "outDir", config))
	if err != nil {
		return
	}

}

func exampleProcessTrim() {

	// Trim to first three pages.
	selectedPages := []string{"-3"}

	config := pdfcpu.NewDefaultConfiguration()

	// Set optional password(s).
	//config.UserPW = "upw"
	//config.OwnerPW = "opw"

	_, err := api.Process(api.TrimCommand("in.pdf", "out.pdf", selectedPages, config))
	if err != nil {
		return
	}

}

func exampleProcessExtractPages() {

	// Extract single-page PDFs for pages 3, 4 and 5.
	selectedPages := []string{"2..4"}

	config := pdfcpu.NewDefaultConfiguration()

	// Set optional password(s).
	//config.UserPW = "upw"
	//config.OwnerPW = "opw"
	log.Println("here at dirOut")
	_, err := api.Process(api.ExtractPagesCommand("mergeProfile.pdf", "output/", selectedPages, config))
	if err != nil {
		return
	}

}

func exampleProcessExtractImages() {

	// Extract all embedded images for first 5 and last 5 pages but not for page 4.
	selectedPages := []string{"-5", "5-", "!4"}

	config := pdfcpu.NewDefaultConfiguration()

	// Set optional password(s).
	//config.UserPW = "upw"
	//config.OwnerPW = "opw"

	_, err := api.Process(api.ExtractImagesCommand("in.pdf", "dirOut", selectedPages, config))
	if err != nil {
		return
	}

}

func exampleProcessListAttachments() {

	config := pdfcpu.NewDefaultConfiguration()

	// Set optional password(s).
	//config.UserPW = "upw"
	//config.OwnerPW = opw"

	list, err := api.Process(api.ListAttachmentsCommand("in.pdf", config))
	if err != nil {
		return
	}

	// Print attachment list.
	for _, l := range list {
		fmt.Println(l)
	}

}

func exampleProcessAddAttachments() {

	config := pdfcpu.NewDefaultConfiguration()

	// Set optional password(s).
	//config.UserPW = "upw"
	//config.OwnerPW = "opw"

	_, err := api.Process(api.AddAttachmentsCommand("in.pdf", []string{"a.csv", "b.jpg", "c.pdf"}, config))
	if err != nil {
		return
	}
}

func exampleProcessRemoveAttachments() {

	config := pdfcpu.NewDefaultConfiguration()

	// Set optional password(s).
	//config.UserPW = "upw"
	//config.OwnerPW = "opw"

	// Not to be confused with the ExtractAttachmentsCommand!

	// Remove all attachments.
	_, err := api.Process(api.RemoveAttachmentsCommand("in.pdf", nil, config))
	if err != nil {
		return
	}

	// Remove specific attachments.
	_, err = api.Process(api.RemoveAttachmentsCommand("in.pdf", []string{"a.csv", "b.jpg"}, config))
	if err != nil {
		return
	}

}

func exampleProcessExtractAttachments() {

	config := pdfcpu.NewDefaultConfiguration()

	// Set optional password(s).
	//config.UserPW = "upw"
	//config.OwnerPW = "opw"

	// Extract all attachments.
	_, err := api.Process(api.ExtractAttachmentsCommand("in.pdf", "dirOut", nil, config))
	if err != nil {
		return
	}

	// Extract specific attachments.
	_, err = api.Process(api.ExtractAttachmentsCommand("in.pdf", "dirOut", []string{"a.csv", "b.pdf"}, config))
	if err != nil {
		return
	}
}

func exampleProcessEncrypt() {

	config := pdfcpu.NewDefaultConfiguration()

	config.UserPW = "upw"
	config.OwnerPW = "opw"

	_, err := api.Process(api.EncryptCommand("in.pdf", "out.pdf", config))
	if err != nil {
		return
	}
}

func exampleProcessDecrypt() {

	config := pdfcpu.NewDefaultConfiguration()

	config.UserPW = "upw"
	config.OwnerPW = "opw"

	_, err := api.Process(api.DecryptCommand("in.pdf", "out.pdf", config))
	if err != nil {
		return
	}
}

func exampleProcessChangeUserPW() {

	config := pdfcpu.NewDefaultConfiguration()

	// supply existing owner pw like so
	config.OwnerPW = "opw"

	pwOld := "pwOld"
	pwNew := "pwNew"

	_, err := api.Process(api.ChangeUserPWCommand("in.pdf", "out.pdf", config, &pwOld, &pwNew))
	if err != nil {
		return
	}
}

func exampleProcessChangeOwnerPW() {

	config := pdfcpu.NewDefaultConfiguration()

	// supply existing user pw like so
	config.UserPW = "upw"

	// old and new owner pw
	pwOld := "pwOld"
	pwNew := "pwNew"

	_, err := api.Process(api.ChangeOwnerPWCommand("in.pdf", "out.pdf", config, &pwOld, &pwNew))
	if err != nil {
		return
	}
}

func exampleProcesslistPermissions() {

	config := pdfcpu.NewDefaultConfiguration()
	config.UserPW = "upw"
	config.OwnerPW = "opw"

	list, err := api.Process(api.ListPermissionsCommand("in.pdf", config))
	if err != nil {
		return
	}

	// Print permissions list.
	for _, l := range list {
		fmt.Println(l)
	}
}

func exampleProcessAddPermissions() {

	config := pdfcpu.NewDefaultConfiguration()
	config.UserPW = "upw"
	config.OwnerPW = "opw"

	config.UserAccessPermissions = pdfcpu.PermissionsAll

	_, err := api.Process(api.AddPermissionsCommand("in.pdf", config))
	if err != nil {
		return
	}

}

func exampleProcessStamp() {

	// Stamp all but the first page.
	selectedPages := []string{"odd,!1"}
	var watermark *pdfcpu.Watermark

	config := pdfcpu.NewDefaultConfiguration()
	// Set optional password(s).
	//config.UserPW = "upw"
	//config.OwnerPW = "opw"

	_, err := api.Process(api.AddWatermarksCommand("in.pdf", "out.pdf", selectedPages, watermark, config))
	if err != nil {
		return
	}

}

func exampleProcessWatermark() {

	// Stamp all but the first page.
	selectedPages := []string{"even"}
	var watermark *pdfcpu.Watermark

	config := pdfcpu.NewDefaultConfiguration()
	// Set optional password(s).
	//config.UserPW = "upw"
	//config.OwnerPW = "opw"

	_, err := api.Process(api.AddWatermarksCommand("in.pdf", "out.pdf", selectedPages, watermark, config))
	if err != nil {
		return
	}

}