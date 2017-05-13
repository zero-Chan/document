package godoc

import (
	"reflect"
)

type BoolSection struct {
	name string
	data bool
}

func CreateBoolSection(name string, data bool) BoolSection {
	sec := BoolSection{
		name: name,
		data: data,
	}
	return sec
}

func NewBoolSection(name string, data bool) *BoolSection {
	sec := CreateBoolSection(name, data)
	return &sec
}

func (sec BoolSection) Type() Type {
	return Bool
}

func (sec BoolSection) Name() string {
	return sec.name
}

func (sec *BoolSection) SetName(name string) {
	sec.name = name
}

func (sec BoolSection) Bool() bool {
	return sec.data
}

func (sec *BoolSection) Unmarshal(data interface{}) error {
	dataVal := reflect.ValueOf(data)
	if dataVal.Kind() != reflect.Ptr || dataVal.IsNil() {
		return ErrorUnmarshalFail{Type: Bool, Value: dataVal}
	}

	rv := dataVal.Elem()
	sec.setValue(&rv)

	return nil
}

func (sec *BoolSection) setValue(rv *reflect.Value) error {
	if rv == nil {
		return ErrorAcceptNilParam{FunctionName: "BoolSection.setValue", NilParam: rv}
	}

	if !rv.CanSet() {
		return ErrorSetValueFail{Type: Bool, KeyName: sec.name, Value: *rv}
	}

	switch rv.Kind() {
	case reflect.Interface:
		rv.Set(reflect.ValueOf(sec.data))

	case reflect.Bool:
		rv.SetBool(sec.data)

	case reflect.Ptr:
		if rv.Type().Elem().Kind() != reflect.Bool {
			return ErrorSetValueFail{Type: Bool, KeyName: sec.name, Value: *rv}
		}

		rv.Set(reflect.ValueOf(&sec.data))
	default:
		return ErrorSetValueFail{Type: Bool, KeyName: sec.name, Value: *rv}
	}

	return nil
}

func (sec *BoolSection) Marshal(data interface{}) error {
	dataVal := reflect.ValueOf(data)
	return sec.getValue(dataVal)
}

func (sec *BoolSection) getValue(rv reflect.Value) error {
	switch rv.Kind() {
	case reflect.Ptr:
		return sec.getValue(rv.Elem())
	case reflect.Bool:
		sec.data = rv.Bool()
	default:
		return ErrorMarshalFail{Type: Bool, InvalidKind: rv.Kind(), ValidKind: reflect.Bool}
	}

	return nil
}
