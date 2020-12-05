package render

import (
    "github.com/areknoster/gochrom/pkg/gochrom"
    "github.com/fogleman/gg"
    "image"
)

type SpectralDiag struct{
    storage gochrom.StateStorage
}

func NewSpectralDiag(storage gochrom.StateStorage) *SpectralDiag {
    return &SpectralDiag{storage: storage}
}

var _ gochrom.Renderer = &SpectralDiag{}

func (s SpectralDiag) Render(w, h int) image.Image {
    ctx := gg.NewContext(w,h)
    ctx.DrawString("hello world", 0.5, 0.5)
    return ctx.Image()
}

