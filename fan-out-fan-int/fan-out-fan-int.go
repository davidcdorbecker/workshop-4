package main

import (
	"fmt"
)

//func main() {
//	i1 := generateWork([]int{0, 2, 6, 8})
//	i2 := generateWork([]int{1, 3, 5, 7})
//
//	out := fanIn(i1, i2)
//
//	for value := range out {
//		fmt.Println("Value:", value)
//	}
//}
//
//func fanIn(inputs ...<-chan int) <-chan int {
//	var wg sync.WaitGroup
//	out := make(chan int)
//
//	wg.Add(len(inputs))
//
//	for _, in := range inputs {
//		go func(ch <-chan int) {
//			for {
//				value, ok := <-ch
//
//				if !ok {
//					wg.Done()
//					break
//				}
//
//				out <- value
//			}
//		}(in)
//	}
//
//	go func() {
//		wg.Wait()
//		close(out)
//	}()
//
//	return out
//}
//
//func generateWork(work []int) <-chan int {
//	ch := make(chan int)
//
//	go func() {
//		defer close(ch)
//
//		for _, w := range work {
//			ch <- w
//		}
//	}()
//
//	return ch
//}
//
func main() {
	work := []int{1, 2, 3, 4, 5, 6, 7, 8}
	in := generateWork(work)

	out1 := fanOut(in)
	out2 := fanOut(in)
	out3 := fanOut(in)
	out4 := fanOut(in)

	for range work {
		select {
		case value := <-out1:
			fmt.Println("Output 1 got:", value)
		case value := <-out2:
			fmt.Println("Output 2 got:", value)
		case value := <-out3:
			fmt.Println("Output 3 got:", value)
		case value := <-out4:
			fmt.Println("Output 4 got:", value)
		}
	}
}

func fanOut(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for data := range in {
			out <- data
		}
	}()

	return out
}

func generateWork(work []int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for _, w := range work {
			ch <- w
		}
	}()

	return ch
}