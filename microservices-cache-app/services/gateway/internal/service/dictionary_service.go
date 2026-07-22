package service

import (
	"context"
	"fmt"
	"log"

	"github.com/rwrwdss/microservices-cache-app/services/gateway/internal/builder"
	"github.com/rwrwdss/microservices-cache-app/services/gateway/internal/repository"
)

const topKeywordsLimit = 100000

type DictionaryService interface {
	LoadDictionary(ctx context.Context) error
}

type dictionaryService struct {
	repository repository.KeywordRepository
	builder    builder.DictionaryBuilder
}

func NewDictionaryService(repository repository.KeywordRepository, builder builder.DictionaryBuilder) DictionaryService {
	return &dictionaryService{repository: repository, builder: builder}
}

func (s *dictionaryService) LoadDictionary(ctx context.Context) error {
	keywords, err := s.repository.LoadTopKeywords(ctx, topKeywordsLimit)
	if err != nil {
		return fmt.Errorf("failed to load top keywords: %w", err)
	}

	if err := s.builder.Build(ctx, keywords); err != nil {
		return fmt.Errorf("failed to build dictionary: %w", err)
	}

	log.Printf("INFO Loaded %d keywords", len(keywords))

	return nil
}
