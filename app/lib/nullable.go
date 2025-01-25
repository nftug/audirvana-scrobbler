package lib

import "github.com/samber/lo"

type Nullable[T any] struct {
	Value T
}

func NewNullable[T any](x *T) Nullable[T] {
	return Nullable[T]{lo.FromPtr(x)}
}

func NewNullableByVal[T any](x T) Nullable[T] {
	return Nullable[T]{x}
}

func (n Nullable[T]) ToCopiedPtr() *T {
	p := lo.EmptyableToPtr(n.Value)
	if p == nil {
		return nil
	}
	return lo.ToPtr(*p)
}
