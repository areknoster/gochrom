package render

import (
    "github.com/areknoster/gochrom/pkg/gochrom"
    "github.com/fogleman/gg"
    "image"
)

type CIA struct{
    storage gochrom.StateStorage
}

func NewCIA(storage gochrom.StateStorage) *CIA {
    return &CIA{storage: storage}
}

var _ gochrom.Renderer = &CIA{}

func (C CIA) Render(w, h int) image.Image {
    ctx := gg.NewContext(w,h)
    ctx.DrawString("hello world", 0.5, 0.5)
    return ctx.Image()
}