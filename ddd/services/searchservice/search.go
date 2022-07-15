package searchservice

import "github.com/gitHomesPhp/repaircat/ddd/aggregate"

type Search struct {
	q                string
	cityCode         string
	undergroundSlug  string
	municipalitySlug string
	page             string

	scCards []*aggregate.ScCard
}

func CreateService(q, cityCode, undergroundSlug, municipalitySlug, page string) *Search {
	return &Search{
		q:                q,
		cityCode:         cityCode,
		undergroundSlug:  undergroundSlug,
		municipalitySlug: municipalitySlug,
		page:             page,
	}
}

func (searchService *Search) GetScCards() []*aggregate.ScCard {
	if searchService.q != "" {
		searchService.getScCardsByQuery()
	}

	if searchService.undergroundSlug != "" {
		searchService.getScCardsByUnderground()
	}

	if searchService.municipalitySlug != "" {
		searchService.getScCardsByMunicipality()
	}

	return searchService.scCards
}

func (searchService *Search) getScCardsByQuery() {

}

func (searchService *Search) getScCardsByUnderground() {

}

func (searchService *Search) getScCardsByMunicipality() {

}
