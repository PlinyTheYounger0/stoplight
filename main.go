package main

import (
	"log"
)

func main() {
	device := "wlx00c0cab5102c"
	promiscuous := true

	err := openPcapReader(device, promiscuous)	
	if err != nil {
		log.Fatal(err)
	}	

}

