package municipality

import "github.com/gitHomesPhp/repaircat/ddd/valueobject"

type Repository interface {
	ListByCity(cityId int) (error, []*valueobject.Municipality)
}
