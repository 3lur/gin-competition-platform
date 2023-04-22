package cache

import "sync"

type Store struct {
}

type CService struct {
	Store Store
}

var once sync.Once
var Cache *CService
