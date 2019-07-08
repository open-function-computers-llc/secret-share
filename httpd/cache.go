package main

import (
	"time"

	"github.com/bluele/gcache"
)

var cache gcache.Cache

func initCache() {
	cache = gcache.New(100).LRU().Build()

	currentTime := time.Now().Format("2006")

	cache.Set("startup", currentTime)
}
