package pfp

import "image"

type Pfp interface {
	Resize(img image.Image) (image.Image, error)
}
