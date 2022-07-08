package service

import "github.com/gitHomesPhp/repaircat/repository/review_external_repository"

func GetScReviewsInfo(scId int) map[string]any {
	reviewInfo := map[string]any{}

	countReviews := review_external_repository.GetCountReviewsSc(scId)

	reviewInfo["count"] = countReviews

	return reviewInfo
}
