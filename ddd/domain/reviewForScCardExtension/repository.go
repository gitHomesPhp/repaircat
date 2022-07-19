package reviewForScCardExtension

import "github.com/gitHomesPhp/repaircat/ddd/aggregate"

type Repository interface {
	ListBySc(page int, scId int) []*aggregate.ReviewForScCardExtension
}
