package interfaces

type T any

type IQueryBuilder[t T] interface {
	Insert(insertData *t) error
	Delete(dataToBeDeleted *t) error
	SelectAll() ([]t, error)
}
