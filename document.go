package godoc

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
	"Time",
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
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Time
)

type Section interface {
	Type() Type
	SetName(string)
	Name() string
	Unmarshal(data interface{}) error
	Marshal(data interface{}) error
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

func (doc Document) Name() string {
	return doc.entity.Name()
}

func (doc *Document) Copy() *Document {
	switch {
	case doc.IsArray():
		asec, _ := doc.Array()
		return NewDocument(asec.Copy())
	case doc.IsBool():
		bsec, _ := doc.Bool()
		return NewDocument(bsec.Copy())
	case doc.IsNumber():
		nsec, _ := doc.Number()
		return NewDocument(nsec.Copy())
	case doc.IsObject():
		osec, _ := doc.Object()
		return NewDocument(osec.Copy())
	case doc.IsString():
		ssec, _ := doc.String()
		return NewDocument(ssec.Copy())
	case doc.IsNil():
		nsec, _ := doc.Nil()
		return NewDocument(nsec.Copy())
	}

	return nil
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
		fallthrough
	case Uint:
		fallthrough
	case Uint8:
		fallthrough
	case Uint16:
		fallthrough
	case Uint32:
		fallthrough
	case Uint64:
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

func (doc Document) IsUint() bool {
	if doc.entity.Type() != Uint {
		return false
	}
	return true
}

func (doc Document) IsUint8() bool {
	if doc.entity.Type() != Uint8 {
		return false
	}
	return true
}

func (doc Document) IsUint16() bool {
	if doc.entity.Type() != Uint16 {
		return false
	}
	return true
}

func (doc Document) IsUint32() bool {
	if doc.entity.Type() != Uint32 {
		return false
	}
	return true
}

func (doc Document) IsUint64() bool {
	if doc.entity.Type() != Uint64 {
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

func (doc Document) IsTime() bool {
	if doc.entity.Type() != Time {
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

func (doc *Document) Time() (*TimeSection, error) {
	if !doc.IsTime() {
		return nil, ErrorInvalidSectionTypeConvert{ErrorType: doc.Type(), Type: Nil}
	}

	sec, ok := doc.entity.(*TimeSection)
	if !ok {
		return nil, ErrorSection2Entity{SectionType: doc.Type(), ConvertEntity: "*TimeSection"}
	}

	return sec, nil
}
