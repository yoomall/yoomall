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

// ok
var ErrOk = ApiError{200, "Ok"}

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
