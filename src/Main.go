package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const directory string = "/proc/"
const stat string = "/stat"

//const metrics string = "metrics.txt"

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func main() {

	metricspage, err := os.OpenFile("/tmp/metrics.txt", os.O_CREATE|os.O_WRONLY, 0600)

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if isNumeric(f.Name()) {
			//if contains any non numberic characters, ignore
			var FullPath = strings.Join([]string{directory, f.Name(), stat}, "")
			//fmt.Println(FullPath, "\t")
			content, err := ioutil.ReadFile(FullPath)
			if err != nil {
				log.Fatal(err)
			}
			//fmt.Printf("File contents:\t%s", content)
			metrics := string(content) // have to convert the byte array to a string
			//append
			//fmt.Println(metricspage)
			metricspage.WriteString(metrics)
		}
		//fmt.Println(directory, f.Name(), stat)
		//statsdir, err := ioutil.ReadDir()
	}
	defer metricspage.Close()

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "hi\nthere\nfriend")
		webpage, _ := ioutil.ReadFile("/tmp/metrics.txt")
		fmt.Fprintf(w, string(webpage))
	})

	log.Fatal(http.ListenAndServe(":9256", nil))
}
