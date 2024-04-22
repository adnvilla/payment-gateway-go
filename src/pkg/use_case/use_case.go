package use_case

import "context"

type UseCase[TRequest any, TResponse any] interface {
	Handle(ctx context.Context, in TRequest) (TResponse, error)
}

type Input interface{}
type Output interface{}
