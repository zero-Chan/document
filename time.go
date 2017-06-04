package godoc

import (
	"reflect"
	"time"
)

type TimeSection struct {
	name string
	data time.Time
}

func CreateTimeSection(name string) TimeSection {
	s := TimeSection{
		name: name,
	}

	return s
}

func NewTimeSection(name string) *TimeSection {
	s := CreateTimeSection(name)
	return &s
}

func CreateUnixSecSection(name string, sec int64) TimeSection {
	s := CreateTimeSection(name)
	s.data = time.Unix(sec, 0)
	return s
}

func NewUnixSecSection(name string, sec int64) *TimeSection {
	s := CreateUnixSecSection(name, sec)
	return &s
}

func CreateUnixNanoSecSection(name string, nsec int64) TimeSection {
	s := CreateTimeSection(name)
	s.data = time.Unix(0, nsec)
	return s
}

func NewUnixNanoSecSection(name string, nsec int64) *TimeSection {
	s := CreateUnixNanoSecSection(name, nsec)
	return &s
}

func (sec TimeSection) Type() Type {
	return Time
}

func (sec TimeSection) Name() string {
	return sec.name
}

func (sec *TimeSection) SetName(name string) {
	sec.name = name
}

func (sec *TimeSection) Copy() *TimeSection {
	ts := NewTimeSection(sec.name)
	ts.data = sec.data

	return ts
}

// TODO : time.method to return

func (sec *TimeSection) Unmarshal(data interface{}) error {
	return nil
}

func (sec *TimeSection) setValue(rv *reflect.Value) error {
	return nil
}

func (sec *TimeSection) Marshal(data interface{}) error {
	return nil
}

func (sec *TimeSection) getValue(rv reflect.Value) error {
	return nil
}
