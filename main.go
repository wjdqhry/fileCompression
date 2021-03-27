package main

import (
	"compress/flate"
	"log"
	"sync"

	"github.com/mholt/archiver"
)

func main() {

	wg := sync.WaitGroup{}

	wg.Add(3)
	go func() {
		newArchive(createZipArchiver(), "/Users/jungbokyo/go/src/fileCompression/bogus", "/Users/jungbokyo/go/src/fileCompression/bokyo.zip")
		wg.Done()
		log.Println("1 done")
	}()

	go func() {
		newArchive(createZipArchiver(), "/Users/jungbokyo/go/src/fileCompression/boguscopy2", "/Users/jungbokyo/go/src/fileCompression/bokyo_copy_2.zip")
		wg.Done()
		log.Println("2 done")
	}()
	go func() {
		newArchive(createZipArchiver(), "/Users/jungbokyo/go/src/fileCompression/boguscopy3", "/Users/jungbokyo/go/src/fileCompression/bokyo_copy_3.zip")

		wg.Done()
		log.Println("3 done")
	}()
	wg.Wait()

	log.Println("all done")
}

func newArchive(zip archiver.Zip, target string, dest string) {
	err := zip.Archive([]string{target}, dest)
	if err != nil {
		log.Println(err)
	}
}

func createZipArchiver() archiver.Zip {
	return archiver.Zip{
		CompressionLevel:       flate.BestCompression,
		MkdirAll:               true,
		SelectiveCompression:   true,
		ContinueOnError:        false,
		OverwriteExisting:      true,
		ImplicitTopLevelFolder: true,
	}
}
