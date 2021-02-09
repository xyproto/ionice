package main

import (
	"io/ioutil"
	"os"

	"github.com/xyproto/gionice"
)

func main() {
	// Make the current process "idle" (level 7)
	gionice.SetIdle(0)

	// Generate I/O activity
	for {
		_ = ioutil.WriteFile("frenetic.dat", []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0644)
		_ = os.Remove("frenetic.dat")
	}
}
