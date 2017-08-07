package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	seed := time.Now().UnixNano()
	rnd := rand.New(rand.NewSource(seed))
	count := 0
	for {
		rNum := int64(math.Floor(rnd.Float64() * 100))
		count++
		fmt.Printf("%d\t", rNum)

		if rNum < 10 {
			fmt.Println("single digit")
		} else if rNum < 51 {
			fmt.Println("not even half way there!")
		} else if rNum == 69 || rNum == 96 {
			fmt.Println("...snicker...")
		} else if rNum > 50 && rNum < 80 {
			fmt.Println("getting closer!")
		} else if rNum < 99 {
			fmt.Println("so, so close!")
		} else {
			fmt.Println("You Made it! MAX OUT!!!")
		}
		if rNum == 99 {
			break
		}
	}
	fmt.Printf("it took %d to get to 99!\n", count)
}
