package xui

import (
	"strconv"
)

// str to int
func atoi(s string) (r int) {
	r, _ = strconv.Atoi(s)
	return
}

// str to bool
func atob(s string) (r bool) {
	r, _ = strconv.ParseBool(s)
	return
}
