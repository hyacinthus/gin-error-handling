package main

import (
	"errors"
	"sync"
)

// We use sync.Map to fake a data store service like mysql or mongodb
var store sync.Map

// data store errors
var (
	ErrNotFound = errors.New("record not found")
	ErrUnknown  = errors.New("unknown server error")
	ErrIDExists = errors.New("the id already exists in db")
)

// DemoData just demo data
type DemoData struct {
	ID     int    `json:"id"`
	UserID int    `json:"-"` // user id is hidden in request and response
	Data   string `json:"data"`
}

// FindDemoByID find demo data from data store
func FindDemoByID(id int) (*DemoData, error) {
	data, ok := store.Load(id)
	if !ok {
		return nil, ErrNotFound
	}
	resp, ok := data.(DemoData)
	if !ok {
		return nil, ErrUnknown
	}
	return &resp, nil
}

// Save save demo data
func (d *DemoData) Save() error {
	_, ok := store.Load(d.ID)
	if ok {
		return ErrIDExists
	}
	store.Store(d.ID, *d)
	return nil
}
