package api

import (
	"fmt"
	"strings"
)

// enumValue is a type that satisfies the spf13/pflag/Value interface so we can create custom flag values that offer limited string-based choices
type enumValue struct {
	allowedValues []string
	value         string
}

func NewEnumValue(allowedValues []string, defaultValue string) *enumValue {
	return &enumValue{
		allowedValues: allowedValues,
		value:         defaultValue,
	}
}

func (v *enumValue) String() string {
	return v.value
}

func (v *enumValue) Set(newValue string) error {
	isIncluded := func(options []string, value string) bool {
		for _, option := range options {
			if value == option {
				return true
			}
		}
		return false
	}
	if !isIncluded(v.allowedValues, newValue) {
		return fmt.Errorf("%s is not included in %s", newValue, strings.Join(v.allowedValues, ","))
	}
	v.value = newValue
	return nil
}

func (v *enumValue) Type() string {
	return "string"
}
