package entity

type Location struct {
	id           int
	city         *City
	address      string
	underground  *Underground
	latitude     string
	longitude    string
	undergrounds []*Underground
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

	locationMap := map[string]any{
		"city":    location.city.ToMap(),
		"address": location.address,
		//"underground": location.underground.ToMap(),
		"latitude":  location.latitude,
		"longitude": location.longitude,
	}

	if len(location.undergrounds) > 0 {
		var undergroundsMap []map[string]any

		for i := 0; i < len(location.undergrounds); i++ {
			undergroundsMap = append(undergroundsMap, location.undergrounds[i].ToMap())
		}

		locationMap["undergrounds"] = undergroundsMap
	}

	return locationMap
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

func (location *Location) AddUnderground(underground *Underground) {
	location.undergrounds = append(location.undergrounds, underground)
}
