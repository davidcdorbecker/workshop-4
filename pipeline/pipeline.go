package main

import (
	"fmt"
	"time"
)

func main() {
	//multiply := func(values []int, multiplier int) []int {
	//	multipliedValues := make([]int, len(values))
	//	for i, v := range values {
	//		multipliedValues[i] = v * multiplier
	//	}
	//	return multipliedValues
	//}
	//
	//add := func(values []int, additive int) []int {
	//	addedValues := make([]int, len(values))
	//	for i, v := range values {
	//		addedValues[i] = v + additive
	//	}
	//	return addedValues
	//}
	//
	//ints := []int{1, 2, 3, 4}
	//for _, v := range add(multiply(ints, 2), 1) {
	//	fmt.Println(v)
	//}

	//start := time.Now()
	//
	//multiply := func(value, multiplier int) int {
	//	time.Sleep(500 * time.Millisecond)
	//	return value * multiplier
	//}
	//add := func(value, additive int) int {
	//	time.Sleep(500 * time.Millisecond)
	//	return value + additive
	//}
	//ints := []int{1, 2, 3, 4, 5, 6, 7, 8}
	//for _, v := range ints {
	//	fmt.Println(multiply(add(multiply(v, 2), 1), 2))
	//}
	//
	//finish := time.Now()
	//fmt.Printf("program takes %v seconds", finish.Sub(start).Seconds())

	start := time.Now()
	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
				case <-done:
					return
				case intStream <- i:
				}
			}
		}()
		return intStream
	}
	multiply := func(
		done <-chan interface{},
		intStream <-chan int,
		multiplier int,
	) <-chan int {
		multipliedStream := make(chan int)
		go func() {
			defer close(multipliedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				default:
					time.Sleep(500 * time.Millisecond)
					multipliedStream <- i*multiplier
				}
			}
		}()
		return multipliedStream
	}
	add := func(
		done <-chan interface{},
		intStream <-chan int,
		additive int,
	) <-chan int {
		addedStream := make(chan int)
		go func() {
			defer close(addedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				default:
					time.Sleep(500 * time.Millisecond)
					addedStream <- i+additive
				}
			}
		}()
		return addedStream
	}
	done := make(chan interface{})
	defer close(done)
	intStream := generator(done, 1, 2, 3, 4, 5, 6, 7, 8)
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)
	for v := range pipeline {
		fmt.Println(v)
	}

	finish := time.Now()
	fmt.Printf("program takes %v seconds", finish.Sub(start).Seconds())
}
