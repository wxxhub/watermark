package test

import (
	"encoding/json"
	"fmt"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/nfnt/resize"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/wxxhub/watermark/draw"
	exif2 "github.com/wxxhub/watermark/exif"
	"golang.org/x/image/font/gofont/goregular"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"time"

	//"golang.org/x/image/font"
	"testing"
)

func TestImage(t *testing.T) {
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	//dc.SetFontFace(font.)
	dc.DrawString("Hello World", 100, 100)
	//dc.DrawStringWrapped("Hello, world! How are you?", 50, 50, 0.5, .05, 1024, 1.5, gg.AlignRight)
	dc.SavePNG("out.png")
}

func TestDrawText(t *testing.T) {
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}

	face := truetype.NewFace(font, &truetype.Options{Size: 60})
	face2 := truetype.NewFace(font, &truetype.Options{Size: 40})
	const S = 1024
	const width = S / 2
	dc := gg.NewContext(S, width)
	dc.SetRGB255(255, 255, 255)
	dc.Clear()

	dc.SetRGB255(0, 0, 0)
	dc.SetFontFace(face)
	dc.DrawStringWrapped(time.Now().Format("2006.01.02 15:04:05"), 0, 50, -0.1, 0, 600, 1.5, gg.AlignLeft)

	dc.DrawLine(0, 0, 100, 100)
	dc.SetRGB255(100, 100, 100)
	dc.SetLineWidth(2)
	dc.Stroke()

	dc.SetRGB255(100, 100, 100)
	dc.SetFontFace(face2)
	dc.DrawStringWrapped(time.Now().Format("2006.01.02 15:04:05"), 0, width/2+20, -0.1, 0, 600, 1.5, gg.AlignLeft)

	dc.SavePNG("text.png")
}

func TestDrawCameraModel(t *testing.T) {
	i, err := draw.CameraModel("LICE-7M4", time.Now().Format("2006.01.02 15:04:05"))
	if err != nil {
		t.Fatal(err)
	}

	file, err := os.Create("test_camera_model.png")
	if err != nil {
		log.Fatal(err)
	}

	png.Encode(file, i)
}

func TestDrawParameter(t *testing.T) {
	i, err := draw.Parameter(nil, "120mm f/2.8 1/100 ISO200", "@running rabbit")
	if err != nil {
		t.Fatal(err)
	}

	file, err := os.Create("test_parameter.png")
	if err != nil {
		log.Fatal(err)
	}

	png.Encode(file, i)
}

func TestExif(t *testing.T) {
	f, err := os.Open("../demo/demo.jpg")
	x, err := exif.Decode(f)
	if err != nil {
		t.Fatal(err)
	}

	//tag, err := x.Get(exif.FNumber)
	tag, err := x.Get(exif.FocalLength)
	if err != nil {
		t.Fatal(err)
	}

	v, err := tag.Rat(0)
	if err != nil {
		t.Fatal(err)
	}

	n, _ := v.Float64()
	t.Log(n)
	t.Log(v)
}

func TestExif3(t *testing.T) {
	f, err := os.Open("../demo/demo.jpg")
	x, err := exif.Decode(f)
	if err != nil {
		t.Fatal(err)
	}

	data, _ := x.MarshalJSON()
	t.Logf("%s\n", data)
	//tag, err := x.Get(exif.FNumber)
	tag, err := x.Get(exif.Make)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(tag.String())

}

func TestExif2(t *testing.T) {
	f, err := os.Open("../demo/demo.jpg")
	e, err := exif2.GetExif(f)
	if err != nil {
		t.Fatal(err)
	}

	info, err := json.Marshal(e)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s", info)
}

func TestDrawParameter2(t *testing.T) {
	f, err := os.Open("../demo/demo.jpg")
	e, err := exif2.GetExif(f)
	if err != nil {
		t.Fatal(err)
	}

	e.FNumber = 4
	i, err := draw.Parameter(nil, fmt.Sprintf("%dmm F/%.1f %s ISO%d", e.FocalLength, e.FNumber, e.ExposureTime, e.ISO), "by 奔跑的兔")
	if err != nil {
		t.Fatal(err)
	}

	file, err := os.Create("test_parameter_2.png")
	if err != nil {
		log.Fatal(err)
	}

	png.Encode(file, i)
}

func TestDrawParameter4(t *testing.T) {
	f, err := os.Open("../demo/demo.jpg")
	e, err := exif2.GetExif(f)
	if err != nil {
		t.Fatal(err)
	}

	markF, err := os.Open("../mark/sony.jpg")
	markE, err := jpeg.Decode(markF)
	if err != nil {
		t.Fatal(err)
	}

	e.FNumber = 4
	i, err := draw.Parameter(markE, fmt.Sprintf("%dmm F/%.1f %s ISO%d", e.FocalLength, e.FNumber, e.ExposureTime, e.ISO), "by 奔跑的兔")
	if err != nil {
		t.Fatal(err)
	}

	file, err := os.Create("test_parameter_2.png")
	if err != nil {
		log.Fatal(err)
	}

	png.Encode(file, i)
}

func TestDrawMarkColumn(t *testing.T) {
	f, err := os.Open("../demo/demo.jpg")
	e, err := exif2.GetExif(f)
	if err != nil {
		t.Fatal(err)
	}

	markF, err := os.Open("../mark/sony.jpg")
	markE, err := jpeg.Decode(markF)
	if err != nil {
		t.Fatal(err)
	}

	e.FNumber = 4
	i, err := draw.MarkColumn(markE, e, "by 奔跑的兔", 2048)
	if err != nil {
		t.Fatal(err)
	}

	file, err := os.Create("test_mark_colum.png")
	if err != nil {
		log.Fatal(err)
	}

	png.Encode(file, i)
}

func TestDrawMarkColumn2(t *testing.T) {
	f, err := os.Open("../demo/demo.jpg")
	if err != nil {
		t.Fatal(err)
	}

	e, err := exif2.GetExif(f)
	if err != nil {
		t.Fatal(err)
	}

	markF, err := os.Open("../mark/sony.jpg")
	markE, err := jpeg.Decode(markF)
	if err != nil {
		t.Fatal(err)
	}

	e.FNumber = 4
	mc, err := draw.MarkColumn(markE, e, "by 奔跑的兔", 2048)
	if err != nil {
		t.Fatal(err)
	}

	f2, err := os.Open("../demo/demo.jpg")
	if err != nil {
		t.Fatal(err)
	}
	ej, err := jpeg.Decode(f2)
	if err != nil {
		t.Fatal(err)
	}

	mcR := resize.Resize(uint(ej.Bounds().Dx()), 0, mc, resize.Lanczos3)

	dc := gg.NewContext(ej.Bounds().Dx(), ej.Bounds().Dy()+mcR.Bounds().Dy())
	dc.SetRGB255(255, 255, 255)
	dc.Clear()

	dc.DrawImage(ej, 0, 0)
	dc.DrawImage(mcR, 0, ej.Bounds().Dy())

	ff, err := os.Create("test_mark_colum.jpg")
	if err != nil {
		t.Fatal(err)
	}

	jpeg.Encode(ff, dc.Image(), nil)
}
