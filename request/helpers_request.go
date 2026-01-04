package request

import (
	"bytes"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"strings"

	goerrors "github.com/TudorHulban/go-errors"
)

func getRequestBoundary(requestHeaders map[string][]string) (string, error) {
	contenType, exists := requestHeaders["Content-Type"]
	if !exists {
		return "",
			goerrors.ErrNilInput{
				InputName: `requestHeaders["Content-Type"]`,
			}
	}

	mediaType, params, err := mime.ParseMediaType(
		strings.Join(
			contenType,
			"",
		),
	)
	if err != nil {
		return "", err
	}

	if strings.HasPrefix(mediaType, "multipart/") {
		if boundary, ok := params["boundary"]; ok {
			return boundary, nil
		}
	}

	return "",
		fmt.Errorf("no boundary found in Content-Type")
}

func parseMultipartForm(formData []byte, requestHeaders map[string][]string) (map[string]string, error) {
	boundary, errCr := getRequestBoundary(requestHeaders)
	if errCr != nil {
		return nil,
			goerrors.ErrValidation{
				Issue:  errCr,
				Caller: "ParseMultipartForm",
			}
	}

	multipartReader := multipart.NewReader(
		bytes.NewReader(formData),
		boundary,
	)

	result := make(map[string]string)

	for {
		part, err := multipartReader.NextPart()
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}
		defer part.Close()

		formField := part.FormName()

		var buf bytes.Buffer

		if _, errRead := buf.ReadFrom(part); errRead != nil {
			return nil,
				errRead
		}

		formValue := buf.String()

		if len(formValue) == 0 {
			continue
		}

		result[formField] = formValue
	}

	return result,
		nil
}
