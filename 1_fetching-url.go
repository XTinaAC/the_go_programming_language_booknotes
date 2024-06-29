/*
	Print the content (or content size) and elapsed time for URLs-fetching.
	
	并发、顺序：concurrent <-> sequential		（时间段内是否同时进行）
	异步、同步：asynchronous <-> synchronous	（函数执行是否立即返回）
	并行、串行：parallel <-> serial			（时间点上是否同时进行）(多核硬件支持)
*/

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	// 顺序式
	sequential_fetch()
	// 并发式
	concurrent_fetch()
}

func sequential_fetch() {
	start := time.Now()
	for idx, url := range os.Args[1:] {
		sub_start := time.Now()

		// Add the url-protocol prefix if it is missing
		const prefix = "http://"
		if !strings.HasPrefix(url, prefix) {
			url = prefix + url
		}

		//【http.Get】makes an HTTP request
		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error during fetching: %v\n", err)
			os.Exit(1)
		}

		// the【Body】field of resp is a readable stream
		respBody := resp.Body
		// read the entire response using【ioutil.ReadAll】
		data, err := ioutil.ReadAll(respBody)
		// close the stream to avoid leakign resources
		respBody.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error during reading: %v\n", err)
			os.Exit(1)
		}

		// time cost of each task
		fmt.Printf("【%d】Status: %s Time: %.2fs elapsed (%s)\n Content: %s ...\n", idx, resp.Status, time.Since(sub_start).Seconds(), url, data[:20])
	}
	
	// total time cost
	fmt.Printf("in total: %.2fs elapsed\n", time.Since(start).Seconds())
}

func concurrent_fetch() {
	start := time.Now()
	// create a【channel】of string using【make】
	ch := make(chan string)

	for idx, url := range os.Args[1:] {
		// start a【goroutine】
		go fetch(url, ch, idx)
	}
	for range os.Args[1:] {
		// receive from【channel】ch
		fmt.Println(<-ch)
	}

	// total time cost
	fmt.Printf("in total: %.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan <- string, idx int) {
	sub_start := time.Now()
		
	// Add the url-protocol prefix if it is missing
	const prefix = "http://"
	if !strings.HasPrefix(url, prefix) {
		url = prefix + url
	}

	//【http.Get】makes an HTTP request
	resp, err := http.Get(url)

	if err != nil {
		// send to【channel】ch
		ch <- fmt.Sprint(err)
		return
	}

	// the【Body】field of resp is a readable stream
	respBody := resp.Body
	// discard the response and report the size instead
	nbytes, err := io.Copy(ioutil.Discard, respBody)
	// close the stream to avoid leakign resources
	respBody.Close()

	// time cost of each task
	ch <- fmt.Sprintf("【%d】Status: %s Time: %.2fs elapsed (%s)\n Count: (%7d bytes)\n", idx, resp.Status, time.Since(sub_start).Seconds(), url, nbytes)
}
