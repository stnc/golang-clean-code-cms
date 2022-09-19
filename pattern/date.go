package main

import (
	"fmt"
	"time"
)

func main() {

	curr_date := time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC)

	oneDayLater := curr_date.AddDate(0, 0, 1)

	oneMonthLater := curr_date.AddDate(0, 1, 0)

	oneYearLater := curr_date.AddDate(1, 0, 0)

	oneYearBack := curr_date.AddDate(-1, 0, 0)

	fmt.Println("Current date: ", curr_date)

	fmt.Println("oneDayLater: ", oneDayLater)

	fmt.Println("oneMonthLater: ", oneMonthLater)

	fmt.Println("oneYearLater: ", oneYearLater)

	fmt.Println("oneYearBack: ", oneYearBack)

}
