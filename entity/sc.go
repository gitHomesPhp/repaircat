package entity

type Sc struct {
	id          int
	name        string
	description string
	phone       string
	email       string
	site        string
	location    *Location
}

func NewSc(name, description, phone, email, site string) *Sc {
	return &Sc{
		name:        name,
		description: description,
		phone:       phone,
		email:       email,
		site:        site,
		location:    nil,
	}
}

func EmptySc() *Sc {
	return &Sc{
		id:          0,
		name:        "",
		description: "",
		phone:       "",
		email:       "",
		site:        "",
	}
}

func (sc *Sc) GetAttributes() (*int, *string, *string, *string, *string, *string) {
	return &sc.id, &sc.name, &sc.description, &sc.phone, &sc.email, &sc.site
}

func (sc *Sc) GetAttributes2() []any {
	return []any{
		&sc.id, &sc.name, &sc.description, &sc.phone, &sc.email, &sc.site,
	}
}

func (sc *Sc) SetId(id int) {
	sc.id = id
}

func (sc *Sc) AddLocation(location *Location) {
	sc.location = location
}

func (sc *Sc) ToMap() map[string]any {
	scMap := map[string]any{
		"id":          sc.id,
		"name":        sc.name,
		"description": sc.description,
		"email":       sc.email,
		"phone":       sc.phone,
		"site":        sc.site,
	}

	if sc.location == nil {
		sc.location = EmptyLocation()
	}

	scMap["location"] = sc.location.ToMap()

	return scMap
}

func (sc *Sc) GetName() string {
	return sc.name
}
func (sc *Sc) GetPhone() string {
	return sc.phone
}
func (sc *Sc) GetDescription() string {
	return sc.description
}
func (sc *Sc) GetEmail() string {
	return sc.email
}
func (sc *Sc) GetSite() string {
	return sc.site
}
func (sc *Sc) GetLocation() *Location {
	return sc.location
}
