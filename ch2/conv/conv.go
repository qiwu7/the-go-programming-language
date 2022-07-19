package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	FEET_TO_METER = 0.3048
)

type Measurement interface {
	String() string
}

type Distance struct {
	number float64
	unit   string
}

func (d Distance) Meter() float64 {
	unit := strings.ToLower(d.unit)
	if unit == "ft" || unit == "feet" {
		return d.number * FEET_TO_METER
	}
	return d.number
}

func (d Distance) Feet() float64 {
	unit := strings.ToLower(d.unit)
	if unit == "meter" || unit == "m" {
		return d.number / FEET_TO_METER
	}
	return d.number
}

func (d Distance) String() string {
	return fmt.Sprintf("%.3gm = %.3gft", d.Meter(), d.Feet())
}

type Temperature struct {
	number float64
	unit   string
}

func (t Temperature) Celsius() float64 {
	unit := strings.ToLower(t.unit)
	if unit == "f" {
		return (t.number - 32) * 5 / 9
	}
	return t.number
}

func (t Temperature) Fahrenheit() float64 {
	unit := strings.ToLower(t.unit)
	if unit == "c" {
		return t.number*9/5 + 32
	}
	return t.number
}

func (t Temperature) String() string {
	return fmt.Sprintf("%.3gC = %.3gF", t.Celsius(), t.Fahrenheit())
}

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			printMeasurement(arg)
		}
	} else {
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			printMeasurement(scan.Text())
		}
	}
}

func printMeasurement(s string) {
	regex := regexp.MustCompile(`(-?\d+(?:\.\d+)?)([A-Za-z]+)`)
	match := regex.FindStringSubmatch(s)
	if match == nil {
		log.Fatalf("Expecting <number><unit>, got %q", s)
	}
	fmt.Println("matches", match)
	n, err := strconv.ParseFloat(match[1], 64)
	if err != nil {
		log.Fatalf("%v isn't a valid number", match[1])
	}
	if match[2] == "" {
		log.Fatalf("%v isn't a valid unit", match[2])
	}
	u := match[2]
	m, err := newMeasurement(n, u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)
}

func newMeasurement(n float64, u string) (Measurement, error) {
	u = strings.ToLower(u)
	switch u {
	case "c", "f":
		return Temperature{
			number: n,
			unit:   u,
		}, nil
	case "ft", "feet", "m", "meter":
		return Distance{
			number: n,
			unit:   u,
		}, nil
	default:
		return nil, fmt.Errorf("unknow conversion: %g%s", n, u)
	}
}
