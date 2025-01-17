package validator

import (
	"strconv"
	"strings"
)

const (
	// Len represents the rule name which will be used to find the default error message.
	Len = "len"
	// LenMsg is the default error message format for len validating rule.
	LenMsg = "%s should %d character"
)

// LenString checks if length of a string equal to give size or not.
func (v *Validator) LenString(s string, size int, field, msg string) *Validator {
	v.Check(len(strings.TrimSpace(s)) == size, field, v.msg(Len, msg, field, size))

	return v
}

// LenInt checks if length of an integer is equal to given size or not.
func (v *Validator) LenInt(i, size int, field, msg string) *Validator {
	v.Check(len(strconv.Itoa(i)) == size, field, v.msg(Len, msg, field, size))

	return v
}

// LenSlice checks if length of a slice is equal to given size or not.
func (v *Validator) LenSlice(s []any, size int, field, msg string) *Validator {
	v.Check(len(s) == size, field, v.msg(Len, msg, field, size))

	return v
}
