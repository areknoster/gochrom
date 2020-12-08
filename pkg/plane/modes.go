package plane

import (
    "github.com/areknoster/gochrom/pkg/geom2d"
    "github.com/areknoster/gochrom/pkg/gochrom"
    "github.com/areknoster/gochrom/pkg/normde"
    "github.com/areknoster/gochrom/pkg/render"
    "github.com/sirupsen/logrus"
    "math"
)

type NullMode struct{}

func NewNullMode() *NullMode {
    return &NullMode{}
}

var _ gochrom.InteractiveMode = &NullMode{}

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

var _ gochrom.InteractiveMode = &SPDEditMode{}

func (s *SPDEditMode) HandleClick(normLoc geom2d.Point) {
    state := s.storage.Get()

    nc := normde.ClampNormRect(normLoc, render.GraphMin, render.GraphMax)
    state.Spectrum.Lambdas[normde.DenormalizeFloat(nc.X, 0, len(state.Spectrum.Lambdas) - 1)] = nc.Y
    s.storage.Set(state)
}

var Nan = math.NaN()

func (s *SPDEditMode) HandleDrag(start geom2d.Point, move geom2d.Vector) {
    if move.X < 0{
        start = start.MoveByVector(move)
        move = move.TimesScalar(-1)
    }
    state := s.storage.Get()
    ncs := normde.ClampNormRect(start, render.GraphMin, render.GraphMax)
    ncf := normde.ClampNormRect(start.MoveByVector(move), render.GraphMin, render.GraphMax)
    startIndex := normde.DenormalizeFloat(ncs.X, 0, len(state.Spectrum.Lambdas)-1)
    finishIndex := normde.DenormalizeFloat(ncf.X, 0, len(state.Spectrum.Lambdas)-1)
    normMv := geom2d.VecBetweenPoints(ncs, ncf)
    for i := startIndex; i <= finishIndex; i++{
        currVec := normMv.TimesScalar(normde.NormRangeIF(i, startIndex, finishIndex))
        currPoint := ncs.MoveByVector(currVec)
        logrus.Debugf("set index: %d to value %f", i, currPoint.Y)
        state.Spectrum.Lambdas[i] = currPoint.Y
    }
    s.storage.Set(state)
}

func (s *SPDEditMode) HandleDragEnd() {}

func (s *SPDEditMode) Name() string {
    return "SPDEditMode"
}
