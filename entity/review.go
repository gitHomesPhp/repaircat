package entity

type Review struct {
	id     int
	text   string
	rating int
	sc     *Sc
}
