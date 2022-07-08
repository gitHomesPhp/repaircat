package entity

type Location struct {
	Id           int
	Address      string
	Latitude     string
	Longitude    string
	City         *City
	Undergrounds []*Underground

	cityId int
}
