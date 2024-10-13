package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var d time.Duration
	flag.DurationVar(&d, "d", 5*time.Second, "duration of timer")
	var m string
	flag.StringVar(&m, "m", "", "message")

	flag.Parse()

	fmt.Printf("starting a timer for %v\n", d.Abs())

	end := time.Now().Add(d)
	tc := time.NewTicker(time.Second)

	stop := time.NewTimer(d)

	for {
		select {
		case t := <-tc.C:
			fmt.Printf("\033[1A\033[K\033[33m%s remaining\033[0m\n", t.Sub(end).Abs().Truncate(time.Second))
		case <-stop.C:
			fmt.Printf("\033[1A\033[K")
			tc.Stop()
			stop.Stop()
			fmt.Println("\033[41m\033[1m TIME IS UP \033[0m", m)
			return
		}
	}
}

// func main() {
// 	var d time.Duration
// 	flag.DurationVar(&d, "d", 5*time.Second, "duration of timer")

// 	flag.Parse()

// 	fmt.Printf("starting a timer for %v\n", d.Abs())

// 	end := time.Now().Add(d)
// 	tc := time.NewTicker(time.Second)

// 	for t := range tc.C {
// 		fmt.Printf("\033[1A\033[K")
// 		fmt.Printf("%s remaining\n", t.Sub(end).Abs())
// 		if end.Compare(t) <= 0 {
// 			tc.Stop()
// 			fmt.Println("\033[0;32mtime is up\033[0m")
// 			break
// 		}
// 	}
// }
