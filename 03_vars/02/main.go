package main

import (
	"fmt"
)

func main() {
	// fmt.Print("Enter your full name: ")
	// in := bufio.NewReader(os.Stdin)
	// line, _ := in.ReadString('\n')
	// fmt.Printf("Hello, %s", line)

	var fullName [3]string
	fmt.Print("Enter First, Middle Last: ")
	fmt.Scan(&fullName[0], &fullName[1], &fullName[2])
	for i := 0; i < len(fullName); i++ {
		fmt.Println(fullName[i])
	}
}
