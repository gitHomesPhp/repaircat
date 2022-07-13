package underground

import "github.com/gitHomesPhp/repaircat/ddd/entity"

type Repository interface {
	ListByCity(cityId int) (error, []*entity.Underground)
}
