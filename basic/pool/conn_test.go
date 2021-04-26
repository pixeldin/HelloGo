package pool

import (
	"fmt"
	"sync"
	"testing"
)

func TestSome(t *testing.T) {
	var sm = new(sync.Map)
	sm.Store("hello", "world")
	if sm != nil {
		sm.Range(func(key, value interface{}) bool {
			fmt.Println(key)
			fmt.Println(value)
			return true
		})
	}
}
