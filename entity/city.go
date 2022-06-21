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

func (city *City) ToMap() map[string]any {
	return map[string]any{
		"label": city.label,
		"code":  city.code,
		"id":    city.id,
	}
}

func (city *City) GetAttributes() []any {
	return []any{
		&city.id, &city.label, &city.code,
	}
}

func (city *City) SetId(id int) {
	city.id = id
}

func (city *City) GetId() int {
	return city.id
}

func (city *City) GetLabel() string {
	return city.label
}

func (city *City) GetCode() string {
	return city.code
}
