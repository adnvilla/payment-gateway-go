package testutils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gin-gonic/gin"
)

func GetTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func MockJsonPost(c *gin.Context, content interface{}, params gin.Params, urlValues url.Values) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")

	if content != nil {
		jsonbytes, err := json.Marshal(content)
		if err != nil {
			panic(err)
		}

		// the request body must be an io.ReadCloser
		// the bytes buffer though doesn't implement io.Closer,
		// so you wrap it in a no-op closer
		c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
	}

	// set path params
	if params != nil {
		c.Params = params
	}

	// set query params
	if urlValues != nil {
		c.Request.URL.RawQuery = urlValues.Encode()
	}
}

func MockJsonGet(c *gin.Context, params gin.Params, urlValues url.Values) {
	c.Request.Method = "GET"
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("user_id", 1)

	// set path params
	if params != nil {
		c.Params = params
	}

	// set query params
	if urlValues != nil {
		c.Request.URL.RawQuery = urlValues.Encode()
	}
}

func MockJsonDelete(c *gin.Context, params gin.Params) {
	c.Request.Method = "DELETE"
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("user_id", 1)

	// set path params
	if params != nil {
		c.Params = params
	}
}

func MockJsonPut(c *gin.Context, content interface{}, params gin.Params) {
	c.Request.Method = "PUT"
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("user_id", 1)

	// set path params
	if params != nil {
		c.Params = params
	}

	if content != nil {
		jsonbytes, err := json.Marshal(content)
		if err != nil {
			panic(err)
		}

		// the request body must be an io.ReadCloser
		// the bytes buffer though doesn't implement io.Closer,
		// so you wrap it in a no-op closer
		c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
	}
}
