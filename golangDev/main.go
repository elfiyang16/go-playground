package main

import (
	"log"
	"net/http"
	"time"
)

const (
	numPollers     = 2                // number of Poller goroutines to launch
	pollInterval   = 60 * time.Second // how often to poll each URL
	statusInterval = 10 * time.Second // how often to log status to stdout
	errTimeout     = 10 * time.Second // back-off timeout on error
)

var urls = []string{
	"http://www.google.com/",
	"http://golang.org/",
	"http://blog.golang.org/",
}

// State represents the last-known state of a URL.
type State struct {
	url    string
	status string
}

// StateMonitor maintains a map that stores the state of the URLs being
// polled, and prints the current state every updateInterval nanoseconds.
// It returns a chan State to which resource state should be sent.
func StateMonitor(updateInterval time.Duration) chan<- State { // 本质是expose a setter for State to be set by the outside
	// 比return instance of State 要好，
	updates := make(chan State)
	urlStatus := make(map[string]string) // 注意Map 一直属于state monitor, 没有被share.  ensuring that it can only be accessed sequentially.
	// This prevents memory corruption issues that might arise from parallel reads and/or writes to a shared map.
	ticker := time.NewTicker(updateInterval)
	// IIFE
	go func() {
		for {
			select { //哪个ready 就做哪个。并行的, blocks until one of its communications is ready to proceed.
			case <-ticker.C: // A time.Ticker is an object that repeatedly sends a value on a channel at a specified interval.
				logState(urlStatus)
			case s := <-updates:
				urlStatus[s.url] = s.status
			}
		}
	}()
	return updates
}

// logState prints a state map.
func logState(s map[string]string) {
	log.Println("Current state:")
	for k, v := range s {
		log.Printf(" %s %s", k, v)
	}
}

// Resource represents an HTTP URL to be polled by this program.
type Resource struct {
	url      string
	errCount int
}

// Poll executes an HTTP HEAD request for url
// and returns the HTTP status string or an error string.
func (r *Resource) Poll() string {
	resp, err := http.Head(r.url)
	if err != nil {
		log.Println("Error", r.url, err)
		r.errCount++
		return err.Error() // rethrow the Error?
	}
	r.errCount = 0
	return resp.Status
}

// Sleep sleeps for an appropriate interval (dependent on error state)
// before sending the Resource to done.
//Classic: a function intended to run inside a goroutine takes a channel,
//upon which it sends its return value (or other indication of completed state).
func (r *Resource) Sleep(done chan<- *Resource) { // takes A send only channel
	time.Sleep(pollInterval + errTimeout*time.Duration(r.errCount))
	done <- r
}

func Poller(in <-chan *Resource, out chan<- *Resource, status chan<- State) {
	// in: receive only; out: send only; status: send only
	for r := range in { // pull requests from the "in" channel
		s := r.Poll()             // return result s
		status <- State{r.url, s} // send state into status, a write only channel, 然后会被Statemonitor 的update 接受
		out <- r                  // throw the request out
	}
}

func main() {
	// Create our input and output channels.
	pending, complete := make(chan *Resource), make(chan *Resource) // default to bilateral

	// Launch the StateMonitor.
	status := StateMonitor(statusInterval) // status: a send only channel

	// Launch some Poller goroutines.
	for i := 0; i < numPollers; i++ { // 对于限定数目launch goroutines
		go Poller(pending, complete, status) // in, out, status
	}

	// Send some Resources to the pending queue.， 因为开始pending== nil, 所以Poller 中for r := range in会block
	// 为什么不能用一个简单的循环呢， 为什么要用goroutine?
	// 因为！unbuffered channel are blocking synchronous
	// 所以如果不在subroutine 中进行的话， 则pending <- &Resource{url: url} will block until Poller can read from pending
	// 但是如果Poller数目小于pending中的requests,则for r := range complete不会运行，因为 main would not yet be receiving from complete.
	go func() {
		for _, url := range urls {
			pending <- &Resource{url: url}
		}
	}()

	for r := range complete {
		// use goroutines to make sure the sleep happens in parrallel
		go r.Sleep(pending) // 循环再把complelte的request, send into pending queue again
	}
}

/* Main --> init pending, complte, status
- launch N goroutines --Poller,
- fill the pending with some requests（这样poller 才不会挨饿）
- 循环把complete(也就是poller 扔给out)的refill 给pending,

Poller：
- 内部循环poll request --> send to state --> throw polled request to out

Statemonitor:
- Init a state channell
- 内不循环receive statu update (from poller) + log state in interval (两个是并行的，select)
- return a state chan to outside（所以main 可以pass 给poller)


Any single Resource pointer may only be sent on either pending or complete at any one time.
This ensures that a Resource is either being handled by a Poller goroutine or sleeping,
 but never both simultaneously.
*
*/
