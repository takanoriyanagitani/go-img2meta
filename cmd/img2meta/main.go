package main

import (
	"io"
	"log"
	"os"

	im "github.com/takanoriyanagitani/go-img2meta"
)

var input io.Reader = os.Stdin

func main() {
	meta, e := im.ReaderToMetadata(input)
	if nil != e {
		log.Printf("%v\n", e)
		return
	}

	jdat, e := meta.ToJson()
	if nil != e {
		log.Printf("%v\n", e)
		return
	}

	_, e = os.Stdout.Write(jdat)
	if nil != e {
		log.Printf("%v\n", e)
		return
	}
}
