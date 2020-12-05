package plane

import (
    "github.com/areknoster/gochrom/pkg/geom2d"
    "github.com/areknoster/gochrom/pkg/gochrom"
)

type NullMode struct{}

func NewNullMode() *NullMode {
    return &NullMode{}
}

var _ gochrom.PlaneMode = &NullMode{}

func (n NullMode) HandleClick(normLoc geom2d.Point) {}

func (n NullMode) HandleDrag(start geom2d.Point, move geom2d.Vector) {}

func (n NullMode) HandleDragEnd() {}

func (n NullMode) Name() string {
    return "Null mode"
}

type SPDEditMode struct {
    storage gochrom.StateStorage
}

func NewSPDEditMode(storage gochrom.StateStorage) *SPDEditMode {
    return &SPDEditMode{storage: storage}
}

var _ gochrom.PlaneMode = &SPDEditMode{}

func (S SPDEditMode) HandleClick(normLoc geom2d.Point) {
    //todo implement
}

func (S SPDEditMode) HandleDrag(start geom2d.Point, move geom2d.Vector) {
    //todo implement
}

func (S SPDEditMode) HandleDragEnd() {
    //todo implement
}

func (S SPDEditMode) Name() string {
    return "SPDEditMode"
}
