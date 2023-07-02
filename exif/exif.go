package exif

import (
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
	"io"
	"time"
)

type Exif struct {
	Make         string
	Model        string
	Artist       string
	ExposureTime string

	ISO         int
	FocalLength int

	DateTime time.Time

	FNumber float64
}

func GetExif(r io.Reader) (*Exif, error) {
	x, err := exif.Decode(r)
	if err != nil {
		return nil, err
	}

	e := &Exif{}
	e.Make, err = getStringTag(x, exif.Make)
	if err != nil {
		return nil, err
	}

	e.Model, err = getStringTag(x, exif.Model)
	if err != nil {
		return nil, err
	}

	e.Artist, err = getStringTag(x, exif.Artist)
	if err != nil {
		return nil, err
	}

	e.ExposureTime, err = getStringTag(x, exif.ExposureTime)
	if err != nil {
		return nil, err
	}

	e.DateTime, err = getDateTimeTag(x, exif.DateTime)
	if err != nil {
		return nil, err
	}

	e.FNumber, err = getFloatTag(x, exif.FNumber)
	if err != nil {
		return nil, err
	}

	e.ISO, err = getIntTag(x, exif.ISOSpeedRatings)
	if err != nil {
		return nil, err
	}

	e.FocalLength, err = getIntTag(x, exif.FocalLength)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func getIntTag(e *exif.Exif, name exif.FieldName) (int, error) {
	tag, err := e.Get(name)
	if err != nil {
		return 0, err
	}
	switch tag.Type {
	case tiff.DTRational:
		r, err := tag.Rat(0)
		if err != nil {
			return 0, err
		}

		v, _ := r.Float64()
		return int(v), nil
	default:
		return tag.Int(0)

	}
}

func getStringTag(e *exif.Exif, name exif.FieldName) (string, error) {
	tag, err := e.Get(name)
	if err != nil {
		return "", err
	}

	switch tag.Type {
	case tiff.DTRational:
		r, err := tag.Rat(0)
		if err != nil {
			return "", err
		}

		return r.RatString(), nil
	default:
		return tag.StringVal()
	}
}

func getDateTimeTag(e *exif.Exif, name exif.FieldName) (time.Time, error) {
	tag, err := e.Get(name)
	if err != nil {
		return time.Time{}, err
	}

	v, err := tag.StringVal()
	if err != nil {
		return time.Time{}, err
	}

	return time.Parse("2006:01:02 15:04:05", v)
}

func getFloatTag(e *exif.Exif, name exif.FieldName) (float64, error) {
	tag, err := e.Get(name)
	if err != nil {
		return 0, err
	}

	r, err := tag.Rat(0)
	if err != nil {
		return 0, err
	}

	v, _ := r.Float64()
	return v, nil
}
