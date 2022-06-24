package entity

type Location struct {
	id          int
	city        *City
	address     string
	underground *Underground
	latitude    string
	longitude   string
}

func NewLocation(city *City, address string, underground *Underground) *Location {
	return &Location{
		city:        city,
		address:     address,
		underground: underground,
	}
}

func EmptyLocation() *Location {
	return &Location{
		id:          0,
		city:        nil,
		address:     "",
		underground: nil,
		latitude:    "",
		longitude:   "",
	}
}

func (location *Location) SetId(id int) {
	location.id = id
}

func (location *Location) GetId() int {
	return location.id
}

func (location *Location) GetAttributes() []any {
	if location.city == nil {
		location.city = EmptyCity()
	}

	if location.underground == nil {
		location.underground = EmptyUnderground()
	}

	return []any{
		&location.id,
		&location.city.label,
		&location.address,
		&location.underground.label,
		&location.latitude,
		&location.longitude,
	}
}

func (location *Location) ToMap() map[string]any {
	if location.city == nil {
		location.city = EmptyCity()
	}

	if location.underground == nil {
		location.underground = EmptyUnderground()
	}

	return map[string]any{
		"city":        location.city.ToMap(),
		"address":     location.address,
		"underground": location.underground.ToMap(),
		"latitude":    location.latitude,
		"longitude":   location.longitude,
	}
}

func (location *Location) GetCity() *City {
	return location.city
}

func (location *Location) GetAddress() string {
	return location.address
}

func (location *Location) GetUnderGround() *Underground {
	return location.underground
}
