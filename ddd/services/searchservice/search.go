package searchservice

import (
	"github.com/gitHomesPhp/repaircat/ddd/aggregate"
	"github.com/gitHomesPhp/repaircat/ddd/domain/scCard/postgresql"
)

type Search struct {
	q                string
	cityCode         string
	undergroundSlug  string
	municipalitySlug string
	page             int

	scCards      []*aggregate.ScCard
	nextPrevious map[string]bool
}

func CreateService(q, cityCode, undergroundSlug, municipalitySlug string, page int) *Search {
	return &Search{
		q:                q,
		cityCode:         cityCode,
		undergroundSlug:  undergroundSlug,
		municipalitySlug: municipalitySlug,
		page:             page,
	}
}

func (searchService *Search) GetScCards() ([]*aggregate.ScCard, map[string]bool) {
	if searchService.q != "" {
		searchService.getScCardsByQuery()
	}

	if searchService.undergroundSlug != "" {
		searchService.getScCardsByUnderground(searchService.page)
	}

	if searchService.municipalitySlug != "" {
		searchService.getScCardsByMunicipality(searchService.page)
	}

	return searchService.scCards, searchService.nextPrevious
}

func (searchService *Search) getScCardsByQuery() {
	scCardRepo := postgresql.GetRepo()

	_, scCards, nextPrevious := scCardRepo.Search(searchService.page, searchService.cityCode, searchService.q)

	searchService.scCards = scCards
	searchService.nextPrevious = nextPrevious
}

func (searchService *Search) getScCardsByUnderground(page int) {
	scCardRepo := postgresql.GetRepo()

	_, scCards, nextPrevious := scCardRepo.ListByFields(page, map[string]string{
		"underground.slug": searchService.undergroundSlug,
		"city.code":        searchService.cityCode,
	})

	searchService.scCards = scCards
	searchService.nextPrevious = nextPrevious
}

func (searchService *Search) getScCardsByMunicipality(page int) {
	scCardRepo := postgresql.GetRepo()

	_, scCards, nextPrevious := scCardRepo.ListByFields(page, map[string]string{
		"municipality.slug": searchService.municipalitySlug,
		"city.code":         searchService.cityCode,
	})

	searchService.scCards = scCards
	searchService.nextPrevious = nextPrevious
}
