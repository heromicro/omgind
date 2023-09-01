package tipes

import (
	"strings"
)

type CacheKey string

func (c CacheKey) Get(suffix string) CacheKey {
	var name strings.Builder
	delim := ":"

	name.WriteString(string(c))
	name.WriteString(delim)
	name.WriteString(suffix)

	return CacheKey(name.String())
}

func (c CacheKey) String() string {
	return string(c)
}

const (
	EndpointsCacheKey CacheKey = "endpoints"
	ProjectsCacheKey  CacheKey = "projects"
	TokenCacheKey     CacheKey = "tokens"
	SourceCacheKey    CacheKey = "sources"
)
