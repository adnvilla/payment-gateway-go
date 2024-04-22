package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaptureOrderUseCase(t *testing.T) {
	u := NewCaptureOrderUseCase()

	out, err := u.Handle(context.Background(), CaptureOrderInput{})

	assert.NoError(t, err)
	assert.NotNil(t, out)
}

func TestCreateOrderUseCase(t *testing.T) {
	u := NewCreateOrderUseCase()

	out, err := u.Handle(context.Background(), CreateOrderInput{})

	assert.NoError(t, err)
	assert.NotNil(t, out)
}
