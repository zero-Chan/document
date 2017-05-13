package godoc

import (
	"reflect"
)

type ArraySection struct {
	name string
	data []*Document
}

func CreateArraySection(name string) ArraySection {
	sec := ArraySection{
		name: name,
		data: make([]*Document, 0),
	}
	return sec
}

func NewArraySection(name string) *ArraySection {
	sec := CreateArraySection(name)
	return &sec
}

func (sec ArraySection) Type() Type {
	return Array
}

func (sec ArraySection) Name() string {
	return sec.name
}

func (sec *ArraySection) SetName(name string) {
	sec.name = name
}

func (sec ArraySection) Len() int {
	return len(sec.data)
}

func (sec ArraySection) Cap() int {
	return cap(sec.data)
}

func (sec *ArraySection) Delete(idx int) {
	sec.data = append(sec.data[:idx], sec.data[idx+1:]...)
}

func (sec *ArraySection) Pop() {
	if len(sec.data) == 0 {
		return
	}
	sec.data = sec.data[:len(sec.data)-1]
}

func (sec *ArraySection) Clear() {
	sec.data = make([]*Document, 0)
}

func (sec *ArraySection) AppendDoc(docs ...*Document) {
	for _, doc := range docs {
		if doc == nil {
			doc = NewDocument(NewNIlSection(""))
		}
		sec.data = append(sec.data, doc)
	}
}

func (sec *ArraySection) Append(vals ...Section) {
	for _, val := range vals {
		if val == nil {
			val = NewNIlSection("")
		}
		sec.data = append(sec.data, NewDocument(val))
	}

}

func (sec *ArraySection) At(idx int) (*Document, error) {
	if idx >= len(sec.data) {
		return nil, ErrorArrayOutOfRange{ArrayKey: sec.name, MaxMember: len(sec.data), InvalidIndex: idx}
	}

	return sec.data[idx], nil
}

func (sec *ArraySection) AppendObject(vals ...*ObjectSection) {
	for _, val := range vals {
		sec.data = append(sec.data, NewDocument(val))
	}
}

func (sec *ArraySection) CanObject(idx int) bool {
	if idx >= len(sec.data) {
		return false
	}

	return sec.data[idx].IsObject()
}

func (sec *ArraySection) AtObject(idx int) (*ObjectSection, error) {
	if idx >= len(sec.data) {
		return nil, ErrorArrayOutOfRange{ArrayKey: sec.name, MaxMember: len(sec.data), InvalidIndex: idx}
	}

	return sec.data[idx].Object()
}

func (sec *ArraySection) AppendString(vals ...string) {
	for _, val := range vals {
		sec.data = append(sec.data, NewDocument(NewStringSection("", val)))
	}
}

func (sec *ArraySection) CanString(idx int) bool {
	if idx >= len(sec.data) {
		return false
	}

	return sec.data[idx].IsString()
}

func (sec *ArraySection) AtString(idx int) (string, error) {
	if idx >= len(sec.data) {
		return "", ErrorArrayOutOfRange{ArrayKey: sec.name, MaxMember: len(sec.data), InvalidIndex: idx}
	}

	strSec, err := sec.data[idx].String()
	if err != nil {
		return "", err
	}

	return strSec.String(), nil
}

func (sec *ArraySection) AppendBool(vals ...bool) {
	for _, val := range vals {
		sec.data = append(sec.data, NewDocument(NewBoolSection("", val)))
	}
}

func (sec *ArraySection) CanBool(idx int) bool {
	if idx >= len(sec.data) {
		return false
	}

	return sec.data[idx].IsBool()
}

func (sec *ArraySection) AtBool(idx int) (bool, error) {
	if idx >= len(sec.data) {
		return false, ErrorArrayOutOfRange{ArrayKey: sec.name, MaxMember: len(sec.data), InvalidIndex: idx}
	}

	bSec, err := sec.data[idx].Bool()
	if err != nil {
		return false, err
	}

	return bSec.Bool(), nil
}

func (sec *ArraySection) AppendArray(vals ...*ArraySection) {
	for _, val := range vals {
		sec.data = append(sec.data, NewDocument(val))
	}
}

func (sec *ArraySection) CanArray(idx int) bool {
	if idx >= len(sec.data) {
		return false
	}

	return sec.data[idx].IsArray()
}

func (sec *ArraySection) AtArray(idx int) (*ArraySection, error) {
	if idx >= len(sec.data) {
		return nil, ErrorArrayOutOfRange{ArrayKey: sec.name, MaxMember: len(sec.data), InvalidIndex: idx}
	}

	arrSec, err := sec.data[idx].Array()
	if err != nil {
		return nil, err
	}

	return arrSec, nil
}

func (sec *ArraySection) AppendNil() {
	sec.data = append(sec.data, NewDocument(NewNIlSection("")))
}

func (sec *ArraySection) CanNil(idx int) bool {
	if idx >= len(sec.data) {
		return false
	}

	return sec.data[idx].IsNil()
}

func (sec *ArraySection) AppendInt(vals ...int) {
	for _, val := range vals {
		sec.data = append(sec.data, NewDocument(NewIntSection("", val)))
	}
}

func (sec *ArraySection) CanInt(idx int) bool {
	if idx >= len(sec.data) {
		return false
	}

	return sec.data[idx].IsInt()
}

func (sec *ArraySection) AtInt(idx int) (int, error) {
	if idx >= len(sec.data) {
		return 0, ErrorArrayOutOfRange{ArrayKey: sec.name, MaxMember: len(sec.data), InvalidIndex: idx}
	}

	numSec, err := sec.data[idx].Number()
	if err != nil {
		return 0, err
	}

	return numSec.Int(), nil
}

func (sec *ArraySection) AppendInt8(vals ...int8) {
	for _, val := range vals {
		sec.data = append(sec.data, NewDocument(NewInt8Section("", val)))
	}
}

func (sec *ArraySection) CanInt8(idx int) bool {
	if idx >= len(sec.data) {
		return false
	}

	return sec.data[idx].IsInt8()
}

func (sec *ArraySection) AtInt8(idx int) (int8, error) {
	if idx >= len(sec.data) {
		return 0, ErrorArrayOutOfRange{ArrayKey: sec.name, MaxMember: len(sec.data), InvalidIndex: idx}
	}

	numSec, err := sec.data[idx].Number()
	if err != nil {
		return 0, err
	}

	return numSec.Int8(), nil
}

func (sec *ArraySection) AppendInt16(vals ...int16) {
	for _, val := range vals {
		sec.data = append(sec.data, NewDocument(NewInt16Section("", val)))
	}
}

func (sec *ArraySection) CanInt16(idx int) bool {
	if idx >= len(sec.data) {
		return false
	}

	return sec.data[idx].IsInt16()
}

func (sec *ArraySection) AtInt16(idx int) (int16, error) {
	if idx >= len(sec.data) {
		return 0, ErrorArrayOutOfRange{ArrayKey: sec.name, MaxMember: len(sec.data), InvalidIndex: idx}
	}

	numSec, err := sec.data[idx].Number()
	if err != nil {
		return 0, err
	}

	return numSec.Int16(), nil
}

func (sec *ArraySection) AppendInt32(vals ...int32) {
	for _, val := range vals {
		sec.data = append(sec.data, NewDocument(NewInt32Section("", val)))
	}
}

func (sec *ArraySection) CanInt32(idx int) bool {
	if idx >= len(sec.data) {
		return false
	}

	return sec.data[idx].IsInt32()
}

func (sec *ArraySection) AtInt32(idx int) (int32, error) {
	if idx >= len(sec.data) {
		return 0, ErrorArrayOutOfRange{ArrayKey: sec.name, MaxMember: len(sec.data), InvalidIndex: idx}
	}

	numSec, err := sec.data[idx].Number()
	if err != nil {
		return 0, err
	}

	return numSec.Int32(), nil
}

func (sec *ArraySection) AppendInt64(vals ...int64) {
	for _, val := range vals {
		sec.data = append(sec.data, NewDocument(NewInt64Section("", val)))
	}
}

func (sec *ArraySection) CanInt64(idx int) bool {
	if idx >= len(sec.data) {
		return false
	}

	return sec.data[idx].IsInt64()
}

func (sec *ArraySection) AtInt64(idx int) (int64, error) {
	if idx >= len(sec.data) {
		return 0, ErrorArrayOutOfRange{ArrayKey: sec.name, MaxMember: len(sec.data), InvalidIndex: idx}
	}

	numSec, err := sec.data[idx].Number()
	if err != nil {
		return 0, err
	}

	return numSec.Int64(), nil
}

func (sec *ArraySection) AppendFloat32(vals ...float32) {
	for _, val := range vals {
		sec.data = append(sec.data, NewDocument(NewFloat32Section("", val)))
	}
}

func (sec *ArraySection) CanFloat32(idx int) bool {
	if idx >= len(sec.data) {
		return false
	}

	return sec.data[idx].IsFloat32()
}

func (sec *ArraySection) AtFloat32(idx int) (float32, error) {
	if idx >= len(sec.data) {
		return 0, ErrorArrayOutOfRange{ArrayKey: sec.name, MaxMember: len(sec.data), InvalidIndex: idx}
	}

	numSec, err := sec.data[idx].Number()
	if err != nil {
		return 0, err
	}

	return numSec.Float32(), nil
}

func (sec *ArraySection) AppendFloat64(vals ...float64) {
	for _, val := range vals {
		sec.data = append(sec.data, NewDocument(NewFloat64Section("", val)))
	}
}

func (sec *ArraySection) CanFloat64(idx int) bool {
	if idx >= len(sec.data) {
		return false
	}

	return sec.data[idx].IsFloat64()
}

func (sec *ArraySection) AtFloat64(idx int) (float64, error) {
	if idx >= len(sec.data) {
		return 0, ErrorArrayOutOfRange{ArrayKey: sec.name, MaxMember: len(sec.data), InvalidIndex: idx}
	}

	numSec, err := sec.data[idx].Number()
	if err != nil {
		return 0, err
	}

	return numSec.Float64(), nil
}

func (sec *ArraySection) Unmarshal(data interface{}) error {
	dataVal := reflect.ValueOf(data)
	if dataVal.IsNil() || dataVal.Kind() != reflect.Ptr {
		return ErrorUnmarshalFail{Type: Array, Value: dataVal}
	}

	rv := dataVal.Elem()

	return sec.setValue(&rv)
}

func (sec *ArraySection) setValue(rv *reflect.Value) error {
	if rv == nil {
		return ErrorAcceptNilParam{FunctionName: "ArraySection.setValue", NilParam: rv}
	}

	if !rv.CanSet() {
		return ErrorSetValueFail{Type: Array, KeyName: sec.name, Value: *rv}
	}

	switch rv.Kind() {
	case reflect.Slice:
		return sec.setSliceValue(rv)
	case reflect.Array:
		return sec.setArrayValue(rv)
	default:
		return ErrorSetValueFail{Type: Array, KeyName: sec.name, Value: *rv}
	}

	return nil
}

func (sec *ArraySection) setSliceValue(rv *reflect.Value) error {
	if rv == nil {
		return ErrorAcceptNilParam{FunctionName: "ArraySection.setSliceValue", NilParam: rv}
	}

	if !rv.CanSet() || rv.Kind() != reflect.Slice {
		return ErrorSetValueFail{Type: Array, KeyName: sec.name, Value: *rv}
	}

	newSlicev := reflect.MakeSlice(rv.Type(), sec.Len(), sec.Len())
	for idx, doc := range sec.data {
		idxv := newSlicev.Index(idx)
		if err := sec.set(doc, &idxv); err != nil {
			return err
		}
	}

	rv.Set(newSlicev)

	return nil
}

func (sec *ArraySection) setArrayValue(rv *reflect.Value) error {
	if rv == nil {
		return ErrorAcceptNilParam{FunctionName: "ArraySection.setArrayValue", NilParam: rv}
	}

	if !rv.CanSet() || rv.Kind() != reflect.Array {
		return ErrorSetValueFail{Type: Array, KeyName: sec.name, Value: *rv}
	}

	for idx := 0; idx < rv.Len(); idx++ {
		doc, err := sec.At(idx)
		if err != nil {
			return err
		}

		idxv := rv.Index(idx)
		if err = sec.set(doc, &idxv); err != nil {
			return err
		}
	}

	return nil
}

func (sec *ArraySection) set(doc *Document, rv *reflect.Value) error {
	switch {
	case doc == nil:
		return ErrorAcceptNilParam{FunctionName: "ArraySection.set", NilParam: doc}
	case rv == nil:
		return ErrorAcceptNilParam{FunctionName: "ArraySection.set", NilParam: rv}
	}

	switch {
	case doc.IsString():
		strSec, err := doc.String()
		if err != nil {
			return err
		}

		if err = strSec.setValue(rv); err != nil {
			return err
		}

	case doc.IsBool():
		bSec, err := doc.Bool()
		if err != nil {
			return err
		}

		if err = bSec.setValue(rv); err != nil {
			return err
		}

	case doc.IsNil():
		nilSec, err := doc.Nil()
		if err != nil {
			return err
		}

		if err = nilSec.setValue(rv); err != nil {
			return err
		}

	case doc.IsNumber():
		numSec, err := doc.Number()
		if err != nil {
			return err
		}

		if err = numSec.setValue(rv); err != nil {
			return err
		}

	case doc.IsObject():
		objSec, err := doc.Object()
		if err != nil {
			return err
		}

		if err = objSec.setValue(rv); err != nil {
			return err
		}

	case doc.IsArray():
		arrSec, err := doc.Array()
		if err != nil {
			return err
		}

		if err = arrSec.setArrayValue(rv); err != nil {
			return err
		}

	default:
		return ErrorSetValueFail{Type: Array, KeyName: sec.name, Value: *rv}
	}

	return nil
}

func (sec *ArraySection) Marshal(data interface{}) error {
	dataVal := reflect.ValueOf(data)
	if dataVal.IsNil() {
		return ErrorAcceptNilParam{FunctionName: "ArraySection.Marshal", NilParam: data}
	}

	return sec.getValue(dataVal)
}

func (sec *ArraySection) getValue(rv reflect.Value) error {
	switch rv.Kind() {
	case reflect.Ptr:
		return sec.getValue(rv.Elem())

	case reflect.Slice:
		sec.getSliceValue(rv)
	case reflect.Array:
		sec.getArrayValue(rv)
	default:
		return ErrorMarshalFail{Type: Object, InvalidKind: rv.Kind(), ValidKind: reflect.Map}
	}

	return nil
}

func (sec *ArraySection) getSliceValue(rv reflect.Value) error {
	if rv.Kind() != reflect.Slice {
		return ErrorGetValueFail{Type: Array, InvalidKind: rv.Kind(), ValidKind: reflect.Slice}
	}

	for idx := 0; idx < rv.Len(); idx++ {
		idxv := rv.Index(idx)
		if err := sec.get(idxv); err != nil {
			return err
		}
	}

	return nil
}

func (sec *ArraySection) getArrayValue(rv reflect.Value) error {
	if rv.Kind() != reflect.Array {
		return ErrorGetValueFail{Type: Array, InvalidKind: rv.Kind(), ValidKind: reflect.Array}
	}

	for idx := 0; idx < rv.Len(); idx++ {
		idxv := rv.Index(idx)
		if err := sec.get(idxv); err != nil {
			return err
		}
	}

	return nil
}

func (sec *ArraySection) get(rv reflect.Value) error {
	switch rv.Kind() {
	case reflect.String:
		sec.AppendString(rv.String())

	case reflect.Interface:
		return sec.get(reflect.ValueOf(rv.Interface()))

	case reflect.Map:
		child := NewObjectSection("")
		if err := child.getMapValue(rv); err != nil {
			return err
		}

		sec.AppendObject(child)

	case reflect.Struct:
		child := NewObjectSection("")
		if err := child.getStructValue(rv); err != nil {
			return err
		}

		sec.AppendObject(child)

	case reflect.Int:
		sec.AppendInt(int(rv.Int()))
	case reflect.Int8:
		sec.AppendInt8(int8(rv.Int()))
	case reflect.Int16:
		sec.AppendInt16(int16(rv.Int()))
	case reflect.Int32:
		sec.AppendInt32(int32(rv.Int()))
	case reflect.Int64:
		sec.AppendInt64(rv.Int())
	case reflect.Float32:
		sec.AppendFloat32(float32(rv.Float()))
	case reflect.Float64:
		sec.AppendFloat64(rv.Float())

	case reflect.Bool:
		sec.AppendBool(rv.Bool())

	case reflect.Array:
		arr := NewArraySection("")
		if err := arr.getValue(rv); err != nil {
			return err
		}
		sec.AppendArray(arr)

	case reflect.Invalid:
		sec.AppendNil()

	default:
		return ErrorGetValueFail{Type: Object, InvalidKind: rv.Kind(), ValidKind: reflect.Interface}
	}

	return nil
}
