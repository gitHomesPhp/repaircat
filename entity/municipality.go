package entity

type Municipality struct {
	id    int
	label string
	city  *City
}

func NewMunicipality(label string, city *City) *Municipality {
	return &Municipality{
		label: label,
		city:  city,
	}
}

func EmptyMunicipality() *Municipality {
	return &Municipality{
		id:    0,
		label: "",
		city:  nil,
	}
}

func (municipality *Municipality) SetId(id int) {
	municipality.id = id
}

func (municipality *Municipality) GetId() int {
	return municipality.id
}

func (municipality *Municipality) GetAttributes() []any {
	return []any{
		&municipality.id,
		&municipality.label,
	}
}
