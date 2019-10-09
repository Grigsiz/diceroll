package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	seed = flag.Int("seed", 0, "seed for random generator. unix(now) be default")
	start = flag.Int("start", 1,"first rand operand")
	end = flag.Int("end", 6,"second rand operand")
	n = flag.Int("n", 1,"cnt")
	norepeat = flag.Bool("norepeat", false,"must not be repeated")
)

// Фукнция должна вернуть число из интервала [l,r]
func randInterval(a,b,n int) (int,bool,bool){
	if a>b {
		return 0,true,false
	} else {
		c := b-a
		if c < n{
			return rand.Intn(c+1) + a,false,false
		} else {
			return rand.Intn(c+1) + a,false,true
		}
	}
}

func is_repeat(a,b int) bool {

}

func main() {
	var v int
	var f, ableTo bool
	flag.Parse()
	if *seed == 0 {
		rand.Seed(time.Now().Unix())
	} else {
		rand.Seed(int64(*seed))
	}
	const N  = *n
	var x [N]int
	for i := 0; i < *n; i++ {
		v, f, ableTo = randInterval(*start, *end, *n)
		if f == true {
			fmt.Print("Error")
		} else {
			// Dice roll 1..6
			if *norepeat && ableTo {

				fmt.Println(v)
			} else {
				fmt.Println(v)
		}
	}
	}
	if ableTo == false && *norepeat{
		fmt.Println("IMPOSSIBLE TO BE UNREPEATABLE , U KNOW")
	}
}
