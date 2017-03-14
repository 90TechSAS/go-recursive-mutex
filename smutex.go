package smutex

import (
	"sync"
	"time"

	"github.com/huandu/goroutine"
)

type SmartMutex struct {
	mutex            sync.Mutex
	internalMutex    sync.Mutex
	currentGoRoutine int64
	lockCount        uint64
}

func (s *SmartMutex) Lock() {
	goRoutineID := goroutine.GoroutineId()

	for {
		s.internalMutex.Lock()
		if s.currentGoRoutine == 0 {
			s.currentGoRoutine = goRoutineID
			break
		} else if s.currentGoRoutine == goRoutineID {
			break
		} else {
			s.internalMutex.Unlock()
			time.Sleep(time.Millisecond)
			continue
		}
	}
	s.lockCount++
	s.internalMutex.Unlock()
}

func (s *SmartMutex) Unlock() {
	s.internalMutex.Lock()
	s.lockCount--
	if s.lockCount == 0 {
		s.currentGoRoutine = 0
	}
	s.internalMutex.Unlock()
}
