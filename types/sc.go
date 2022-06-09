package types

type Sc struct {
	id          int
	Name        string
	Description string
	Email       string
	Phone       string
	Site        string
}

func NewSc() *Sc {
	sc := Sc{}
	return &sc
}

func (sc *Sc) SetId(id int) {
	sc.id = id
}

func (sc *Sc) ToMap() map[string]any {
	return map[string]any{
		"id":          sc.id,
		"name":        sc.Name,
		"description": sc.Description,
		"email":       sc.Email,
		"phone":       sc.Phone,
		"site":        sc.Site,
	}
}
