package utils

import (
	"bytes"
	"io"
	"net"
	"net/http"
	"time"

	"yoomall/src/core/helper/response"

	"github.com/gin-gonic/gin"
)

func ProxyRequest(ctx *gin.Context) {

	url := ctx.Query("url")
	req, err := http.NewRequest(ctx.Request.Method, url, nil)
	if err != nil {
		ctx.JSON(500, map[string]any{"error": err.Error()})
		return
	}

	// clone headers
	for k, v := range ctx.Request.Header {
		req.Header[k] = v
	}

	// clone query
	for k, v := range ctx.Request.URL.Query() {
		req.URL.Query()[k] = v
	}

	// clone body
	if ctx.Request.Body != nil {
		body, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(500, map[string]any{"error": err.Error()})
			return
		}
		req.Body = io.NopCloser(bytes.NewBuffer(body))
	}

	conn := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
		},
	}

	resp, err := conn.Do(req)
	if err != nil {
		response.Error(response.ErrInternalError, err.Error()).Done(ctx)
		return
	}

	defer resp.Body.Close()

	// handle 302
	if resp.StatusCode == 302 {
		ctx.Redirect(resp.StatusCode, resp.Header.Get("Location"))
		return
	}

	body := make([]byte, 0)
	if resp.Body != nil {
		body, _ = io.ReadAll(resp.Body)
	}

	ctx.Writer.Header().Set("Content-Type", resp.Header.Get("Content-Type"))

	// clone resp headers
	for k, v := range resp.Header {
		ctx.Writer.Header()[k] = v
	}

	ctx.Writer.WriteHeader(resp.StatusCode)
	ctx.Writer.Write(body)
}
