package option

import "encoding/json"

type SlicePatch[T comparable] struct {
	value     []T
	isDefined bool
}

func SlicePatchSome[T comparable](value []T) SlicePatch[T] {
	copied := append([]T{}, value...)
	return SlicePatch[T]{value: copied, isDefined: true}
}

func SlicePatchEmpty[T comparable]() SlicePatch[T] {
	return SlicePatch[T]{value: []T{}, isDefined: true}
}

func SlicePatchUndefined[T comparable]() SlicePatch[T] {
	return SlicePatch[T]{value: nil, isDefined: false}
}

func (o SlicePatch[T]) IsUndefined() bool { return !o.isDefined }

func (o SlicePatch[T]) TryUnwrap() ([]T, bool) {
	if !o.isDefined {
		return nil, false
	}
	return append([]T{}, o.value...), true
}

func (o SlicePatch[T]) Unwrap() []T {
	if !o.isDefined {
		return nil
	}
	return append([]T{}, o.value...)
}

func (o SlicePatch[T]) MarshalJSON() ([]byte, error) {
	if !o.isDefined {
		return nil, nil
	}
	return json.Marshal(o.value)
}

func (o *SlicePatch[T]) UnmarshalJSON(data []byte) error {
	if data == nil {
		*o = SlicePatchUndefined[T]()
		return nil
	}

	if string(data) == "null" {
		*o = SlicePatchEmpty[T]()
		return nil
	}

	var decoded []T
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	*o = SlicePatchSome(decoded)
	return nil
}
