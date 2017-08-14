package unique_finder

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"strings"
	"testing"
)

func TestPhoneMapAddMember(t *testing.T) {
	finder := NewFinder()
	finder.Add("0980000001", "2016-09-01", "2016-10-01")
	finder.Add("0980000001", "2016-11-01", "2016-12-01")
	finder.Add("0980000002", "2016-09-01", "2016-10-01")
	finder.Add("0980000003", "2016-09-01", "2016-10-01")
	if len(finder.PhoneMap) != 3 {
		t.Log("Phone map should contains 3 members")
		t.Fail()
	}
}

func TestListPhoneWithActivate(t *testing.T) {
	finder := NewFinder()
	finder.Add("0980000001", "2016-09-01", "2016-11-01")
	finder.Add("0980000001", "2016-11-01", "2016-12-01")
	listPhone := finder.ListPhoneWithActivatedTime()
	if listPhone[0] != "0980000001,2016-09-01" {
		t.Fail()
	}
}

func TestWriteCsv(t *testing.T) {
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)
	w := csv.NewWriter(writer)

	finder := NewFinder()
	finder.Add("0980000001", "2016-09-01", "2016-11-01")
	finder.Add("0980000001", "2016-11-01", "2016-12-01")
	finder.WriteToCSV(w)

	expected := `0980000001,2016-09-01` //write with end of line

	if !strings.Contains(buffer.String(), expected) {
		t.Log("output should contains", expected)
		t.Fail()
	}
}
