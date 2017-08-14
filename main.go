package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/duythinht/activated_phones/unique_finder"
)

func main() {

	in, out := parseArgs()

	fin, err := os.Open(in)
	defer func() { _ = fin.Close() }()
	checkErr(err)

	fmt.Println("Read csv from", in, "and output to", out)

	reader := getCSVReader(fin)
	finder := makeUniquePhoneFinder(reader)

	outputResult(out, finder)
}

func parseArgs() (in string, out string) {
	flag.StringVar(&in, "in", "sample.csv", "enter path of input csv")
	flag.StringVar(&out, "out", "result.csv", "enter path of input csv")
	flag.Parse()
	return
}

func makeUniquePhoneFinder(r *csv.Reader) *unique_finder.Finder {
	finder := unique_finder.NewFinder()
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		checkErr(err)
		finder.Add(record[0], record[1], record[2])
	}
	return finder
}

func outputResult(out string, finder *unique_finder.Finder) {
	f, err := os.Open(out)

	if os.IsNotExist(err) {
		f, err = os.Create(out)
		checkErr(err)
	}

	defer func() { _ = f.Close() }()

	w := csv.NewWriter(f)

	_ = w.Write([]string{"PHONE_NUMBER", "REAL_ACTIVATION_DATE"}) //Write header
	finder.WriteToCSV(w)
}

func getCSVReader(f *os.File) *csv.Reader {

	r := csv.NewReader(f)

	_, err := r.Read() //skip header

	checkErr(err)

	return r
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
