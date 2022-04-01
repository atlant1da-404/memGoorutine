package server

import (
	"testing"
)

func TestAddToCache(t *testing.T) {

	model := Storage{
		RequestId: "3123",
		Position:  5,
		Type:      "start",
	}

	addToCache(model, model.Position)

	if len(positionCache) == 0 {
		t.Errorf("cache position is not added")
	}
}

func TestRun(t *testing.T) {

	model := Storage{
		RequestId: "3123",
		Ping:      true,
	}

	run(model)

	if memStorage[model.RequestId].Ping {
		t.Errorf("run is not continue")
	}
}

func TestWait(t *testing.T) {

	model := Storage{
		RequestId: "3123",
		Ping:      false,
	}

	wait(model, 5)

	if !memStorage[model.RequestId].Ping {
		t.Errorf("wait is not stop")
	}
}
