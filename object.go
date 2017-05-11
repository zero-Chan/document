package document

import (
	"fmt"
	"reflect"
)

type ObjectSection struct {
	name string
	data map[string]*Document
}

func CreateObjectSection(name string) ObjectSection {
	s := ObjectSection{
		name: name,
		data: make(map[string]*Document),
	}

	return s
}

func NewObjectSection(name string) *ObjectSection {
	s := CreateObjectSection(name)
	return &s
}

func (o ObjectSection) Type() Type {
	return Object
}

func (o *ObjectSection) Clear() {
	o.data = make(map[string]*Document)
}

func (sec *ObjectSection) Keys() []string {
	keys := make([]string, 0)
	for k, _ := range sec.data {
		keys = append(keys, k)
	}
	return keys
}

func (sec *ObjectSection) At(key string) (*Document, error) {
	v, exist := sec.data[key]
	if !exist {
		return nil, ErrorObjectNotExistKey{ObjectKey: sec.name, Key: key}
	}

	if v == nil {
		return nil, ErrorAcceptNilParam{FunctionName: "ObjectSection.At", NilParam: v}
	}

	return v, nil
}

func (sec *ObjectSection) Set(key string, doc *Document) error {
	if doc == nil {
		return ErrorAcceptNilParam{FunctionName: "ObjectSection.Set", NilParam: doc}
	}
	sec.data[key] = doc
	return nil
}

func (o *ObjectSection) Object(key string) (*ObjectSection, error) {
	v, exist := o.data[key]
	if !exist {
		return nil, ErrorObjectNotExistKey{ObjectKey: o.name, Key: key}
	}

	if v == nil {
		return nil, ErrorAcceptNilParam{FunctionName: "ObjectSection.Object", NilParam: v}
	}

	objSec, err := v.Object()
	if err != nil {
		return nil, err
	}

	return objSec, nil
}

func (o *ObjectSection) SetObject(key string, val *ObjectSection) {
	o.data[key] = NewDocument(val)
}

func (o *ObjectSection) Int(key string) (int, error) {
	v, exist := o.data[key]
	if !exist {
		return 0, ErrorObjectNotExistKey{ObjectKey: o.name, Key: key}
	}

	if v == nil {
		return 0, ErrorAcceptNilParam{FunctionName: "ObjectSection.Int", NilParam: v}
	}

	numSec, err := v.Number()
	if err != nil {
		return 0, err
	}

	return numSec.Int(), nil
}

func (o *ObjectSection) SetInt(key string, val int) {
	o.data[key] = NewDocument(NewIntSection(key, val))
}

func (o *ObjectSection) Int8(key string) (int8, error) {
	v, exist := o.data[key]
	if !exist {
		return 0, ErrorObjectNotExistKey{ObjectKey: o.name, Key: key}
	}

	if v == nil {
		return 0, ErrorAcceptNilParam{FunctionName: "ObjectSection.Int8", NilParam: v}
	}

	numSec, err := v.Number()
	if err != nil {
		return 0, err
	}

	return numSec.Int8(), nil
}

func (o *ObjectSection) SetInt8(key string, val int8) {
	o.data[key] = NewDocument(NewInt8Section(key, val))
}

func (o *ObjectSection) Int16(key string) (int16, error) {
	v, exist := o.data[key]
	if !exist {
		return 0, ErrorObjectNotExistKey{ObjectKey: o.name, Key: key}
	}

	if v == nil {
		return 0, ErrorAcceptNilParam{FunctionName: "ObjectSection.Int16", NilParam: v}
	}

	numSec, err := v.Number()
	if err != nil {
		return 0, err
	}

	return numSec.Int16(), nil
}

func (o *ObjectSection) SetInt16(key string, val int16) {
	o.data[key] = NewDocument(NewInt16Section(key, val))
}

func (o *ObjectSection) Int32(key string) (int32, error) {
	v, exist := o.data[key]
	if !exist {
		return 0, ErrorObjectNotExistKey{ObjectKey: o.name, Key: key}
	}

	if v == nil {
		return 0, ErrorAcceptNilParam{FunctionName: "ObjectSection.Int32", NilParam: v}
	}

	numSec, err := v.Number()
	if err != nil {
		return 0, err
	}

	return numSec.Int32(), nil
}

func (o *ObjectSection) SetInt32(key string, val int32) {
	o.data[key] = NewDocument(NewInt32Section(key, val))
}

func (o *ObjectSection) Int64(key string) (int64, error) {
	v, exist := o.data[key]
	if !exist {
		return 0, ErrorObjectNotExistKey{ObjectKey: o.name, Key: key}
	}

	if v == nil {
		return 0, ErrorAcceptNilParam{FunctionName: "ObjectSection.Int64", NilParam: v}
	}

	numSec, err := v.Number()
	if err != nil {
		return 0, err
	}

	return numSec.Int64(), nil
}

func (o *ObjectSection) SetInt64(key string, val int64) {
	o.data[key] = NewDocument(NewInt64Section(key, val))
}

func (sec *ObjectSection) Float32(key string) (float32, error) {
	v, exist := sec.data[key]
	if !exist {
		return 0, ErrorObjectNotExistKey{ObjectKey: sec.name, Key: key}
	}

	if v == nil {
		return 0, ErrorAcceptNilParam{FunctionName: "ObjectSection.Float32", NilParam: v}
	}

	numSec, err := v.Number()
	if err != nil {
		return 0, err
	}

	return numSec.Float32(), nil
}

func (sec *ObjectSection) SetFloat32(key string, val float32) {
	sec.data[key] = NewDocument(NewFloat32Section(key, val))
}

func (sec *ObjectSection) Float64(key string) (float64, error) {
	v, exist := sec.data[key]
	if !exist {
		return 0, ErrorObjectNotExistKey{ObjectKey: sec.name, Key: key}
	}

	if v == nil {
		return 0, ErrorAcceptNilParam{FunctionName: "ObjectSection.Float64", NilParam: v}
	}

	numSec, err := v.Number()
	if err != nil {
		return 0, err
	}

	return numSec.Float64(), nil
}

func (sec *ObjectSection) SetFloat64(key string, val float64) {
	sec.data[key] = NewDocument(NewFloat64Section(key, val))
}

func (sec *ObjectSection) String(key string) (string, error) {
	v, exist := sec.data[key]
	if !exist {
		return "", ErrorObjectNotExistKey{ObjectKey: sec.name, Key: key}
	}

	if v == nil {
		return "", ErrorAcceptNilParam{FunctionName: "ObjectSection.String", NilParam: v}
	}

	strSec, err := v.String()
	if err != nil {
		return "", err
	}

	return strSec.data, nil
}

func (sec *ObjectSection) SetString(key, val string) {
	sec.data[key] = NewDocument(NewStringSection(key, val))
}

func (sec *ObjectSection) Array(key string) (*ArraySection, error) {
	v, exist := sec.data[key]
	if !exist {
		return nil, ErrorObjectNotExistKey{ObjectKey: sec.name, Key: key}
	}

	if v == nil {
		return nil, ErrorAcceptNilParam{FunctionName: "ObjectSection.Array", NilParam: v}
	}

	arrSec, err := v.Array()
	if err != nil {
		return nil, err
	}

	return arrSec, nil
}

func (sec *ObjectSection) SetArray(key string, val *ArraySection) {
	sec.data[key] = NewDocument(NewArraySection(key))
}

func (sec *ObjectSection) Bool(key string) (bool, error) {
	v, exist := sec.data[key]
	if !exist {
		return false, ErrorObjectNotExistKey{ObjectKey: sec.name, Key: key}
	}

	if v == nil {
		return false, ErrorAcceptNilParam{FunctionName: "ObjectSection.Bool", NilParam: v}
	}

	bSec, err := v.Bool()
	if err != nil {
		return false, err
	}

	return bSec.data, nil
}

func (sec *ObjectSection) SetBool(key string, val bool) {
	sec.data[key] = NewDocument(NewBoolSection(key, val))
}

func (sec *ObjectSection) Unmarshal(data interface{}) error {
	dataVal := reflect.ValueOf(data)
	if dataVal.IsNil() || dataVal.Kind() != reflect.Ptr {
		return ErrorUnmarshalFail{Type: Object, Value: dataVal}
	}

	rv := dataVal.Elem()

	switch rv.Kind() {
	case reflect.Struct:
		return sec.setStructValue(&rv)
	case reflect.Map:
		return sec.setMapValue(&rv)
	default:
		return ErrorSetValueFail{Type: Object, KeyName: sec.name, Value: rv}
	}
}

func (sec *ObjectSection) setStructValue(rv *reflect.Value) error {
	if rv == nil {
		return ErrorAcceptNilParam{FunctionName: "ObjectSection.setStructValue", NilParam: rv}
	}

	if !rv.CanSet() {
		return ErrorSetValueFail{Type: Object, KeyName: sec.name, Value: *rv}
	}

	rt := rv.Type()
	n := rt.NumField()

	for idx := 0; idx < n; idx++ {
		vf := rv.Field(idx)
		tf := rt.Field(idx)
		tag := tf.Tag.Get("document")
		// if tag empty or "-" ignore
		if tag == "-" || tag == "omitempty" {
			continue
		}

		if len(tag) == 0 {
			tag = tf.Name
		}

		doc, exist := sec.data[tag]
		if !exist {
			continue
		}

		switch doc.Type() {
		case String:
			strSec, err := doc.String()
			if err != nil {
				return err
			}

			if err = strSec.setValue(&vf); err != nil {
				return err
			}

		case Int:
			numSec, err := doc.Number()
			if err != nil {
				return err
			}

			if err = numSec.setValue(&vf); err != nil {
				return err
			}

			// TODO

		}
	}

	return nil
}

func (sec *ObjectSection) setMapValue(rv *reflect.Value) error {
	if rv == nil {
		return ErrorAcceptNilParam{FunctionName: "ObjectSection.setStructValue", NilParam: rv}
	}

	if !rv.CanSet() {
		return ErrorSetValueFail{Type: Object, KeyName: sec.name, Value: *rv}
	}

	for key, doc := range sec.data {
		switch doc.Type() {
		case Object:
			child, err := sec.Object(key)
			if err != nil {
				return err
			}

			m := make(map[string]interface{})
			if err = child.Unmarshal(&m); err != nil {
				return err
			}
			rv.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(m))

		case Int:
			d, err := sec.Int(key)
			if err != nil {
				return err
			}
			rv.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(d))

		case Int8:
			d, err := sec.Int8(key)
			if err != nil {
				return err
			}
			rv.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(d))

		case Int16:
			d, err := sec.Int16(key)
			if err != nil {
				return err
			}
			rv.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(d))

		case Int32:
			d, err := sec.Int32(key)
			if err != nil {
				return err
			}
			rv.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(d))

		case Int64:
			d, err := sec.Int64(key)
			if err != nil {
				return err
			}
			rv.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(d))

		case Float32:
			d, err := sec.Float32(key)
			if err != nil {
				return err
			}
			rv.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(d))

		case Float64:
			d, err := sec.Float64(key)
			if err != nil {
				return err
			}
			rv.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(d))

		case String:
			d, err := sec.String(key)
			if err != nil {
				return err
			}
			rv.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(d))

		case Bool:
			d, err := sec.Bool(key)
			if err != nil {
				return err
			}
			rv.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(d))

		case Array:
			d, err := sec.Array(key)
			if err != nil {
				return err
			}
			arr := make([]interface{}, 0)
			if err = d.Unmarshal(&arr); err != nil {
				return err
			}
			rv.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(arr))

			// TODO : case
		}
	}

	return nil
}

func (sec *ObjectSection) set(key string, doc *Document, rv *reflect.Value) error {
	//	switch doc.Type() {
	//	case String:
	//		d, err := sec.String()
	//		if err != nil {
	//			return err
	//		}

	//	}

	return nil
}

func (sec *ObjectSection) Marshal(data interface{}) error {
	dataVal := reflect.ValueOf(data)
	if dataVal.IsNil() {
		return ErrorAcceptNilParam{FunctionName: "ObjectSection.Marshal", NilParam: data}
	}

	return sec.getValue(dataVal)
}

func (sec *ObjectSection) getValue(rv reflect.Value) error {
	switch rv.Kind() {
	case reflect.Ptr:
		return sec.getValue(rv.Elem())

	case reflect.Map:
		sec.getMapValue(rv)
	case reflect.Struct:
		sec.getStructValue(rv)
	default:
		return ErrorMarshalFail{Type: Object, InvalidKind: rv.Kind(), ValidKind: reflect.Map}
	}

	return nil
}

func (sec *ObjectSection) getStructValue(rv reflect.Value) error {
	return nil
}

func (sec *ObjectSection) getMapValue(rv reflect.Value) error {
	if rv.Kind() != reflect.Map {
		return ErrorGetValueFail{Type: Object, InvalidKind: rv.Kind(), ValidKind: reflect.Map}
	}

	for _, keyv := range rv.MapKeys() {
		if keyv.Kind() != reflect.String {
			return fmt.Errorf("Please marshal from map[string]XXX format. invalidKeyType: %s", keyv.Kind().String())
		}

		valv := rv.MapIndex(keyv)
		err := sec.get(keyv.String(), valv)
		if err != nil {
			return err
		}
	}

	return nil
}

func (sec *ObjectSection) get(key string, rv reflect.Value) error {
	switch rv.Kind() {
	case reflect.String:
		sec.SetString(key, rv.String())

	case reflect.Interface:
		return sec.get(key, reflect.ValueOf(rv.Interface()))

	case reflect.Map:
		child := NewObjectSection(key)
		if err := child.getMapValue(rv); err != nil {
			return err
		}
		sec.SetObject(key, child)

	case reflect.Struct:
		child := NewObjectSection(key)
		if err := child.getStructValue(rv); err != nil {
			return err
		}
		sec.SetObject(key, child)

	case reflect.Int:
		sec.SetInt(key, int(rv.Int()))
	case reflect.Int8:
		sec.SetInt8(key, int8(rv.Int()))
	case reflect.Int16:
		sec.SetInt16(key, int16(rv.Int()))
	case reflect.Int32:
		sec.SetInt32(key, int32(rv.Int()))
	case reflect.Int64:
		sec.SetInt64(key, rv.Int())
	case reflect.Float32:
		sec.SetFloat32(key, float32(rv.Float()))
	case reflect.Float64:
		sec.SetFloat64(key, rv.Float())

	case reflect.Bool:
		sec.SetBool(key, rv.Bool())

	case reflect.Array:
		//		arr := NewArraySection(key)
		// arr.setValue
	}

	return nil
}
