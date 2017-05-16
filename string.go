package godoc

import (
	"reflect"
)

type StringSection struct {
	name string
	data string
}

func CreateStringSection(name string, data string) StringSection {
	s := StringSection{
		name: name,
		data: data,
	}

	return s
}

func NewStringSection(name string, data string) *StringSection {
	s := CreateStringSection(name, data)
	return &s
}

func (sec StringSection) Type() Type {
	return String
}

func (sec StringSection) Name() string {
	return sec.name
}

func (sec *StringSection) SetName(name string) {
	sec.name = name
}

func (sec *StringSection) Copy() *StringSection {
	return NewStringSection(sec.name, sec.data)
}

func (sec StringSection) String() string {
	return sec.data
}

func (sec *StringSection) Unmarshal(data interface{}) error {
	dataVal := reflect.ValueOf(data)
	if dataVal.Kind() != reflect.Ptr || dataVal.IsNil() {
		return ErrorUnmarshalFail{Type: String, Value: dataVal}
	}

	rv := dataVal.Elem()
	return sec.setValue(&rv)
}

func (sec *StringSection) setValue(rv *reflect.Value) error {
	if rv == nil {
		return ErrorAcceptNilParam{FunctionName: "StringSection.setValue", NilParam: rv}
	}

	if !rv.CanSet() {
		return ErrorSetValueFail{Type: String, KeyName: sec.name, Value: *rv}
	}

	switch rv.Kind() {
	case reflect.Interface:
		rv.Set(reflect.ValueOf(sec.data))

	case reflect.String:
		rv.SetString(sec.data)
	case reflect.Ptr:
		if rv.Type().Elem().Kind() != reflect.String {
			return ErrorSetValueFail{Type: String, KeyName: sec.name, Value: *rv}
		}

		rv.Set(reflect.ValueOf(&sec.data))
	default:
		return ErrorSetValueFail{Type: String, KeyName: sec.name, Value: *rv}
	}

	return nil
}

func (sec *StringSection) Marshal(data interface{}) error {
	dataVal := reflect.ValueOf(data)
	return sec.getValue(dataVal)
}

func (sec *StringSection) getValue(rv reflect.Value) error {
	switch rv.Kind() {
	case reflect.Ptr:
		return sec.getValue(rv.Elem())
	case reflect.String:
		sec.data = rv.String()
	default:
		return ErrorMarshalFail{Type: String, InvalidKind: rv.Kind(), ValidKind: reflect.String}
	}

	return nil
}
