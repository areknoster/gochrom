package state

import (
	"github.com/areknoster/gochrom/pkg/gochrom"
	"sync"
)

type StateStorage struct {
	state   gochrom.State
	stateMx *sync.Mutex
	refreshers []func()
}

var _ gochrom.StateStorage = &StateStorage{}

func NewStateStorage() *StateStorage {
	state := gochrom.State{}

	ss := &StateStorage{
		stateMx: &sync.Mutex{},
	}
	ss.state = state
	return ss
}

func (ss *StateStorage) Get() gochrom.State {
	ss.stateMx.Lock()
	defer ss.stateMx.Unlock()
	return ss.state
}

func (ss *StateStorage) Set(state gochrom.State) {
	ss.stateMx.Lock()
	ss.state = state
	ss.stateMx.Unlock()
	ss.refresh()
}

func (ss *StateStorage) refresh(){
	for _, refresher := range ss.refreshers {
		refresher()
	}
}

func (ss *StateStorage)AddRefresh(ref func()){
	ss.refreshers = append(ss.refreshers, ref)
}