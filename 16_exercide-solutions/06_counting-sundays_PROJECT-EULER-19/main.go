/*
https://projecteuler.net/problem=19

You are given the following information, but you may prefer to do some research for yourself.

1 Jan 1900 was a Monday.
Thirty days has September,
April, June and November.
All the rest have thirty-one,
Saving February alone,
Which has twenty-eight, rain or shine.
And on leap years, twenty-nine.
A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
*/

package main

import "fmt"

func main() {
	var answerCounter int
	var year int
	var firstSunday int = 7
	for year = 1900; year <= 2000; year++ {
		var daysInYear int

		// Check if leap year
		if year%100 == 0 && year%400 == 0 {
			daysInYear = 366
		} else if year%100 != 0 && year%4 == 0 {
			daysInYear = 366
		} else {
			daysInYear = 365
		}

		var listOfSundays []int
		listOfSundays = append(listOfSundays, firstSunday)

		// Create a list of days that are Sundays in that year out of daysInYear(365 or 366)
		nextSunday := firstSunday + 7
		for nextSunday <= daysInYear {
			listOfSundays = append(listOfSundays, nextSunday)
			nextSunday += 7
		}

		var listFirstDaysOfMonth [12]int
		var listdaysInMonths [12]int

		// Define days in each month
		if daysInYear == 365 {
			listdaysInMonths = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
		} else {
			listdaysInMonths = [12]int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
		}

		var counter int = 0

		listFirstDaysOfMonth[0] = 1 // First day of January
		counter = 1                 // Stores prev month first day value

		for i := 1; i < len(listdaysInMonths); i++ {
			listFirstDaysOfMonth[i] = listdaysInMonths[i-1] + counter
			counter = listFirstDaysOfMonth[i]
		}

		if year == 1900 {
			answerCounter = 0
		} else {
			// Check if first days of months match with any sundays in that year
			for i := 0; i < 12; i++ {
				checkSundayOnFirst := BinarySearch(listOfSundays, listFirstDaysOfMonth[i])
				if checkSundayOnFirst == -1 {
					continue
				} else {
					answerCounter += 1
				}
			}
		}

		// Calculate first sunday for next year
		lastSundayOfYear := listOfSundays[len(listOfSundays)-1]
		firstSunday = 7 - (daysInYear - lastSundayOfYear)

		//fmt.Println(year, ": ", lastSundayOfYear, ": ", listFirstDaysOfMonth, "|||", listOfSundays) 	// for testing
	}

	fmt.Println(answerCounter)
}

// https://gist.github.com/emre/6340638#file-golang_binary_search-go
// Used Binary search to gain search efficiency as compared to sequential search.
func BinarySearch(target_map []int, value int) int {

	start_index := 0
	end_index := len(target_map) - 1

	for start_index <= end_index {
		median := (start_index + end_index) / 2

		if target_map[median] < value {
			start_index = median + 1
		} else {
			end_index = median - 1
		}
	}

	if start_index == len(target_map) || target_map[start_index] != value {
		return -1
	} else {
		return start_index
	}
}
