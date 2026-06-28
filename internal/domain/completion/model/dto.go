package model

import "github.com/go-ozzo/ozzo-validation/v4"

type CompletionRequestDTO struct {
	Model    string       `json:"model"`
	Messages []MessageDTO `json:"messages"`
}

type MessageDTO struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func (r CompletionRequestDTO) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Model, validation.Required),
		validation.Field(&r.Messages, validation.Required, validation.Length(1, 0)),
	)
}