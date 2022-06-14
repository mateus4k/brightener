package handler

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/disintegration/imaging"
)

var (
	outputPathName = "output"
)

func Handle(inputDir string, brightnessPercentage float64) {
	wg := sync.WaitGroup{}

	files, err := getFiles(inputDir)
	if err != nil {
		log.Fatal(err)
	}

	outputPath, err := getOutputPath(inputDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		wg.Add(1)

		go processFile(inputDir, f.Name(), outputPath, brightnessPercentage, &wg)
	}

	wg.Wait()
}

func getFiles(dir string) ([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, fmt.Errorf("no files were found at %s", dir)
	}

	return files, nil
}

func getOutputPath(inputDir string) (string, error) {
	outputPath := fmt.Sprintf("%s/%s", inputDir, outputPathName)

	_, err := os.Stat(outputPath)
	if err != nil {
		if !os.IsNotExist(err) {
			return "", err
		}

		err := os.Mkdir(outputPath, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	return outputPath, nil
}

func processFile(inputPath, filename, outputPath string, brightnessPercentage float64, wg *sync.WaitGroup) {
	defer wg.Done()

	sourceImage, err := imaging.Open(fmt.Sprintf("%s/%s", inputPath, filename))
	if err != nil {
		log.Fatal(err)
	}

	updatedImage := imaging.AdjustBrightness(sourceImage, brightnessPercentage)

	output := fmt.Sprintf("%s/%s", outputPath, filename)

	if err = imaging.Save(updatedImage, output); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done:", output)
}
