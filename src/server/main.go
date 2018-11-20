package main

import (
	"fmt"
	"github.com/hhrutter/pdfcpu/pkg/pdfcpu"
	"github.com/jung-kurt/gofpdf"
	"log"
	"net/http"
	"strings"
	"github.com/hhrutter/pdfcpu/pkg/api"
)

func main()  {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/pdf", createPdf)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}


}

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello bro " + message
	w.Write([]byte(message))


}

func createPdf(w http.ResponseWriter, r *http.Request) {
	log.Println("here")
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world 5")
	err := pdf.OutputFileAndClose("hello.pdf")
	log.Println(err)
	w.Write([]byte("pdf"))
	exampleProcessMerge()
}


func exampleProcessMerge() {

	// Concatenate this sequence of PDF files:
	filenamesIn := []string{"hello.pdf", "hello.pdf", "hello.pdf"}

	_, err := api.Process(api.MergeCommand(filenamesIn, "merge.pdf", pdfcpu.NewDefaultConfiguration()))
	if err != nil {
		return
	}

}

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

	_, err := api.Process(api.OptimizeCommand("in.pdf", "out.pdf", config))
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
	selectedPages := []string{"3..5"}

	config := pdfcpu.NewDefaultConfiguration()

	// Set optional password(s).
	//config.UserPW = "upw"
	//config.OwnerPW = "opw"

	_, err := api.Process(api.ExtractPagesCommand("in.pdf", "dirOut", selectedPages, config))
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
//package main
//import (
//	"net/http"
//	"strings"
//)
//func sayHello(w http.ResponseWriter, r *http.Request) {
//	message := r.URL.Path
//	message = strings.TrimPrefix(message, "/")
//	message = "Hello " + message
//	w.Write([]byte(message))
//}
//func main() {
//	http.HandleFunc("/", sayHello)
//	if err := http.ListenAndServe(":8080", nil); err != nil {
//		panic(err)
//	}
//}