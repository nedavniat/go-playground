//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
	"sync"
)

func producer(stream Stream, tweets chan *Tweet, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(tweets)
			return 
		}
		tweets <- tweet
	}
}

func consumer(tweets chan *Tweet, wg *sync.WaitGroup) {
	defer wg.Done()
	for t := range tweets {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(2)
	
	stream := GetMockStream()
	tweets := make(chan *Tweet, len(stream.tweets)) 

	// Producer
	go producer(stream, tweets, &wg)

	// Consumer
	go consumer(tweets, &wg)

	wg.Wait()
	fmt.Printf("Process took %s\n", time.Since(start))
}

//waitGroup 

//context
