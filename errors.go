package herr

import (
	"net/http"
)

var ErrNotFound = Error{
	CodeField:    http.StatusNotFound,
	StatusField:  http.StatusText(http.StatusNotFound),
	ReasonField:  NotFoundReason,
	MessageField: "the requested resource could not be found",
}

var ErrUnauthorized = Error{
	CodeField:    http.StatusUnauthorized,
	StatusField:  http.StatusText(http.StatusUnauthorized),
	ReasonField:  UnauthorizedReason,
	MessageField: "the request could not be authorized",
}

var ErrForbidden = Error{
	CodeField:    http.StatusForbidden,
	StatusField:  http.StatusText(http.StatusForbidden),
	ReasonField:  ForbiddenReason,
	MessageField: "the requested action was forbidden",
}

var ErrInternalServerError = Error{
	CodeField:    http.StatusInternalServerError,
	StatusField:  http.StatusText(http.StatusInternalServerError),
	ReasonField:  InternalServerReason,
	MessageField: "an internal server error occurred, please contact the system administrator",
}

var ErrBadRequest = Error{
	CodeField:    http.StatusBadRequest,
	StatusField:  http.StatusText(http.StatusBadRequest),
	ReasonField:  BadRequestReason,
	MessageField: "the request was malformed or contained invalid parameters",
}

var ErrUnsupportedMediaType = Error{
	CodeField:    http.StatusUnsupportedMediaType,
	StatusField:  http.StatusText(http.StatusUnsupportedMediaType),
	ReasonField:  UnsupportedMediaTypeReason,
	MessageField: "the request is using an unknown content type",
}

var ErrConflict = Error{
	CodeField:    http.StatusConflict,
	StatusField:  http.StatusText(http.StatusConflict),
	ReasonField:  ConflictReason,
	MessageField: "the resource could not be created due to a conflict",
}

var ErrPaymentRequired = Error{
	CodeField:    http.StatusPaymentRequired,
	StatusField:  http.StatusText(http.StatusPaymentRequired),
	ReasonField:  PaymentRequiredReason,
	MessageField: "the request can not be processed until the client makes a payment",
}

var ErrInvalidJSON = ErrBadRequest.WithDescription("Invalid JSON data").WithReason(InvalidJSONReason)
