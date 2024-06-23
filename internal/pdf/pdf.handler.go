package pdf

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// UploadHandler godoc
// @Summary Upload a PDF file and extract text
// @Description Uploads a PDF file, extracts its text, and saves the extracted text
// @Tags PDF
// @Accept multipart/form-data
// @Produce json
// @Param pdf formData file true "PDF file"
// @Param Authorization header string true "Auth Token" default(Bearer <token>)
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/upload [post]
func UploadHandler(c *gin.Context) {
	file, err := c.FormFile("pdf")
	if err != nil {
		c.Set("Response", "Invalid file")
		c.Status(http.StatusBadRequest)
		return
	}

	if err := os.MkdirAll("pdfs", os.ModePerm); err != nil {
		c.Set("Response", "Could not create directory")
		c.Status(http.StatusInternalServerError)
		return
	}

	filePath := filepath.Join("pdfs", file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.Set("Response", "Could not save file")
		c.Status(http.StatusInternalServerError)
		return
	}

	pdfReader := PDFReader{}
	pages, err := pdfReader.ExtractPages(filePath)
	if err != nil {
		c.Set("Response", fmt.Sprintf("Error extracting pages: %v", err))
		c.Status(http.StatusInternalServerError)
		return
	}

	text := pdfReader.ArrayForText(pages)

	textFilePath := filepath.Join("pdfs", file.Filename+".txt")
	if err := ioutil.WriteFile(textFilePath, []byte(text), 0644); err != nil {
		c.Set("Response", "Could not save extracted text")
		c.Status(http.StatusInternalServerError)
		return
	}

	extractedText, err := ioutil.ReadFile(textFilePath)
	if err != nil {
		c.Set("Response", "Could not read extracted text")
		c.Status(http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message":        "File uploaded and processed successfully",
		"text_file":      textFilePath,
		"extracted_text": string(extractedText),
	}

	c.Set("Response", response)
	c.Status(http.StatusOK)
}
