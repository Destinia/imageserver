package bmp

import (
	"testing"

	"github.com/Destinia/imageserver"
	imageserver_image "github.com/Destinia/imageserver/image"
	imageserver_image_test "github.com/Destinia/imageserver/image/_test"
)

var _ imageserver_image.Encoder = &Encoder{}

func TestEncoder(t *testing.T) {
	imageserver_image_test.TestEncoder(t, &Encoder{}, "bmp")
}

func TestEncoderChange(t *testing.T) {
	c := (&Encoder{}).Change(imageserver.Params{})
	if c {
		t.Fatal("not false")
	}
}
