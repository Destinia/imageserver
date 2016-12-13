package tiff

import (
	"testing"

	"github.com/Destinia/imageserver"
	imageserver_image_test "github.com/Destinia/imageserver/image/_test"
	_ "github.com/Destinia/imageserver/image/jpeg"
	"github.com/Destinia/imageserver/testdata"
)

func Benchmark(b *testing.B) {
	for _, tc := range []struct {
		name string
		im   *imageserver.Image
	}{
		{"Small", testdata.Small},
		{"Medium", testdata.Medium},
		{"Large", testdata.Large},
		{"Huge", testdata.Huge},
	} {
		b.Run(tc.name, func(b *testing.B) {
			imageserver_image_test.BenchmarkEncoder(b, &Encoder{}, tc.im, imageserver.Params{})
		})
	}
}
