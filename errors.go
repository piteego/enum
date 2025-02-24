package enum

import (
	"errors"
	"fmt"
)

var (
	// ErrNotRegisteredYet is returned when the given Enum type is not registered in the internal registry.
	// Register the Enum values to fix this error.
	ErrNotRegisteredYet = errors.New("enum not registered yet")
	// ErrInvalidValue is returned when the given value is not one of the registered values of the given Enum type.
	ErrInvalidValue = errors.New("invalid enum value")
)

func errNotRegisteredYet(enumName string) error {
	return fmt.Errorf("[Enum] %q %w", enumName, ErrNotRegisteredYet)
}

func errInvalidValue[V interface{ ~string | Enum }](enumName string, expected []V, got V) error {
	return fmt.Errorf("[Enum] %w for %s: must be one of %v, got %v", ErrInvalidValue, enumName, expected, got)
}
