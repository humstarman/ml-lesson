package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var n = flag.Int("n", 100, "numbers")
var w2 = flag.Int("2", 1, "w2")
var w1 = flag.Int("1", 1, "w1")
var w0 = flag.Int("0", 0, "w0")

func init() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var x2a []int
	var x1a []int
	var ya []int
	x0 := 1
	for i := 0; i < *n; i++ {
		x2 := rand.Intn(200) - 100
		x1 := rand.Intn(200) - 100
		x2a = append(x2a, x2)
		x1a = append(x1a, x1)
		y := *w2*x2 + *w1*x1 + *w0*x0
		if y > 0 {
			ya = append(ya, 1)
		} else {
			ya = append(ya, -1)
		}
	}
	a := 0
	b := 0
	c := 0
	for {
		e := 0
		for i := 0; i < *n; i++ {
			x2 := x2a[i]
			x1 := x1a[i]
			y := ya[i]
			h := a*x2 + b*x1 + c*x0
			if h*y <= 0 {
				a += y * x2
				b += y * x1
				c += y * x0
				e++
			}
		}
		if e == 0 {
			break
		}
	}
	fmt.Printf("origin w2:%v w1:%v w0:%v\n", *w2, *w1, *w0)
	fmt.Printf("current w2:%v w1:%v w0:%v\n", a, b, c)
}
