package main

import (
    "fyne.io/fyne"
    "fyne.io/fyne/app"
    "fyne.io/fyne/layout"
    "github.com/areknoster/gochrom/pkg/plane"
    "github.com/areknoster/gochrom/pkg/render"
    "github.com/areknoster/gochrom/pkg/state"
    "github.com/sirupsen/logrus"
)

type Config struct{
    title string
    CIASize fyne.Size
    SPDSize fyne.Size
}

func main(){
    logrus.SetLevel(logrus.DebugLevel)
    cfg := Config{
        title:      "GoChrom",
        CIASize: fyne.Size{600, 400},
        SPDSize: fyne.Size{600, 300},

    }

    fyneApp := app.New()
    window :=fyneApp.NewWindow(cfg.title)

    ss := state.NewStateStorage()

    CIARenderer := render.NewCIE(ss)
    CIAPlane := plane.NewPlane(CIARenderer, cfg.CIASize)
    ss.AddRefresh(CIAPlane.Refresh)

    SPDRenderer := render.NewSPD(ss)
    SPDPlane, SPDSetMode := plane.NewInteractivePlane(SPDRenderer, cfg.SPDSize)
    ss.AddRefresh(SPDPlane.Refresh)
    SPDSetMode(plane.NewSPDEditMode(ss))


    container := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), CIAPlane, SPDPlane)
    window.SetContent(container)
    window.SetFixedSize(true)
    window.ShowAndRun()


}

