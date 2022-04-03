package storage

import (
	"testing"
)

var (
	testCaseOne = Storage{
		RequestId: "31n1",
		Position:  5,
		Type:      "start",
	}
	testCaseTwo = Storage{
		RequestId: "1nm1mn",
		Ping:      true,
	}
	testCaseThree = Storage{
		RequestId: "31n1mn2",
		Ping:      false,
	}
)

const (
	errPositionCacheEmpty = "position cache empty"
	errRunNotContinue     = "run is not continue"
	errWaitNotStop        = "wait is not stop"
)

const (
	positionCacheEmpty = 0
)

func TestAddToCache(t *testing.T) {

	model := testCaseOne

	AddToCache(model, model.Position)

	if len(PositionCache) == positionCacheEmpty {
		t.Errorf(errPositionCacheEmpty)
	}
}

func TestRun(t *testing.T) {

	model := testCaseTwo

	Run(model)

	if MemStorage[model.RequestId].Ping {
		t.Errorf(errRunNotContinue)
	}
}

func TestWait(t *testing.T) {

	model := testCaseThree

	Wait(model, 5)

	if !MemStorage[model.RequestId].Ping {
		t.Errorf(errWaitNotStop)
	}
}
