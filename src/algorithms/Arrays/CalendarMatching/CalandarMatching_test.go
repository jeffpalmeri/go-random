package calendarmatching

import "testing"

// go test -v -run TestCalendarMatching ./... -count=1
func TestCalendarMatching(t *testing.T) {
	// calendar1 := [["9:00", "10:30"], ["12:00", "13:00"], ["16:00", "18:00"]]
	// dailyBounds1 := ["9:00", "20:00"]
	// calendar2 := [["10:00", "11:30"], ["12:30", "14:30"], ["14:30", "15:00"], ["16:00", "17:00"]]
	// dailyBounds2 := ["10:00", "18:30"]
	// meetingDuration := 30

	calendar1 := []StringMeeting{
		{"9:00", "10:30"},
		{"12:00", "13:30"},
		{"16:00", "18:00"},
	}
	calendar2 := []StringMeeting{
		{"10:00", "11:30"},
		{"12:30", "14:30"},
		{"14:30", "15:00"},
		{"16:00", "17:00"},
	}
	dailyBounds1 := StringMeeting{"9:00", "20:00"}
	dailyBounds2 := StringMeeting{"10:00", "18:30"}
	meetingDuration := 30

	CalendarMatching(calendar1, dailyBounds1, calendar2, dailyBounds2, meetingDuration)
}
