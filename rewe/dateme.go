// Golang program to show
// the use of time.Now() Function
package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	dt := time.Now()
	year := dt.Format("2006")
	year_toint, _ := strconv.Atoi(year)
	year_prev := year_toint - 1
	mont := dt.Format("01")
	mont_toint, _ := strconv.Atoi(mont)
	mont_prev := mont_toint - 01

	fmt.Println("Current year: ", dt.Format("2006"))
	fmt.Println("Current month: ", dt.Format("01"))
	fmt.Printf("Print a previous Year: %04d\n", year_prev)
	fmt.Printf("Print a previous Month: %02d\n", mont_prev)
}
