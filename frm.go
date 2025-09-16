package forms

import (
	"regexp"
	"time"
)

//Form represents details about a HTML form.
type Form struct {
	Action uint8
	Fields []Field
	Err    error
}

//Field represents attributes for each HTML form field.
type Field struct {
	Name, Err, Placeholder   string
	Options                  []Option
	Max, Min, MaxLen, MinLen int
	Step                     float32
	Regex                    *regexp.Regexp
	V8                       func(*Field, ...string)
	Value                    interface{}
	Required, Disable, Focus bool
	//Size                     uint8
}

//Option represents attributes for an option within a select box, list box or data list.
type Option struct {
	Label, Value string
	Selected     bool
}

//Bytes returns f.Value as a byte slice.
func (f *Field) Bytes() (o []byte) {
	if f == nil || f.Value == nil {
		return
	}

	o, _ = f.Value.([]byte)
	return
}

//Str returns f.Value as a string.
func (f *Field) Str() (o string) {
	if f == nil || f.Value == nil {
		return
	}

	o, _ = f.Value.(string)
	return
}

//Strs returns f.Value as a string slice.
func (f *Field) Strs() (o []string) {
	if f == nil || f.Value == nil {
		return
	}

	o, _ = f.Value.([]string)
	return
}

//Float returns f.Value as a float32.
func (f *Field) Float() (o float32) {
	if f == nil || f.Value == nil {
		return
	}

	o, _ = f.Value.(float32)
	return
}

//Float64 returns f.Value as a float64.
func (f *Field) Float64() (o float64) {
	if f == nil || f.Value == nil {
		return
	}

	o, _ = f.Value.(float64)
	return
}

//Uint returns f.Value as an unsigned integer.
func (f *Field) Uint() (o uint) {
	if f == nil || f.Value == nil {
		return
	}

	o, _ = f.Value.(uint)
	return
}

//Uints returns f.Value as an unsigned integer slice.
func (f *Field) Uints() (o []uint) {
	if f == nil || f.Value == nil {
		return
	}

	o, _ = f.Value.([]uint)
	return
}

//Checked returns f.Value as a boolean.
func (f *Field) Checked() (o bool) {
	if f == nil || f.Value == nil {
		return
	}

	o, _ = f.Value.(bool)
	return
}

//Time returns f.Value as a datetime.
func (f *Field) Time() (o time.Time) {
	if f == nil || f.Value == nil {
		return
	}

	o, _ = f.Value.(time.Time)
	return
}
