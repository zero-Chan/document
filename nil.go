package document

type NilSection struct {
	name string
	data interface{}
}

func CreateNilSection(name string) NilSection {
	sec := NilSection{
		name: name,
		data: nil,
	}
	return sec
}

func NewNIlSection(name string) *NilSection {
	sec := CreateNilSection(name)
	return &sec
}

func (sec NilSection) Type() Type {
	return Nil
}
