package pdf

import (
	"fmt"
	"os"
	"strings"

	"github.com/dslipak/pdf"
)

type PDFReader struct{}

func (p *PDFReader) ExtractPages(pdfPath string) ([]string, error) {
	file, err := os.Open(pdfPath)
	if err != nil {
		return nil, fmt.Errorf("error opening PDF: %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("error getting file info: %v", err)
	}

	r, err := pdf.NewReader(file, fileInfo.Size())
	if err != nil {
		return nil, fmt.Errorf("error creating PDF reader: %v", err)
	}

	var pagesText []string
	numPages := r.NumPage()
	for i := 1; i <= numPages; i++ {
		page := r.Page(i)
		if page.V.IsNull() {
			continue
		}

		content, err := page.GetPlainText(nil)
		if err != nil {
			return nil, fmt.Errorf("error extracting text from page %d: %v", i, err)
		}

		pagesText = append(pagesText, content)
	}

	return pagesText, nil
}

func (p *PDFReader) ArrayForText(array []string) string {
	return strings.Join(array, " ")
}
