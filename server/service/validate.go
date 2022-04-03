package service

import "atlant1da-404/server/storage"

func (s Service) CheckInStorage() bool {

	if !storage.CheckInStorage(s.model) {
		return false
	}

	return true
}
