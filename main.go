package main

import (
	"compress/flate"
	"log"

	"github.com/mholt/archiver"
)

func main() {
	zip := archiver.Zip{
		CompressionLevel:       flate.BestCompression,
		MkdirAll:               true,
		SelectiveCompression:   true,
		ContinueOnError:        false,
		OverwriteExisting:      true,
		ImplicitTopLevelFolder: true,
	}
	err := zip.Archive([]string{"bogus"}, "./bokyo.zip")
	if err != nil {
		log.Fatal(err)
	}
	//err = archiver.Extract()
}
