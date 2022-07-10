package entity

type Sc struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
	Site        string    `json:"site"`
	Location    *Location `json:"location"`

	locationId int
}
