package json

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/zero-Chan/document"
)

type JsonTransCoder struct {
}

func CreateJsonTransCoder() JsonTransCoder {
	coder := JsonTransCoder{}
	return coder
}

func NewJsonTransCoder() *JsonTransCoder {
	coder := CreateJsonTransCoder()
	return &coder
}

// Unmarshal: make []byte -> doc
// []byte -> interface{} 		[ json.unmarshal ]
// interface{} -> doc	[ doc.marshal ]

//    To unmarshal JSON into an interface value, Unmarshal stores one of these in the interface value:
//       bool, for JSON booleans
//       float64, for JSON numbers
//       string, for JSON strings
//       []interface{}, for JSON arrays
//       map[string]interface{}, for JSON objects
//       nil for JSON null
func (coder *JsonTransCoder) Unmarshal(data []byte, v interface{}) (*document.Document, error) {
	err := json.Unmarshal(data, v)
	if err != nil {
		return nil, err
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return nil, fmt.Errorf("non-pointer(%s)", rv.Type().String())
	}

	vv := rv.Elem()
	switch vv.Kind() {
	case reflect.String:
		strSec := document.NewStringSection("", vv.String())
		err := strSec.Marshal(vv.String())
		if err != nil {
			return nil, err
		}
		return document.NewDocument(strSec), nil

	case reflect.Map:
		objSec := document.NewObjectSection("")
		err := objSec.Marshal(vv.Interface())
		if err != nil {
			return nil, err
		}

		return document.NewDocument(objSec), nil

	case reflect.Struct:
		objSec := document.NewObjectSection("")
		err := objSec.Marshal(vv.Interface())
		if err != nil {
			return nil, err
		}

		return document.NewDocument(objSec), nil
	}

	return nil, nil
}

// Marshal: make doc to json format string
// doc -> interface{}	[ doc.unmarshal ]
// interface{} -> []byte	 [ json.marshal ]
func (coder *JsonTransCoder) Marshal(doc document.Document) ([]byte, error) {
	switch doc.Type() {
	case document.String:
		strSec, err := doc.String()
		if err != nil {
			return nil, err
		}

		var v string
		strSec.Unmarshal(&v)
		return json.Marshal(v)

	case document.Object:
		objSec, err := doc.Object()
		m := make(map[string]interface{})
		err = objSec.Unmarshal(&m)
		if err != nil {
			return nil, err
		}

		return json.Marshal(m)

	case document.Number:
		// TODO
	}

	return nil, nil
}
