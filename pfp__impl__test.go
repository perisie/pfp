package pfp

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"image/jpeg"
	"os"
	"testing"
)

func Test(t *testing.T) {
	b, err := os.ReadFile("img/x_profile.jpeg")
	assert.Nil(t, err)
	assert.True(t, len(b) > 0)
	r := bytes.NewReader(b)
	image, err := jpeg.Decode(r)
	assert.Nil(t, err)
	assert.NotNil(t, image)
	var pfp Pfp
	pfp, err = Pfp__impl__new()
	assert.Nil(t, err)
	image__resize, err := pfp.Resize(image)
	assert.Nil(t, err)
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, image__resize, nil)
	assert.Nil(t, err)
	err = os.WriteFile("img/x_profile__out.jpeg", buf.Bytes(), 0600)
	assert.Nil(t, err)
}
