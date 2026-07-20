package service

type SearchService interface {
	Search(query string) ([]string, error)
}

type searchService struct{}

func NewSearchService() SearchService {
	return &searchService{}
}

func (s *searchService) Search(query string) ([]string, error) {
	return []string{}, nil
}
