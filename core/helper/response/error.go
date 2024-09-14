package response

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var httpCodeArray = []int{200, 400, 401, 403, 404, 500}

var ErrNotFound = ApiError{404, "Not Found"}
var ErrInternalError = ApiError{500, "Internal Error"}
var ErrInvalidArgument = ApiError{400, "Invalid Argument"}
var ErrNotAuthorized = ApiError{401, "Not Authorized"}
var ErrBadRequest = ApiError{400, "Bad Request"}
var ErrForbidden = ApiError{403, "Forbidden"}

// ok
var ErrOk = ApiError{200, "Ok"}

// custom code
var ErrNotAuth = ApiError{401, "Not Auth"}
var ErrNotPermission = ApiError{403, "Not Permission"}

func (e ApiError) Error() string {
	return e.Message
}

// is http code
func (e ApiError) IsHttpCode() bool {
	for _, v := range httpCodeArray {
		if v == e.Code {
			return true
		}
	}
	return false
}
