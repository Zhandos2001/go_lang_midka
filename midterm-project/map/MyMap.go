package my_map

import (
	"sync"
)

type MyMap struct {
	mx sync.Mutex
	m map[string]string
}

// Method map with initial test data
func MakeMap() *MyMap {
	return &MyMap{
		m: map[string] string {
			"year": "2022",
			"month": "October",
			"day": "24",
		},
	}
}

// Method for getting value from map by it's key
func (r *MyMap) Get(key string) (string, bool) {
	r.mx.Lock()
	defer r.mx.Unlock()
	value, ok := r.m[key]
	return value, ok
}

// Method for setting key: value pair
func (r *MyMap) Set(key string, val string) {
	r.mx.Lock()
	defer r.mx.Unlock()
	r.m[key] = val
}

// Map that we will use
var Map = MakeMap()
