package service

import "log/slog"

type Service struct {
	repo   Repositorier
	logger *slog.Logger
}

func NewService(repo Repositorier, logger *slog.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}
