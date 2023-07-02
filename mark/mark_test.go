package mark

import (
	"image"
	"image/jpeg"
	"log"
	"os"
	"testing"
)

func TestGetEmbedImage(t *testing.T) {
	fs, err := sony.Open("sony.jpg")
	if err != nil {
		t.Fatal(err)
	}

	image, s, err := image.Decode(fs)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(s)
	file, err := os.Create("test_read.jpg")
	if err != nil {
		log.Fatal(err)
	}

	jpeg.Encode(file, image, nil)
}
