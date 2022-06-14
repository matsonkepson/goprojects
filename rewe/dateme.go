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
	yyyy := dt.Format("2006")
	yyyy_toint, _ := strconv.Atoi(yyyy)
	yyyy_prev := yyyy_toint - 1
	mm := dt.Format("01")
	mm_toint, _ := strconv.Atoi(mm)
	mm_prev := mm_toint - 01

	fmt.Println("Current year: ", dt.Format("2006"))
	fmt.Println("Current month: ", dt.Format("01"))
	fmt.Printf("Print a previous Year: %04d\n", yyyy_prev)
	fmt.Printf("Print a previous Month: %02d\n", mm_prev)
}
