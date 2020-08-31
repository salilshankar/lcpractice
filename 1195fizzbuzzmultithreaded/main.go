//Contains all ways you can use to synchronize goroutines. Either use a waitgroup, or a counter.
//Personally, the counter method seems more intuitive.

// Solution for leetcode #1195. Fizz Buzz Multithreaded using Golang

package main

import (
	"fmt"
)

type fizzbuzz struct {
	index int
	value string
}

//var wg sync.WaitGroup

func main() {
	strs := fizzBuzz(100)

	for _, str := range strs {
		fmt.Println(str)
	}

}

func fizzBuzz(amount int) []string {

	fbch := make(chan fizzbuzz, amount)

	arr := make([]string, amount)

	//wg.Add(4)

	go number(fbch, amount)
	go fizz(fbch, amount)
	go buzz(fbch, amount)
	go fizzBUZZ(fbch, amount)

	//wg.Wait()

	close(fbch)

	// ctr := 1

	// for fbz := range fbch {
	// 	arr[fbz.index-1] = fbz.value
	// 	if ctr == amount {
	// 		break
	// 	}
	// 	ctr++
	// }

	var fbz fizzbuzz

	for ctr := 1; ctr <= amount; ctr++ {
		fbz = <-fbch
		arr[fbz.index-1] = fbz.value
	}

	// for fbz := range fbch {
	// 	arr[fbz.index-1] = fbz.value
	// }

	return arr

}

func fizz(fbch chan fizzbuzz, amount int) {
	for i := 3; i <= amount; i += 3 {
		str := fizzbuzz{
			i,
			"fizz",
		}
		if i%5 != 0 {
			fbch <- str
		}
	}
	//wg.Done()
}

func buzz(fbch chan fizzbuzz, amount int) {
	for i := 5; i <= amount; i += 5 {
		str := fizzbuzz{
			i,
			"buzz",
		}
		if i%3 != 0 {
			fbch <- str
		}
	}
	//wg.Done()
}

func fizzBUZZ(fbch chan fizzbuzz, amount int) {
	for i := 15; i <= amount; i += 15 {
		str := fizzbuzz{
			i,
			"fizzbuzz",
		}
		fbch <- str
	}
	//wg.Done()
}

func number(fbch chan fizzbuzz, amount int) {
	for i := 1; i <= amount; i++ {
		if (i%5) != 0 && (i%3) != 0 {
			str := fizzbuzz{
				i,
				fmt.Sprintf("%v", i),
			}
			fbch <- str
		}
	}
	//wg.Done()
}
