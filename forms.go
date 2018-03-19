package forms

import (
	"time"
	"regexp"
)

type Form struct {
	Action uint8
	Fields []Field
	Error  error
	Expiry time.Time
}

type Field struct {
	Name, Error, Value, Placeholder       string
	Options                               []Option
	MaxLen, MinLen                        int
	Min, Max                              int
	Step, ValueFloat32                    float32
	Regex                                 *regexp.Regexp
	V8                                    func(*Field, ...string)
	ValueUint                             uint
	ValueUintSlice                        []uint
	Required, Disable, AutoFocus, Checked bool
	//size                                  uint8
}

type Option struct {
	Label, Value string
	Selected     bool
}

func Fetch(getForm func(uint8)Form, formIDs ...uint8) (f []Form) {
	for _, id := range formIDs {
		f = append(f, getForm(id))
	}
	return f
}
