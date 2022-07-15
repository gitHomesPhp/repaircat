package valueobject

type Municipality struct {
	Id    int    `json:"id"`
	Label string `json:"label"`
	Slug  string `json:"slug"`
}
