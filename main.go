package main

import (
	"compress/flate"
	"log"
	"os"
	"sync"

	"github.com/mholt/archiver"
)

func main() {
	pwd := "/data/ai_data_2020_01/016.도로주행영상/Training/"
	zipList := []string{
		"1.Training_bb/도심로/야간일몰_맑음_120_전방",
		"1.Training_bb/도심로/주간일출_맑음_120_전방",
		"1.Training_bb/도심로/야간일몰_맑음_30_전방",
		"1.Training_bb/도심로/주간일출_맑음_45_전방",
		"1.Training_bb/도심로/야간일몰_맑음_60_전방",
		"1.Training_bb/도심로/주간일출_맑음_60_전방",
		"1.Training_bb/도심로/주간일출_강우_120_전방",
		"1.Training_bb/도심로/주간일출_맑음_120_측방_전측방",
		"1.Training_bb/도심로/주간일출_강우_30_전방",
		"1.Training_bb/도심로/주간일출_맑음_50_측방_후측방",
		"1.Training_bb/도심로/주간일출_강우_60_전방",
		"1.Training_bb/도심로/도심로_Labeling",
		"1.Training_bb/자동차전용도로/야간일몰_맑음_120_전방",
		"1.Training_bb/자동차전용도로/야간일몰_맑음_30_전방",
		"1.Training_bb/자동차전용도로/주간일출_강우_45_전방",
		"1.Training_bb/자동차전용도로/야간일몰_맑음_60_전방",
		"1.Training_bb/자동차전용도로/자동차전용도로_Labeling",
		"1.Training_bb/자동차전용도로/주간일출_맑음_120_전방",
		"1.Training_bb/자동차전용도로/주간일출_안개_30_전방",
		"1.Training_bb/자동차전용도로/주간일출_맑음_45_전방",
		"1.Training_bb/자동차전용도로/주간일출_맑음_60_전방",
		"1.Training_bb/자동차전용도로/주간일출_안개_120_전방",
		"1.Training_bb/자동차전용도로/야간일몰_안개_45_전방",
		"1.Training_bb/자동차전용도로/주간일출_안개_45_전방",
		"1.Training_bb/자동차전용도로/주간일출_안개_60_전방",
		"1.Training_fs/도심로/도심로_Labeling",
		"1.Training_fs/도심로/주간일출_맑음_60_전방",
		"1.Training_fs/자동차전용도로/자동차전용도로_Labeling",
		"1.Training_fs/자동차전용도로/주간일출_맑음_60_전방",
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
