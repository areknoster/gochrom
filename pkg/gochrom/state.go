package gochrom

type State struct {
}

type StateStorage interface {
    Get() State
    Set(State)
}
