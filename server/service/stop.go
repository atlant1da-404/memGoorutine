package service

import (
	"atlant1da-404/server/storage"
	"atlant1da-404/server/utils"
	"errors"
)

func (s Service) Stop() error {

	if value, ok := storage.MemStorage[s.model.RequestId]; ok {

		if value.Position < utils.IndexOfSix {
			storage.Wait(s.model, value.Position)
			storage.AddToCache(s.model, value.Position)
			return nil
		}

		return errors.New(utils.ErrMoreThanSix)
	}

	return errors.New(utils.ErrNotFound)
}
