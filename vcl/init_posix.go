// +build linux !cgo, darwin !cgo

package vcl

import (
	"fmt"
)

func showError(err interface{}) {
	fmt.Println(err)
}
