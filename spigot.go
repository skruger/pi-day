package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// PiDigits returns the first n digits of pi (including the leading "3")
// using the Rabinowitz–Wagon uncompressed spigot algorithm.
func PiDigits(n int) string {
	if n <= 0 {
		return ""
	}
	// Need about 10n/3 mixed-radix digits.
	length := 10*n/3 + 1
	a := make([]int, length)
	for i := range a {
		a[i] = 2
	}

	var out strings.Builder
	nines := 0
	predigit := 0
	havePre := false
	emit := func(d int) {
		out.WriteByte(byte('0' + d))
	}

	for j := 0; j < n+1; j++ {
		q := 0
		for i := length; i > 0; i-- {
			x := 10*a[i-1] + q*i
			a[i-1] = x % (2*i - 1)
			q = x / (2*i - 1)
		}
		a[0] = q % 10
		q /= 10

		switch {
		case q == 9:
			nines++
		case q == 10:
			if havePre {
				emit(predigit + 1)
			}
			for k := 0; k < nines; k++ {
				emit(0)
			}
			predigit, havePre, nines = 0, true, 0
		default:
			if havePre {
				emit(predigit)
			}
			for k := 0; k < nines; k++ {
				emit(9)
			}
			predigit, havePre, nines = q, true, 0
		}
	}
	// The final predigit is a guard digit; discard it.
	s := out.String()
	if len(s) > n {
		s = s[:n]
	}
	return s
}

func main() {
	n := flag.Int("n", 50, "number of digits of pi to compute (including the leading 3)")
	flag.Parse()
	if *n < 1 {
		fmt.Fprintln(os.Stderr, "n must be >= 1")
		os.Exit(1)
	}
	s := PiDigits(*n)
	if len(s) > 1 {
		fmt.Printf("%s.%s\n", s[:1], s[1:])
	} else {
		fmt.Println(s)
	}
}
