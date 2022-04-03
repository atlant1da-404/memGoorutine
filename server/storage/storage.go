package storage

import "atlant1da-404/server/utils"

type Storage struct {
	RequestId string `json:"request_id"`
	Type      string `json:"type"`
	Position  int    `json:"position"`
	Ping      bool   `json:"-"`
}

var (
	MemStorage        = make(map[string]Storage, utils.SizeOfStorage)
	PositionCache     = make(map[string]int, utils.SizeOfStorage)
	PreviousRequestId string
	PreviousType      string
)

func Run(model Storage) {
	MemStorage[model.RequestId] = Storage{model.RequestId, model.Type, model.Position, false}
}

func Wait(model Storage, position int) {
	MemStorage[model.RequestId] = Storage{model.RequestId, model.Type, position, true}
}

func AddToCache(model Storage, position int) {
	PositionCache[model.RequestId] = position
}

func CheckInCache(model Storage) bool {

	if _, ok := PositionCache[model.RequestId]; ok {
		return true
	}

	return false
}

func DeleteFromCache(model Storage) {
	delete(PositionCache, model.RequestId)
}

func CheckInStorage(model Storage) bool {

	if _, ok := MemStorage[model.RequestId]; ok {
		return true
	}

	return false
}

func UpdateItem(model Storage) {
	MemStorage[model.RequestId] = model
	delete(PositionCache, model.RequestId)
}
