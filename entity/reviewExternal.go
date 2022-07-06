package entity

const gis2 = 1

type ReviewExternal struct {
	id     int
	text   string
	rating int
	sc     *Sc
}
