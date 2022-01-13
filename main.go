package main

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	_ "net/http/pprof" // This bootstaps the Go profiling module
)

func main() {
	profilerAddr := "localhost:9000"

	// Start a goroutine that runs the profiling server...
	go func() {
		if err := http.ListenAndServe(profilerAddr, nil); err != nil {
			panic(err)
		}
	}()

	// Start another goroutine to do some work
	go func() {
		var i int

		for {
			data := fmt.Sprintf("Data %d", i)

			hash := sha256.Sum256([]byte(data))
			_ = hash

			if i%100000 == 0 {
				fmt.Printf(".")
			}

			i++
		}
	}()

	fmt.Printf("INFO: Starting profile on http://%s/debug/pprof\n", profilerAddr)
	fmt.Printf("Press ENTER to stop application....")
	fmt.Scanln()
}
