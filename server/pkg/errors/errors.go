package errors

import "errors"

var ErrNotFound = errors.New("not found")
var ErrNegativeAmount = errors.New("cannot top up negative amount")
var ParseErrorCode = "parse_error"
var NotFoundCode = "not_found"
var ServerErrorCode = "server_error"
var ValidationErrorCode = "validation_error"
var InvalidFormat = "invalid card format: %s"
var InvalidJson = "invalid json"
