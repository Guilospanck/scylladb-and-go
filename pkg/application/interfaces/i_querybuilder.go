package interfaces

type IQueryBuilder interface {
	SelectAll() (any, error)
}
