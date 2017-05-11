package document

type ArraySection struct {
	name string
	data []*Document
}

func CreateArraySection(name string) ArraySection {
	sec := ArraySection{
		name: name,
		data: make([]*Document, 0),
	}
	return sec
}

func NewArraySection(name string) *ArraySection {
	sec := CreateArraySection(name)
	return &sec
}

func (sec ArraySection) Type() Type {
	return Array
}

func (sec *ArraySection) Unmarshal(data interface{}) error {

	return nil
}

func (sec *ArraySection) Marshal(data interface{}) error {
	return nil
}
