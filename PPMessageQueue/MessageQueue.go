package PPMessageQueue

import (
	"errors"
	"sync"
	"time"
)

type MessageQueue struct {
	L        sync.Mutex
	Datas    chan interface{}
	WaitTime time.Duration
	Capacity int
}

func NewPPQueue(size int, t time.Duration) *MessageQueue {

	return &MessageQueue{
		L:        sync.Mutex{},
		Datas:    make(chan interface{}, size),
		WaitTime: t,
		Capacity: size,
	}

}

func (ppq *MessageQueue) push(data interface{}) error {
	go ppq.PushData(data)
	return nil
}

func (ppq *MessageQueue) getCapacity() int {
	return ppq.Capacity
}

func (ppq *MessageQueue) Pop() interface{} {

	if ppq.getSize() == 0 {
		return nil
	}

	return <-ppq.Datas
}

func (ppq *MessageQueue) Lock() error {

	err := ppq.L.Lock

	if err != nil {
		return errors.New("Lock failed")
	}
	return nil
}

func (ppq *MessageQueue) Unlock() error {

	err := ppq.L.Unlock

	if err != nil {
		return errors.New("UNLock failed")
	}
	return nil

}

func (ppq *MessageQueue) getSize() int {
	return len(ppq.Datas)
}

func (ppq *MessageQueue) PushData(data interface{}) error {

	if ppq.getSize() == ppq.getCapacity() {
		return errors.New("Memory limit exceeded")
	}

	lockerr := ppq.Lock()
	ppq.Datas <- data
	unlockerr := ppq.Unlock()

	if lockerr != nil || unlockerr != nil {
		return errors.New("Push error")
	}

	return nil

}

func (ppq *MessageQueue) PopAllData() ([]interface{}, error) {

	ans := make([]interface{}, 0)

	for {

		select {

		case msg := <-ppq.Datas:
			ans = append(ans, msg)
		case <-time.After(ppq.WaitTime):
			return ans, errors.New("Time Out")
		default:
			if ppq.getSize() == 0 {
				return ans, nil
			}

		}

	}

	return ans, nil

}
