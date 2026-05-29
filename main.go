package main

import (
	"log"
	"fmt"
	"time"
)

func main() {
	fmt.Printf("%v: Stoplight Started\n", time.Now())
	device := "wlp3s0"
	promiscuous := true

	err := openPcapReader(device, promiscuous)	
	if err != nil {
		log.Fatal(err)
	}	

}

