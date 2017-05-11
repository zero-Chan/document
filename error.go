package document

import (
	"fmt"
	"reflect"
)

const (
	AcceptNilParam = "AcceptNilParam"

	ObjectNotExistKey         = "ObjectNotExistKey"
	InvalidSectionTypeConvert = "InvalidSectionTypeConvert"
	Section2Entity            = "Section2Entity"

	UnmarshalFail = "UnmarshalFail"
	SetValueFail  = "SetValueFail"

	MarshalFail  = "MarshalFail"
	GetValueFail = "GetValueFail"
)

type ErrorAcceptNilParam struct {
	FunctionName string
	NilParam     interface{}
}

func (e ErrorAcceptNilParam) Error() string {
	return fmt.Sprintf("ErrorCode[%s] FunctionName[%s] NilParam[%T]", AcceptNilParam, e.FunctionName, e.NilParam)
}

type ErrorObjectNotExistKey struct {
	ObjectKey string
	Key       string
}

func (e ErrorObjectNotExistKey) Error() string {
	return fmt.Sprintf("ErrorCode[%s] Object[%s] Key[%s]", ObjectNotExistKey, e.ObjectKey, e.Key)
}

type ErrorInvalidSectionTypeConvert struct {
	ErrorType Type
	Type      Type
}

func (e ErrorInvalidSectionTypeConvert) Error() string {
	return fmt.Sprintf("ErrorCode[%s] ErrorType[%s] Type[%s]", InvalidSectionTypeConvert, e.ErrorType.String(), e.Type.String())
}

type ErrorSection2Entity struct {
	SectionType   Type
	ConvertEntity string
}

func (e ErrorSection2Entity) Error() string {
	return fmt.Sprintf("ErrorCode[%s] SectionType[%s] ConvertEntity[%s]", Section2Entity, e.SectionType.String(), e.ConvertEntity)
}

type ErrorUnmarshalFail struct {
	Type  Type
	Value reflect.Value
}

func (e ErrorUnmarshalFail) Error() string {
	switch {
	case e.Value.Kind() != reflect.Ptr:
		return fmt.Sprintf("ErrorCode[%s] Type[%s] InvalidValueType[%s] ShouldType[%s]", UnmarshalFail, e.Type.String(), e.Value.Kind().String(), reflect.Ptr.String())
	case e.Value.IsNil():
		return fmt.Sprintf("ErrorCode[%s] Type[%s] Value[nil]", UnmarshalFail, e.Type.String())

	}

	return ""
}

type ErrorSetValueFail struct {
	Type    Type
	KeyName string
	Value   reflect.Value
}

func (e ErrorSetValueFail) Error() string {
	return fmt.Sprintf("ErrorCode[%s] Type[%s] KeyName[%s] CanSet[%t] SetValueType[%s]", SetValueFail, e.Type.String(), e.KeyName, e.Value.CanSet(), e.Value.Type().String())
}

type ErrorMarshalFail struct {
	Type        Type
	InvalidKind reflect.Kind
	ValidKind   reflect.Kind
}

func (e ErrorMarshalFail) Error() string {
	return fmt.Sprintf("ErrorCode[%s] Type[%s] InvalidValueType[%s] validType[%s]", MarshalFail, e.Type.String(), e.InvalidKind.String(), e.ValidKind.String())
}

type ErrorGetValueFail struct {
	Type        Type
	InvalidKind reflect.Kind
	ValidKind   reflect.Kind
}

func (e ErrorGetValueFail) Error() string {
	return fmt.Sprintf("ErrorCode[%s] Type[%s] InvalidKind[%s] ValidKind[%s]", GetValueFail, e.Type.String(), e.InvalidKind.String(), e.ValidKind.String())
}
