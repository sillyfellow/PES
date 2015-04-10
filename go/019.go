/*
 *You are given the following information, but you may prefer to do some research for yourself.
 *
 *1 Jan 1900 was a Monday.
 *Thirty days has September,
 *April, June and November.
 *All the rest have thirty-one,
 *Saving February alone,
 *Which has twenty-eight, rain or shine.
 *And on leap years, twenty-nine.
 *A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
 *How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
 */

package main

import "fmt"

type DateCustomStruct struct {
	d int
	m int
	y int
}

func (date *DateCustomStruct) dateToDay() int {
	mmap := map[int]int{1: 0, 2: 3, 3: 3, 4: 6, 5: 1, 6: 4, 7: 6, 8: 2, 9: 5, 10: 0, 11: 3, 12: 5}
	d := date.d
	m := mmap[date.m]
	y := (date.y % 100)
	c := 6 - 2*((date.y-y)%400)/100
	adjust := 0
	if y%4 == 0 && date.m < 3 {
		adjust = -1
	}
	return (adjust + d + m + y + (y / 4) + c) % 7
}

func findFirstOfMonthOnSpecificCount(specificDay int, start, end *DateCustomStruct) int {
	count := 0
	m, y := start.m, start.y
	months := 12
	for y <= end.y {
		if y == end.y {
			months = end.m
		}
		for m <= months {
			date := &DateCustomStruct{1, m, y}
			day := date.dateToDay()
			if day == specificDay {
				count += 1
			}
			m = m + 1
		}
		y = y + 1
		m = 1 // new year, start at jan.
	}
	return count
}

func main() {
	start := &DateCustomStruct{1, 1, 1901}
	end := &DateCustomStruct{31, 12, 2000}
	fmt.Println(findFirstOfMonthOnSpecificCount(0, start, end))
}
