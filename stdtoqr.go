package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rsc/rsc/qr"
)

func main() {
	var file = flag.String("file", "qr.png", "file to output qr code to")
	flag.Parse()
	bio := bufio.NewReader(os.Stdin)
	line, _, err := bio.ReadLine()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	code, err := qr.Encode(string(line), qr.Q)

	q := code.PNG()

	ioutil.WriteFile(*file, q, 0600)
}
