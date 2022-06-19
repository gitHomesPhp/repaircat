package entity

type Review struct {
	id       int
	author   string
	common   string
	positive string
	negative string
	star     int
	sc       *Sc
}
