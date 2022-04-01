package server

import (
	"errors"
	"time"
)

const (
	errorRepeat   = "элемент повторялся"
	moreThan6     = "элемент уже прошел 6 шаг"
	errorNotFound = "элемент не существует"
	timeOut       = 10
	indexOf6      = 6
	delay         = 1
)

func starter(model Storage) error {

	if previousItem == model.RequestId {
		return errors.New(errorRepeat)
	}

	go func() {

		for model.Position <= timeOut {

			if memStorage[model.RequestId].Ping && memStorage[model.RequestId].Type == stop {
				return
			}

			if exist := checkInCache(model); exist {
				model.Position = positionCache[model.RequestId]
				deleteFromCache(model)
			}

			model.Position += delay
			run(model)
			time.Sleep(1 * time.Second)
		}

	}()

	previousItem = model.RequestId
	return nil
}

func stopper(model Storage) error {

	if value, ok := memStorage[model.RequestId]; ok {

		if value.Position < indexOf6 {
			wait(model, value.Position)
			addToCache(model, value.Position)
			return nil
		}

		return errors.New(moreThan6)
	}

	return errors.New(errorNotFound)
}
