package builder

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/rwrwdss/microservices-cache-app/services/gateway/internal/model"
)

type textDictionaryBuilder struct {
	path string
}

func NewTextDictionaryBuilder(path string) DictionaryBuilder {
	return &textDictionaryBuilder{path: path}
}

func (b *textDictionaryBuilder) Build(ctx context.Context, keywords []model.Keyword) error {
	file, err := os.Create(b.path)
	if err != nil {
		return fmt.Errorf("failed to create dictionary file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, keyword := range keywords {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if _, err := writer.WriteString(keyword.Text + "\n"); err != nil {
			return fmt.Errorf("failed to write keyword %q: %w", keyword.Text, err)
		}
	}

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("failed to flush dictionary file: %w", err)
	}

	return nil
}
