package api

import "net/http"

type Response struct {
	Errors []string `json:"errors"`
}

type ErrorCarrier interface {
	AddErrorMessage(err string)
	AddError(err error)
	GetCode() int
}

func (r *Response) AddErrorMessage(err string) {
	r.Errors = append(r.Errors, err)
}

func (r *Response) AddError(err error) {
	r.Errors = append(r.Errors, err.Error())
}

func (r *Response) GetCode() int {
	if len(r.Errors) > 0 {
		return http.StatusBadRequest
	} else {
		return http.StatusOK
	}
}
