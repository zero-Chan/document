package document

import (
	//	"fmt"

	"reflect"
)

type NumberSection struct {
	name string
	data float64
	t    Type
}

func CreateIntSection(name string, data int) NumberSection {
	sec := NumberSection{
		name: name,
		data: float64(data),
		t:    Int,
	}
	return sec
}

func NewIntSection(name string, data int) *NumberSection {
	sec := CreateIntSection(name, data)
	return &sec
}

func CreateInt8Section(name string, data int8) NumberSection {
	sec := NumberSection{
		name: name,
		data: float64(data),
		t:    Int8,
	}
	return sec
}

func NewInt8Section(name string, data int8) *NumberSection {
	sec := CreateInt8Section(name, data)
	return &sec
}

func CreateInt16Section(name string, data int16) NumberSection {
	sec := NumberSection{
		name: name,
		data: float64(data),
		t:    Int16,
	}
	return sec
}

func NewInt16Section(name string, data int16) *NumberSection {
	sec := CreateInt16Section(name, data)
	return &sec
}

func CreateInt32Section(name string, data int32) NumberSection {
	sec := NumberSection{
		name: name,
		data: float64(data),
		t:    Int32,
	}
	return sec
}

func NewInt32Section(name string, data int32) *NumberSection {
	sec := CreateInt32Section(name, data)
	return &sec
}

func CreateInt64Section(name string, data int64) NumberSection {
	sec := NumberSection{
		name: name,
		data: float64(data),
		t:    Int64,
	}
	return sec
}

func NewInt64Section(name string, data int64) *NumberSection {
	sec := CreateInt64Section(name, data)
	return &sec
}

func CreateFloat32Section(name string, data float32) NumberSection {
	sec := NumberSection{
		name: name,
		data: float64(data),
		t:    Float32,
	}
	return sec
}

func NewFloat32Section(name string, data float32) *NumberSection {
	sec := CreateFloat32Section(name, data)
	return &sec
}

func CreateFloat64Section(name string, data float64) NumberSection {
	sec := NumberSection{
		name: name,
		data: data,
		t:    Float64,
	}
	return sec
}

func NewFloat64Section(name string, data float64) *NumberSection {
	sec := CreateFloat64Section(name, data)
	return &sec
}

func (sec NumberSection) Type() Type {
	return sec.t
}

func (sec NumberSection) Name() string {
	return sec.name
}

func (sec NumberSection) Int() int {
	return int(sec.data)
}

func (sec *NumberSection) SetInt(data int) {
	sec.data = float64(data)
	sec.t = Int
}

func (sec NumberSection) Int8() int8 {
	return int8(sec.data)
}

func (sec *NumberSection) SetInt8(data int8) {
	sec.data = float64(data)
	sec.t = Int8
}

func (sec NumberSection) Int16() int16 {
	return int16(sec.data)
}

func (sec *NumberSection) SetInt16(data int16) {
	sec.data = float64(data)
	sec.t = Int16
}

func (sec NumberSection) Int32() int32 {
	return int32(sec.data)
}

func (sec *NumberSection) SetInt32(data int32) {
	sec.data = float64(data)
	sec.t = Int32
}

func (sec NumberSection) Int64() int64 {
	return int64(sec.data)
}

func (sec *NumberSection) SetInt64(data int64) {
	sec.data = float64(data)
	sec.t = Int32
}

func (sec NumberSection) Float32() float32 {
	return float32(sec.data)
}

func (sec *NumberSection) SetFloat32(data float32) {
	sec.data = float64(data)
	sec.t = Float32
}

func (sec NumberSection) Float64() float64 {
	return sec.data
}

func (sec *NumberSection) SetFloat64(data float64) {
	sec.data = data
	sec.t = Float64
}

func (sec *NumberSection) Unmarshal(data interface{}) error {
	dataVal := reflect.ValueOf(data)
	if dataVal.Kind() != reflect.Ptr || dataVal.IsNil() {
		return ErrorUnmarshalFail{Type: Number, Value: dataVal}
	}

	rv := dataVal.Elem()
	sec.setValue(&rv)

	return nil
}

func (sec *NumberSection) setValue(rv *reflect.Value) error {
	if rv == nil {
		return ErrorAcceptNilParam{FunctionName: "NumberSection.setValue", NilParam: rv}
	}

	if !rv.CanSet() {
		return ErrorSetValueFail{Type: Number, KeyName: sec.name, Value: *rv}
	}

	switch rv.Kind() {
	case reflect.Interface:
		rv.Set(reflect.ValueOf(sec.data))
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		rv.SetInt(int64(sec.data))
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		rv.SetFloat(sec.data)

	case reflect.Ptr:
		switch rv.Type().Kind() {
		case reflect.Int:
			fallthrough
		case reflect.Int8:
			fallthrough
		case reflect.Int16:
			fallthrough
		case reflect.Int32:
			fallthrough
		case reflect.Int64:
			fallthrough
		case reflect.Float32:
			fallthrough
		case reflect.Float64:
			rv.Set(reflect.ValueOf(&sec.data))
		default:
			return ErrorSetValueFail{Type: Number, KeyName: sec.name, Value: *rv}
		}

	default:
		return ErrorSetValueFail{Type: Number, KeyName: sec.name, Value: *rv}
	}

	return nil
}

func (sec *NumberSection) Marshal(data interface{}) error {
	dataVal := reflect.ValueOf(data)
	return sec.getValue(dataVal)
}

func (sec *NumberSection) getValue(rv reflect.Value) error {
	switch rv.Kind() {
	case reflect.Ptr:
		return sec.getValue(rv.Elem())
	case reflect.Int:
		sec.SetInt(int(rv.Int()))
	case reflect.Int8:
		sec.SetInt8(int8(rv.Int()))
	case reflect.Int16:
		sec.SetInt16(int16(rv.Int()))
	case reflect.Int32:
		sec.SetInt32(int32(rv.Int()))
	case reflect.Int64:
		sec.SetInt64(rv.Int())
	case reflect.Float32:
		sec.SetFloat32(float32(rv.Float()))
	case reflect.Float64:
		sec.SetFloat64(rv.Float())
	default:
		return ErrorMarshalFail{Type: Number, InvalidKind: rv.Kind(), ValidKind: reflect.Float64}
	}

	return nil
}
