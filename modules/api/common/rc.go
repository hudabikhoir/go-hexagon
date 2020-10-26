package common

import "net/http"

type ResponseCode struct {
	RC         string   `json:"rc"`
	Message    string   `json:"message"`
	Messages   []string `json:"messages"`
	HttpStatus int      `json:"http_status"`
}

//GetErrorMessage return desired response based on RC (Response Code).
func GetErrorMessage(rc string, msgs []string) ResponseCode {
	var (
		response ResponseCode
	)

	switch rc {
	default:
		response = ResponseCode{
			RC:         rc,
			Message:    "bad request",
			HttpStatus: http.StatusBadRequest,
		}
	case "00":
		response = ResponseCode{
			RC:         rc,
			Message:    "success",
			HttpStatus: http.StatusOK,
		}
	case "02":
		response = ResponseCode{
			RC:         rc,
			Message:    "maintenance",
			HttpStatus: http.StatusServiceUnavailable,
		}

	// 50 ==> client request related error
	case "50A":
		response = ResponseCode{
			RC:         rc,
			Message:    "request timeout",
			HttpStatus: http.StatusRequestTimeout,
		}
	case "50B":
		response = ResponseCode{
			RC:         rc,
			Message:    "request limit exceed",
			HttpStatus: http.StatusTooManyRequests,
		}
	case "50C":
		response = ResponseCode{
			RC:         rc,
			Message:    "empty body",
			HttpStatus: http.StatusBadRequest,
		}
	case "50D":
		response = ResponseCode{
			RC:         rc,
			Message:    "something when wrong",
			HttpStatus: http.StatusInternalServerError,
		}
	case "50E":
		response = ResponseCode{
			RC:         rc,
			Message:    "stack error",
			HttpStatus: http.StatusInternalServerError,
		}

	// 51 ==> auth related error
	case "51A":
		response = ResponseCode{
			RC:         rc,
			Message:    "missing authentication token",
			HttpStatus: http.StatusUnauthorized,
		}
	case "51B":
		response = ResponseCode{
			RC:         rc,
			Message:    "invalid credential",
			HttpStatus: http.StatusUnauthorized,
		}
	case "51C":
		response = ResponseCode{
			RC:         rc,
			Message:    "invalid signature",
			HttpStatus: http.StatusUnauthorized,
		}
	case "51D":
		response = ResponseCode{
			RC:         rc,
			Message:    "invalid or expired token",
			HttpStatus: http.StatusUnauthorized,
		}
	case "51E":
		response = ResponseCode{
			RC:         rc,
			Message:    "invalid roles",
			HttpStatus: http.StatusForbidden,
		}

	// 52 ===> data related error
	case "52A":
		response = ResponseCode{
			RC:         rc,
			Message:    "inactive data",
			HttpStatus: http.StatusBadRequest,
		}
	case "52B":
		response = ResponseCode{
			RC:         rc,
			Message:    "route or data not found or unavailable",
			HttpStatus: http.StatusNotFound,
		}
	case "52C":
		response = ResponseCode{
			RC:         rc,
			Message:    "request validation mismatch",
			Messages:   msgs,
			HttpStatus: http.StatusUnprocessableEntity,
		}
	case "52D":
		response = ResponseCode{
			RC:         rc,
			Message:    "duplicate data",
			HttpStatus: http.StatusBadRequest,
		}
	case "52E":
		response = ResponseCode{
			RC:         rc,
			Message:    "invalid or expired otp",
			HttpStatus: http.StatusBadRequest,
		}
	case "52F":
		response = ResponseCode{
			RC:         rc,
			Message:    "invalid amount",
			HttpStatus: http.StatusBadRequest,
		}
	case "52G":
		response = ResponseCode{
			RC:         rc,
			Message:    "invalid file type",
			HttpStatus: http.StatusBadRequest,
		}
	case "52H":
		response = ResponseCode{
			RC:         rc,
			Message:    "invalid csv column header",
			HttpStatus: http.StatusBadRequest,
		}
	case "52I":
		response = ResponseCode{
			RC:         rc,
			Message:    "data has been modified",
			HttpStatus: http.StatusBadRequest,
		}
	case "52L":
		response = ResponseCode{
			RC:         rc,
			Message:    "file too large",
			HttpStatus: http.StatusBadRequest,
		}

	// InternalError or unexpected error
	case "53A":
		response = ResponseCode{
			RC:         rc,
			Message:    "fail update data",
			HttpStatus: http.StatusInternalServerError,
		}
	case "53S":
		response = ResponseCode{
			RC:         rc,
			Message:    "something went wrong",
			HttpStatus: http.StatusInternalServerError,
		}
	}

	return response
}
