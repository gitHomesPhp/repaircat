package aggregate

import (
	"github.com/gitHomesPhp/repaircat/ddd/entity"
	"github.com/gitHomesPhp/repaircat/ddd/valueobject"
)

type ReviewForScCardExtension struct {
	Review  *entity.ReviewExternal
	Visitor *valueobject.ExternalVisitor
}

func NewReviewForScCardExtension() *ReviewForScCardExtension {
	review := &entity.ReviewExternal{
		Id:        0,
		Text:      "",
		Rating:    0,
		VisitorId: 0,
		ScId:      0,
	}

	visitor := &valueobject.ExternalVisitor{
		Id:     0,
		Source: "",
	}

	return &ReviewForScCardExtension{
		Review:  review,
		Visitor: visitor,
	}
}
