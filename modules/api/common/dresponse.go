package common

//DefaultResponse default payload response
type DefaultResponse struct {
	Code       string      `json:"code"`
	Message    string      `json:"message"`
	Messages   []string    `json:"messages,omitempty"`
	Payload    interface{} `json:"payload,omitempty"`
	HttpStatus int         `json:"http_status,omitempty"`
}

//NewInternalServerErrorResponse default internal server error response
func NewInternalServerErrorResponse() DefaultResponse {
	return DefaultResponse{
		"500",
		"Internal server error",
		[]string{},
		nil,
		500,
	}
}

//NewUnathorizedErrorResponse default internal server error response
func NewUnathorizedErrorResponse() DefaultResponse {
	return DefaultResponse{
		"401",
		"Unauthorized",
		[]string{},
		nil,
		500,
	}
}

func NewErrorResponse(rc string, msgs []string) DefaultResponse {
	resCode := GetErrorMessage(rc, msgs)
	return DefaultResponse{
		resCode.RC,
		resCode.Message,
		resCode.Messages,
		nil,
		resCode.HttpStatus,
	}
}

//NewNotFoundResponse default not found error response
func NewNotFoundResponse() DefaultResponse {
	return DefaultResponse{
		"404",
		"Not found",
		[]string{},
		nil,
		404,
	}
}

//NewBadRequestResponse default not found error response
func NewBadRequestResponse() DefaultResponse {
	return DefaultResponse{
		"400",
		"Bad request",
		[]string{},
		nil,
		400,
	}
}

//NewConflictResponse default not found error response
func NewConflictResponse() DefaultResponse {
	return DefaultResponse{
		"409",
		"Data has been modified",
		[]string{},
		nil,
		409,
	}
}
