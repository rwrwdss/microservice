package searchclient

import "context"

type SearchClient interface {
	Search(ctx context.Context, query string, limit int) ([]string, error)
}
