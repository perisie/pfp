package pfp

import (
	"golang.org/x/image/draw"
	"image"
)

type Pfp__impl struct {
}

func (p *Pfp__impl) Resize(img image.Image) (image.Image, error) {
	sz := img.Bounds().Dx()
	if img.Bounds().Dy() < sz {
		sz = img.Bounds().Dy()
	}
	rect := image.Rect(0, 0, sz, sz)
	img__crop := image.NewRGBA(image.Rect(0, 0, sz, sz))
	for y := 0; y < sz; y++ {
		for x := 0; x < sz; x++ {
			img__crop.Set(x, y, img.At(x, y))
		}
	}
	sz = 500
	rect = image.Rect(0, 0, sz, sz)
	img__out := image.NewRGBA(rect)
	draw.CatmullRom.Scale(img__out, rect, img__crop, img__crop.Bounds(), draw.Over, nil)
	return img__out, nil
}

func Pfp__impl__new() (*Pfp__impl, error) {
	return &Pfp__impl{}, nil
}
