package server

import (
	"sync"
	"testing"
	"time"
)

func TestStarter(t *testing.T) {

	var wg sync.WaitGroup

	wg.Add(1)

	model := Storage{
		RequestId: "3123",
		Type:      "start",
	}

	ch := make(chan error, 1)

	go func() {
		ch <- starter(model)
	}()

	if <-ch != nil {
		t.Error(<-ch)
	}
	close(ch)

	if len(memStorage) > 0 {
		t.Errorf("start does not add")
	}
}

func TestStop(t *testing.T) {

	var wg sync.WaitGroup

	wg.Add(2)

	model := Storage{
		RequestId: "3123",
		Type:      "start",
	}

	go starter(model)
	go stopper(model)
}

func TestRepeat(t *testing.T) {

	var wg sync.WaitGroup

	wg.Add(2)

	model := Storage{
		RequestId: "3123",
		Type:      "start",
	}

	err := make(chan error, 1)
	go starter(model)

	time.Sleep(2 * time.Second)

	go func() {
		err <- starter(model)
	}()

	if <-err == nil {
		t.Errorf("repeat not work")
	}
}

func TestChange(t *testing.T) {

	var wg sync.WaitGroup

	wg.Add(1)

	model := Storage{
		RequestId: "2312",
		Type:      "start",
		Position:  0,
	}

	go starter(model)
	time.Sleep(3 * time.Second)

	if memStorage[model.RequestId].Position == 0 {
		t.Errorf("position doesn't change")
	}
}

func TestStopAfter6(t *testing.T) {

	var wg sync.WaitGroup

	wg.Add(1)

	model := Storage{
		RequestId: "3123",
		Type:      "stop",
		Position:  6,
	}

	err := make(chan error)

	go func() {
		err <- stopper(model)
	}()

	time.Sleep(2 * time.Second)

	ers := <-err
	close(err)

	if ers == nil {
		t.Errorf("stop after 6")
	}
}

func TestStopTo6(t *testing.T) {

	var wg sync.WaitGroup

	wg.Add(1)

	model := Storage{
		RequestId: "32132",
		Type:      "stop",
		Position:  3,
	}

	memStorage[model.RequestId] = model

	err := make(chan error)

	go func() {
		err <- stopper(model)
	}()

	time.Sleep(1 * time.Second)

	errs := <-err
	close(err)

	if errs != nil {
		t.Error(errs)
	}

	if model.Position != 3 {
		t.Errorf("position changed")
	}
}
