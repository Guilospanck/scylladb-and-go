package interfaces

type T any

type IQueryBuilder[t T] interface {
	SelectAll() ([]t, error)
}
