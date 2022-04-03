package service

import (
	"atlant1da-404/server/storage"
	"sync"
	"testing"
	"time"
)

var (
	testCaseOne = storage.Storage{
		RequestId: "3123",
		Type:      "start",
	}

	testCaseTwo = storage.Storage{
		RequestId: "3513",
		Type:      "stop",
	}

	testCaseThree = storage.Storage{
		RequestId: "312jhb1",
		Type:      "start",
	}

	testCaseFour = storage.Storage{
		RequestId: "312bn12nb1n2",
		Type:      "start",
		Position:  0,
	}

	testCaseFive = storage.Storage{
		RequestId: "132n1mnmn1112",
		Type:      "stop",
		Position:  6,
	}

	testCaseSix = storage.Storage{
		RequestId: "1nm2mn1mnm1",
		Type:      "start",
		Position:  randomNumber,
	}
)

const (
	startNotCorrect        = "start is not work correctly"
	stopNotCorrect         = "stop is not work correctly"
	repeatNotCorrect       = "repeat is not work correctly"
	positionNotCorrect     = "position doesn't change"
	stopAfterSixNotCorrect = "stop after 6"
	stopToSixNotCorrect    = "stop to 6 not work correctly"
)

const (
	memStorageFull  = 10
	sizeOfChan      = 1
	memStorageEmpty = 0
	zeroPosition    = 0
	randomNumber    = 3
)

func TestStarter(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(1)

	model := testCaseOne

	ch := make(chan error, sizeOfChan)

	go func() {
		ch <- Add(model)
	}()

	if <-ch != nil {
		t.Error(<-ch)
	}
	close(ch)

	if len(storage.MemStorage) > memStorageEmpty {
		t.Errorf(startNotCorrect)
	}
}

func TestStop(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(2)

	model := testCaseTwo

	go Add(model)
	go Stop(model)

	if model.Position == memStorageFull {
		t.Errorf(stopNotCorrect)
	}
}

func TestRepeat(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(2)

	model := testCaseThree

	chErr := make(chan error, sizeOfChan)
	go Add(model)

	time.Sleep(3 * time.Second)

	go func() {
		chErr <- Add(model)
	}()

	err := <-chErr

	if err == nil {
		t.Errorf(repeatNotCorrect)
	}
}

func TestChange(t *testing.T) {

	var wg sync.WaitGroup

	wg.Add(1)

	model := testCaseFour

	go Add(model)
	time.Sleep(3 * time.Second)

	if storage.MemStorage[model.RequestId].Position == zeroPosition {
		t.Errorf(positionNotCorrect)
	}
}

func TestStopAfter6(t *testing.T) {

	var wg sync.WaitGroup

	wg.Add(1)

	model := testCaseFive

	err := make(chan error)

	go func() {
		err <- Stop(model)
	}()

	time.Sleep(2 * time.Second)

	ers := <-err
	close(err)

	if ers == nil {
		t.Errorf(stopAfterSixNotCorrect)
	}
}

func TestStopTo6(t *testing.T) {

	var wg sync.WaitGroup

	wg.Add(1)

	model := testCaseSix

	storage.MemStorage[model.RequestId] = model

	err := make(chan error)

	go func() {
		err <- Stop(model)
	}()

	time.Sleep(3 * time.Second)

	errs := <-err
	close(err)

	if errs != nil {
		t.Error(errs)
	}

	if model.Position != randomNumber {
		t.Errorf(stopToSixNotCorrect)
	}
}
