package entity

type Underground struct {
	id    int
	label string
	city  *City
}

func EmptyUnderground() *Underground {
	return &Underground{
		id:    0,
		label: "",
		city:  nil,
	}
}

func (underground *Underground) GetAttributes() []any {
	return []any{
		&underground.id, &underground.label,
	}
}

func (underground *Underground) ToMap() map[string]any {
	undergroundMap := map[string]any{
		"id":    underground.id,
		"label": underground.label,
	}

	if underground.city != nil {
		undergroundMap["city"] = underground.city
	}

	return undergroundMap
}

func (underground *Underground) SetId(id int) {
	underground.id = id
}

func (underground *Underground) GetId() int {
	return underground.id
}
