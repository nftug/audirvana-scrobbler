package option

type Patch[T comparable] struct {
	option    Option[T]
	isDefined bool
}

func PatchSome[T comparable](value T) Patch[T] {
	return Patch[T]{option: Some(value), isDefined: true}
}

func PatchNone[T comparable]() Patch[T] {
	return Patch[T]{option: None[T](), isDefined: true}
}

func PatchUndefined[T comparable]() Patch[T] {
	return Patch[T]{option: None[T](), isDefined: false}
}

func (p Patch[T]) IsUndefined() bool { return !p.isDefined }

func (p Patch[T]) IsNull() bool { return p.isDefined && p.option.IsNone() }

func (p Patch[T]) IsSome() bool { return p.isDefined && p.option.IsSome() }

func (p Patch[T]) TryUnwrapOption() (Option[T], bool) {
	return p.option, p.isDefined
}

func (p Patch[T]) TryUnwrap() (*T, bool) {
	if !p.isDefined {
		return nil, false
	}
	return p.option.Unwrap(), true
}

func (p Patch[T]) MarshalJSON() ([]byte, error) {
	if !p.isDefined {
		return nil, nil
	}
	return p.option.MarshalJSON()
}

func (p *Patch[T]) UnmarshalJSON(data []byte) error {
	// Commit-style decode: only mutate receiver after successful decode.
	// If a field is present (even if null), encoding/json calls UnmarshalJSON with data != nil.
	if data == nil {
		*p = PatchUndefined[T]()
		return nil
	}

	var opt Option[T]
	if err := opt.UnmarshalJSON(data); err != nil {
		return err
	}

	*p = Patch[T]{option: opt, isDefined: true}
	return nil
}
