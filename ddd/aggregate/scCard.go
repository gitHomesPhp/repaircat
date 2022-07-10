package aggregate

import (
	"github.com/gitHomesPhp/repaircat/ddd/entity"
	"github.com/gitHomesPhp/repaircat/ddd/valueobject"
)

type ScCard struct {
	Sc         *entity.Sc              `json:"sc"`
	ReviewInfo *valueobject.ReviewInfo `json:"review_info"`
}

func NewScCard() *ScCard {
	var undergrounds []*entity.Underground

	city := &entity.City{
		Code:  "",
		Label: "",
	}

	location := &entity.Location{
		Address:      "",
		Latitude:     "",
		Longitude:    "",
		Undergrounds: undergrounds,
		City:         city,
	}

	sc := &entity.Sc{
		Id:          0,
		Name:        "",
		Description: "",
		Phone:       "",
		Email:       "",
		Site:        "",
		Location:    location,
	}

	reviewInfo := &valueobject.ReviewInfo{}

	return &ScCard{
		Sc:         sc,
		ReviewInfo: reviewInfo,
	}
}
