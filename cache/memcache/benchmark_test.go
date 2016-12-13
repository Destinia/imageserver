package memcache

import (
	"testing"

	"github.com/Destinia/imageserver"
	cachetest "github.com/Destinia/imageserver/cache/_test"
	"github.com/Destinia/imageserver/testdata"
)

func BenchmarkGet(b *testing.B) {
	for _, tc := range []struct {
		name string
		im   *imageserver.Image
	}{
		{"Small", testdata.Small},
		{"Medium", testdata.Medium},
		{"Large", testdata.Large},
	} {
		b.Run(tc.name, func(b *testing.B) {
			cch := newTestCache(b)
			cachetest.BenchmarkGet(b, cch, 1, tc.im) // memcached is unstable with more parallelism
		})
	}
}
