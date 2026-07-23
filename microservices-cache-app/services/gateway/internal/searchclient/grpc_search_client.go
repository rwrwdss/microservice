package searchclient

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/rwrwdss/microservices-cache-app/services/gateway/internal/searchclient/keywordspb"
)

type grpcSearchClient struct {
	client keywordspb.SearchServiceClient
}

func NewGRPCSearchClient(target string) (SearchClient, error) {
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to create search-core client: %w", err)
	}

	return &grpcSearchClient{client: keywordspb.NewSearchServiceClient(conn)}, nil
}

func (c *grpcSearchClient) Search(ctx context.Context, query string, limit int) ([]string, error) {
	response, err := c.client.Search(ctx, &keywordspb.SearchRequest{
		Query: query,
		Limit: int32(limit),
	})
	if err != nil {
		return nil, fmt.Errorf("search-core request failed: %w", err)
	}

	results := response.GetResults()
	if results == nil {
		results = []string{}
	}

	return results, nil
}
