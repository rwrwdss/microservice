package repository

import (
	"context"
	"fmt"

	"github.com/rwrwdss/microservices-cache-app/services/gateway/internal/model"
	"gorm.io/gorm"
)

type keywordRow struct {
	Text     string `gorm:"column:text"`
	Requests int    `gorm:"column:requests"`
}

func (keywordRow) TableName() string {
	return "keywords"
}

type postgresKeywordRepository struct {
	db *gorm.DB
}

func NewPostgresKeywordRepository(db *gorm.DB) KeywordRepository {
	return &postgresKeywordRepository{db: db}
}

func (r *postgresKeywordRepository) LoadTopKeywords(ctx context.Context, limit int) ([]model.Keyword, error) {
	var rows []keywordRow

	err := r.db.WithContext(ctx).
		Order("requests DESC").
		Limit(limit).
		Find(&rows).Error
	if err != nil {
		return nil, fmt.Errorf("failed to load top keywords: %w", err)
	}

	keywords := make([]model.Keyword, 0, len(rows))
	for _, row := range rows {
		keywords = append(keywords, model.Keyword{
			Text:     row.Text,
			Requests: row.Requests,
		})
	}

	return keywords, nil
}
