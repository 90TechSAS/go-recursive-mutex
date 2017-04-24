# go-recursive-mutex (recmutex)
Recmutex is a tiny mutex lib working with goroutine's id for recursive mutex

## Installation

```bash
go get -u github.com/huandu/goroutine
go get -u github.com/90TechSAS/go-recursive-mutex
```

## Recursive locking with native mutex

This is a simple example of recursive locking with native mutex

```go
package main

import (
	"sync"
)

type Struct struct{ sync.Mutex }

func main() {
	var s Struct
	s.A()
}

func (s *Struct) A() {
	s.Lock()
	s.B()
	s.Unlock()
}

func (s *Struct) B() {
	s.Lock()
	s.Unlock()
}
```

Obviously, this code doesn't work and make a beautiful error: ```fatal error: all goroutines are asleep - deadlock!```

## Multi locking with recmutex

Recmutex allow you to do recursive mutex with Go!

```go
package main

import (
	"github.com/90TechSAS/go-recursive-mutex"
)

type Struct struct{ recmutex.RecursiveMutex }

func main() {
	var s Struct
	s.A()
}

func (s *Struct) A() {
	s.Lock()
	s.B()
	s.Unlock()
}

func (s *Struct) B() {
	s.Lock()
	s.Unlock()
}
```

## Benchmarking

```bash
$ go run examples/benchmark_narive-vs-recursive.go 
[Native Mutex] Elapsed: 4.374499979s / 2285975 locks/sec / counter: 10000000
[Recursive Mutex] Elapsed: 3.854341015s / 2594477 locks/sec / counter: 10000000
```