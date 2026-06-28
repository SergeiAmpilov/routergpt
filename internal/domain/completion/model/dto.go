package model

type CompletionRequestDTO struct {
	Model    string       `json:"model"`
	Messages []MessageDTO `json:"messages"`
}

type MessageDTO struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
