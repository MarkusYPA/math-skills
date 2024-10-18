package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"unicode"
)

// readData returns a string as a slice of float64s
func readData(s string) (data []float64) {
	lines := []string{""}
	indx := 0

	for _, r := range s {
		if unicode.IsNumber(r) || r == '-' {
			lines[indx] += string(r)
		}

		if r == '\n' {
			lines = append(lines, "")
			indx++
		}
	}

	for _, s := range lines {
		if len(s) == 0 {
			continue
		}
		num, e := strconv.ParseFloat(s, 64)
		if e != nil {
			continue
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

// mean returns the mean of a slice of float64s
func mean(d []float64) float64 {
	sum := 0.0
	for _, v := range d {
		sum += v
	}
	return sum / float64(len(d))
}

// median returns the median of a slice of float64s
func median(d []float64) float64 {
	ds := bubSort(d)
	if len(ds)%2 == 0 {
		return (ds[len(d)/2] + ds[(len(d)/2)-1]) / 2
	} else {
		return ds[len(d)/2]
	}
}

// bubSort is a bubble sort function that returns an new
// slice of float64s d, sorted from smallest to largest
func bubSort(d []float64) []float64 {
	ds := make([]float64, len(d))
	copy(ds, d)

	for i := 0; i < len(ds)-1; i++ {
		for j := i + 1; j < len(ds); j++ {
			if ds[i] > ds[j] {
				ds[i], ds[j] = ds[j], ds[i]
			}
		}
	}
	return ds
}

// variance returns the variance of a slice of float64s
func variance(d []float64) float64 {
	sumOfSqOfDiff := 0.0
	avg := mean(d)
	for _, f := range d {
		sumOfSqOfDiff += (f - avg) * (f - avg)
	}
	return sumOfSqOfDiff / float64(len(d))
}

// sqrt calculates the square root of a float64 according to
// the Babylonian method a.k.a. Newton's method
func sqrt(x float64) float64 {
	if x < 0 {
		return -1.0
	}
	if x == 0 {
		return 0.0
	}

	// start with a guess
	guess := x
	const tolerance = 1e-10 // precision is 0.0000000001

	for {
		nextG := 0.5 * (guess + x/guess)  // (x+1)/2 on the first try
		if abs(guess-nextG) < tolerance { // Stop when the change is smaller than the tolerance
			break
		}
		guess = nextG
	}
	return guess
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

	if len(args) != 1 {
		fmt.Println("Provide datafile")
		return
	}

	dataFile, err := os.Open(args[0])
	if err != nil {
		log.Fatalln(err.Error())
	}
	dataBytes, err := io.ReadAll(dataFile)
	if err != nil {
		log.Fatalln(err.Error())
	}
	data := readData(string(dataBytes))

	if len(data) == 0 {
		log.Fatalln("No valid data found")
	}

	fmt.Println("Average:", roundToInt(mean(data)))
	fmt.Println("Median:", roundToInt(median(data)))
	fmt.Println("Variance:", roundToInt(variance(data)))
	fmt.Println("Standard deviation:", roundToInt(sqrt(variance(data))))
}
