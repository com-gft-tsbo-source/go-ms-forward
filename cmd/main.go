package main

import (
	"os"

	"github.com/com-gft-tsbo-source/go-ms-forward/msforward"
)

// ###########################################################################
// ###########################################################################
// MAIN
// ###########################################################################
// ###########################################################################

var usage []byte = []byte("ms-forward: [OPTIONS] ")

func main() {

	var ms msforward.MsForward
	msforward.InitFromArgs(&ms, os.Args, nil)
	ms.Run()
}
