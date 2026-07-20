package service

import (
	"context"
	"log"

	"github.com/rwrwdss/microservices-cache-app/services/gateway/internal/repository"
)

const topKeywordsLimit = 100000

type DictionaryService interface {
	LoadDictionary(ctx context.Context) error
}

type dictionaryService struct {
	repository repository.KeywordRepository
}

func NewDictionaryService(repository repository.KeywordRepository) DictionaryService {
	return &dictionaryService{repository: repository}
}

func (s *dictionaryService) LoadDictionary(ctx context.Context) error {
	keywords, err := s.repository.LoadTopKeywords(ctx, topKeywordsLimit)
	if err != nil {
		return err
	}

	log.Printf("INFO Loaded %d keywords", len(keywords))

	return nil
}
