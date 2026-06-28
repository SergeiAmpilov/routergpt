package config

type Config interface {
	GetConfig() *AppConfig
}