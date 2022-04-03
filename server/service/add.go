package service

import (
	"atlant1da-404/server/storage"
	"atlant1da-404/server/utils"
	"errors"
	"time"
)

func (s Service) Add() error {

	if storage.PreviousRequestId == s.model.RequestId && storage.PreviousType == s.model.Type {
		return errors.New(utils.ErrRepeat)
	}

	go func() {

		for s.model.Position != utils.TimeOut {

			if storage.MemStorage[s.model.RequestId].Ping && storage.MemStorage[s.model.RequestId].Type == utils.Stop {
				break
			}

			if exist := storage.CheckInCache(s.model); exist {
				s.model.Position = storage.PositionCache[s.model.RequestId]
				storage.DeleteFromCache(s.model)
			}

			s.model.Position += utils.Delay
			storage.Run(s.model)
			time.Sleep(1 * time.Second)
		}

	}()

	storage.PreviousRequestId = s.model.RequestId
	storage.PreviousType = s.model.Type
	return nil
}
