package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
)

func main() {
	//checkStatus := func(
	//	done <-chan interface{},
	//	urls ...string,
	//) <-chan *http.Response {
	//	responses := make(chan *http.Response)
	//	go func() {
	//		defer close(responses)
	//		for _, url := range urls {
	//			resp, err := http.Get(url)
	//			if err != nil {
	//				fmt.Println(err)
	//				continue
	//			}
	//			select {
	//			case <-done:
	//				return
	//			case responses <- resp:
	//			}
	//		}
	//	}()
	//	return responses
	//}
	//done := make(chan interface{})
	//defer close(done)
	//urls := []string{"https://www.google.com", "https://badhost"}
	//for response := range checkStatus(done, urls...) {
	//	fmt.Printf("Response: %v\n", response.Status)
	//}

	//type Result struct {
	//	Error error
	//	Response *http.Response
	//}
	//checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
	//	results := make(chan Result)
	//	go func() {
	//		defer close(results)
	//		for _, url := range urls {
	//			var result Result
	//			resp, err := http.Get(url)
	//			result = Result{Error: err, Response: resp}
	//			select {
	//			case <-done:
	//				return
	//			case results <- result:
	//			}
	//		}
	//	}()
	//	return results
	//}
	//done := make(chan interface{})
	//defer close(done)
	//urls := []string{"https://www.google.com", "https://badhost"}
	//for result := range checkStatus(done, urls...) {
	//	if result.Error != nil {
	//		fmt.Printf("error: %v", result.Error)
	//		continue
	//	}
	//	fmt.Printf("Response: %v\n", result.Response.Status)
	//}

	//type Result struct {
	//	Error    error
	//	Response *http.Response
	//}
	//
	//checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
	//	results := make(chan Result)
	//	go func() {
	//		defer close(results)
	//		for _, url := range urls {
	//			var result Result
	//			resp, err := http.Get(url)
	//			result = Result{Error: err, Response: resp}
	//			select {
	//			case <-done:
	//				return
	//			case results <- result:
	//			}
	//		}
	//	}()
	//	return results
	//}
	//
	//done := make(chan interface{})
	//defer close(done)
	//errCount := 0
	//urls := []string{"a", "https://www.google.com", "b", "c", "d", "e"}
	//for result := range checkStatus(done, urls...) {
	//	if result.Error != nil {
	//		fmt.Printf("error: %v\n", result.Error)
	//		errCount++
	//		if errCount >= 3 {
	//			fmt.Println("Too many errors, breaking!")
	//			break
	//		}
	//		continue
	//	}
	//	fmt.Printf("Response: %v\n", result.Response.Status)
	//}

	type Result struct {
		Error    error
		Response *http.Response
	}
	checkStatus := func(ctx context.Context, urls ...string) <-chan Result {
		results := make(chan Result)
		wg := sync.WaitGroup{}

		go func() {
			wg.Wait()
			close(results)
		}()

		for _, url := range urls {
			wg.Add(1)
			go func(ctx context.Context, url string) {
				defer wg.Done()
				var result Result
				resp, err := http.Get(url)
				result = Result{Error: err, Response: resp}
				select {
				case <-ctx.Done():
					fmt.Println(ctx.Err())
					return
				case results <- result:
				}
			}(ctx, url)
		}
		return results
	}

	urls := []string{"https://www.google.com", "https://badhost", "https://www.facebook.com"}
	ctx, cancel := context.WithCancel(context.Background())
	for result := range checkStatus(ctx, urls...) {
		if result.Error != nil {
			fmt.Printf("error: %v\n", result.Error)
			cancel()
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}
