package api

import (
	"gitee.com/ying32/govcl/vcl/dylib"
)

var (
	libvcl = dylib.NewLazyDLL("libvcl.dylib")
)
