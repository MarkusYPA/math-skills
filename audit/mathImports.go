package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"unicode"

	"github.com/montanaflynn/stats"
)

// readData returns a string as a slice of float64s
func readData(s string) (data []float64) {
	lines := []string{""}
	indx := 0
	for _, r := range s {
		if unicode.IsNumber(r) {
			lines[indx] += string(r)
		}

		if r == '\n' {
			lines = append(lines, "")
			indx++
		}
	}

	for _, s := range lines {
		num, e := strconv.ParseFloat(s, 64)
		if e != nil {
			fmt.Println(e)
			return nil
		}
		if s != "" {
			data = append(data, num)
		}
	}

	return
}

// main prints out statistical values of a data set given as an argument
// using well tested and reliable packages
func main() {
	args := os.Args[1:]
	var dataFile *os.File

	if len(args) == 1 {
		df, err := os.Open(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		dataFile = df
	} else {
		fmt.Println("Provide datafile")
		return
	}

	dataBytes, err := io.ReadAll(dataFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	data := readData(string(dataBytes))

	// Using a well known and tested package to create
	// reliable results for comparison
	avg, e1 := stats.Mean(data)
	med, e2 := stats.Median(data)
	vari, e3 := stats.Variance(data)
	stdd, e4 := stats.StandardDeviation(data)

	if e1 != nil || e4 != nil || e3 != nil || e2 != nil {
		fmt.Println("ERROR")
		return
	}

	// rounding with math.Round()
	fmt.Println("Average:", math.Round(avg))
	fmt.Println("Median:", math.Round(med))
	fmt.Println("Variance:", math.Round(vari))
	fmt.Println("Standard deviation:", math.Round(stdd))
}
