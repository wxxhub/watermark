package draw

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/nfnt/resize"
	"github.com/wxxhub/watermark/exif"
	"github.com/wxxhub/watermark/ttf"
	"image"
	"image/color"
	"sync"
)

const (
	Height  = 100
	Padding = 100
)

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
		mainSize = Height / 2.5 * 1.5
		subSize  = Height / 2.5

		mainFace = truetype.NewFace(getFont(), &truetype.Options{Size: mainSize})
		subFace  = truetype.NewFace(getFont(), &truetype.Options{Size: subSize})
	)

	dc := gg.NewContext(width, Height+2*Padding)
	dc.SetColor(backgroundColor)
	dc.Clear()

	// draw main
	dc.SetColor(mainColor)
	dc.SetFontFace(mainFace)
	dc.DrawStringWrapped(model, 0, Padding, -0.1, 0, float64(width), 1.5, gg.AlignLeft)

	// draw sub
	dc.SetColor(subColor)
	dc.SetFontFace(subFace)
	dc.DrawStringWrapped(photoTime, 0, Height/2+Padding, -0.1, 0, float64(width), 1.5, gg.AlignLeft)

	return dc.Image(), nil
}

func Parameter(sign image.Image, param, author string) (image.Image, error) {
	var (
		width = Height * 11
		// (Height-100)/2.5*1.5
		mainSize = Height / 2.5 * 1.5
		subSize  = Height / 2.5

		mainFace = truetype.NewFace(getFont(), &truetype.Options{Size: mainSize})
		subFace  = truetype.NewFace(getFont(), &truetype.Options{Size: subSize})
	)

	dc := gg.NewContext(width, Height+2*Padding)
	dc.SetColor(backgroundColor)
	dc.Clear()

	rSign := resize.Resize(0, Height, sign, resize.Lanczos3)

	rSignW := float64(rSign.Bounds().Dx())

	dc.DrawImage(rSign, 0, Padding)

	dc.SetColor(subColor)
	dc.SetLineWidth(5)
	dc.DrawLine(rSignW+20, Padding, rSignW+20, Height+Padding)
	dc.Stroke()

	// draw main
	dc.SetColor(mainColor)
	dc.SetFontFace(mainFace)
	dc.DrawStringWrapped(param, rSignW+40, Padding, 0, 0, float64(width), 1.5, gg.AlignLeft)

	// draw sub
	dc.SetColor(subColor)
	dc.SetFontFace(subFace)
	dc.DrawStringWrapped(author, rSignW+40, Height/2+Padding, 0, 0, float64(width), 1.5, gg.AlignLeft)

	return dc.Image(), nil
}

func MarkColumn(sign image.Image, exif *exif.Exif, author string, width int) (image.Image, error) {
	mI, err := CameraModel(exif.Model, exif.DateTime.Format("2006.01.02 15:04:05"))
	if err != nil {
		return nil, err
	}

	pI, err := Parameter(sign, fmt.Sprintf("%dmm F/%.1f %s ISO%d", exif.FocalLength, exif.FNumber, exif.ExposureTime, exif.ISO), author)
	if err != nil {
		return nil, err
	}

	oW := mI.Bounds().Dx() + pI.Bounds().Dx()
	if oW > width {
		width = oW + 1024
	} else {

	}
	dc := gg.NewContext(width, Height+2*Padding)
	dc.SetRGB255(255, 255, 255)
	dc.Clear()

	dc.DrawImage(mI, 0, 0)
	dc.DrawImage(pI, width-pI.Bounds().Dx(), 0)

	return dc.Image(), nil
}
