package service

import "atlant1da-404/server/storage"

func (s Service) Update() {

	s.model.Ping = false
	s.model.Position = 0

	storage.UpdateItem(s.model)
}
