package enum

import (
	"fmt"
	"reflect"
	"slices"
)

type (
	numeric interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
	}
	identifier interface {
		// EnumUid must be implemented to be an Enum.
		// It must return a unique identifier for the concrete Enum type.
		EnumUid() string
	}
	// Enum represents a numeric enum.
	Enum interface {
		numeric
		identifier
	}
)

// typeName returns the type name of the given Enum type using
//
//   - reflect.TypeOf(e).PkgPath() if the descriptive argument is true
//
//   - fmt.Sprintf and %T verb if the descriptive argument is false.
func typeName[E Enum](e E, descriptive bool) string {
	if !descriptive {
		return fmt.Sprintf("%T", e)
	}
	reflectType := reflect.TypeOf(e)
	return fmt.Sprintf("%s.%s", reflectType.PkgPath(), reflectType.Name())
}

// New returns the Enum value of the given description if it is one of the registered values of the given Enum type.
func New[E Enum](desc string) (*E, error) {
	var e E
	anyMember, registered := enumSet.Get(e.EnumUid())
	if !registered {
		return nil, errNotRegisteredYet(typeName(e, true))
	}
	member := anyMember.(*setMember[E])
	for enum, description := range member.description {
		if description == desc {
			return &enum, nil
		}
	}
	oneofStrings := make([]string, len(member.oneof))
	for i := range member.oneof {
		oneofStrings[i] = member.description[member.oneof[i]]
	}
	return nil, errInvalidValue(member.shortName, oneofStrings, desc)
}

// Is checks if the given enum value is one of the target values of the given Enum type.
func Is[E Enum](enum E, target E, or ...E) bool {
	if enum == target {
		return true
	}
	for i := range or {
		if enum == or[i] {
			return true
		}
	}
	return false
}

// Validate checks if the given enum value is one of the registered values of the given Enum type.
func Validate[E Enum](e E) error {
	anyMember, registered := enumSet.Get(e.EnumUid())
	if !registered {
		return errNotRegisteredYet(typeName(e, true))
	}
	member := anyMember.(*setMember[E])
	for i := range member.oneof {
		if e == member.oneof[i] {
			return nil
		}
	}
	return errInvalidValue(member.shortName, member.oneof, e)
}

// String returns the description of the given Enum value or an empty string if the description is not found.
// It panics if the given Enum type is not registered yet.
func String[E Enum](e E) string {
	anyMember, registered := enumSet.Get(e.EnumUid())
	if !registered {
		panic(errNotRegisteredYet(typeName(e, true)))
	}
	member := anyMember.(*setMember[E])
	if desc, found := member.description[e]; found {
		return desc
	}
	return ""
}

// Strings returns the descriptions of the registered values of the given Enum type.
// It panics if the given Enum type is not registered yet.
func Strings[E Enum]() []string {
	var e E
	anyMember, registered := enumSet.Get(e.EnumUid())
	if !registered {
		panic(errNotRegisteredYet(typeName(e, true)))
	}
	member := anyMember.(*setMember[E])
	descriptions := make([]string, len(member.oneof))
	for i := range member.oneof {
		descriptions[i] = member.description[member.oneof[i]]
	}
	return descriptions
}

// Values returns the registered values of the given Enum type except the optional "but" values.
// It panics if the given Enum type is not registered yet.
func Values[E Enum](but ...E) []E {
	var e E
	anyMember, registered := enumSet.Get(e.EnumUid())
	if !registered {
		panic(errNotRegisteredYet(typeName(e, true)))
	}
	member := anyMember.(*setMember[E])
	if len(but) == 0 {
		return member.oneof
	}
	result := make([]E, 0, len(member.oneof))
	for i := range member.oneof {
		if !Is(member.oneof[i], but[0], but[1:]...) {
			result = append(result, member.oneof[i])
		}
	}
	return slices.Clip(result)
}
