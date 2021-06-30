package herr

import "sync"

const (
	InvalidParametersReason    = "INVALID_PARAMETERS"
	ValidationReason           = "VALIDATION_ERROR"
	InternalServerReason       = "INTERNAL_SERVER_ERROR"
	NotFoundReason             = "NOT_FOUND"
	PaymentRequiredReason      = "PAYMENT_REQUIRED"
	BadRequestReason           = "BAD_REQUEST"
	ForbiddenReason            = "FORBIDDEN"
	UnauthorizedReason         = "UNAUTHORIZED"
	ConflictReason             = "CONFLICT"
	UnsupportedMediaTypeReason = "UNSUPPORTED_MEDIA_TYPE"
	InvalidJSONReason          = "INVALID_JSON"
)

type ReasonList struct {
	list []string
	mu   sync.RWMutex
}

func (r *ReasonList) Get() []string {
	r.mu.RLock()
	list := r.list
	r.mu.RUnlock()
	return list
}

func (r *ReasonList) Add(reasons ...string) {
	r.mu.Lock()
	r.list = append(r.list, reasons...)
	r.mu.Unlock()
}

var Reasons = ReasonList{list: []string{
	InvalidParametersReason,
	ValidationReason,
	InternalServerReason,
	NotFoundReason,
	PaymentRequiredReason,
	BadRequestReason,
	ForbiddenReason,
	UnauthorizedReason,
	ConflictReason,
	UnsupportedMediaTypeReason,
}}
