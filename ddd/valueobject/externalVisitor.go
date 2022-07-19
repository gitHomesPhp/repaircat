package valueobject

type ExternalVisitor struct {
	Id     int    `json:"-"`
	Source string `json:"source"`
	Name   string `json:"name"`
}
