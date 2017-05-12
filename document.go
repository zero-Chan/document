package document

type Type int

var typeString = []string{
	"Object",
	"Number",
	"String",
	"Array",
	"Bool",
	"Nil",
	"Int",
	"Int8",
	"Int16",
	"Int32",
	"Int64",
	"Float32",
	"Float64",
}

func (t Type) String() string {
	return typeString[t]
}

const (
	Object Type = iota
	Number
	String
	Array
	Bool
	Nil
	Int
	Int8
	Int16
	Int32
	Int64
	Float32
	Float64
)

type Section interface {
	Type() Type
}

type Document struct {
	entity Section
}

func CreateDocument(entity Section) Document {
	if entity == nil {
		entity = NewNIlSection("")
	}

	doc := Document{
		entity: entity,
	}

	return doc
}

func NewDocument(entity Section) *Document {
	doc := CreateDocument(entity)
	return &doc
}

func (doc Document) Type() Type {
	return doc.entity.Type()
}

func (doc Document) IsObject() bool {
	if doc.entity.Type() != Object {
		return false
	}

	return true
}

func (doc Document) IsNumber() bool {
	var ok bool = false
	switch doc.entity.Type() {
	case Int:
		fallthrough
	case Int8:
		fallthrough
	case Int16:
		fallthrough
	case Int32:
		fallthrough
	case Int64:
		fallthrough
	case Float32:
		fallthrough
	case Float64:
		ok = true
	default:
		ok = false
	}

	return ok
}

func (doc Document) IsInt() bool {
	if doc.entity.Type() != Int {
		return false
	}
	return true
}

func (doc Document) IsInt8() bool {
	if doc.entity.Type() != Int8 {
		return false
	}
	return true
}

func (doc Document) IsInt16() bool {
	if doc.entity.Type() != Int16 {
		return false
	}
	return true
}

func (doc Document) IsInt32() bool {
	if doc.entity.Type() != Int32 {
		return false
	}
	return true
}

func (doc Document) IsInt64() bool {
	if doc.entity.Type() != Int64 {
		return false
	}
	return true
}

func (doc Document) IsFloat32() bool {
	if doc.entity.Type() != Float32 {
		return false
	}
	return true
}

func (doc Document) IsFloat64() bool {
	if doc.entity.Type() != Float64 {
		return false
	}
	return true
}

func (doc Document) IsString() bool {
	if doc.entity.Type() != String {
		return false
	}

	return true
}

func (doc Document) IsBool() bool {
	if doc.entity.Type() != Bool {
		return false
	}

	return true
}

func (doc Document) IsArray() bool {
	if doc.entity.Type() != Array {
		return false
	}

	return true
}

func (doc Document) IsNil() bool {
	if doc.entity.Type() != Nil {
		return false
	}

	return true
}

func (doc *Document) Object() (*ObjectSection, error) {
	if !doc.IsObject() {
		return nil, ErrorInvalidSectionTypeConvert{ErrorType: doc.Type(), Type: Object}
	}

	sec, ok := doc.entity.(*ObjectSection)
	if !ok {
		return nil, ErrorSection2Entity{SectionType: doc.Type(), ConvertEntity: "*ObjectSection"}
	}

	return sec, nil
}

func (doc *Document) Number() (*NumberSection, error) {
	if !doc.IsNumber() {
		return nil, ErrorInvalidSectionTypeConvert{ErrorType: doc.Type(), Type: Number}
	}

	sec, ok := doc.entity.(*NumberSection)
	if !ok {
		return nil, ErrorSection2Entity{SectionType: doc.Type(), ConvertEntity: "*NumberSection"}
	}

	return sec, nil
}

func (doc *Document) String() (*StringSection, error) {
	if !doc.IsString() {
		return nil, ErrorInvalidSectionTypeConvert{ErrorType: doc.Type(), Type: String}
	}

	sec, ok := doc.entity.(*StringSection)
	if !ok {
		return nil, ErrorSection2Entity{SectionType: doc.Type(), ConvertEntity: "*StringSection"}
	}

	return sec, nil
}

func (doc *Document) Array() (*ArraySection, error) {
	if !doc.IsArray() {
		return nil, ErrorInvalidSectionTypeConvert{ErrorType: doc.Type(), Type: Array}
	}

	sec, ok := doc.entity.(*ArraySection)
	if !ok {
		return nil, ErrorSection2Entity{SectionType: doc.Type(), ConvertEntity: "*ArraySection"}
	}

	return sec, nil
}

func (doc *Document) Bool() (*BoolSection, error) {
	if !doc.IsBool() {
		return nil, ErrorInvalidSectionTypeConvert{ErrorType: doc.Type(), Type: Bool}
	}

	sec, ok := doc.entity.(*BoolSection)
	if !ok {
		return nil, ErrorSection2Entity{SectionType: doc.Type(), ConvertEntity: "*BoolSection"}
	}

	return sec, nil
}

func (doc *Document) Nil() (*NilSection, error) {
	if !doc.IsNil() {
		return nil, ErrorInvalidSectionTypeConvert{ErrorType: doc.Type(), Type: Nil}
	}

	sec, ok := doc.entity.(*NilSection)
	if !ok {
		return nil, ErrorSection2Entity{SectionType: doc.Type(), ConvertEntity: "*NilSection"}
	}

	return sec, nil
}
