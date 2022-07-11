package scCardExtension

import "github.com/gitHomesPhp/repaircat/ddd/aggregate"

type Repository interface {
	Get(id int) (error, *aggregate.ScCardExtension)
}
