package herr

import "net/http"

type Option = func(w http.ResponseWriter)

type Writer interface {
	Write(w http.ResponseWriter, r *http.Request, status int, body interface{}, options ...Option)
	OK(w http.ResponseWriter, r *http.Request, body interface{}, options ...Option)
	Created(w http.ResponseWriter, r *http.Request, body interface{}, options ...Option)
	Error(w http.ResponseWriter, r *http.Request, err error, options ...Option)
}
