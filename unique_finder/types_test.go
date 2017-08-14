package unique_finder

import "testing"

var __dump = `
0987000001,2016-03-01,2016-05-01
0987000002,2016-02-01,2016-03-01
0987000001,2016-01-01,2016-03-01
0987000001,2016-12-01,
0987000002,2016-03-01,2016-05-01
0987000003,2016-01-01,2016-01-10
0987000001,2016-09-01,2016-12-01
0987000002,2016-05-01,
0987000001,2016-06-01,2016-09-01
`

func TestNewHistoryFromString(t *testing.T) {
	history := NewHistory("2017-09-07", "2017-10-15")
	if history.ActivedAt.Year() != 2017 {
		t.Log("Activated year should be 2017")
		t.Log("Actually get:", history.ActivedAt.Year())
		t.Fail()
	}

	if history.ActivedAt.Month().String() != "September" {
		t.Log("Activated month should be September")
		t.Log("Actually get:", history.ActivedAt.Month().String())
		t.Fail()
	}

	if history.ActivedAt.Day() != 7 {
		t.Log("Activated day should be 7")
		t.Log("Actually get:", history.ActivedAt.Day())
		t.Fail()
	}

	//One more sample for deactive day
	if history.DeactivatedAt.Day() != 15 {
		t.Log("Deactivated day should be 15")
		t.Log("Actually get:", history.ActivedAt.Day())
		t.Fail()
	}
}

func TestHistoryDeactiveTimeEmpty(t *testing.T) {
	history := NewHistory("2017-09-07", "")
	if history.DeactivatedAt != nil {
		t.Log("Deactivated Day should be nil")
		t.Fail()
	}
}

func TestInsertToSortedActivatedTime(t *testing.T) {
	data := [][]string{
		[]string{"2016-03-01", "2016-05-01"},
		[]string{"2016-12-01", ""},
		[]string{"2016-09-01", "2016-12-01"},
		[]string{"2016-06-01", "2016-09-01"},
	}

	phoneHistory := NewPhoneHistory("0980000001")
	histories := make([]*History, 0)

	for _, historyPair := range data {
		h := NewHistory(historyPair[0], historyPair[1])
		histories = append(histories, h)
		phoneHistory.Insert(h)
	}

	for _, h := range phoneHistory.History {
		t.Log(h.ActivedAt.String())
	}

	if phoneHistory.History[0].ActivedAt.Equal(*histories[2].ActivedAt) {
		t.Log("Highest activate time should be at head")
		t.Fail()
	}

	if phoneHistory.History[3].ActivedAt.Equal(*histories[1].ActivedAt) {
		t.Log("Lowest activate time should be at tail")
		t.Fail()
	}
}

func TestGetLatestActivatedTime(t *testing.T) {
	data := [][]string{
		[]string{"2016-03-01", "2016-05-01"},
		[]string{"2016-12-01", ""},
		[]string{"2016-09-01", "2016-12-01"},
		[]string{"2016-06-01", "2016-09-01"},
	}

	phoneHistory := NewPhoneHistory("0980000001")
	histories := make([]*History, 0)

	for _, historyPair := range data {
		h := NewHistory(historyPair[0], historyPair[1])
		histories = append(histories, h)
		phoneHistory.Insert(h)
	}

	for _, h := range phoneHistory.History {
		t.Log(h.ActivedAt.String())
	}

	if !phoneHistory.LatestActivate().Equal(*histories[3].ActivedAt) {
		t.Log("Last activate is incorrect")
		t.Fail()
	}
}
