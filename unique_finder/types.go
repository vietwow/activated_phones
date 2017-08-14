package unique_finder

import (
	"sort"
	"time"
)

const (
	TIME_LAYOUT = "2006-01-02"
)

type PhoneHistory struct {
	PhoneNumber string
	History     []*History
}

type History struct {
	ActivedAt     *time.Time
	DeactivatedAt *time.Time
}

func NewPhoneHistory(number string) *PhoneHistory {
	return &PhoneHistory{
		PhoneNumber: number,
		History:     make([]*History, 0),
	}
}

func NewHistory(activatedAt string, deactivatedAt string) *History {
	history := new(History)

	activeTime, err := time.Parse(TIME_LAYOUT, activatedAt)
	if err != nil {
		panic(err)
	}

	history.ActivedAt = &activeTime

	if deactivatedAt == "" {
		return history
	}

	deactiveTime, err := time.Parse(TIME_LAYOUT, deactivatedAt)
	if err != nil {
		panic(err)
	}
	history.DeactivatedAt = &deactiveTime
	return history
}

func (ph *PhoneHistory) LatestActivate() *time.Time {
	for i := 0; i < len(ph.History)-1; i++ {
		if !ph.History[i].ActivedAt.Equal(*ph.History[i+1].DeactivatedAt) {
			return ph.History[i].ActivedAt
		}
	}

	return ph.History[len(ph.History)-1].ActivedAt
}

func (ph *PhoneHistory) Insert(history *History) {
	i := sort.Search(len(ph.History), func(i int) bool {
		return history.ActivedAt.After(*(ph.History[i].ActivedAt))
	})
	ph.History = append(ph.History[:i], append([]*History{history}, ph.History[i:]...)...)
}
