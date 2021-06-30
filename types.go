package herr

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Item struct {
	Name        string `json:"name,omitempty"`
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}

// Error.
type Error struct {
	CodeField        int                    `json:"code,omitempty"`        // The Code of staus.
	StatusField      string                 `json:"status,omitempty"`      // Status is the status text
	ReasonField      string                 `json:"reason,omitempty"`      // The reason of the error.
	RequestIDField   string                 `json:"request_id,omitempty"`  // RequestIDField is request ID.
	MessageField     string                 `json:"message,omitempty"`     // Message of error.
	DescriptionField string                 `json:"description,omitempty"` // Description of error.
	DebugField       string                 `json:"debug,omitempty"`       // Debug info
	DetailsField     map[string]interface{} `json:"details,omitempty"`     // Details
	ItemsField       []Item                 `json:"items,omitempty"`       // Items

	err        error
	errWrapped bool
}

func (e *Error) Unwrap() error {
	return e.err
}

func (e *Error) Wrap(err error) {
	e.errWrapped = true
	e.err = err
}

func (e Error) WithWrap(err error) *Error {
	e.errWrapped = true
	e.err = err
	return &e
}

func (e Error) Is(err error) bool {
	switch te := err.(type) {
	case Error:
		return e.MessageField == te.MessageField &&
			e.StatusField == te.StatusField &&
			e.CodeField == te.CodeField &&
			e.ReasonField == te.ReasonField
	case *Error:
		return e.MessageField == te.MessageField &&
			e.StatusField == te.StatusField &&
			e.CodeField == te.CodeField &&
			e.ReasonField == te.ReasonField
	default:
		return false
	}
}

func (e Error) As(target interface{}) bool {
	switch te := target.(type) {
	case *Error:
		*te = e
		return true
	default:
		return false
	}
}

func (e Error) Status() string {
	return e.StatusField
}

func (e Error) Error() string {
	switch {
	case e.err != nil && e.MessageField == "":
		return e.err.Error()
	case e.err != nil && e.errWrapped:
		return e.MessageField + ": " + e.err.Error()
	}
	return e.MessageField
}

func (e Error) Message() string {
	return e.MessageField
}

func (e Error) RequestID() string {
	return e.RequestIDField
}

func (e Error) Reason() string {
	return e.ReasonField
}

func (e Error) Debug() string {
	return e.DebugField
}

func (e Error) Details() map[string]interface{} {
	return e.DetailsField
}

func (e Error) Detail(key string) (interface{}, bool) {
	if e.DetailsField == nil {
		return nil, false
	}
	value, ok := e.DetailsField[key]
	return value, ok
}

func (e Error) Items() []Item {
	return e.ItemsField
}

func (e Error) StatusCode() int {
	return e.CodeField
}

// With clones the error.
func (e Error) With() *Error {
	return &e
}

// WithStatusCode clones the error with status and status code.
func (e Error) WithStatusCode(code int) *Error {
	e.CodeField = code
	e.StatusField = http.StatusText(code)
	return &e
}

// WithRequestID clones the error with request ID.
func (e Error) WithRequestID(requestID string) *Error {
	e.RequestIDField = requestID
	return &e
}

// WithRequestID clones the error with request ID.
func (e Error) WithRequest(r *http.Request) *Error {
	id := r.Header.Get("x-request-id")
	if id == "" {
		id = r.Header.Get("request-Id")
	}
	if id == "" {
		id = r.Header.Get("req_id")
	}
	e.RequestIDField = id
	return &e
}

// WithReason clones the error with reason.
func (e Error) WithReason(reason string) *Error {
	e.ReasonField = reason
	return &e
}

// WithError clones the error with.
func (e Error) WithError(message string) *Error {
	e.MessageField = message
	return &e
}

// WithErrorf clones the error with formated error message.
func (e Error) WithErrorf(message string, args ...interface{}) *Error {
	err := fmt.Errorf(message, args...)
	if err := errors.Unwrap(err); err != nil {
		e.errWrapped = false
		e.err = err
	}
	return e.WithError(e.Error())
}

// WithDescription clones the error with.
func (e Error) WithDescription(description string) *Error {
	e.DescriptionField = description
	return &e
}

// WithErrorf clones the error with formated error message.
func (e Error) WithDescriptionf(description string, args ...interface{}) *Error {
	return e.WithDescription(fmt.Sprintf(description, args...))
}

// WithDebugf clones the error with.
func (e Error) WithDebugf(debug string, args ...interface{}) *Error {
	return e.WithDebug(fmt.Sprintf(debug, args...))
}

// WithDebug clones the error with formated debug info.
func (e Error) WithDebug(debug string) *Error {
	e.DebugField = debug
	return &e
}

// WithItems clones the error with items.
func (e Error) WithItems(items ...Item) *Error {
	e.ItemsField = append(e.ItemsField, items...)
	return &e
}

// WithDetail clones the error with detail.
func (e Error) WithDetail(key string, detail interface{}) *Error {
	if e.DetailsField == nil {
		e.DetailsField = map[string]interface{}{}
	}
	e.DetailsField[key] = detail
	return &e
}

// WithDetailf clones the error with formated additional detail.
func (e Error) WithDetailf(key string, message string, args ...interface{}) *Error {
	if e.DetailsField == nil {
		e.DetailsField = map[string]interface{}{}
	}
	e.DetailsField[key] = fmt.Sprintf(message, args...)
	return &e
}

// SetStatusCode sets the status and status code.
func (e *Error) SetStatusCode(code int) *Error {
	e.CodeField = code
	e.StatusField = http.StatusText(code)
	return e
}

// SetRequestID sets the error request ID.
func (e *Error) SetRequestID(requestID string) *Error {
	e.RequestIDField = requestID
	return e
}

// SetRequestID sets the error request ID.
func (e *Error) SetRequest(r *http.Request) *Error {
	id := r.Header.Get("x-request-id")
	if id == "" {
		id = r.Header.Get("request-Id")
	}
	if id == "" {
		id = r.Header.Get("req_id")
	}
	e.RequestIDField = id
	return e
}

// SetReason sets the error reason.
func (e *Error) SetReason(reason string) *Error {
	e.ReasonField = reason
	return e
}

// SetError sets the error.
func (e *Error) SetError(message string) *Error {
	e.MessageField = message
	return e
}

// SetErrorf sets the error formated error message.
func (e *Error) SetErrorf(message string, args ...interface{}) *Error {
	err := fmt.Errorf(message, args...)
	if err := errors.Unwrap(err); err != nil {
		e.errWrapped = false
		e.err = err
	}
	return e.SetError(err.Error())
}

// SetDescription sets the description.
func (e Error) SetDescription(description string) *Error {
	e.DescriptionField = description
	return &e
}

// Setdescription sets the error formated error message.
func (e *Error) SetDescriptionf(description string, args ...interface{}) *Error {
	return e.SetDescription(fmt.Sprintf(description, args...))
}

// SetDebug sets the error formated debug info.
func (e *Error) SetDebug(debug string) *Error {
	e.DebugField = debug
	return e
}

// SetDebugf sets the debug info.
func (e *Error) SetDebugf(debug string, args ...interface{}) *Error {
	return e.SetDebug(fmt.Sprintf(debug, args...))
}

// SetItems sets the items.
func (e *Error) SetItems(items ...Item) *Error {
	e.ItemsField = append(e.ItemsField, items...)
	return e
}

// SetDetail sets details.
func (e *Error) SetDetail(key string, detail interface{}) *Error {
	if e.DetailsField == nil {
		e.DetailsField = map[string]interface{}{}
	}
	e.DetailsField[key] = detail
	return e
}

// SetDetailf sets the formated additional detail.
func (e *Error) SetDetailf(key string, message string, args ...interface{}) *Error {
	if e.DetailsField == nil {
		e.DetailsField = map[string]interface{}{}
	}
	e.DetailsField[key] = fmt.Sprintf(message, args...)
	return e
}

func (e Error) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "rid=%s\n", e.RequestIDField)
			fmt.Fprintf(s, "msg=%s\n", e.MessageField)
			fmt.Fprintf(s, "reason=%s\n", e.ReasonField)
			fmt.Fprintf(s, "details=%+v\n", e.DetailsField)
			fmt.Fprintf(s, "debug=%s\n", e.DebugField)
			fmt.Fprintf(s, "err=%s\n", e.err)
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, e.Error())
	case 'q':
		fmt.Fprintf(s, "%q", e.MessageField)
	}
}

func ToError(err error) *Error {
	e := &Error{
		CodeField:    http.StatusInternalServerError,
		DetailsField: map[string]interface{}{},
	}
	e.Wrap(err)

	if c := MessageCarrier(nil); errors.As(err, &c) {
		e.MessageField = c.Message()
	}
	if c := ReasonCarrier(nil); errors.As(err, &c) {
		e.ReasonField = c.Reason()
	}
	if c := RequestIDCarrier(nil); errors.As(err, &c) && c.RequestID() != "" {
		e.RequestIDField = c.RequestID()
	}
	if c := DetailsCarrier(nil); errors.As(err, &c) && c.Details() != nil {
		e.DetailsField = c.Details()
	}
	if c := StatusCarrier(nil); errors.As(err, &c) && c.Status() != "" {
		e.StatusField = c.Status()
	}
	if c := StatusCodeCarrier(nil); errors.As(err, &c) && c.StatusCode() != 0 {
		e.CodeField = c.StatusCode()
	}
	if c := DebugCarrier(nil); errors.As(err, &c) {
		e.DebugField = c.Debug()
	}
	if c := ItemsCarrier(nil); errors.As(err, &c) {
		e.ItemsField = c.Items()
	}

	if e.StatusField == "" {
		e.StatusField = http.StatusText(e.StatusCode())
	}

	return e
}
