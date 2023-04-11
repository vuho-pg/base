package api

import (
	"github.com/vuho-pg/base/pkg/util/errors"
	"github.com/vuho-pg/base/pkg/util/pagination"
	"net/http"
)

type ResponseMeta struct {
	*pagination.Page
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

type Response interface {
	GetMeta() ResponseMeta
	GetCode() int
	GetData() any
}

type response struct {
	Meta ResponseMeta `json:"meta"`
	Data any          `json:"data"`
}

func (r *response) GetMeta() ResponseMeta {
	return r.Meta
}

func (r *response) GetCode() int {
	return r.Meta.Code
}

func (r *response) GetData() any {
	return r.Data
}

func NewResponse(data any, page *pagination.Page, code int, msg string, err error) Response {
	resp := &response{
		Meta: ResponseMeta{
			Page:    page,
			Code:    code,
			Message: msg,
		},
		Data: data,
	}
	if err != nil {
		resp.Meta.Error = err.Error()
	}
	return resp
}

func Success(data any, message string) Response {
	return NewResponse(data, nil, http.StatusOK, message, nil)
}

func SuccessPagination(data any, p pagination.Page, message string) Response {
	return NewResponse(data, &p, http.StatusOK, message, nil)
}

func BadRequest(message string) Response {
	return NewResponse(nil, nil, http.StatusBadRequest, "", errors.BadRequest(message))
}

func InternalError(err error) Response {
	return NewResponse(nil, nil, http.StatusInternalServerError, "", err)
}

func FromError(err error) Response {
	errWrap, ok := err.(errors.Error)
	if ok {
		switch errWrap.Type {
		case errors.ErrBadRequest:
			return BadRequest(err.Error())
		case errors.ErrInternal:
			return InternalError(err)
		default:
			return InternalError(err)
		}
	}
	return InternalError(err)
}
