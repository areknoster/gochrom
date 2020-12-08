package gochrom

import "github.com/areknoster/gochrom/pkg/spectrum"

type State struct {
    Spectrum *spectrum.Data
    IsFramed bool
}

type StateStorage interface {
    Get() State
    Set(State)
}
