package forms

import "regexp"

type Form struct {
	Action uint8
	Fields []Field
	Error  error
}

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

type Option struct {
	Label, Value string
	Selected     bool
}

func (f *Field)Str()string{
	return f.Value.(string)
}

func (f *Field)StrSlice()[]string{
	return f.Value.([]string)
}

func (f *Field)Float()float32{
	return f.Value.(float32)
}

func (f *Field)Float64()float64{
	return f.Value.(float64)
}

func (f *Field)Uint()uint{
	return f.Value.(uint)
}

func (f *Field)UintSlice()[]uint{
	return f.Value.([]uint)
}

func (f *Field)Checked()bool{
	return f.Value.(bool)
}