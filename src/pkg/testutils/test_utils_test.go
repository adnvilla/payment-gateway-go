package testutils

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

type StructTest struct {
	A int
}

func TestMockJsonPost(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	MockJsonPost(ctx, StructTest{
		A: 125,
	})
}

func TestMockJsonPostNil(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	MockJsonPost(ctx, nil)
}

func TestMockJsonGet(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	params := url.Values{}
	params.Add("p1", "v1")
	MockJsonGet(ctx, params)
}

func TestMockJsonGetNil(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	MockJsonGet(ctx, nil)
}
