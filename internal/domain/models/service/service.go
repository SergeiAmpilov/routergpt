package service

import (
	"context"
	"fmt"

	"routergpt/internal/domain/models/model"
	"routergpt/internal/domain/models/repository"
)

type modelsService struct {
	repository repository.ModelsRepository
}

func NewModelsService(repository repository.ModelsRepository) ModelsService {
	return &modelsService{
		repository: repository,
	}
}

func (s *modelsService) CreateModel(ctx context.Context, req model.CreateAIModelRequest) (*model.AIModel, error) {
	// Check if model with same name already exists
	existingModel, err := s.repository.GetByName(ctx, req.Name)
	if err == nil && existingModel != nil {
		return nil, fmt.Errorf("model with name %s already exists", req.Name)
	}

	aiModel := &model.AIModel{
		Name:     req.Name,
		Provider: req.Provider,
		Version:  req.Version,
	}

	err = s.repository.Create(ctx, aiModel)
	if err != nil {
		return nil, fmt.Errorf("failed to create model: %w", err)
	}

	return aiModel, nil
}