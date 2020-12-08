package normde

import (
	"image"

	"fyne.io/fyne"
	"github.com/bamiaux/rez"
	"github.com/sirupsen/logrus"
)

func ResizeImage(toResize image.Image, size fyne.Size)  image.Image{
	resized := image.NewRGBA(image.Rectangle{
		Max: image.Point{size.Width, size.Height},
	})
	err := rez.Convert(resized, toResize, rez.NewBicubicFilter())
	if err!= nil{
		logrus.Panicf("could not convert image size: %s", err.Error())
		return nil
	}
	return resized
}
