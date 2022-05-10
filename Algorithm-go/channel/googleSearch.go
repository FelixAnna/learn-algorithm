package channel

import (
	"fmt"
	"math/rand"
	"time"
)

const timeoutSec = 3

type Result string
type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(10) * int(time.Second)))
		return Result(fmt.Sprintf("%v result for \"%v\"\n", kind, query))
	}
}
func searchInReplics(query string, replicas ...string) Result {
	ch := make(chan Result)
	searchReplicas := func(i int) { ch <- fakeSearch(replicas[i])(query) }
	for i := range replicas {
		go searchReplicas(i)
	}

	return <-ch //only return first result
}

func SearchAll(keywords string) []Result {
	results := make([]Result, 0)
	c := make(chan Result)
	start := time.Now()
	go func() {
		c <- searchInReplics(keywords, "Web1", "Web2", "Web3")
	}()
	go func() {
		c <- searchInReplics(keywords, "Image1", "Image2", "Image3")
	}()
	go func() {
		c <- searchInReplics(keywords, "Video1", "Video2", "Video3")
	}()

	timeout := time.After(timeoutSec * time.Second)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return nil
		}
	}

	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
	return results
}
