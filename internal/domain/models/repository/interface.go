package repository

import (
	"context"
	"routergpt/internal/domain/models/model"
)

type ModelsRepository interface {
	Create(ctx context.Context, aiModel *model.AIModel) error
	GetByName(ctx context.Context, name string) (*model.AIModel, error)
}