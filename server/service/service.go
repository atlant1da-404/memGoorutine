package service

import "atlant1da-404/server/storage"

type Services interface {
	Add() error
	Stop() error
	Update()
	CheckInStorage() bool
}

type Service struct {
	model storage.Storage
	Services
}

func NewService(model storage.Storage) *Service {
	return &Service{
		model: model,
	}
}
