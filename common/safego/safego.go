package safego

import (
	"fmt"
	"os"
)

type Handler func(err interface{})

var DefaultHandler = func(err interface{}) {
	fmt.Fprintf(os.Stderr, "recovered: %s\n%s", err, CallStack(3))
}

// Go run the f with a goroutine and keep it away from panic.
// it will use DefaultHandler if argument handler is nil
func Go(f func(), handler ...Handler) {
	handle := DefaultHandler
	switch len(handler) {
	case 1:
		handle = handler[0]
	}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				handle(r)
			}
		}()

		f()
	}()
}
