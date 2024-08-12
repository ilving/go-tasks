package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func main() {
	rq, _ := http.NewRequest(http.MethodGet, "https://releases.ubuntu.com/22.04.3/ubuntu-22.04.3-desktop-amd64.iso", nil)

	c, err := http.DefaultClient.Do(rq)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Body.Close()

	buf := make([]byte, 4096)
	bl := 0
	for i := 0; ; i++ {
		n, err := c.Body.Read(buf)
		bl += n
		if err != nil {
			break
		}

		if i%50_000 == 0 {
			PrintMemUsage()
			fmt.Printf("downloaded: %d mb\n", bToMb(uint64(bl)))
		}
	}

	PrintMemUsage()

	fmt.Printf("%d\n", bl)
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
