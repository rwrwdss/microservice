package repository

import (
	"context"

	"github.com/rwrwdss/microservices-cache-app/services/gateway/internal/model"
)

type KeywordRepository interface {
	LoadTopKeywords(ctx context.Context, limit int) ([]model.Keyword, error)
}
