package draw

import (
	"github.com/fogleman/gg"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/wxxhub/watermark/ttf"
	"image"
	"image/color"
	"sync"
)

const Height = 400

var (
	backgroundColor = color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}
	mainColor = color.RGBA{
		R: 0,
		G: 0,
		B: 0,
		A: 255,
	}

	subColor = color.RGBA{
		R: 100,
		G: 100,
		B: 100,
		A: 200,
	}
)

//var font = gomedium.TTF
//var font = gomono.TTF
//var font = gomonobold.TTF
//var font = goregular.TTF
//var font = gosmallcapsitalic.TTF

var font *truetype.Font
var fontOnce sync.Once

func getFont() *truetype.Font {
	fontOnce.Do(func() {
		fontT, err := freetype.ParseFont(ttf.WenKaiRegular)
		if err != nil {
			panic(err)
		}
		font = fontT
	})

	return font
}

func CameraModel(model, photoTime string) (image.Image, error) {
	var (
		width = Height * 5
		// (Height-100)/2.5*1.5
		mainSize = (Height - 100) / 2.5 * 1.5
		subSize  = (Height - 100) / 2.5

		mainFace = truetype.NewFace(getFont(), &truetype.Options{Size: mainSize})
		subFace  = truetype.NewFace(getFont(), &truetype.Options{Size: subSize})
	)

	dc := gg.NewContext(width, Height)
	dc.SetColor(backgroundColor)
	dc.Clear()

	// draw main
	dc.SetColor(mainColor)
	dc.SetFontFace(mainFace)
	dc.DrawStringWrapped(model, 0, 50, -0.1, 0.1, float64(width), 1.5, gg.AlignLeft)

	// draw sub
	dc.SetColor(subColor)
	dc.SetFontFace(subFace)
	dc.DrawStringWrapped(photoTime, 0, Height/2, -0.1, 0, float64(width), 1.5, gg.AlignLeft)

	return dc.Image(), nil
}

func Parameter(sign image.Image, param, author string) (image.Image, error) {
	var (
		width = Height * 8
		// (Height-100)/2.5*1.5
		mainSize = (Height - 100) / 2.5 * 1.5
		subSize  = (Height - 100) / 2.5

		mainFace = truetype.NewFace(getFont(), &truetype.Options{Size: mainSize})
		subFace  = truetype.NewFace(getFont(), &truetype.Options{Size: subSize})
	)

	dc := gg.NewContext(width, Height)
	dc.SetColor(backgroundColor)
	dc.Clear()

	// draw main
	dc.SetColor(mainColor)
	dc.SetFontFace(mainFace)
	dc.DrawStringWrapped(param, 0, 50, -0.01, 0.1, float64(width), 1.5, gg.AlignLeft)

	// draw sub
	dc.SetColor(subColor)
	dc.SetFontFace(subFace)
	dc.DrawStringWrapped(author, 0, Height/2, -0.01, 0, float64(width), 1.5, gg.AlignLeft)

	return dc.Image(), nil
}
