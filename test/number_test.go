package test

import (
	"testing"

	"github.com/zero-Chan/godoc"
)

func TestNumberSection(t *testing.T) {
	// int
	sec := godoc.NewIntSection("", 10)
	var buf int
	err := sec.Unmarshal(&buf)
	if err != nil {
		t.Errorf("sec.Unmarshal(%T) fail: %s", &buf, err)
		t.FailNow()
	}

	if buf != 10 {
		t.Errorf("[%d] not equal [%d]", buf, 10)
		t.FailNow()
	}

	buf = 20
	err = sec.Marshal(buf)
	if err != nil {
		t.Errorf("sec.Marshal(%T) fail: %s", buf, err)
		t.FailNow()
	}

	if sec.Int() != 20 {
		t.Errorf("[%d] not equal [%d]", sec.Int(), 20)
		t.FailNow()
	}

	var buf2 int
	var buf2Ptr *int = &buf2

	// *int
	err = sec.Unmarshal(buf2Ptr)
	if err != nil {
		t.Errorf("sec.Unmarshal(%T) fail: %s", buf2Ptr, err)
		t.FailNow()
	}

	if buf2 != 20 || *buf2Ptr != 20 {
		t.Errorf("[%d] not equal [%d]", buf2, 20)
		t.FailNow()
	}

	*buf2Ptr = 30
	err = sec.Marshal(buf2Ptr)
	if err != nil {
		t.Errorf("sec.Marshal(%T) fail: %s", buf2Ptr, err)
		t.FailNow()
	}

	if sec.Int() != 30 {
		t.Errorf("[%d] not equal [%d]", sec.Int(), 30)
		t.FailNow()
	}

	// to float
	var bufFloat float32
	err = sec.Unmarshal(&bufFloat)
	if err != nil {
		t.Errorf("sec.Marshal(%T) fail: %s", bufFloat, err)
		t.FailNow()
	}

	if bufFloat != 30.0 {
		t.Errorf("[%f] not equal [%d]", bufFloat, 30)
		t.FailNow()
	}

	bufFloat = 30.1234
	err = sec.Marshal(bufFloat)
	if err != nil {
		t.Errorf("sec.Marshal(%T) fail: %s", bufFloat, err)
		t.FailNow()
	}

	if sec.Int() != 30 {
		t.Errorf("[%d] not equal [%d]", sec.Int(), 30)
		t.FailNow()
	}

	if sec.Float32() != 30.1234 {
		t.Errorf("[%f] not equal [%f]", sec.Float32(), 30.1234)
		t.FailNow()
	}

	// interface{}
	var contain interface{}
	err = sec.Unmarshal(&contain)
	if err != nil {
		t.Errorf("sec.Unmarshal(%T) fail: %s", &contain, err)
		t.FailNow()
	}

	ctnFloat, ok := contain.(float32)
	if !ok {
		t.Errorf("interface[%T] can not convert to [float32]", contain)
		t.FailNow()
	}

	if float32(ctnFloat) != 30.1234 {
		t.Errorf("[%f] not equal [%f]", ctnFloat, 30.1234)
		t.FailNow()
	}

	contain = "Hello"
	err = sec.Marshal(contain)
	if err == nil {
		t.Errorf("sec.Marshal(%T) OK: type[string] can convert to stringSection", contain)
		t.FailNow()
	}

	// other type can not convert marshal to NumberSection
	var i string = "Hello_2"
	err = sec.Marshal(&i)
	if err == nil {
		t.Errorf("sec.Marshal(%T) OK: type[string] can convert to stringSection", contain)
		t.FailNow()
	}
}
