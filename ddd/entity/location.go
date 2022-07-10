package entity

type Location struct {
	Id           int            `json:"-"`
	Address      string         `json:"address"`
	Latitude     string         `json:"latitude"`
	Longitude    string         `json:"longitude"`
	City         *City          `json:"city"`
	Undergrounds []*Underground `json:"undergrounds"`

	cityId int
}
