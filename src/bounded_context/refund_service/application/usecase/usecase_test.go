package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRefundUseCase(t *testing.T) {
	u := NewCreateRefundUseCase()

	out, err := u.Handle(context.Background(), CreateRefundInput{})

	assert.NoError(t, err)
	assert.NotNil(t, out)
}

func TestGetRefundUseCase(t *testing.T) {
	u := NewGetRefundUseCase()

	out, err := u.Handle(context.Background(), GetRefundInput{})

	assert.NoError(t, err)
	assert.NotNil(t, out)
}
