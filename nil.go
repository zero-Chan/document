package godoc

import (
	"reflect"
)

type NilSection struct {
	name string
	data *int
}

func CreateNilSection(name string) NilSection {
	sec := NilSection{
		name: name,
		data: nil,
	}
	return sec
}

func NewNIlSection(name string) *NilSection {
	sec := CreateNilSection(name)
	return &sec
}

func (sec NilSection) Type() Type {
	return Nil
}

func (sec NilSection) Name() string {
	return sec.name
}

func (sec *NilSection) SetName(name string) {
	sec.name = name
}

func (sec *NilSection) Copy() *NilSection {
	return NewNIlSection(sec.name)
}

func (sec *NilSection) Unmarshal(data interface{}) error {
	dataVal := reflect.ValueOf(data)
	if dataVal.Kind() != reflect.Ptr || dataVal.IsNil() {
		return ErrorUnmarshalFail{Type: Nil, Value: dataVal}
	}

	rv := dataVal.Elem()
	return sec.setValue(&rv)
}

func (sec *NilSection) setValue(rv *reflect.Value) error {
	if rv == nil {
		return ErrorAcceptNilParam{FunctionName: "NilSection.setValue", NilParam: rv}
	}

	if !rv.CanSet() {
		return ErrorSetValueFail{Type: Nil, KeyName: sec.name, Value: *rv}
	}

	var v *int = nil
	switch rv.Kind() {
	case reflect.Interface:
		rv.Set(reflect.ValueOf(v))
	case reflect.Invalid:
		rv.Set(reflect.ValueOf(v))
	case reflect.Ptr:
		rv.Set(reflect.ValueOf(v))
	default:
		return ErrorSetValueFail{Type: Nil, KeyName: sec.name, Value: *rv}
	}

	return nil
}

func (sec *NilSection) Marshal(data interface{}) error {
	dataVal := reflect.ValueOf(data)
	return sec.getValue(dataVal)
}

func (sec *NilSection) getValue(rv reflect.Value) error {
	switch rv.Kind() {
	case reflect.Ptr:
		return sec.getValue(rv.Elem())
	case reflect.Invalid:
		sec.data = nil
	default:
		return ErrorMarshalFail{Type: Nil, InvalidKind: rv.Kind(), ValidKind: reflect.Invalid}
	}

	return nil
}
