package model

import "github.com/go-ozzo/ozzo-validation/v4"

type CreateAIModelRequest struct {
	Name     string `json:"name"`
	Provider string `json:"provider"`
	Version  string `json:"version"`
}

func (req CreateAIModelRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required, validation.Length(1, 100)),
		validation.Field(&req.Provider, validation.Required, validation.Length(1, 100)),
		validation.Field(&req.Version, validation.Required, validation.Length(1, 50)),
	)
}