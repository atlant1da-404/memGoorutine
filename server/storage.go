package server

type Storage struct {
	RequestId string `json:"request_id"`
	Type      string `json:"type"`
	Position  int    `json:"position"`
	Ping      bool   `json:"-"`
}

const (
	sizeOfStorage = 500
)

var (
	memStorage    = make(map[string]Storage, sizeOfStorage)
	positionCache = make(map[string]int, sizeOfStorage)
	previousItem  string
)

func run(model Storage) {
	memStorage[model.RequestId] = Storage{model.RequestId, model.Type, model.Position, false}
}

func wait(model Storage, position int) {
	memStorage[model.RequestId] = Storage{model.RequestId, model.Type, position, true}
}

func addToCache(model Storage, position int) {
	positionCache[model.RequestId] = position
}

func checkInCache(model Storage) bool {

	if _, ok := positionCache[model.RequestId]; ok {
		return true
	}

	return false
}

func deleteFromCache(model Storage) {
	delete(positionCache, model.RequestId)
}
