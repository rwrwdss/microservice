package builder

import (
	"context"

	"github.com/rwrwdss/microservices-cache-app/services/gateway/internal/model"
)

type DictionaryBuilder interface {
	Build(ctx context.Context, keywords []model.Keyword) error
}
