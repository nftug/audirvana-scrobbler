package lib

import "github.com/samber/lo"

type Nullable[T any] struct {
	value T
}

func NewNullable[T any](x *T) Nullable[T] {
	return Nullable[T]{lo.FromPtr(x)}
}

func NewNullableByVal[T any](x T) Nullable[T] {
	return Nullable[T]{x}
}

func (n Nullable[T]) Unwrap() *T {
	p := lo.EmptyableToPtr(n.value)
	if p == nil {
		return nil
	}
	return lo.ToPtr(*p)
}

func (n Nullable[T]) Raw() T {
	return n.value
}
