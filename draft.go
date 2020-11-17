package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk("./task", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			file, _ := os.Open(path)
			cr := csv.NewReader(file)
			rowIndex := 0
			for row, err := cr.Read(); err == nil; row, err = cr.Read() {
				if rowIndex == 4 {
					fmt.Println(row[2])
				}
				rowIndex++
			}
		}
		return nil
	})
}
