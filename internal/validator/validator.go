package validator

import (
	"strings"
	"unicode/utf8"
)

type Validator struct {
	FieldErrors map[string]string
}

func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0
}

func (v *Validator) AddFieldError(key, message string) {
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}
	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = message
	}
}

func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func MaxChar(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

func MinNum(value int, n int) bool {
	return value >= n
}

func PermittedInt(value int, permitedValue ...int) bool {
	for i := range permitedValue {
		if value == permitedValue[i] {
			return true
		}
	}
	return false
}
