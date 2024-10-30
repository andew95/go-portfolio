package uploadService

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
)

func SaveFileV1(file multipart.File, fileName string) error {
	tempFile, err := os.Create("public/" + fileName)
	if err != nil {
		return err
	}
	defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	tempFile.Write(fileBytes)
	return nil
}

func SaveFileV2(file multipart.File, fileName string) error {
	// Create a buffer to store the request body
	var buf bytes.Buffer

	// Create a new multipart writer with the buffer
	w := multipart.NewWriter(&buf)
	fw, err := w.CreateFormFile("/public", fileName)
	if err != nil {
		return err
	}

	if _, err := io.Copy(fw, file); err != nil {
		return err
	}

	// Close the multipart writer to finalize the request
	w.Close()
	return nil
}
