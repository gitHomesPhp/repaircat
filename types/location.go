package types

type Location struct {
	id          int
	City        string
	Address     string
	Underground string
}

func NewLocation() *Location {
	return &Location{}
}

func BuildLocation(city string, address string, underground string) *Location {
	return &Location{
		City:        city,
		Address:     address,
		Underground: underground,
	}
}

func (location *Location) SetId(id int) {
	location.id = id
}

func (location *Location) ToMap() map[string]any {
	return map[string]any{
		"id":          location.id,
		"city":        location.City,
		"address":     location.Address,
		"underground": location.Underground,
	}
}
