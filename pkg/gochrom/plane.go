package gochrom

import (
    "github.com/areknoster/gochrom/pkg/geom2d"
    "image"
)

type Plane interface {
    Refresh()
}

type Renderer interface {
    Render(w, h int) image.Image
}

type InteractiveMode interface {
    HandleClick(normLoc geom2d.Point)
    HandleDrag(start geom2d.Point, move geom2d.Vector)
    HandleDragEnd()
    Name() string
}
