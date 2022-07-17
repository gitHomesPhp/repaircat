package entity

import "github.com/gitHomesPhp/repaircat/ddd/valueobject"

type Location struct {
	Id             int                         `json:"-"`
	Address        string                      `json:"address"`
	Latitude       string                      `json:"latitude"`
	Longitude      string                      `json:"longitude"`
	City           *City                       `json:"city"`
	Undergrounds   []*Underground              `json:"undergrounds"`
	Municipalities []*valueobject.Municipality `json:"municipalities"`

	cityId int
}
