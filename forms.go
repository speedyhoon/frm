package forms

import (
	"regexp"
	"time"
)

//Form represents details within a HTML form.
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

//Str returns f.Value as a string.
func (f *Field) Str() string {
	return f.Value.(string)
}

//StrSlice returns f.Value as a string slice.
func (f *Field) StrSlice() []string {
	return f.Value.([]string)
}

//Float returns f.Value as a float32.
func (f *Field) Float() float32 {
	return f.Value.(float32)
}

//Float64 returns f.Value as a float64.
func (f *Field) Float64() float64 {
	return f.Value.(float64)
}

//Uint returns f.Value as an unsigned integer.
func (f *Field) Uint() uint {
	return f.Value.(uint)
}

//UintSlice returns f.Value as an unsigned integer slice.
func (f *Field) UintSlice() []uint {
	return f.Value.([]uint)
}

//Checked returns f.Value as a boolean.
func (f *Field) Checked() bool {
	return f.Value.(bool)
}

//Time returns f.Value as a datetime.
func (f *Field) Time() time.Time {
	return f.Value.(time.Time)
}
