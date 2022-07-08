package scCard

type Repository interface {
	Get()
	Add()
	List(page int)
}
