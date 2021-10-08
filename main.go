package main

import (
	// "bufio"
	"fmt"
	"math/rand"
	// "os"
	"time"
	// "myapp/mylogger"
)

func main() {
	// basic for loop
	for i := 0; i <= 10; i++ {
		fmt.Println("i is", i)
	}

	// 'while' for loop
	rand.Seed(time.Now().UnixNano())
	i := 1000
	for i > 100 {
		// get a random number between 1 and 1001
		i = rand.Intn(1000)

		fmt.Println("i is ", i)
	}

	// infinite loop (in more detail)
	// read input from user 5 times and write to log
	// reader := bufio.NewReader(os.Stdin)
	// ch := make(chan string)

	// go mylogger.ListenForLog(ch)

	// fmt.Println("Enter something")

	// for i := 0; i < 5; i++ {
	// 	fmt.Print("-> ")
	// 	input, _ := reader.ReadString('\n')
	// 	ch <- input
	// 	time.Sleep(time.Second)
	// }

	// nested loops
	for i := 1; i <= 10; i++ {
		fmt.Print("i is ", i)
		
		for j := 1; j <= 3; j++ {
			fmt.Print("   j: ", j)
		}

		fmt.Println()
	}
}