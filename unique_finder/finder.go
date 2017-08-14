package unique_finder

import (
	"bytes"
	"encoding/csv"
	"log"
)

type Finder struct {
	PhoneMap map[string]*PhoneHistory
}

func NewFinder() *Finder {
	return &Finder{
		make(map[string]*PhoneHistory, 0),
	}
}

func (f *Finder) Add(phoneNumber string, activeAt string, deactiveAt string) {
	history := NewHistory(activeAt, deactiveAt)
	if ph, ok := f.PhoneMap[phoneNumber]; ok {
		ph.Insert(history)
	} else { //create a new once phone history
		newPh := NewPhoneHistory(phoneNumber)
		newPh.Insert(history)
		f.PhoneMap[phoneNumber] = newPh
	}
}

func (f *Finder) ListPhoneWithActivatedTime() []string {
	result := make([]string, 0)
	var rowBuffer bytes.Buffer
	for k, v := range f.PhoneMap {
		rowBuffer.WriteString(k)
		rowBuffer.WriteString(",")
		rowBuffer.WriteString(v.LatestActivate().Format(TIME_LAYOUT))
		result = append(result, rowBuffer.String())
		rowBuffer.Reset()
	}
	return result
}

func (f *Finder) WriteToCSV(w *csv.Writer) {
	for k, v := range f.PhoneMap {
		err := w.Write([]string{k, v.LatestActivate().Format(TIME_LAYOUT)})
		if err != nil {
			log.Fatal(err)
		}
	}
	w.Flush()
}
