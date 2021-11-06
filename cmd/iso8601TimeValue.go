package cmd

import (
	"fmt"

	"github.com/relvacode/iso8601"
)

// iso8601TimeValue is a type that satisfies the spf13/pflag/Value interface so we can create custom flag values with it
type iso8601TimeValue iso8601.Time

func newBoolValue(val iso8601.Time, p *iso8601.Time) *iso8601TimeValue {
	*p = val
	return (*iso8601TimeValue)(p)
}

func (b *iso8601TimeValue) Set(s string) error {
	v, err := iso8601.ParseString(s)
	*&b.Time = v
	return err
}

func (b *iso8601TimeValue) String() string {
	return fmt.Sprintf("%v", b.Time.UTC().Format("2006-01-02T15:04:05Z"))
}

func (b *iso8601TimeValue) Type() string {
	return "iso8601.Time"
}
