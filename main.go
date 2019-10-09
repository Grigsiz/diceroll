package main
/*

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
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

func createBuffer()  {
	file, err := os.Create("some.dat")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
}

func writeBuffer(a int)  {
	file, err := os.Open("some.data")
	writer := bufio.NewWriter(file)
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}
	defer file.Close()

	_, _ = writer.WriteString(string(a))
	_, _ = writer.WriteString("\n")

}

func readBufferAndCheck(a int) int {
	file, err := os.Open("some.data")
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return 0
	}
	defer file.Close()
	fl := 0
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return 0
			}
		}
		if line == string(a){
			fl += 1
		}
		return fl
	}
	return 0
}

func main() {
	createBuffer()
	var v int
	var f, ableTo bool
	flag.Parse()
	if *seed == 0 {
		rand.Seed(time.Now().Unix())
	} else {
		rand.Seed(int64(*seed))
	}

	for i := 0; i < *n; i++ {
		v, f, ableTo = randInterval(*start, *end, *n)
		if f == true {
			fmt.Print("Error")
		} else {
			// Dice roll 1..6
			if *norepeat && ableTo {
				if readBufferAndCheck(v) != 0 {
					writeBuffer(v)
					fmt.Println(v)
				} else {
					for readBufferAndCheck(v) == 0{
						v, f, ableTo = randInterval(*start, *end, *n)
					}
					writeBuffer(v)
					fmt.Println(v)
				}
			} else {
				fmt.Println(v)
			}
		}
	}
	if ableTo == false && *norepeat{
		fmt.Println("IMPOSSIBLE TO BE UNREPEATABLE , U KNOW")
	}
}
*/
