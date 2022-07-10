package scCard

import "github.com/gitHomesPhp/repaircat/ddd/aggregate"

type Repository interface {
	Get()
	Add()
	List(page int) (error, []*aggregate.ScCard, map[string]bool)
}
