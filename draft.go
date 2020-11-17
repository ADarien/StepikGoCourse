package main

import (
	"archive/zip"
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	const url = "https://github.com/semyon-dev/stepik-go/raw/master/work_with_files/task_csv_1/task.zip"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		log.Fatal(err)
	}

	for _, zipFile := range zipReader.File {
		file, err := zipFile.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		csvReader := csv.NewReader(file)
		rows, err := csvReader.ReadAll()
		if err != nil || len(rows) != 10 || len(rows[4]) != 10 {
			continue
		}

		fmt.Println(rows[4][2])
		break
	}
}
