package test

import (
	"testing"

	"github.com/zero-Chan/document"
)

func TestStringSection(t *testing.T) {
	// string
	sec := document.NewStringSection("Key", "Val_1")
	var buf string
	err := sec.Unmarshal(&buf)
	if err != nil {
		t.Errorf("sec.Unmarshal(%T) fail: %s", &buf, err)
		t.FailNow()
	}

	if buf != "Val_1" {
		t.Errorf("[%s] not equal [%s]", buf, "Val_1")
		t.FailNow()
	}

	buf = "Val_2"
	err = sec.Marshal(buf)
	if err != nil {
		t.Errorf("sec.Marshal(%T) fail: %s", buf, err)
		t.FailNow()
	}

	if sec.String() != "Val_2" {
		t.Errorf("[%s] not equal [%s]", sec.String(), "Val_2")
		t.FailNow()
	}

	var buf2 string
	var buf2Ptr *string = &buf2

	// *string
	err = sec.Unmarshal(buf2Ptr)
	if err != nil {
		t.Errorf("sec.Unmarshal(%T) fail: %s", buf2Ptr, err)
		t.FailNow()
	}

	if buf2 != "Val_2" || *buf2Ptr != "Val_2" {
		t.Errorf("[%s] not equal [%s]", buf2, "Val_2")
		t.FailNow()
	}

	*buf2Ptr = "Val_3"
	err = sec.Marshal(buf2Ptr)
	if err != nil {
		t.Errorf("sec.Marshal(%T) fail: %s", buf2Ptr, err)
		t.FailNow()
	}

	if sec.String() != "Val_3" {
		t.Errorf("[%s] not equal [%s]", sec.String(), "Val_3")
		t.FailNow()
	}

	// interface{}
	var contain interface{}
	err = sec.Unmarshal(&contain)
	if err != nil {
		t.Errorf("sec.Unmarshal(%T) fail: %s", &contain, err)
		t.FailNow()
	}

	ctnStr, ok := contain.(string)
	if !ok {
		t.Errorf("interface[%T] can not convert to [string]", contain)
		t.FailNow()
	}

	if ctnStr != "Val_3" {
		t.Errorf("[%s] not equal [%s]", ctnStr, "Val_3")
		t.FailNow()
	}

	contain = "Val_4"
	err = sec.Marshal(contain)
	if err != nil {
		t.Errorf("sec.Marshal(%T) fail: %s", contain, err)
		t.FailNow()
	}

	if sec.String() != "Val_4" {
		t.Errorf("[%s] not equal [%s]", sec.String(), "Val_4")
		t.FailNow()
	}

	contain = 1
	err = sec.Marshal(contain)
	if err == nil {
		t.Errorf("sec.Marshal(%T) OK: type[int] can convert to stringSection", contain)
		t.FailNow()
	}

	// other type can not convert marshal to stringSection
	var i int = 10
	sec.Marshal(&i)
	if err == nil {
		t.Errorf("sec.Marshal(%T) OK: type[int] can convert to stringSection", contain)
		t.FailNow()
	}
}
