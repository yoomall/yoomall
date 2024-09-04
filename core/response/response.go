package response

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiJsonResponse struct {
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
	HttpCode int         `json:"-"`
}

func NewApiJsonResponse(code int, message string, data interface{}, httpCode int) *ApiJsonResponse {
	return &ApiJsonResponse{
		Code:     code,
		Message:  message,
		Data:     data,
		HttpCode: httpCode,
	}
}

func (a *ApiJsonResponse) ToJson() string {
	b, _ := json.Marshal(a)
	return string(b)
}

// with ctx
func (a *ApiJsonResponse) WithCtx(ctx *gin.Context) {
	ctx.JSON(a.HttpCode, a)
}

// success
func Success(data interface{}) *ApiJsonResponse {
	return NewApiJsonResponse(http.StatusOK, "ok", data, http.StatusOK)
}

// error
func Error(error ApiError, message string) *ApiJsonResponse {
	if message == "" {
		message = error.Error()
	}
	var httpCode int = http.StatusBadRequest
	if error.IsHttpCode() {
		httpCode = error.Code
	}
	return NewApiJsonResponse(error.Code, message, nil, httpCode)
}
