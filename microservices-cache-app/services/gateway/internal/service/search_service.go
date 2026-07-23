package service

import (
	"context"

	"github.com/rwrwdss/microservices-cache-app/services/gateway/internal/searchclient"
)

const defaultSearchLimit = 10

type SearchService interface {
	Search(ctx context.Context, query string) ([]string, error)
}

type searchService struct {
	client searchclient.SearchClient
}

func NewSearchService(client searchclient.SearchClient) SearchService {
	return &searchService{client: client}
}

func (s *searchService) Search(ctx context.Context, query string) ([]string, error) {
	return s.client.Search(ctx, query, defaultSearchLimit)
}
