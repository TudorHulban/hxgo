package request

import (
	goerrors "github.com/TudorHulban/go-errors"
)

type RequestData struct {
	Content map[string]string
}

type ParamsNewRequestData struct {
	RequestHeaders map[string][]string
	BodyRaw        []byte
}

func NewRequestData(params *ParamsNewRequestData) (*RequestData, error) {
	data, errParse := parseMultipartForm(
		params.BodyRaw,
		params.RequestHeaders,
	)
	if errParse != nil {
		return nil,
			goerrors.ErrInvalidInput{
				Issue:      errParse,
				Caller:     "NewRequestData",
				InputName:  "request data",
				InputValue: params,
			}
	}

	return &RequestData{
			Content: data,
		},
		nil
}
