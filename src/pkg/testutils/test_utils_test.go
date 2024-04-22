package testutils

import (
	"net/http/httptest"
	"testing"
)

type StructTest struct {
	A int
}

func TestMockJson(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	MockJsonPost(ctx, StructTest{
		A: 125,
	})
}

func TestMockJsonNil(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	MockJsonPost(ctx, nil)
}
