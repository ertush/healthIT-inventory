package util

import (
	"time"

	"github.com/patrickmn/go-cache"
)

func New() *cache.Cache {

	var c *cache.Cache = cache.New(24*time.Hour, 24*time.Hour)

	return c
}
