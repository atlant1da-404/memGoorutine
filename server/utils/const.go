package utils

import (
	"encoding/json"
	"net/http"
)

const (
	ErrRepeat      = "элемент повторялся"
	ErrMoreThanSix = "элемент уже прошел 6 шаг"
	ErrNotFound    = "элемент не существует"
)

const (
	Stop          = "stop"
	Start         = "start"
	ReqId         = "request_id"
	TimeOut       = 10
	IndexOfSix    = 6
	Delay         = 1
	SizeOfStorage = 500
	ChSize        = 100
)

const (
	Hardcode  = "localhost:8000"
	CreateUrl = "/create"
	ListUrl   = "/list"
	DeleteUrl = "/delete"
	UpdateUrl = "/update"
	OK        = "OK!"
)

func HttpError(w http.ResponseWriter, err error) {
	_ = json.NewEncoder(w).Encode(err.Error())
}
