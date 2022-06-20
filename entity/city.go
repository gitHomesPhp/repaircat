package entity

type City struct {
	id    int
	label string
	code  string
}

func NewCity(label, code string) *City {
	return &City{
		label: label,
		code:  code,
	}
}

func EmptyCity() *City {
	return &City{
		id:    0,
		label: "",
		code:  "",
	}
}

func (city *City) GetAttributes() []any {
	return []any{
		&city.id, &city.label, &city.code,
	}
}

func (city *City) GetLabel() string {
	return city.label
}

func (city *City) GetCode() string {
	return city.code
}
