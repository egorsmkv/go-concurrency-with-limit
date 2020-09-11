package main

import (
	"fmt"
	"time"
)

func main() {
	timeStart := time.Now()
	results := Process()
	timeEnd := time.Now()

	timeDiff := timeEnd.Sub(timeStart)

	fmt.Printf("Number of results: %d\n", len(results))
	fmt.Println(results)

	fmt.Println()
	fmt.Println("Time start:", timeStart)
	fmt.Println("Time end:", timeEnd)
	fmt.Println("Time difference:", timeDiff)
}
