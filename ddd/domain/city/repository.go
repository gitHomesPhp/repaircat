package city

import "github.com/gitHomesPhp/repaircat/ddd/entity"

type Repository interface {
	List() (error, []*entity.City)
}
