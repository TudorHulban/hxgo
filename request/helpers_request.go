package request

import (
	"bytes"
	"errors"
	"io"
	"mime"
	"mime/multipart"
	"strings"

	goerrors "github.com/TudorHulban/go-errors"
)

func getRequestBoundary(requestHeaders map[string][]string) (string, error) {
	contentType, exists := requestHeaders["Content-Type"]
	if !exists {
		return "",
			goerrors.ErrNilInput{
				InputName: `requestHeaders["Content-Type"]`,
			}
	}

	mediaType, params, err := mime.ParseMediaType( // TODO: not cheap. overkill? alternatives: manual boundary scan or strings.Index("boundary=")
		strings.Join( // TODO: allocates. Content-Type has exactly one value so joining might not be needed.
			contentType,
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
		errors.New("no boundary found in Content-Type")
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

		formField := part.FormName()

		var buf bytes.Buffer // TODO: pre-size buffer using Content-Length (if present), sync.Pool, read directly into a []byte slice

		if _, errRead := buf.ReadFrom(part); errRead != nil {
			_ = part.Close()

			return nil,
				errRead
		}

		_ = part.Close()

		formValue := buf.String()

		if len(formValue) == 0 {
			continue
		}

		result[formField] = formValue
	}

	return result,
		nil
}
