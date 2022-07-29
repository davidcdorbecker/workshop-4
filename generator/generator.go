package main

import (
	"fmt"
	"sync"
	"time"
)

//func main() {
//	ch := generator()
//
//	for i := 0; i < 5; i++ {
//		value := <-ch
//		fmt.Println("Value:", value)
//	}
//}
//
//func generator() <-chan int {
//	ch := make(chan int)
//
//	go func() {
//		for i := 0; ; i++ {
//			ch <- i
//		}
//	}()
//
//	return ch
//}

type genericFn func()

func main() {

	ch := generator()
	for str := range ch {
		fmt.Println(str)
	}
}

func generator() <-chan string {
	ch := make(chan string)
	wg := sync.WaitGroup{}

	veryExpensiveFns := func() []genericFn {
		var fns []genericFn
		for i := 1; i <= 3; i++ {
			func(delay time.Duration) {
				fns = append(fns, func() {
					time.Sleep(delay)
				})
			}(time.Duration(i) * time.Second)
		}
		return fns
	}()

	for _, fn := range veryExpensiveFns {
		wg.Add(1)
		go func(fn genericFn) {
			defer wg.Done()
			start := time.Now()
			fn()
			finish := time.Now()
			ch <- fmt.Sprintf("after %v", finish.Sub(start).Seconds())
		}(fn)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}
