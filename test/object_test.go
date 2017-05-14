package test

import (
	"testing"

	"github.com/zero-Chan/godoc"
)

func TestObjectSectionUnmarshal(t *testing.T) {
	// obj
	sec := godoc.NewObjectSection("")
	sec.SetBool("BoolKey", true)
	sec.SetString("StringKey", "Val_1")
	sec.SetInt("IntKey", 10)
	sec.SetFloat32("F32Key", 123.123)
	sec.SetNil("NilKey")

	arr := godoc.NewArraySection("")
	arr.AppendInt(20)
	arr.AppendString("Val_1-1")
	sec.SetArray("ArrayKey", arr)

	obj := godoc.NewObjectSection("")
	obj.SetString("ChildString", "Val_1_1")
	sec.SetObject("ObjectKey", obj)

	buf := make(map[string]interface{})
	err := sec.Unmarshal(&buf)
	if err != nil {
		t.Errorf("sec.Unmarshal(%T) fail: %s", &buf, err)
		t.FailNow()
	}

	// bool
	boolv, exist := buf["BoolKey"]
	if !exist {
		t.Errorf("BoolKey not exist")
		t.FailNow()
	}

	bv, ok := boolv.(bool)
	if !ok {
		t.Errorf("Bool type instancing fail. invalidType[%T]. Type[bool]", boolv)
		t.FailNow()
	} else if bv != true {
		t.Errorf("Bool type Get Invalid value. invalidValue[%t]. value[%t]", boolv, true)
		t.FailNow()
	}

	// string
	stringv, exist := buf["StringKey"]
	if !exist {
		t.Errorf("StringKey not exist")
		t.FailNow()
	}

	sv, ok := stringv.(string)
	if !ok {
		t.Errorf("String type instancing fail. invalidType[%T]. Type[string]", stringv)
		t.FailNow()
	} else if sv != "Val_1" {
		t.Errorf("String type Get Invalid value. invalidValue[%s]. value[%s]", stringv, "Val_1")
		t.FailNow()
	}

	// int
	intv, exist := buf["IntKey"]
	if !exist {
		t.Errorf("IntKey not exist")
		t.FailNow()
	}

	iv, ok := intv.(int)
	if !ok {
		t.Errorf("Int type instancing fail. invalidType[%T]. Type[int]", intv)
		t.FailNow()
	} else if iv != 10 {
		t.Errorf("Int type Get Invalid value. invalidValue[%d]. value[%d]", iv, 10)
		t.FailNow()
	}

	// float
	floatv, exist := buf["F32Key"]
	if !exist {
		t.Errorf("F32Key not exist")
		t.FailNow()
	}

	fv, ok := floatv.(float32)
	if !ok {
		t.Errorf("Float type instancing fail. invalidType[%T]. Type[float32]", floatv)
		t.FailNow()
	} else if fv != 123.123 {
		t.Errorf("Float type Get Invalid value. invalidValue[%f]. value[%f]", fv, 123.123)
		t.FailNow()
	}

	nilv, exist := buf["NilKey"]
	if !exist {
		t.Errorf("NilKey not exist")
		t.FailNow()
	}

	nv, ok := nilv.(*int)
	if !ok {
		t.Errorf("*Int type instancing fail. invalidType[%T]. Type[*int]", nilv)
		t.FailNow()
	}

	if nv != nil {
		t.Error("NIl type Get Invalid value. invalidValue[", nilv, "], value[nil]")
		t.FailNow()
	}

	// arrray
	arrayv, exist := buf["ArrayKey"]
	if !exist {
		t.Errorf("ArrayKey not exist")
		t.FailNow()
	}

	arrv, ok := arrayv.([]interface{})
	if !ok {
		t.Errorf("Array type instancing fail. invalidType[%T]. Type[[]interface{}]", arrayv)
		t.FailNow()
	}

	if len(arrv) != 2 {
		t.Errorf("Array type Get Invalid length. invalidLen[%d]. len[%d]", len(arrv), 2)
		t.FailNow()
	}

	arr0v := arrv[0]
	arriv, ok := arr0v.(int)
	if !ok {
		t.Errorf("Array int type instancing fail. invalidType[%T]. Type[int]", arr0v)
		t.FailNow()
	} else if arriv != 20 {
		t.Errorf("Int type Get Invalid value. invalidValue[%d]. value[%d]", arriv, 20)
		t.FailNow()
	}

	arr1v := arrv[1]
	arrsv, ok := arr1v.(string)
	if !ok {
		t.Errorf("Array int type instancing fail. invalidType[%T]. Type[int]", arr1v)
		t.FailNow()
	} else if arrsv != "Val_1-1" {
		t.Errorf("Int type Get Invalid value. invalidValue[%s]. value[%s]", arrsv, "Val_1-1")
		t.FailNow()
	}

	// object
	objectv, exist := buf["ObjectKey"]
	if !exist {
		t.Errorf("ObjectKey not exist")
		t.FailNow()
	}

	objv, ok := objectv.(map[string]interface{})
	if !ok {
		t.Errorf("Object type instancing fail. invalidType[%T]. Type[map[string]interface{}]", objv)
		t.FailNow()
	}

	chlidstrv, exist := objv["ChildString"]
	if !exist {
		t.Errorf("ChildString not exist")
		t.FailNow()
	}

	csv, ok := chlidstrv.(string)
	if !ok {
		t.Errorf("Array int type instancing fail. invalidType[%T]. Type[string]", csv)
		t.FailNow()
	} else if csv != "Val_1_1" {
		t.Errorf("Int type Get Invalid value. invalidValue[%s]. value[%s]", csv, "Val_1_1")
		t.FailNow()
	}
}

func TestObjectSectionMarshal(t *testing.T) {
	buf := make(map[string]interface{})

	buf["StringKey"] = "Val_1"
	buf["IntKey"] = 10
	buf["BoolKey"] = true
	buf["NilKey"] = nil

	arrbuf := make([]interface{}, 2)
	arrbuf[0] = "Val_1-1"
	arrbuf[1] = 20
	buf["ArrayKey"] = arrbuf

	childbuf := make(map[string]interface{})
	childbuf["ChildString"] = "Val_1_1"
	buf["ObjectKey"] = childbuf

	sec := godoc.NewObjectSection("")
	err := sec.Marshal(buf)
	if err != nil {
		t.Errorf("sec.Marshal(%T) fail: %s", buf, err)
		t.FailNow()
	}

	// string
	if v, err := sec.String("StringKey"); err != nil {
		t.Errorf("Get StringKey error: %s", err)
		t.FailNow()
	} else if v != "Val_1" {
		t.Errorf("StringKey Get Invalid value. invalidValue[%s]. value[%s]", v, "Val_1")
		t.FailNow()
	}

	// int
	if v, err := sec.Int("IntKey"); err != nil {
		t.Errorf("Get IntKey error: %s", err)
		t.FailNow()
	} else if v != 10 {
		t.Errorf("IntKey Get Invalid value. invalidValue[%d]. value[%d]", v, 10)
		t.FailNow()
	}

	// bool
	if v, err := sec.Bool("BoolKey"); err != nil {
		t.Errorf("Get BoolKey error: %s", err)
		t.FailNow()
	} else if v != true {
		t.Errorf("BoolKey Get Invalid value. invalidValue[%t]. value[%t]", v, true)
		t.FailNow()
	}

	// array
	arrv, err := sec.Array("ArrayKey")
	if err != nil {
		t.Errorf("Get ArrayKey error: %s", err)
		t.FailNow()
	}

	if arrv.Len() != 2 {
		t.Errorf("Invalid arrv len. invalidLen[%d], validLen[%d]", arrv.Len(), 2)
		t.FailNow()
	}

	if v, err := arrv.AtString(0); err != nil {
		t.Errorf("Get Array[0] StringVal fail: %s", err)
		t.FailNow()
	} else if v != "Val_1-1" {
		t.Errorf("StringVal Get Invalid value. invalidValue[%s]. value[%s]", v, "Val_1-1")
		t.FailNow()
	}

	if v, err := arrv.AtInt(1); err != nil {
		t.Errorf("Get Array[0] IntVal fail: %s", err)
		t.FailNow()
	} else if v != 20 {
		t.Errorf("IntVal Get Invalid value. invalidValue[%d]. value[%d]", v, 20)
		t.FailNow()
	}

	// object
	objv, err := sec.Object("ObjectKey")
	if err != nil {
		t.Errorf("Get ObjectKey error: %s", err)
		t.FailNow()
	}

	if v, err := objv.String("ChildString"); err != nil {
		t.Errorf("Get ChildString fail: %s", err)
		t.FailNow()
	} else if v != "Val_1_1" {
		t.Errorf("ChildString Get Invalid value. invalidValue[%s]. value[%s]", v, "Val_1_1")
		t.FailNow()
	}
}
