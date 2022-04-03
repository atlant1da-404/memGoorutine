package service

import (
	"atlant1da-404/server/storage"
	"atlant1da-404/server/utils"
	"errors"
	"time"
)

func Add(model storage.Storage) error {

	if storage.PreviousRequestId == model.RequestId && storage.PreviousType == model.Type {
		return errors.New(utils.ErrRepeat)
	}

	go func() {

		for model.Position <= utils.TimeOut {

			if storage.MemStorage[model.RequestId].Ping && storage.MemStorage[model.RequestId].Type == utils.Stop {
				break
			}

			if exist := storage.CheckInCache(model); exist {
				model.Position = storage.PositionCache[model.RequestId]
				storage.DeleteFromCache(model)
			}

			model.Position += utils.Delay
			storage.Run(model)
			time.Sleep(1 * time.Second)
		}

	}()

	storage.PreviousRequestId = model.RequestId
	storage.PreviousType = model.Type
	return nil
}

func Stop(model storage.Storage) error {

	if value, ok := storage.MemStorage[model.RequestId]; ok {

		if value.Position < utils.IndexOfSix {
			storage.Wait(model, value.Position)
			storage.AddToCache(model, value.Position)
			return nil
		}

		return errors.New(utils.ErrMoreThanSix)
	}

	return errors.New(utils.ErrNotFound)
}

func CheckInStorage(model storage.Storage) bool {

	if !storage.CheckInStorage(model) {
		return false
	}

	return true
}

func Update(model storage.Storage) {

	model.Ping = false
	model.Position = 0

	storage.UpdateItem(model)
}