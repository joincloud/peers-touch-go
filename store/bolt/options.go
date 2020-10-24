package bolt

import (
	"context"

	"github.com/joincloud/peers-touch-go/store"
)

type buckets struct{}

func CreateBuckets(bucketNames []string) store.Option {
	return func(o *store.Options) {
		o.Context = context.WithValue(o.Context, buckets{}, bucketNames)
	}
}
