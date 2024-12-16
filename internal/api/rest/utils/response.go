package utils

import "encoding/json"

// ResponseSuccess with data
type ResponseSuccess struct {
	Data any `json:"data"`
}

func NewResponseSuccess(data any) *ResponseSuccess {
	return &ResponseSuccess{data}
}

// ResponseBad for error with code
type ResponseBad struct {
	ErrCode string `json:"errorCode"`
}

func (rb *ResponseBad) Error() string {
	return rb.ErrCode
}

func (rb *ResponseBad) Marshal() string {
	data, _ := json.Marshal(rb)
	return string(data)
}

func NewResponseBad(errCode string) *ResponseBad {
	return &ResponseBad{ErrCode: errCode}
}

// ResponseBadValidation for validation error
type ResponseBadValidation struct {
	IsBadValidation bool   `json:"isBadValidation"`
	Message         string `json:"message"`
}

func (rb *ResponseBadValidation) Marshal() string {
	data, _ := json.Marshal(rb)
	return string(data)
}

func (rb *ResponseBadValidation) Error() string {
	return rb.Message
}

func NewResponseBadValidation(message string) *ResponseBadValidation {
	return &ResponseBadValidation{IsBadValidation: true, Message: message}
}
