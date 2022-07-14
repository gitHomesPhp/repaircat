package postgresql

import (
	"context"
	"fmt"
	"github.com/gitHomesPhp/repaircat/ddd/aggregate"
	"github.com/gitHomesPhp/repaircat/ddd/entity"
	"github.com/gitHomesPhp/repaircat/ddd/valueobject"
	"github.com/jackc/pgx/v4"
	"os"
)

type ScCardRepository struct {
	scCards    []*aggregate.ScCard
	mapScCards map[int]*aggregate.ScCard
}

func GetRepo() *ScCardRepository {
	return &ScCardRepository{
		scCards:    []*aggregate.ScCard{},
		mapScCards: map[int]*aggregate.ScCard{},
	}
}

func (ScCardRepository *ScCardRepository) List(page int) (error, []*aggregate.ScCard, map[string]bool) {
	const COUNT = 10

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return err, nil, nil
	}

	rows, _ := conn.Query(context.Background(), SelectScCardsByCursor, COUNT, COUNT*page-COUNT)

	var scIds []int
	var locationIds []int
	locationIdsMap := map[int]int{}

	for rows.Next() {
		scCard := aggregate.NewScCard()
		var locationId int

		attrs := []any{
			&scCard.Sc.Id,
			&scCard.Sc.Name,
			&scCard.Sc.Description,
			&scCard.Sc.Phone,
			&scCard.Sc.Email,
			&scCard.Sc.Site,
			&scCard.Sc.Location.Address,
			&scCard.Sc.Location.Latitude,
			&scCard.Sc.Location.Longitude,
			&scCard.Sc.Location.City.Label,

			&locationId,
		}

		rows.Scan(attrs...)

		scIds = append(scIds, scCard.Sc.Id)
		locationIds = append(locationIds, locationId)

		locationIdsMap[locationId] = scCard.Sc.Id

		ScCardRepository.scCards = append(ScCardRepository.scCards, scCard)
		ScCardRepository.mapScCards[scCard.Sc.Id] = scCard
	}

	rows, _ = conn.Query(context.Background(), GetScCardInfo, scIds)

	for rows.Next() {
		var scId int
		reviewInfo := &valueobject.ReviewInfo{
			Count:  0,
			Rating: 0,
		}
		rows.Scan(
			&scId,
			&reviewInfo.Count,
			&reviewInfo.Rating,
		)

		ScCardRepository.mapScCards[scId].ReviewInfo = reviewInfo
	}

	rows, _ = conn.Query(context.Background(), GetUndergrounds, locationIds)

	for rows.Next() {
		var locId int
		underground := &entity.Underground{
			Label: "",
		}

		rows.Scan(&underground.Label, &locId)

		ScCardRepository.mapScCards[locationIdsMap[locId]].Sc.Location.Undergrounds = append(
			ScCardRepository.mapScCards[locationIdsMap[locId]].Sc.Location.Undergrounds,
			underground,
		)
	}

	from := ScCardRepository.scCards[0].Sc.Id
	to := ScCardRepository.scCards[len(ScCardRepository.scCards)-1].Sc.Id

	var previous bool
	var next bool

	conn.QueryRow(context.Background(), PreviousNextQuery, from, to).Scan(&previous, &next)

	defer conn.Close(context.Background())

	return nil, ScCardRepository.scCards, map[string]bool{
		"previous": previous,
		"next":     next,
	}
}

func (ScCardRepository *ScCardRepository) Search(page int, query string) (error, []*aggregate.ScCard, map[string]bool) {
	const COUNT = 10

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return err, nil, nil
	}

	var locationIds []int

	rows, _ := conn.Query(context.Background(), FindLocationIds, query)

	for rows.Next() {
		var locationId int
		rows.Scan(&locationId)

		locationIds = append(locationIds, locationId)
	}

	previous := true
	next := false

	defer conn.Close(context.Background())

	return nil, ScCardRepository.scCards, map[string]bool{
		"previous": previous,
		"next":     next,
	}
}

func (ScCardRepository *ScCardRepository) Add() {

}

func (ScCardRepository *ScCardRepository) Get() {

}
