package testutils

import (
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type StructTest struct {
	A int
}

func TestMockJsonPost(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	urlValues := url.Values{}
	urlValues.Add("p1", "v1")
	params := []gin.Param{
		{
			Key:   "id",
			Value: "value",
		},
	}
	MockJsonPost(ctx, StructTest{
		A: 125,
	}, params, urlValues)
}

func TestMockJsonPostNil(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	MockJsonPost(ctx, nil, nil, nil)
}

func TestMockJsonGet(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	urlValues := url.Values{}
	urlValues.Add("p1", "v1")
	params := []gin.Param{
		{
			Key:   "id",
			Value: "value",
		},
	}
	MockJsonGet(ctx, params, urlValues)
}

func TestMockJsonGetNil(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	MockJsonGet(ctx, nil, nil)
}

func TestMockJsonPut(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	params := []gin.Param{
		{
			Key:   "id",
			Value: "value",
		},
	}
	MockJsonPut(ctx, StructTest{
		A: 125,
	}, params)
}

func TestMockJsonPutNil(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	MockJsonPut(ctx, nil, nil)
}

func TestMockJsonDelete(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	params := []gin.Param{
		{
			Key:   "id",
			Value: "value",
		},
	}
	MockJsonDelete(ctx, params)
}

func TestMockJsonDeleteNil(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	MockJsonDelete(ctx, nil)
}

func TestFailPostMarshal(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	params := []gin.Param{
		{
			Key:   "id",
			Value: "value",
		},
	}

	assert.Panics(t, func() {
		MockJsonPost(ctx, make(chan int), params, nil)
	})
}

func TestFailPutMarshal(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	params := []gin.Param{
		{
			Key:   "id",
			Value: "value",
		},
	}

	assert.Panics(t, func() {
		MockJsonPut(ctx, make(chan int), params)
	})
}
