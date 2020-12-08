package state

import (
	"github.com/areknoster/gochrom/pkg/gochrom"
	"github.com/areknoster/gochrom/pkg/spectrum"
	"sync"
)

type StateStorage struct {
	state   gochrom.State
	stateMx *sync.Mutex
	refreshers []func()
}

var _ gochrom.StateStorage = &StateStorage{}

func NewStateStorage() *StateStorage {
	state := gochrom.State{
		Spectrum: spectrum.NewData(),
	}

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
	ss.Refresh()
}

func (ss *StateStorage) Refresh(){
	for _, refresher := range ss.refreshers {
		refresher()
	}
}

func (ss *StateStorage) AddRefresh(ref func()){
	ss.refreshers = append(ss.refreshers, ref)
}