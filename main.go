package main

import (
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

func main() {
	go func() {
		w := app.NewWindow(app.Size(710, 710))
		err := run(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(w *app.Window) error {
	ops := new(op.Ops)
	for {
		switch e := w.NextEvent().(type) {
		case app.FrameEvent:
			bg := color.NRGBA{R: 200, G: 250, B: 110, A: 255}
			paint.Fill(ops, bg)

			w := color.NRGBA{R: 255, G: 255, B: 255, A: 255}

			for i := range 10 {
				// border
				rb := clip.Rect{
					Min: image.Pt(50+i*30-2, 50+i*30-2),
					Max: image.Pt(400+i*30+2, 400+i*30+2),
				}
				paint.FillShape(ops, w, rb.Op())
				fg := color.NRGBA{R: 100 - uint8(i)*10, G: 50, B: 10 + uint8(i)*10, A: 255}
				r := clip.Rect{
					Min: image.Pt(50+i*30, 50+i*30),
					Max: image.Pt(400+i*30, 400+i*30),
				}
				paint.FillShape(ops, fg, r.Op())
			}

			e.Frame(ops)
		case app.DestroyEvent:
			return e.Err
		}
	}
}
