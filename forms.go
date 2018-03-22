package forms

import "regexp"

type Form struct {
	Action uint8
	Fields []Field
	Error  string
}

type Field struct {
	Name, Error, Value, Placeholder       string
	Options                               []Option
	Max, Min, MaxLen, MinLen              int
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
