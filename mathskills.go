package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"
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

// roundToInt rounds a floa64 to an integer
func roundToInt(f float64) int {
	diff := f - float64(int(f))

	if diff >= 0.5 {
		return int(f) + 1
	} else if diff <= -0.5 {
		return int(f) - 1
	} else {
		return int(f)
	}
}

// average returns the mean of a slice of float64s
func average(d []float64) float64 {
	sum := 0.0
	for _, v := range d {
		sum += v
	}
	return sum / float64(len(d))
}

// median returns the median of a slice of float64s
func median(d []float64) float64 {
	if len(d)%2 == 0 {
		return (d[len(d)/2] + d[(len(d)/2)-1]) / 2
	} else {
		return d[len(d)/2]
	}
}

// variance returns the variance of a slice of float64s
func variance(d []float64) float64 {
	sumOfSqOfDiff := 0.0
	avg := average(d)
	for _, f := range d {
		sumOfSqOfDiff += (f - avg) * (f - avg)
	}
	return sumOfSqOfDiff / float64(len(d))
}

// sqrt calculates the square root of a float64 according to
// the Babylonian method a.k.a. Newton's method
func sqrt(x float64) float64 {
	if x < 0 {
		return -1
	}

	// start with a guess
	z := x
	const tolerance = 1e-10 // how precise you want the result to be
	for {
		nextZ := 0.5 * (z + x/z)
		if abs(z-nextZ) < tolerance { // Stop when the change is smaller than the tolerance
			break
		}
		z = nextZ
	}
	return z
}

// abs calculates the absolute value of a float64
func abs(f float64) float64 {
	if f < 0 {
		f *= -1
	}
	return f
}

// main prints out statistical values of a data set given as an argument
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

	fmt.Println("Average:", roundToInt(average(data)))
	fmt.Println("Median:", roundToInt(median(data)))
	fmt.Println("Variance:", roundToInt(variance(data)))
	fmt.Println("Standard deviation:", roundToInt(sqrt(variance(data))))
}
