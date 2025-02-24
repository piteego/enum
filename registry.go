package enum

import (
	"github.com/orcaman/concurrent-map/v2"
	"log"
)

var enumSet = cmap.New[any]()

type setMember[E Enum] struct {
	uid         string
	shortName   string
	fullName    string
	oneof       []E
	description map[E]string
}

func Register[E Enum](description map[E]string) {
	if len(description) == 0 {
		panic("[Enum] no description provided to be registered")
	}
	var e E
	uid := e.EnumUid()
	// check if the enum type is already registered

	if _, registered := enumSet.Get(uid); registered {
		log.Printf("[Enum] %q is already registered with uid %q.", typeName(e, true), uid)
		return
	}
	oneof := make([]E, 0, len(description))
	for key := range description {
		oneof = append(oneof, key)
	}
	// register the enum type as a member of the enumSet
	member := new(setMember[E])
	member.uid = uid
	member.shortName = typeName(e, false)
	member.fullName = typeName(e, true)
	member.oneof = oneof
	member.description = description
	enumSet.Set(uid, member)
	log.Printf("[Enum] Successfully registered %q as %v.", member.fullName, member.description)
}
