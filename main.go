/*Had to be written by Ricky lol*/
package main

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"strings"
	"time"
)

var (
	A, B        float64
	clearScreen = "\x1b[2J"
	homeCursor  = "\x1b[H"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Print(clearScreen + homeCursor)
		fmt.Println("You are an idiot sandwich. ~ Gordon Ramsay")
		os.Exit(0)
	}()

	fmt.Print(clearScreen)
	for {
		z := make([]float64, 1760)
		b := make([]byte, 1760)
		for i := range b {
			b[i] = ' '
		}

		for j := 0.0; j < 6.28; j += 0.07 {
			for i := 0.0; i < 6.28; i += 0.02 {
				sinA := math.Sin(A)
				cosA := math.Cos(A)
				sinB := math.Sin(B)
				cosB := math.Cos(B)
				sini := math.Sin(i)
				cosi := math.Cos(i)
				sinj := math.Sin(j)
				cosj := math.Cos(j)

				h := cosj + 2
				D := 1 / (sini*h*sinA + sinj*cosA + 5)
				t := sini*h*cosA - sinj*sinA

				x := int(40 + 30*D*(cosi*h*cosB-t*sinB))
				y := int(12 + 15*D*(cosi*h*sinB+t*cosB))
				o := x + 80*y
				N := int(8 * ((sinj*sinA-sini*cosj*cosA)*cosB - sini*cosj*sinA - sinj*cosA - cosi*cosj*sinB))

				if 0 <= y && y < 22 && 0 <= x && x < 80 && D > z[o] {
					z[o] = D
					if N > 0 {
						b[o] = ".,-~:;=!*#$@"[N%12]
					} else {
						b[o] = '.'
					}
				}
			}
		}

		fmt.Print(homeCursor)
		var builder strings.Builder
		for k := 0; k < 1760; k++ {
			if k%80 == 0 {
				builder.WriteByte('\n')
			}
			builder.WriteByte(b[k])
		}
		fmt.Print(builder.String())

		A += 0.04
		B += 0.02

		time.Sleep(time.Millisecond * 16)
	}
}
