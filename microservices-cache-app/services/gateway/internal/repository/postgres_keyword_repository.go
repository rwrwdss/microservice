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

func Migrate(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).AutoMigrate(&keywordRow{})
}

func Seed(ctx context.Context, db *gorm.DB) error {
	var count int64
	if err := db.WithContext(ctx).Model(&keywordRow{}).Count(&count).Error; err != nil {
		return fmt.Errorf("failed to count keywords: %w", err)
	}

	if count > 0 {
		return nil
	}

	seedRows := []keywordRow{
		{Text: "iphone", Requests: 500},
		{Text: "ipad", Requests: 400},
		{Text: "ipod", Requests: 300},
		{Text: "imac", Requests: 200},
		{Text: "intel", Requests: 100},
	}

	if err := db.WithContext(ctx).Create(&seedRows).Error; err != nil {
		return fmt.Errorf("failed to seed keywords: %w", err)
	}

	return nil
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
