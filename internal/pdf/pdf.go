package pdf

import (
	"fmt"
	"os"
	"strings"

	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
)

type PDFReader struct{}

func (p *PDFReader) ExtractPages(pdfPath string) ([]string, error) {
	uniPdfLicenseKey := os.Getenv("API_KEY")
	err := license.SetMeteredKey(uniPdfLicenseKey)
	if err != nil {
		return nil, fmt.Errorf("error setting license key: %v", err)
	}

	file, err := os.Open(pdfPath)
	if err != nil {
		return nil, fmt.Errorf("error opening PDF: %v", err)
	}
	defer file.Close()

	pdfReader, err := model.NewPdfReader(file)
	if err != nil {
		return nil, fmt.Errorf("error creating PDF reader: %v", err)
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return nil, fmt.Errorf("error getting number of pages: %v", err)
	}

	var pagesText []string
	for i := 1; i <= numPages; i++ {
		page, err := pdfReader.GetPage(i)
		if err != nil {
			return nil, fmt.Errorf("error getting page %d: %v", i, err)
		}

		ex, err := extractor.New(page)
		if err != nil {
			return nil, fmt.Errorf("error creating text extractor for page %d: %v", i, err)
		}

		text, err := ex.ExtractText()
		if err != nil {
			return nil, fmt.Errorf("error extracting text from page %d: %v", i, err)
		}

		pagesText = append(pagesText, text)
	}

	return pagesText, nil
}

func (p *PDFReader) ArrayForText(array []string) string {
	return strings.Join(array, " ")
}
