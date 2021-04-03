package main

import (
	"compress/flate"
	"log"
	"os"
	"sync"

	"github.com/mholt/archiver"
)

func main() {
	pwd := "/data/sftp/1차_구축사업/unit5/07랜드마크_이미지_AI데이터/files/01.데이터/"
	zipList := []string{
"/1.Training/원시데이터/대구광역시/랜드마크_이미지_대구광역시_001",
"/1.Training/원시데이터/대구광역시/랜드마크_이미지_대구광역시_002",
"/1.Training/원시데이터/대구광역시/랜드마크_이미지_대구광역시_003",
"/1.Training/원시데이터/대구광역시/랜드마크_이미지_대구광역시_004",
"/1.Training/원시데이터/대구광역시/랜드마크_이미지_대구광역시_005",
"/1.Training/원시데이터/대구광역시/랜드마크_이미지_대구광역시_006",
"/1.Training/원시데이터/대구광역시/랜드마크_이미지_대구광역시_007",
"/1.Training/원시데이터/대구광역시/랜드마크_이미지_대구광역시_008",
"/1.Training/원시데이터/대전광역시/랜드마크_이미지_대전광역시_002",
"/1.Training/원시데이터/대전광역시/랜드마크_이미지_대전광역시_003",
"/1.Training/원시데이터/대전광역시/랜드마크_이미지_대전광역시_001",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_001",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_002",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_003",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_004",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_005",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_006",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_007",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_008",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_009",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_010",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_011",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_012",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_013",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_014",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_015",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_016",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_017",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_018",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_019",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_020",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_021",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_022",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_023",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_024",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_025",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_026",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_027",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_028",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_029",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_030",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_031",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_032",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_033",
"/1.Training/원시데이터/부산광역시/랜드마크_이미지_부산광역시_034",

}
	for _, origin := range zipList {
		if _, err := os.Stat(pwd + origin); err != nil {
			log.Fatal("경로가 잘못되었습니다")
		} else {
			log.Println(origin, "okay!!")
		}
	}

	wg := sync.WaitGroup{}
	wg.Add(len(zipList))

	for _, target := range zipList {
		go func(t string) {
			target := pwd + t
			newArchive(createZipArchiver(), target, target+".zip")
			wg.Done()
			log.Println(target, "complete!")
		}(target)
	}

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
