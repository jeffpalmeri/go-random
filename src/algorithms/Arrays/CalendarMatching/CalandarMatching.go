package calendarmatching

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type StringMeeting struct {
	Start string
	End   string
}

func CalendarMatching(
	calendar1 []StringMeeting, dailyBounds1 StringMeeting,
	calendar2 []StringMeeting, dailyBounds2 StringMeeting,
	meetingDuration int,
) []StringMeeting {
	a := mergeIntervals(calendar1, calendar2)
	fmt.Println("Merged: ", a)

	res := [][]int{}
	for i := 0; i < len(a)-1; i++ {
		if a[i+1][0]-a[i][0] <= meetingDuration {
			res = append(res, []int{a[i][1], a[i+1][0]})
		}
	}
	fmt.Println("Times: ", res)
	return calendar1
}

func mergeIntervals(a, b []StringMeeting) [][]int {
	a2 := convertStringsToHours(a)
	b2 := convertStringsToHours(b)
	newArr := append(a2, b2...)
	if len(newArr) == 0 {
		return newArr

	}
	sort.Slice(newArr, func(i, j int) bool {
		return newArr[i][0] < newArr[j][0]
	})
	// Ok, the intevals are finally combined and sorted at this point
	// Need to actually merge them next
	ok := [][]int{newArr[0]}
	fmt.Println("Sorted: ", newArr)
	j := 0
	for i := 1; i < len(newArr); i++ {
		if ok[j][1] >= newArr[i][0] {
			ok[j][1] = max(newArr[i][0], newArr[i][1])
		} else {
			ok = append(ok, newArr[i])
			j++
		}
	}

	return ok
}

func hourValue(a string) int {
	colonIndex := strings.Index(a, ":")
	if colonIndex == -1 {
		return -1
	}
	intVersion, err := strconv.Atoi(a[:colonIndex])
	if err != nil {
		return -1
	}
	afterColon := a[colonIndex+1:]
	fmt.Println("After colon: ", afterColon)
	switch afterColon {
	case "00":
		fmt.Println("00 case")
		// Add nothing
	case "15":

	case "30":
		fmt.Println("30 case")
		// Add 0.50 to the number
	case "45":
	}
	return intVersion
}

func convertStringsToHours(a []StringMeeting) [][]int {
	a2 := make([][]int, len(a))
	for i, val := range a {
		a2[i] = []int{hourValue(val.Start), hourValue(val.End)}
	}
	return a2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

