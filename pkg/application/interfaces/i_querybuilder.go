package interfaces

import "context"

type T any

type IQueryBuilder[t T] interface {
	Insert(ctx context.Context, insertData *t) error
	Delete(ctx context.Context, dataToBeDeleted *t) error
	DeleteAllFromPartitioningKey(ctx context.Context, dataToBeDeleted *t) error
	Select(ctx context.Context, dataToGet *t) ([]t, error)
	Get(ctx context.Context, dataToGet *t) (*t, error)
	SelectAll(ctx context.Context) ([]t, error)
}
