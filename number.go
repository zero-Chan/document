package godoc

import (
	"reflect"
)

type NumberSection struct {
	name  string
	idata int64
	fdata float64
	udata uint64
	t     Type
}

func CreateIntSection(name string, data int) NumberSection {
	sec := NumberSection{
		name:  name,
		idata: int64(data),
		t:     Int,
	}
	return sec
}

func NewIntSection(name string, data int) *NumberSection {
	sec := CreateIntSection(name, data)
	return &sec
}

func CreateInt8Section(name string, data int8) NumberSection {
	sec := NumberSection{
		name:  name,
		idata: int64(data),
		t:     Int8,
	}
	return sec
}

func NewInt8Section(name string, data int8) *NumberSection {
	sec := CreateInt8Section(name, data)
	return &sec
}

func CreateInt16Section(name string, data int16) NumberSection {
	sec := NumberSection{
		name:  name,
		idata: int64(data),
		t:     Int16,
	}
	return sec
}

func NewInt16Section(name string, data int16) *NumberSection {
	sec := CreateInt16Section(name, data)
	return &sec
}

func CreateInt32Section(name string, data int32) NumberSection {
	sec := NumberSection{
		name:  name,
		idata: int64(data),
		t:     Int32,
	}
	return sec
}

func NewInt32Section(name string, data int32) *NumberSection {
	sec := CreateInt32Section(name, data)
	return &sec
}

func CreateInt64Section(name string, data int64) NumberSection {
	sec := NumberSection{
		name:  name,
		idata: int64(data),
		t:     Int64,
	}
	return sec
}

func NewInt64Section(name string, data int64) *NumberSection {
	sec := CreateInt64Section(name, data)
	return &sec
}

func CreateFloat32Section(name string, data float32) NumberSection {
	sec := NumberSection{
		name:  name,
		fdata: float64(data),
		t:     Float32,
	}

	return sec
}

func NewFloat32Section(name string, data float32) *NumberSection {
	sec := CreateFloat32Section(name, data)
	return &sec
}

func CreateFloat64Section(name string, data float64) NumberSection {
	sec := NumberSection{
		name:  name,
		fdata: data,
		t:     Float64,
	}
	return sec
}

func NewFloat64Section(name string, data float64) *NumberSection {
	sec := CreateFloat64Section(name, data)
	return &sec
}

func CreateUintSection(name string, data uint) NumberSection {
	sec := NumberSection{
		name:  name,
		udata: uint64(data),
		t:     Uint,
	}
	return sec
}

func NewUintSection(name string, data uint) *NumberSection {
	sec := CreateUintSection(name, data)
	return &sec
}

func CreateUint8Section(name string, data uint8) NumberSection {
	sec := NumberSection{
		name:  name,
		udata: uint64(data),
		t:     Uint8,
	}
	return sec
}

func NewUint8Section(name string, data uint8) *NumberSection {
	sec := CreateUint8Section(name, data)
	return &sec
}

func CreateUint16Section(name string, data uint16) NumberSection {
	sec := NumberSection{
		name:  name,
		udata: uint64(data),
		t:     Uint16,
	}
	return sec
}

func NewUint16Section(name string, data uint16) *NumberSection {
	sec := CreateUint16Section(name, data)
	return &sec
}

func CreateUint32Section(name string, data uint32) NumberSection {
	sec := NumberSection{
		name:  name,
		udata: uint64(data),
		t:     Uint32,
	}
	return sec
}

func NewUint32Section(name string, data uint32) *NumberSection {
	sec := CreateUint32Section(name, data)
	return &sec
}

func CreateUint64Section(name string, data uint64) NumberSection {
	sec := NumberSection{
		name:  name,
		udata: uint64(data),
		t:     Uint64,
	}
	return sec
}

func NewUint64Section(name string, data uint64) *NumberSection {
	sec := CreateUint64Section(name, data)
	return &sec
}

func (sec NumberSection) Type() Type {
	return sec.t
}

func (sec NumberSection) Name() string {
	return sec.name
}

func (sec *NumberSection) SetName(name string) {
	sec.name = name
}

func (sec NumberSection) Int() int {
	v := int(sec.idata)
	switch sec.Type() {
	case Int:
		fallthrough
	case Int8:
		fallthrough
	case Int16:
		fallthrough
	case Int32:
		fallthrough
	case Int64:
		v = int(sec.idata)
	case Uint:
		fallthrough
	case Uint8:
		fallthrough
	case Uint16:
		fallthrough
	case Uint32:
		fallthrough
	case Uint64:
		v = int(sec.udata)
	case Float32:
		fallthrough
	case Float64:
		v = int(sec.fdata)
	}

	return v
}

func (sec *NumberSection) SetInt(data int) {
	sec.idata = int64(data)
	sec.t = Int
}

func (sec NumberSection) Int8() int8 {
	v := int8(sec.idata)
	switch sec.Type() {
	case Int:
		fallthrough
	case Int8:
		fallthrough
	case Int16:
		fallthrough
	case Int32:
		fallthrough
	case Int64:
		v = int8(sec.idata)
	case Uint:
		fallthrough
	case Uint8:
		fallthrough
	case Uint16:
		fallthrough
	case Uint32:
		fallthrough
	case Uint64:
		v = int8(sec.udata)
	case Float32:
		fallthrough
	case Float64:
		v = int8(sec.fdata)
	}

	return v
}

func (sec *NumberSection) SetInt8(data int8) {
	sec.idata = int64(data)
	sec.t = Int8
}

func (sec NumberSection) Int16() int16 {
	v := int16(sec.idata)
	switch sec.Type() {
	case Int:
		fallthrough
	case Int8:
		fallthrough
	case Int16:
		fallthrough
	case Int32:
		fallthrough
	case Int64:
		v = int16(sec.idata)
	case Uint:
		fallthrough
	case Uint8:
		fallthrough
	case Uint16:
		fallthrough
	case Uint32:
		fallthrough
	case Uint64:
		v = int16(sec.udata)
	case Float32:
		fallthrough
	case Float64:
		v = int16(sec.fdata)
	}

	return v
}

func (sec *NumberSection) SetInt16(data int16) {
	sec.idata = int64(data)
	sec.t = Int16
}

func (sec NumberSection) Int32() int32 {
	v := int32(sec.idata)
	switch sec.Type() {
	case Int:
		fallthrough
	case Int8:
		fallthrough
	case Int16:
		fallthrough
	case Int32:
		fallthrough
	case Int64:
		v = int32(sec.idata)
	case Uint:
		fallthrough
	case Uint8:
		fallthrough
	case Uint16:
		fallthrough
	case Uint32:
		fallthrough
	case Uint64:
		v = int32(sec.udata)
	case Float32:
		fallthrough
	case Float64:
		v = int32(sec.fdata)
	}

	return v
}

func (sec *NumberSection) SetInt32(data int32) {
	sec.idata = int64(data)
	sec.t = Int32
}

func (sec NumberSection) Int64() int64 {
	v := int64(sec.idata)
	switch sec.Type() {
	case Int:
		fallthrough
	case Int8:
		fallthrough
	case Int16:
		fallthrough
	case Int32:
		fallthrough
	case Int64:
		v = int64(sec.idata)
	case Uint:
		fallthrough
	case Uint8:
		fallthrough
	case Uint16:
		fallthrough
	case Uint32:
		fallthrough
	case Uint64:
		v = int64(sec.udata)
	case Float32:
		fallthrough
	case Float64:
		v = int64(sec.fdata)
	}

	return v
}

func (sec *NumberSection) SetInt64(data int64) {
	sec.idata = int64(data)
	sec.t = Int64
}

func (sec NumberSection) Uint() uint {
	v := uint(sec.idata)
	switch sec.Type() {
	case Int:
		fallthrough
	case Int8:
		fallthrough
	case Int16:
		fallthrough
	case Int32:
		fallthrough
	case Int64:
		v = uint(sec.idata)
	case Uint:
		fallthrough
	case Uint8:
		fallthrough
	case Uint16:
		fallthrough
	case Uint32:
		fallthrough
	case Uint64:
		v = uint(sec.udata)
	case Float32:
		fallthrough
	case Float64:
		v = uint(sec.fdata)
	}

	return v
}

func (sec *NumberSection) SetUint(data uint) {
	sec.udata = uint64(data)
	sec.t = Uint
}

func (sec NumberSection) Uint8() uint8 {
	v := uint8(sec.udata)
	switch sec.Type() {
	case Int:
		fallthrough
	case Int8:
		fallthrough
	case Int16:
		fallthrough
	case Int32:
		fallthrough
	case Int64:
		v = uint8(sec.idata)
	case Uint:
		fallthrough
	case Uint8:
		fallthrough
	case Uint16:
		fallthrough
	case Uint32:
		fallthrough
	case Uint64:
		v = uint8(sec.udata)
	case Float32:
		fallthrough
	case Float64:
		v = uint8(sec.fdata)
	}

	return v
}

func (sec *NumberSection) SetUint8(data uint8) {
	sec.udata = uint64(data)
	sec.t = Uint8
}

func (sec NumberSection) Uint16() uint16 {
	v := uint16(sec.udata)
	switch sec.Type() {
	case Int:
		fallthrough
	case Int8:
		fallthrough
	case Int16:
		fallthrough
	case Int32:
		fallthrough
	case Int64:
		v = uint16(sec.idata)
	case Uint:
		fallthrough
	case Uint8:
		fallthrough
	case Uint16:
		fallthrough
	case Uint32:
		fallthrough
	case Uint64:
		v = uint16(sec.udata)
	case Float32:
		fallthrough
	case Float64:
		v = uint16(sec.fdata)
	}

	return v
}

func (sec *NumberSection) SetUint16(data uint16) {
	sec.udata = uint64(data)
	sec.t = Uint16
}

func (sec NumberSection) Uint32() uint32 {
	v := uint32(sec.idata)
	switch sec.Type() {
	case Int:
		fallthrough
	case Int8:
		fallthrough
	case Int16:
		fallthrough
	case Int32:
		fallthrough
	case Int64:
		v = uint32(sec.idata)
	case Uint:
		fallthrough
	case Uint8:
		fallthrough
	case Uint16:
		fallthrough
	case Uint32:
		fallthrough
	case Uint64:
		v = uint32(sec.udata)
	case Float32:
		fallthrough
	case Float64:
		v = uint32(sec.fdata)
	}

	return v
}

func (sec *NumberSection) SetUint32(data uint32) {
	sec.udata = uint64(data)
	sec.t = Uint32
}

func (sec NumberSection) Uint64() uint64 {
	v := uint64(sec.idata)
	switch sec.Type() {
	case Int:
		fallthrough
	case Int8:
		fallthrough
	case Int16:
		fallthrough
	case Int32:
		fallthrough
	case Int64:
		v = uint64(sec.idata)
	case Uint:
		fallthrough
	case Uint8:
		fallthrough
	case Uint16:
		fallthrough
	case Uint32:
		fallthrough
	case Uint64:
		v = uint64(sec.udata)
	case Float32:
		fallthrough
	case Float64:
		v = uint64(sec.fdata)
	}

	return v
}

func (sec *NumberSection) SetUint64(data uint64) {
	sec.udata = uint64(data)
	sec.t = Uint64
}

func (sec NumberSection) Float32() float32 {
	v := float32(sec.fdata)
	switch sec.Type() {
	case Int:
		fallthrough
	case Int8:
		fallthrough
	case Int16:
		fallthrough
	case Int32:
		fallthrough
	case Int64:
		v = float32(sec.idata)
	case Uint:
		fallthrough
	case Uint8:
		fallthrough
	case Uint16:
		fallthrough
	case Uint32:
		fallthrough
	case Uint64:
		v = float32(sec.udata)
	case Float32:
		fallthrough
	case Float64:
		v = float32(sec.fdata)
	}

	return v
}

func (sec *NumberSection) SetFloat32(data float32) {
	sec.fdata = float64(data)
	sec.t = Float32
}

func (sec NumberSection) Float64() float64 {
	v := float64(sec.fdata)
	switch sec.Type() {
	case Int:
		fallthrough
	case Int8:
		fallthrough
	case Int16:
		fallthrough
	case Int32:
		fallthrough
	case Int64:
		v = float64(sec.idata)
	case Uint:
		fallthrough
	case Uint8:
		fallthrough
	case Uint16:
		fallthrough
	case Uint32:
		fallthrough
	case Uint64:
		v = float64(sec.udata)
	case Float32:
		fallthrough
	case Float64:
		v = float64(sec.fdata)
	}

	return v
}

func (sec *NumberSection) SetFloat64(data float64) {
	sec.fdata = float64(data)
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
		if err := sec.setInterfaceValue(rv); err != nil {
			return err
		}
	case reflect.Ptr:
		if err := sec.setPtrValue(rv); err != nil {
			return err
		}
	default:
		if err := sec.set(rv); err != nil {
			return err
		}
	}

	return nil
}

func (sec *NumberSection) setInterfaceValue(rv *reflect.Value) error {
	switch sec.Type() {
	case Int:
		rv.Set(reflect.ValueOf(sec.Int()))
	case Int8:
		rv.Set(reflect.ValueOf(sec.Int8()))
	case Int16:
		rv.Set(reflect.ValueOf(sec.Int16()))
	case Int32:
		rv.Set(reflect.ValueOf(sec.Int32()))
	case Int64:
		rv.Set(reflect.ValueOf(sec.Int64()))
	case Float32:
		rv.Set(reflect.ValueOf(sec.Float32()))
	case Float64:
		rv.Set(reflect.ValueOf(sec.Float64()))
	case Uint:
		rv.Set(reflect.ValueOf(sec.Uint()))
	case Uint8:
		rv.Set(reflect.ValueOf(sec.Uint8()))
	case Uint16:
		rv.Set(reflect.ValueOf(sec.Uint16()))
	case Uint32:
		rv.Set(reflect.ValueOf(sec.Uint32()))
	case Uint64:
		rv.Set(reflect.ValueOf(sec.Uint64()))
	default:
		return ErrorSetValueFail{Type: sec.Type(), KeyName: sec.name, Value: *rv}
	}

	return nil
}

func (sec *NumberSection) setPtrValue(rv *reflect.Value) error {
	if rv == nil {
		return ErrorAcceptNilParam{FunctionName: "NumberSection.setPtrValue", NilParam: rv}
	}

	if !rv.CanSet() || rv.Kind() != reflect.Ptr {
		return ErrorSetValueFail{Type: Number, KeyName: sec.name, Value: *rv}
	}

	switch rv.Type().Kind() {
	case reflect.Int:
		v := sec.Int()
		rv.Set(reflect.ValueOf(&v))
	case reflect.Int8:
		v := sec.Int8()
		rv.Set(reflect.ValueOf(&v))
	case reflect.Int16:
		v := sec.Int16()
		rv.Set(reflect.ValueOf(&v))
	case reflect.Int32:
		v := sec.Int32()
		rv.Set(reflect.ValueOf(&v))
	case reflect.Int64:
		v := sec.Int64()
		rv.Set(reflect.ValueOf(&v))
	case reflect.Float32:
		v := sec.Float32()
		rv.Set(reflect.ValueOf(&v))
	case reflect.Float64:
		v := sec.Float64()
		rv.Set(reflect.ValueOf(&v))
	case reflect.Uint:
		v := sec.Uint()
		rv.Set(reflect.ValueOf(&v))
	case reflect.Uint8:
		v := sec.Uint8()
		rv.Set(reflect.ValueOf(&v))
	case reflect.Uint16:
		v := sec.Uint16()
		rv.Set(reflect.ValueOf(&v))
	case reflect.Uint32:
		v := sec.Uint32()
		rv.Set(reflect.ValueOf(&v))
	case reflect.Uint64:
		v := sec.Uint64()
		rv.Set(reflect.ValueOf(&v))
	default:
		return ErrorSetValueFail{Type: sec.Type(), KeyName: sec.name, Value: *rv}
	}

	return nil
}

func (sec *NumberSection) set(rv *reflect.Value) error {
	switch rv.Kind() {
	case reflect.Int:
		rv.Set(reflect.ValueOf(sec.Int()))
	case reflect.Int8:
		rv.Set(reflect.ValueOf(sec.Int8()))
	case reflect.Int16:
		rv.Set(reflect.ValueOf(sec.Int16()))
	case reflect.Int32:
		rv.Set(reflect.ValueOf(sec.Int32()))
	case reflect.Int64:
		rv.Set(reflect.ValueOf(sec.Int64()))
	case reflect.Float32:
		rv.Set(reflect.ValueOf(sec.Float32()))
	case reflect.Float64:
		rv.Set(reflect.ValueOf(sec.Float64()))
	case reflect.Uint:
		rv.Set(reflect.ValueOf(sec.Uint()))
	case reflect.Uint8:
		rv.Set(reflect.ValueOf(sec.Uint8()))
	case reflect.Uint16:
		rv.Set(reflect.ValueOf(sec.Uint16()))
	case reflect.Uint32:
		rv.Set(reflect.ValueOf(sec.Uint32()))
	case reflect.Uint64:
		rv.Set(reflect.ValueOf(sec.Uint64()))
	default:
		return ErrorSetValueFail{Type: sec.Type(), KeyName: sec.name, Value: *rv}
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
	case reflect.Uint:
		sec.SetUint(uint(rv.Uint()))
	case reflect.Uint8:
		sec.SetUint8(uint8(rv.Uint()))
	case reflect.Uint16:
		sec.SetUint16(uint16(rv.Uint()))
	case reflect.Uint32:
		sec.SetUint32(uint32(rv.Uint()))
	case reflect.Uint64:
		sec.SetUint64(uint64(rv.Uint()))
	default:
		return ErrorMarshalFail{Type: sec.Type(), InvalidKind: rv.Kind(), ValidKind: reflect.Float64}
	}

	return nil
}
