package entity

type Sc struct {
	Id          int `json:"id"`
	Name        string
	Description string
	Phone       string
	Email       string
	Site        string
	Location    *Location

	locationId int
}
