package entity

type Location struct {
	id          int
	city        string
	address     string
	underground string
}

func NewLocation(city, address, underground string) *Location {
	return &Location{
		city:        city,
		address:     address,
		underground: underground,
	}
}

func EmptyLocation() *Location {
	return &Location{
		id:          0,
		city:        "",
		address:     "",
		underground: "",
	}
}

func (location *Location) SetId(id int) {
	location.id = id
}

func (location *Location) GetId() int {
	return location.id
}

func (location *Location) GetAttributes() (*int, *string, *string, *string) {
	return &location.id, &location.city, &location.address, &location.underground
}

func (location *Location) GetAttributes2() []any {
	return []any{
		&location.id, &location.city, &location.address, &location.underground,
	}
}

func (location *Location) ToMap() map[string]any {
	return map[string]any{
		"city":        location.city,
		"address":     location.address,
		"underground": location.underground,
	}
}

func (location *Location) GetCity() string {
	return location.city
}

func (location *Location) GetAddress() string {
	return location.address
}

func (location *Location) GetUnderGround() string {
	return location.underground
}
