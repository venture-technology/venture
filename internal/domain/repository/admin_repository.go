package repository

import (
	"context"
)

type IAdminRepository interface {
	NewApiKey(ctx context.Context, id, name string) error
	GetApiKey(ctx context.Context, id string) (string, error)
}
