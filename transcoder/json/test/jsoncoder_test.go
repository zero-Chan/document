package test

import (
	"testing"

	"github.com/zero-Chan/godoc"
	"github.com/zero-Chan/godoc/transcoder/json"
)

func TestJsonCoderMarshal(t *testing.T) {
	coder := transcoder.CreateJsonTransCoder()

	objSec := godoc.NewObjectSection("")
	objbuf, err := coder.Marshal(godoc.CreateDocument(objSec))
	if err != nil {
		t.Errorf("JsonCoder marshal Object fail: %s", err)
		t.FailNow()
	}

	if string(objbuf) != "{}" {
		t.Errorf("Invalid Json marshal object")
		t.FailNow()
	}

	arrSec := godoc.NewArraySection("")
	arrbuf, err := coder.Marshal(godoc.CreateDocument(arrSec))
	if err != nil {
		t.Errorf("JsonCoder marshal Array fail: %s", err)
		t.FailNow()
	}

	if string(arrbuf) != "[]" {
		t.Errorf("Invalid Json marshal array. Invalidbuf[ '%s' ], validbuf[ '%s' ]", string(arrbuf), "[]")
		t.FailNow()
	}

	strSec := godoc.NewStringSection("", "stringval")
	strbuf, err := coder.Marshal(godoc.CreateDocument(strSec))
	if err != nil {
		t.Errorf("JsonCoder marshal String fail: %s", err)
		t.FailNow()
	}

	if string(strbuf) != "\"stringval\"" {
		t.Errorf("Invalid Json marshal string. invalidbuf[ '%s' ], validbuf[ '%s' ]", string(strbuf), "stringval")
		t.FailNow()
	}

	intSec := godoc.NewIntSection("", 10)
	intbuf, err := coder.Marshal(godoc.CreateDocument(intSec))
	if err != nil {
		t.Errorf("JsonCoder marshal int fail: %s", err)
		t.FailNow()
	}

	if string(intbuf) != "10" {
		t.Errorf("Invalid Json marshal string. invalidbuf[ '%s' ], validbuf[ '%s' ]", string(intbuf), "10")
		t.FailNow()
	}

	floatSec := godoc.NewFloat32Section("", 123.123)
	floatbuf, err := coder.Marshal(godoc.CreateDocument(floatSec))
	if err != nil {
		t.Errorf("JsonCoder marshal float fail: %s", err)
		t.FailNow()
	}

	if string(floatbuf) != "123.123" {
		t.Errorf("Invalid Json marshal string. invalidbuf[ '%s' ], validbuf[ '%s' ]", string(floatbuf), "123.123")
		t.FailNow()
	}

	boolSec := godoc.NewBoolSection("", true)
	boolbuf, err := coder.Marshal(godoc.CreateDocument(boolSec))
	if err != nil {
		t.Errorf("JsonCoder marshal bool fail: %s", err)
		t.FailNow()
	}

	if string(boolbuf) != "true" {
		t.Errorf("Invalid Json marshal string. invalidbuf[ '%s' ], validbuf[ '%s' ]", string(floatbuf), "123.123")
		t.FailNow()
	}

	nilSec := godoc.NewNIlSection("")
	nilbuf, err := coder.Marshal(godoc.CreateDocument(nilSec))
	if err != nil {
		t.Errorf("JsonCoder marshal nil fail: %s", err)
		t.FailNow()
	}

	if string(nilbuf) != "null" {
		t.Errorf("Invalid Json marshal nil. invalidbuf[ '%s' ], validbuf[ '%s' ]", string(nilbuf), "null")
		t.FailNow()
	}
}

func TestJsonCoderUnmarshal(t *testing.T) {
	coder := transcoder.CreateJsonTransCoder()

	objbuf := []byte(`
		{}
	`)

	objdoc, err := coder.Unmarshal(objbuf)
	if err != nil {
		t.Errorf("JsonCoder Unmarshal obj fail: %s", err)
		t.FailNow()
	}

	if objdoc.Type() != godoc.Object {
		t.Errorf("Invalid Json UnMarshal object. invalidType[%s], validType[%s]", objdoc.Type().String(), godoc.Object.String())
		t.FailNow()
	}

	arrbuf := []byte(`
		[]
	`)

	arrdoc, err := coder.Unmarshal(arrbuf)
	if err != nil {
		t.Error("JsonCoder Unmarshal array fail : %s", err)
		t.FailNow()
	}

	if arrdoc.Type() != godoc.Array {
		t.Errorf("Invalid Json UnMarshal array. invalidType[%s], validType[%s]", arrdoc.Type().String(), godoc.Array.String())
		t.FailNow()
	}

	strbuf := []byte(`
		"stringv"
	`)

	strdoc, err := coder.Unmarshal(strbuf)
	if err != nil {
		t.Error("JsonCoder Unmarshal string fail : %s", err)
		t.FailNow()
	}

	if strdoc.Type() != godoc.String {
		t.Errorf("Invalid Json UnMarshal string. invalidType[%s], validType[%s]", strdoc.Type().String(), godoc.String.String())
		t.FailNow()
	}

	numbuf := []byte(`
		10
	`)

	numdoc, err := coder.Unmarshal(numbuf)
	if err != nil {
		t.Error("JsonCoder Unmarshal num fail : %s", err)
		t.FailNow()
	}

	if numdoc.Type() != godoc.Float64 {
		t.Errorf("Invalid Json UnMarshal float64. invalidType[%s], validType[%s]", numdoc.Type().String(), godoc.Float64.String())
		t.FailNow()
	}

	boolbuf := []byte(`
		true
	`)

	booldoc, err := coder.Unmarshal(boolbuf)
	if err != nil {
		t.Error("JsonCoder Unmarshal bool fail : %s", err)
		t.FailNow()
	}

	if booldoc.Type() != godoc.Bool {
		t.Errorf("Invalid Json UnMarshal bool. invalidType[%s], validType[%s]", booldoc.Type().String(), godoc.Bool.String())
		t.FailNow()
	}

	nilbuf := []byte(`
		null
	`)

	nildoc, err := coder.Unmarshal(nilbuf)
	if err != nil {
		t.Error("JsonCoder Unmarshal nil fail : %s", err)
		t.FailNow()
	}

	if nildoc.Type() != godoc.Nil {
		t.Errorf("Invalid Json UnMarshal bool. invalidType[%s], validType[%s]", nildoc.Type().String(), godoc.Nil.String())
		t.FailNow()
	}
}
