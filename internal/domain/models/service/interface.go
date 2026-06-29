package service

import (
	"context"
	"routergpt/internal/domain/models/model"
)

type ModelsService interface {
	CreateModel(ctx context.Context, req model.CreateAIModelRequest) (*model.AIModel, error)
}