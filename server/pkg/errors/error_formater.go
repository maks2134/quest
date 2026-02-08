package errors

type Error struct {
	ErrorMsg    string                 `json:"error,omitempty"`
	Message     string                 `json:"message,omitempty"`
	Errors      map[string]interface{} `json:"errors,omitempty"`
	ErrorDetail []ErrorDetail          `json:"error_detail,omitempty"`
	StatusCode  int                    `json:"-"`
}

func (e *Error) Error() string {
	if e.Message != "" {
		return e.Message
	}
	if len(e.ErrorDetail) > 0 {
		return e.ErrorDetail[0].Detail
	}
	if e.ErrorMsg != "" {
		return e.ErrorMsg
	}
	return "unknown error"
}

type ErrorDetail struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
	Attr   string `json:"attr,omitempty"`
}

func NewError(statusCode int, details ...ErrorDetail) *Error {
	e := Error{
		Errors:      make(map[string]interface{}),
		ErrorDetail: make([]ErrorDetail, 0),
		StatusCode:  statusCode,
	}
	if len(details) > 0 {
		e.Message = details[0].Detail
		e.ErrorDetail = details
		for _, detail := range details {
			if detail.Attr != "" {
				e.Errors[detail.Attr] = detail.Detail
			} else {
				e.Errors["base"] = detail.Detail
			}
		}
	}

	return &e
}

func NewSimpleError(statusCode int, message string) *Error {
	return &Error{
		ErrorMsg:   message,
		Message:    message,
		Errors:     map[string]interface{}{"base": message},
		StatusCode: statusCode,
	}
}
