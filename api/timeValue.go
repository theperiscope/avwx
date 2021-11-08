package api

import (
	"fmt"
	"time"
)

// timeValue is a type that satisfies the spf13/pflag/Value interface so we can create custom flag values with it
type timeValue time.Time

func (b *timeValue) Set(s string) error {
	v, err := time.Parse("2006-01-02T15:04:05Z", s)
	*b = (timeValue)(v)
	return err
}

func (b *timeValue) String() string {
	return fmt.Sprintf("%v", time.Time(*b).UTC().Format("2006-01-02T15:04:05Z"))
}

func (b *timeValue) Type() string {
	return "time.Time"
}
