package api

import (
	"git.oschina.net/ying32/govcl/vcl/dylib"
)

var (
	libvcl = dylib.NewLazyDLL("libvcl.so")
)
