package entity

type ReviewExternal struct {
	Id        int    `json:"-"`
	Text      string `json:"text"`
	Rating    int    `json:"rating"`
	VisitorId int    `json:"-"`
	ScId      int    `json:"-"`
}
