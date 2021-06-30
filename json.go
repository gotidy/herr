package herr

import (
	"encoding/json"
	"net/http"
)

// Logger is the errors logger.
type Logger interface {
	Error(message string, requestID string, err error)
}

// LoggerFunc implements the Logger interface for the logging function.
type LoggerFunc func(message string, requestID string, err error)

// Error logs the error.
func (o LoggerFunc) Error(message string, requestID string, err error) {
	o(message, requestID, err)
}

type Raw = []byte

var _ = Writer(JSONWriter{})

// JSONWriter is the JSON response writer.
type JSONWriter struct {
	log Logger
}

// NewJSONWriter creates new JSONWriter.
func NewJSONWriter(log Logger) *JSONWriter {
	return &JSONWriter{log: log}
}

// Write response.
func (j JSONWriter) Write(w http.ResponseWriter, r *http.Request, status int, body interface{}, options ...Option) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	var err error
	if bytes, ok := body.(Raw); ok {
		_, err = w.Write(bytes)
	} else {
		err = json.NewEncoder(w).Encode(body)
	}
	if err != nil {
		j.log.Error("writing response error", requestID(w, r), err)
	}
}

// Write OK (200) response with data encoded to JSON.
func (j JSONWriter) OK(w http.ResponseWriter, r *http.Request, body interface{}, options ...Option) {
	j.Write(w, r, http.StatusOK, body, options...)
}

// Write Created (201) response with data encoded to JSON.
func (j JSONWriter) Created(w http.ResponseWriter, r *http.Request, body interface{}, options ...Option) {
	j.Write(w, r, http.StatusCreated, body, options...)
}

type jsonError struct {
	Error interface{}
}

// Write Error response with error encoded to JSON.
func (j JSONWriter) Error(w http.ResponseWriter, r *http.Request, err error, options ...Option) {
	j.log.Error("request error", requestID(w, r), err)

	errorData := ToError(err).SetRequestID(requestID(w, r))

	j.Write(w, r, errorData.StatusCode(), jsonError{Error: errorData})
}

func requestID(w http.ResponseWriter, r *http.Request) string {
	if requestID := w.Header().Get("Request-Id"); requestID != "" {
		return requestID
	}
	if requestID := w.Header().Get("X-Request-Id"); requestID != "" {
		return requestID
	}
	if requestID := r.Header.Get("Request-Id"); requestID != "" {
		return requestID
	}
	if requestID := r.Header.Get("X-Request-Id"); requestID != "" {
		return requestID
	}
	return ""
}
