package main

import (
	"log"
	"fmt"
	"time"
)

func main() {
	fmt.Printf("%v: Stoplight Started\n", time.Now())
	device := "wlx00c0cab5102c"
	promiscuous := true

	err := openPcapReader(device, promiscuous)	
	if err != nil {
		log.Fatal(err)
	}	

}

