package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const directory string = "/proc/"
const stat string = "/stat"

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func main() {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if isNumeric(f.Name()) {
			//if contains any non numberic characters, break
			var FullPath = strings.Join([]string{directory, f.Name(), stat}, "")
			fmt.Println(FullPath, "\t")
			content, err := ioutil.ReadFile(FullPath)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("File contents:\t%s", content)
		} else {
			continue
		}
		//fmt.Println(directory, f.Name(), stat)
		//statsdir, err := ioutil.ReadDir()
	}
}
