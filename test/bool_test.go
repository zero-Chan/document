package test

import (
	"testing"

	"github.com/zero-Chan/document"
)

func TestBoolSection(t *testing.T) {
	// bool
	sec := document.NewBoolSection("Key", true)
	var buf bool
	err := sec.Unmarshal(&buf)
	if err != nil {
		t.Errorf("sec.Unmarshal(%T) fail: %s", &buf, err)
		t.FailNow()
	}

	if buf != true {
		t.Errorf("[%t] not equal [%t]", buf, true)
		t.FailNow()
	}

	buf = false
	err = sec.Marshal(buf)
	if err != nil {
		t.Errorf("sec.Marshal(%T) fail: %s", buf, err)
		t.FailNow()
	}

	if sec.Bool() != false {
		t.Errorf("[%t] not equal [%t]", sec.Bool(), false)
		t.FailNow()
	}

	var buf2 bool
	var buf2Ptr *bool = &buf2

	// *string
	err = sec.Unmarshal(buf2Ptr)
	if err != nil {
		t.Errorf("sec.Unmarshal(%T) fail: %s", buf2Ptr, err)
		t.FailNow()
	}

	if buf2 != false || *buf2Ptr != false {
		t.Errorf("[%s] not equal [%s]", buf2, false)
		t.FailNow()
	}

	*buf2Ptr = true
	err = sec.Marshal(buf2Ptr)
	if err != nil {
		t.Errorf("sec.Marshal(%T) fail: %s", buf2Ptr, err)
		t.FailNow()
	}

	if sec.Bool() != true {
		t.Errorf("[%t] not equal [%t]", sec.Bool(), true)
		t.FailNow()
	}

	// interface{}
	var contain interface{}
	err = sec.Unmarshal(&contain)
	if err != nil {
		t.Errorf("sec.Unmarshal(%T) fail: %s", &contain, err)
		t.FailNow()
	}

	ctnBool, ok := contain.(bool)
	if !ok {
		t.Errorf("interface[%T] can not convert to [bool]", contain)
		t.FailNow()
	}

	if ctnBool != true {
		t.Errorf("[%t] not equal [%t]", ctnBool, true)
		t.FailNow()
	}

	contain = false
	err = sec.Marshal(contain)
	if err != nil {
		t.Errorf("sec.Marshal(%T) fail: %s", contain, err)
		t.FailNow()
	}

	if sec.Bool() != false {
		t.Errorf("[%t] not equal [%t]", sec.Bool(), false)
		t.FailNow()
	}

	// other type can not convert marshal to boolSection
	var i int = 10
	err = sec.Marshal(&i)
	if err == nil {
		t.Errorf("sec.Marshal(%T) OK: type[int] can convert to boolSection", contain)
		t.FailNow()
	}
}
