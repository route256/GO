package main

import (
	"fmt"
	"time"
)

func main() {
	// simple()
	// sleepAndDuration()
	// parseAndFormat()
	// timeMath()
	test()
}

func simple() {
	t := time.Now()
	fmt.Println(t)
}

func sleepAndDuration() {
	t := time.Now()
	time.Sleep(time.Second * 2)
	elapsed := time.Since(t)

	fmt.Println(elapsed)
}

func parseAndFormat() {
	fmt.Println(
		time.Now().Format("2006-01-02 PM 03:04:05 Z07:00"),
	)

	loc := time.FixedZone("Moscow", 3*60*60)
	moscowTime, err := time.ParseInLocation(
		"2006-01-02 15:04:05",
		"2022-09-24 12:49:01",
		loc,
	)
	fmt.Println(moscowTime, err)

	targetLoc := time.FixedZone("NewYork", -4*60*60)
	newYorkTime := moscowTime.In(targetLoc)
	fmt.Println(newYorkTime)

	fmt.Println(moscowTime.Unix())
	fmt.Println(newYorkTime.Unix())

	fmt.Println(newYorkTime.Hour())
}

func timeMath() {
	now := time.Now()

	twoHoursAnd20minLater := now.Add(time.Hour*2 + time.Minute*20)
	tenMinutesBefore := now.Add(-time.Minute * 10)

	fmt.Println("now + 2h 20m:", twoHoursAnd20minLater)
	fmt.Println("         now:", now)
	fmt.Println("   now - 10m:", tenMinutesBefore)

	after2months15days := now.AddDate(0, 2, 15)
	oneYearBefore := now.AddDate(-1, 0, 0)

	fmt.Println()

	fmt.Println("         1 year before:", oneYearBefore.Format("2006-01-02"))
	fmt.Println("                   now:", now.Format("2006-01-02"))
	fmt.Println("after 2 months 15 days:", after2months15days.Format("2006-01-02"))

	fmt.Println()

	fmt.Println("             now - tenMinutesBefore:", now.Sub(tenMinutesBefore))
	fmt.Println("      now is after tenMinutesBefore:", now.After(tenMinutesBefore))
	fmt.Println("now is before twoHoursAnd20minLater:", now.Before(twoHoursAnd20minLater))

	fmt.Println()

	fmt.Println("      time.Minute < time.Hour:", time.Minute < time.Hour)
	fmt.Println("time.Second*120 > time.Minute:", time.Second*120 > time.Minute)
}

func test() {
	loc := time.FixedZone("Moscow", 3*60*60)
	moscowTime, err := time.ParseInLocation(
		"2006-01-02 15:04:05",
		"2022-01-31 12:49:01",
		loc,
	)
	fmt.Println(moscowTime, err)
	moscowTime = moscowTime.AddDate(1, 0, 0)
	fmt.Println(moscowTime)
}
