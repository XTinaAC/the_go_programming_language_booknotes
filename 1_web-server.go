/*
	A minimal "echo" server
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mutex sync.Mutex
var count int

func main() {
	// connect a handler function to all incoming URLs
	http.HandleFunc("/", requestHandler)
	http.HandleFunc("/count", requestCounter)
	
	// start a server listening for incoming requests on port 8000
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// echo the Path component of the requested URL
func requestHandler(w http.ResponseWriter, r *http.Request) {
	// avoid the race condition using【mu.Lock】and【mu.Unlock】
	mutex.Lock()
	count++
	mutex.Unlock()

	// %q quoted <-> "string" %s or 'rune' %c
	fmt.Fprintf(w, "\nMethod, URL, Proto, Host, RemoteAddress => %q %q %q %q %q\n", r.Method, r.URL, r.Proto, r.Host, r.RemoteAddr)
	for key, value := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", key, value)
	}
	/*
	 【ParseForm】populates【r.Form】and【r.PostForm】.
		1）For all requests, ParseForm parses the【raw query】from the URL and updates【r.Form】；
		2）For POST, PUT, and PATCH requests, it also reads the【request body】
			parses it as a form，and puts the results into both【r.PostForm】and【r.Form】. 
			【Request body】 parameters take precedence over【URL query】string values in r.Form.
	  https://pkg.go.dev/net/http#Request.ParseForm

	 Like the【for】、【switch】statements, an【if】may include an optional【simple statement】
	 （a short variable declaration、an increment/assignment statement、a function call）
	 that can be used to set a value before it is tested
	*/
	if err := r.ParseForm(); err != nil {
	    log.Print(err)
	}
	for key, value := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", key, value)
	}
}

// echo the number of requests so far
func requestCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	fmt.Fprintf(w, "Number of Requests: %d\n", count)
	mutex.Unlock()
}