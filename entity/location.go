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

func NewLocation(city *City, address string, undergrounds []*Underground, latitude, longitude string) *Location {
	return &Location{
		city:         city,
		address:      address,
		undergrounds: undergrounds,
		latitude:     latitude,
		longitude:    longitude,
	}
}

func EmptyLocation() *Location {
	return &Location{
		id:           0,
		city:         nil,
		address:      "",
		latitude:     "",
		longitude:    "",
		undergrounds: []*Underground{},
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

	return []any{
		&location.id,
		&location.city.label,
		&location.address,
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
		"city":      location.city.ToMap(),
		"address":   location.address,
		"latitude":  location.latitude,
		"longitude": location.longitude,
	}

	var undergroundsMap []map[string]any

	if len(location.undergrounds) > 0 {
		for i := 0; i < len(location.undergrounds); i++ {
			undergroundsMap = append(undergroundsMap, location.undergrounds[i].ToMap())
		}
	} else {
		undergroundsMap = []map[string]any{}
	}

	locationMap["undergrounds"] = undergroundsMap

	return locationMap
}

func (location *Location) GetCity() *City {
	return location.city
}

func (location *Location) GetAddress() string {
	return location.address
}

func (location *Location) GetLatitude() string {
	return location.latitude
}

func (location *Location) GetLongitude() string {
	return location.longitude
}

func (location *Location) GetUnderGround() *Underground {
	return location.underground
}

func (location *Location) GetUndergrounds() []*Underground {
	return location.undergrounds
}

func (location *Location) AddUnderground(underground *Underground) {
	location.undergrounds = append(location.undergrounds, underground)
}
