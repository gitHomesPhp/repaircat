package entity

const gis2 = 1

type ReviewExternal struct {
	id        int
	Text      string
	Rating    int
	VisitorId int
	Sc        *Sc
}
